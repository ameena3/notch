package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"database/sql"

	pb "github.com/ameena3/notch_project/back_service/gen/product"
	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	Client interface {
		IndexProducts(ctx context.Context) ([]*pb.Product, error)
		GetProduct(ctx context.Context, name string) (*pb.Product, error)
		GetProductBySKU(ctx context.Context, sku int64) (*pb.Product, error)
		CreateCart(ctx context.Context, cart *pb.Cart) (*pb.Cart, error)
		GetCart(ctx context.Context, id int64) (*pb.Cart, error)
		GetCartProducts(ctx context.Context, id int64) (*pb.CartProducts, error)
		UpdateCart(ctx context.Context, cart *pb.Cart) (*pb.Cart, error)
		AddProductToCart(ctx context.Context, p *pb.Product, cart *pb.Cart) (*pb.Cart, error)
		RemoveProductFromCart(ctx context.Context, p *pb.Product, cart *pb.Cart) (*pb.Cart, error)
	}

	client struct {
		db *sql.DB
	}
)

// Make sure that the interface is implenented correctly
var (
	_            Client = (*client)(nil)
	inMemoryCart        = make([]*pb.Cart, 0)
)

func NewClient() *client {
	dsn := "server=" + "sql" + ";user id=" + "sa" + ";password=" + "notch@12345" + ";database=" + "notch"
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &client{
		db: db,
	}
}

func (c *client) IndexProducts(ctx context.Context) (result []*pb.Product, err error) {
	cmd :=
		`SELECT  
	 	id AS SKU, 
	 	name as Name, 
	 	description as Description, 
	 	Price as Price  
	FROM  
	 	"dbo"."usersproduct" 
	 		FOR JSON AUTO`
	rows, err := c.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	if cols == nil {
		return nil, nil
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		var d interface{}
		vals[i] = &d
		if i != 0 {
			fmt.Print("\t")
		}
		fmt.Print(cols[i])
	}
	fmt.Println()
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(vals) > 0 {
			datas := fmt.Sprintf("%v", *vals[0].(*interface{}))
			var results []*pb.Product
			data := []byte(datas)
			err = json.Unmarshal(data, &results)
			if err != nil {
				return nil, err
			}
			result = append(result, results...)
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return
}

func (c *client) GetProduct(ctx context.Context, name string) (result *pb.Product, err error) {
	results, err := c.IndexProducts(ctx)
	for _, product := range results {
		if strings.EqualFold(product.Name, name) {
			result = product
		}
	}
	if result == nil {
		err = status.Errorf(codes.NotFound, "product not found")
	}
	return
}

func (c *client) GetProductBySKU(ctx context.Context, sku int64) (result *pb.Product, err error) {
	sqlq := `
	SELECT  
		id AS SKU, 
 		name as Name, 
 		description as Description, 
 		Price as Price  
 	FROM  
 		"dbo"."usersproduct" 
 	WHERE id  = ?  FOR JSON AUTO , WITHOUT_ARRAY_WRAPPER 
 	 `
	var datas string
	err = c.db.QueryRow(sqlq, sku).Scan(
		&datas,
	)
	if err != nil {
		return nil, err
	}
	data := []byte(datas)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	if result == nil {
		err = status.Errorf(codes.NotFound, "product not found")
	}
	return
}

func (c *client) CreateCart(ctx context.Context, cart *pb.Cart) (result *pb.Cart, err error) {
	// Get the last cart id
	var lastCartID sql.NullInt32
	err = c.db.QueryRow(`SELECT MAX(id) FROM "dbo"."cart"`).Scan(&lastCartID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	lastCartID.Int32++

	sqlq := `
	INSERT INTO 
		"dbo"."cart" (
			"id",
			"name"
		) 
	VALUES (
		?,
		?
	)
	`
	_, err = c.db.ExecContext(ctx, sqlq, lastCartID.Int32, cart.Name)
	if err != nil {
		return nil, err
	}
	cart.CartId = int64(lastCartID.Int32)
	return cart, nil
}

func (c *client) GetCart(ctx context.Context, id int64) (*pb.Cart, error) {
	sqlq := `
	SELECT  
		id AS cart_id, 
		name as name,
		totalprice as total_price 
	FROM  
		"dbo"."cart" 
	WHERE id  = ?  FOR JSON AUTO , WITHOUT_ARRAY_WRAPPER 
	 `
	var datas string
	err := c.db.QueryRow(sqlq, id).Scan(
		&datas,
	)
	if err != nil {
		return nil, err
	}
	data := []byte(datas)
	var result *pb.Cart
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	if result == nil {
		err = status.Errorf(codes.NotFound, "cart not found")
	}
	return result, err
}

func (c *client) GetCartProducts(ctx context.Context, id int64) (*pb.CartProducts, error) {
	sqlq := `
	SELECT
    p.id as SKU ,
    p.name as Name ,
    p.description as Description,
    p.price as Price
 FROM  
 cart as c  
 join productcart as pc on pc.cartid = c.id  
 join usersproduct as p on p.id = pc.productid
    WHERE
        c.id = ? 
    FOR JSON AUTO 
	 `
	var datas string
	err := c.db.QueryRow(sqlq, id).Scan(
		&datas,
	)
	if err != nil {
		return nil, err
	}
	data := []byte(datas)
	var result []*pb.Product
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	if result == nil {
		err = status.Errorf(codes.NotFound, "cart not found")
	}
	return &pb.CartProducts{
		Products: result,
	}, err
}

func (c *client) UpdateCart(ctx context.Context, cart *pb.Cart) (*pb.Cart, error) {
	sqlq := `
	UPDATE 
		"dbo"."cart" 
	SET 
		"totalprice" = ? 
	WHERE id  = ? 
	`
	_, err := c.db.ExecContext(ctx, sqlq, cart.TotalPrice, cart.CartId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *client) AddProductToCart(ctx context.Context, product *pb.Product, cart *pb.Cart) (*pb.Cart, error) {
	sqlq := `
	INSERT INTO 
		"dbo"."productcart" 
	(
		productID,
		cartID
	)
	VALUES (
		 ? ,
		? 
	)
	`
	_, err := c.db.ExecContext(ctx, sqlq, product.SKU, cart.CartId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *client) RemoveProductFromCart(ctx context.Context, product *pb.Product, cart *pb.Cart) (*pb.Cart, error) {
	sqlq := `
	DELETE TOP(1) FROM 
		"dbo"."productcart" 
	WHERE productID  = ?  AND cartID = ? 
	`
	_, err := c.db.ExecContext(ctx, sqlq, product.SKU, cart.CartId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *client) DeleteCart(ctx context.Context, id int64) error {
	sqlq := `
	DELETE FROM 
		"dbo"."cart" 
	WHERE id  = ? 
	`
	_, err := c.db.ExecContext(ctx, sqlq, id)
	if err != nil {
		return err
	}
	return nil
}
