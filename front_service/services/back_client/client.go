package client

import (
	"context"
	"log"

	pb "github.com/ameena3/notch_project/back_service/gen/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	Client interface {
		IndexProducts(ctx context.Context) ([]*pb.Product, error)
		CreateCart(ctx context.Context, name string) (*pb.Cart, error)
		AddProductToCart(ctx context.Context, cartID int64, productID int64) (*pb.Cart, error)
		GetCartProducts(ctx context.Context, cartID int64) (*pb.CartProducts, error)
		RemoveProductFromCart(ctx context.Context, cartID int64, productID int64) (*pb.Cart, error)
		GetCart(ctx context.Context, cartID int64) (*pb.Cart, error)
	}

	client struct {
		client pb.ProductServiceClient
	}
)

// Make sure that the interface is implenented correctly
var _ Client = (*client)(nil)

func NewClient() *client {
	// Set up a connection to the server.
	conn, err := grpc.Dial("backservice:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &client{
		client: pb.NewProductServiceClient(conn),
	}
}

func (c *client) IndexProducts(ctx context.Context) ([]*pb.Product, error) {
	// Contact the server and print out its response.
	r, err := c.client.IndexProducts(ctx, &pb.IndexProductsRequest{})
	if err != nil {
		return nil, err
	}
	return r.Products, nil
}

func (c *client) CreateCart(ctx context.Context, name string) (*pb.Cart, error) {
	// Contact the server and print out its response.
	r, err := c.client.CreateCart(ctx, &pb.CreateCartRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *client) AddProductToCart(ctx context.Context, cartID int64, productID int64) (*pb.Cart, error) {
	// Contact the server and print out its response.
	r, err := c.client.AddProductToCart(ctx, &pb.AddProductToCartRequest{
		CartId:     cartID,
		ProductSku: productID,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *client) GetCartProducts(ctx context.Context, cartID int64) (*pb.CartProducts, error) {
	// Contact the server and print out its response.
	r, err := c.client.GetCartProducts(ctx, &pb.GetCartProductsRequest{
		Id: cartID,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *client) RemoveProductFromCart(ctx context.Context, cartID int64, productID int64) (*pb.Cart, error) {
	// Contact the server and print out its response.
	r, err := c.client.RemoveProductFromCart(ctx, &pb.RemoveProductFromCartRequest{
		CartId:    cartID,
		ProductId: productID,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *client) GetCart(ctx context.Context, cartID int64) (*pb.Cart, error) {
	// Contact the server and print out its response.
	r, err := c.client.GetCart(ctx, &pb.GetCartRequest{
		Id: cartID,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}
