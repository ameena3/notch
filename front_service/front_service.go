package frontservice

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"

	pb "github.com/ameena3/notch_project/back_service/gen/product"
	client "github.com/ameena3/notch_project/front_service/services/back_client"
)

type (
	FrontService struct {
		backClient client.Client
		maps       cartpMap
	}
	homePage struct {
		Title    string
		CartID   string
		Products []*pb.Product
	}
	checkoutPage struct {
		Title      string
		Price      float32
		Tax        float32
		TotalPrice float32
	}
	cartpMap struct {
		mp map[string]string
		sync.Mutex
	}
)

func NewFrontService(backClient client.Client) *FrontService {
	return &FrontService{
		backClient: backClient,
		maps: cartpMap{
			mp: map[string]string{},
		},
	}
}

func (fs *FrontService) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method, "checkout here") //get request method
	if r.Method == "GET" {
		idtoLook, err := strconv.ParseInt(r.URL.Query().Get("cart_id"), 10, 64)
		if err != nil {
			log.Println(err)
		}
		dataG, err := fs.backClient.GetCart(r.Context(), idtoLook)
		if err != nil {
			log.Println(err)
		}
		t, err := template.ParseFiles("front_service/pages/checkout.gtpl")
		if err != nil {
			log.Println(err)
		}
		data := checkoutPage{
			Title:      "Checkout Page",
			Price:     dataG.TotalPrice,
			Tax:        dataG.TotalPrice * .10,
			TotalPrice: dataG.TotalPrice + (dataG.TotalPrice * .10),
		}
		t.Execute(w, data)
	} else {
		r.ParseForm()
		cartID := r.URL.Query().Get("cart_id")
		for k, v := range r.URL.Query() {
			log.Printf("key : %s, value query : %v \n", k, v)
		}
		cartid, err := strconv.ParseInt(cartID, 10, 64)
		if err != nil {
			log.Println(err)
		}
		for k, _ := range r.Form {
			// Form has cart id as one of the keys
			if k == "cart_id" {
				continue
			}
			productid, err := strconv.ParseInt(k, 10, 64)
			if err != nil {
				log.Println(err)
			}
			_, err = fs.backClient.RemoveProductFromCart(r.Context(), cartid, productid)
			if err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, "/checkout?cart_id"+cartID, http.StatusFound)
	}
}

func (fs *FrontService) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method, "home here") //get request method
	if r.Method == "GET" {
		idtoLook, err := strconv.ParseInt(r.URL.Query().Get("cart_id"), 10, 64)
		if err != nil {
			log.Println(err)
		}
		dataG, err := fs.backClient.GetCartProducts(r.Context(), idtoLook)
		if err != nil {
			log.Println(err)
		}
		data := homePage{
			Title:    "Home Page",
			CartID:   r.URL.Query().Get("cart_id"),
			Products: dataG.Products,
		}
		t, err := template.ParseFiles("front_service/pages/home.gtpl")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, data)
	} else {
		r.ParseForm()
		cartID := r.URL.Query().Get("cart_id")
		for k, v := range r.URL.Query() {
			log.Printf("key : %s, value query : %v \n", k, v)
		}
		cartid, err := strconv.ParseInt(cartID, 10, 64)
		if err != nil {
			log.Println(err)
		}
		for k, _ := range r.Form {
			// Form has cart id as one of the keys
			if k == "cart_id" {
				continue
			}
			productid, err := strconv.ParseInt(k, 10, 64)
			if err != nil {
				log.Println(err)
			}
			_, err = fs.backClient.RemoveProductFromCart(r.Context(), cartid, productid)
			if err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, "/checkout?cart_id="+cartID, http.StatusFound)
	}
}

func (fs *FrontService) CartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method, "cart here") //get request method
	if r.Method == "GET" {
		dataG, err := fs.backClient.IndexProducts(r.Context())
		if err != nil {
			log.Println(err)
		}
		data := homePage{
			Title:    "Cart Page",
			CartID:   r.URL.Query().Get("cart_id"),
			Products: dataG,
		}
		t, err := template.ParseFiles("front_service/pages/cart.gtpl")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, data)
	} else {
		r.ParseForm()
		cartID := r.URL.Query().Get("cart_id")
		for k, v := range r.URL.Query() {
			log.Printf("key : %s, value query : %v \n", k, v)
		}
		cartid, err := strconv.ParseInt(cartID, 10, 64)
		if err != nil {
			log.Println(err)
		}
		for k, _ := range r.Form {
			// Form has cart id as one of the keys
			if k == "cart_id" {
				continue
			}
			productid, err := strconv.ParseInt(k, 10, 64)
			if err != nil {
				log.Println(err)
			}
			_, err = fs.backClient.AddProductToCart(r.Context(), cartid, productid)
			if err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, "/home?cart_id="+cartID, http.StatusFound)
	}
}

func (fs *FrontService) LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, err := template.ParseFiles("front_service/pages/login.gtpl")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		if r.Form["username"][0] == "admin" || r.Form["password"][0] == "admin" {
			c, err := fs.backClient.CreateCart(r.Context(), r.Form["username"][0])
			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/cart?cart_id="+fmt.Sprint(c.CartId), http.StatusFound)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
