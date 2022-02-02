minikube:
	@minikube start -p notchdev --kubernetes-version=v1.21.0 

services: minikube 
	@kubectl apply -f ./services
	@kubectl rollout status deployment sql
	./db_setup/database_setup.sh

test:
	@go fmt ./...
	@go test -v -cover ./...

generate:
	@make -C back_service generate
	@go run github.com/golang/mock/mockgen@v1.6.0 -source ./back_service/services/database_client/client.go  -destination ./back_service/services/database_client_mock/client_mock.go

docker-images: generate
	@make -C back_service docker-image
	@make -C front_service docker-image

instructions: services
	@echo "Run the following commands to get started:"