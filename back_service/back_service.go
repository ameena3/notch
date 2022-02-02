package backservice

import (
	"context"

	pb "github.com/ameena3/notch_project/back_service/gen/product"
	dbClient "github.com/ameena3/notch_project/back_service/services/database_client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	server struct {
		pb.UnimplementedProductServiceServer
		dbClient dbClient.Client
	}
	ServerParams struct {
		DbClient dbClient.Client
	}
)

// Make sure that the interface is implenented correctly
var _ pb.ProductServiceServer = (*server)(nil)

func NewServer(params ServerParams) (*server, error) {
	if params.DbClient == nil {
		return nil, status.Errorf(codes.InvalidArgument, "dbClient must be provided")
	}
	return &server{
		dbClient: params.DbClient,
	}, nil
}

func (s *server) IndexProducts(ctx context.Context, payload *pb.IndexProductsRequest) (*pb.IndexProductsResponse, error) {
	results, err := s.dbClient.IndexProducts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get products: %v", err)
	}
	return &pb.IndexProductsResponse{
		Products: results,
	}, nil
}
func (s *server) GetProduct(ctx context.Context, payload *pb.GetProductRequest) (*pb.Product, error) {
	result, err := s.dbClient.GetProduct(ctx, payload.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
	}
	return result, nil
}

func (s *server) CreateCart(ctx context.Context, payload *pb.CreateCartRequest) (result *pb.Cart, err error) {
	result, err = s.dbClient.CreateCart(ctx, &pb.Cart{
		Name: payload.Name,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create cart: %v", err)
	}
	return
}

func (s *server) AddProductToCart(ctx context.Context, payload *pb.AddProductToCartRequest) (result *pb.Cart, err error) {
	p, err := s.dbClient.GetProductBySKU(ctx, payload.GetProductSku())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
	}
	c, err := s.dbClient.GetCart(ctx, payload.GetCartId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %v", err)
	}
	// Update the cart with the total price
	c.TotalPrice = c.TotalPrice + p.Price
	c, err = s.dbClient.UpdateCart(ctx, c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update cart: %v", err)
	}
	c, err = s.dbClient.AddProductToCart(ctx, p, c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add product to cart: %v", err)
	}
	return c, nil
}

func (s *server) RemoveProductFromCart(ctx context.Context, payload *pb.RemoveProductFromCartRequest) (*pb.Cart, error) {
	p, err := s.dbClient.GetProductBySKU(ctx, payload.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
	}
	c, err := s.dbClient.GetCart(ctx, payload.GetCartId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %v", err)
	}
	// Update the cart with the total price
	c.TotalPrice = c.TotalPrice - p.Price
	c, err = s.dbClient.UpdateCart(ctx, c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update cart: %v", err)
	}
	c, err = s.dbClient.RemoveProductFromCart(ctx, p, c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to remove product from cart: %v", err)
	}
	return c, nil
}

func (s *server) GetCart(ctx context.Context, payload *pb.GetCartRequest) (c *pb.Cart, err error) {
	c, err = s.dbClient.GetCart(ctx, payload.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %v", err)
	}
	return
}

func (s *server) GetCartProducts(ctx context.Context, payload *pb.GetCartProductsRequest) (*pb.CartProducts, error) {
	c, err := s.dbClient.GetCartProducts(ctx, payload.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart products: %v", err)
	}
	return c, nil
}
