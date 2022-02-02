package backservice

import (
	"context"
	"testing"

	"github.com/ameena3/notch_project/back_service/gen/product"
	mock_client "github.com/ameena3/notch_project/back_service/services/database_client_mock"
	"github.com/golang/mock/gomock"
)

func Test_index_products(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_db_client := mock_client.NewMockClient(ctrl)
	mock_db_client.EXPECT().IndexProducts(gomock.Any()).Return(nil, nil)
	ser := &server{
		dbClient: mock_db_client,
	}
	res, err := ser.IndexProducts(context.TODO(), &product.IndexProductsRequest{
		Count: 0,
	})
	if err != nil || res.Products != nil {
		t.Errorf("Expected nil, nil, but got %v, %v", res, err)
	}
}

func Test_get_products(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_db_client := mock_client.NewMockClient(ctrl)
	ctx := context.TODO()
	name := "testname"
	mock_db_client.EXPECT().GetProduct(ctx, name).Return(nil, nil)
	ser := &server{
		dbClient: mock_db_client,
	}
	res, err := ser.GetProduct(ctx, &product.GetProductRequest{
		Name: name,
	})
	if err != nil || res != nil {
		t.Errorf("Expected nil, nil, but got %v, %v", res, err)
	}
}

func Test_create_cart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_db_client := mock_client.NewMockClient(ctrl)
	ctx := context.TODO()
	name := "testname"
	mock_db_client.EXPECT().CreateCart(ctx, &product.Cart{
		Name: name,
	}).Return(nil, nil)
	ser := &server{
		dbClient: mock_db_client,
	}
	res, err := ser.CreateCart(ctx, &product.CreateCartRequest{
		Name: name,
	})
	if err != nil || res != nil {
		t.Errorf("Expected nil, nil, but got %v, %v", res, err)
	}
}

func Test_add_product_to_cart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_db_client := mock_client.NewMockClient(ctrl)
	ctx := context.TODO()
	name := "testname"
	sku := int64(10.0)
	cartID := int64(1)
	mock_db_client.EXPECT().GetProductBySKU(ctx, sku).Return(&product.Product{
		SKU:         sku,
		Name:        name,
		Description: "testdescription",
		Price:       10,
	}, nil)
	mock_db_client.EXPECT().GetCart(ctx, cartID).Return(&product.Cart{
		CartId:     cartID,
		TotalPrice: 0,
		Name:       name,
	}, nil)
	mock_db_client.EXPECT().UpdateCart(ctx, &product.Cart{
		CartId:     cartID,
		TotalPrice: 10,
		Name:       name,
	}).Return(&product.Cart{
		CartId:     cartID,
		TotalPrice: 10,
		Name:       name,
	}, nil)
	mock_db_client.EXPECT().AddProductToCart(ctx, &product.Product{
		SKU:         sku,
		Name:        name,
		Description: "testdescription",
		Price:       10,
	}, &product.Cart{
		CartId:     cartID,
		TotalPrice: 10,
		Name:       name,
	})
	ser := &server{
		dbClient: mock_db_client,
	}
	res, err := ser.AddProductToCart(ctx, &product.AddProductToCartRequest{
		ProductSku: sku,
		CartId:     cartID,
	})
	if err != nil || res != nil {
		t.Errorf("Expected nil, nil, but got %v, %v", res, err)
	}
}

func Test_remove_product_from_cart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_db_client := mock_client.NewMockClient(ctrl)
	ctx := context.TODO()
	name := "testname"
	sku := int64(10.0)
	cartID := int64(1)
	mock_db_client.EXPECT().GetProductBySKU(ctx, sku).Return(&product.Product{
		SKU:         sku,
		Name:        name,
		Description: "testdescription",
		Price:       10,
	}, nil)
	mock_db_client.EXPECT().GetCart(ctx, cartID).Return(&product.Cart{
		CartId:     cartID,
		TotalPrice: 10,
		Name:       name,
	}, nil)
	mock_db_client.EXPECT().UpdateCart(ctx, &product.Cart{
		CartId:     cartID,
		TotalPrice: 0,
		Name:       name,
	}).Return(&product.Cart{
		CartId:     cartID,
		TotalPrice: 0,
		Name:       name,
	}, nil)
	mock_db_client.EXPECT().RemoveProductFromCart(ctx, &product.Product{
		SKU:         sku,
		Name:        name,
		Description: "testdescription",
		Price:       10,
	}, &product.Cart{
		CartId:     cartID,
		TotalPrice: 0,
		Name:       name,
	})
	ser := &server{
		dbClient: mock_db_client,
	}
	res, err := ser.RemoveProductFromCart(ctx, &product.RemoveProductFromCartRequest{
		CartId:    cartID,
		ProductId: sku,
	})
	if err != nil || res != nil {
		t.Errorf("Expected nil, nil, but got %v, %v", res, err)
	}
}

func Test_get_cart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_db_client := mock_client.NewMockClient(ctrl)
	ctx := context.TODO()
	name := "testname"
	cartID := int64(1)
	mock_db_client.EXPECT().GetCart(ctx, cartID).Return(&product.Cart{
		CartId:     cartID,
		TotalPrice: 0,
		Name:       name,
	}, nil)
	ser := &server{
		dbClient: mock_db_client,
	}
	res, err := ser.GetCart(ctx, &product.GetCartRequest{
		Id: cartID,
	})
	if err != nil || res == nil {
		t.Errorf("Expected nil, nil, but got %v, %v", res, err)
	}
}

func Test_get_cart_products(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_db_client := mock_client.NewMockClient(ctrl)
	ctx := context.TODO()
	name := "testname"
	cartID := int64(1)
	mock_db_client.EXPECT().GetCartProducts(ctx, cartID).Return(&product.CartProducts{
		Products: []*product.Product{
			{
				SKU:         0,
				Name:        name,
				Description: "",
				Price:       0,
			},
		},
	}, nil)
	ser := &server{
		dbClient: mock_db_client,
	}
	res, err := ser.GetCartProducts(ctx, &product.GetCartProductsRequest{
		Id: cartID,
	})
	if err != nil || res == nil {
		t.Errorf("Expected nil, nil, but got %v, %v", res, err)
	}
}
