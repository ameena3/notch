// Code generated by MockGen. DO NOT EDIT.
// Source: ./back_service/services/database_client/client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	reflect "reflect"

	product "github.com/ameena3/notch_project/back_service/gen/product"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddProductToCart mocks base method.
func (m *MockClient) AddProductToCart(ctx context.Context, p *product.Product, cart *product.Cart) (*product.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProductToCart", ctx, p, cart)
	ret0, _ := ret[0].(*product.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProductToCart indicates an expected call of AddProductToCart.
func (mr *MockClientMockRecorder) AddProductToCart(ctx, p, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProductToCart", reflect.TypeOf((*MockClient)(nil).AddProductToCart), ctx, p, cart)
}

// CreateCart mocks base method.
func (m *MockClient) CreateCart(ctx context.Context, cart *product.Cart) (*product.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", ctx, cart)
	ret0, _ := ret[0].(*product.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockClientMockRecorder) CreateCart(ctx, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockClient)(nil).CreateCart), ctx, cart)
}

// GetCart mocks base method.
func (m *MockClient) GetCart(ctx context.Context, id int64) (*product.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", ctx, id)
	ret0, _ := ret[0].(*product.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockClientMockRecorder) GetCart(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockClient)(nil).GetCart), ctx, id)
}

// GetCartProducts mocks base method.
func (m *MockClient) GetCartProducts(ctx context.Context, id int64) (*product.CartProducts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartProducts", ctx, id)
	ret0, _ := ret[0].(*product.CartProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartProducts indicates an expected call of GetCartProducts.
func (mr *MockClientMockRecorder) GetCartProducts(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartProducts", reflect.TypeOf((*MockClient)(nil).GetCartProducts), ctx, id)
}

// GetProduct mocks base method.
func (m *MockClient) GetProduct(ctx context.Context, name string) (*product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, name)
	ret0, _ := ret[0].(*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockClientMockRecorder) GetProduct(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockClient)(nil).GetProduct), ctx, name)
}

// GetProductBySKU mocks base method.
func (m *MockClient) GetProductBySKU(ctx context.Context, sku int64) (*product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductBySKU", ctx, sku)
	ret0, _ := ret[0].(*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductBySKU indicates an expected call of GetProductBySKU.
func (mr *MockClientMockRecorder) GetProductBySKU(ctx, sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductBySKU", reflect.TypeOf((*MockClient)(nil).GetProductBySKU), ctx, sku)
}

// IndexProducts mocks base method.
func (m *MockClient) IndexProducts(ctx context.Context) ([]*product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexProducts", ctx)
	ret0, _ := ret[0].([]*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexProducts indicates an expected call of IndexProducts.
func (mr *MockClientMockRecorder) IndexProducts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexProducts", reflect.TypeOf((*MockClient)(nil).IndexProducts), ctx)
}

// RemoveProductFromCart mocks base method.
func (m *MockClient) RemoveProductFromCart(ctx context.Context, p *product.Product, cart *product.Cart) (*product.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProductFromCart", ctx, p, cart)
	ret0, _ := ret[0].(*product.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveProductFromCart indicates an expected call of RemoveProductFromCart.
func (mr *MockClientMockRecorder) RemoveProductFromCart(ctx, p, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProductFromCart", reflect.TypeOf((*MockClient)(nil).RemoveProductFromCart), ctx, p, cart)
}

// UpdateCart mocks base method.
func (m *MockClient) UpdateCart(ctx context.Context, cart *product.Cart) (*product.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, cart)
	ret0, _ := ret[0].(*product.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockClientMockRecorder) UpdateCart(ctx, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockClient)(nil).UpdateCart), ctx, cart)
}
