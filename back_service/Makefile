generate:
	@protoc --go_out=./gen  --go-grpc_out=./gen  ./gen/proto/product.proto

docker-image:
	@docker build -t ameena3/back_service .