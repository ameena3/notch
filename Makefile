minikube:
	@minikube start -p notchdev --kubernetes-version=v1.21.0 

services: minikube
	@kubectl apply -f ./services
	@kubectl rollout status deployment sql
	./db_setup/database_setup.sh

instructions: services
	@echo "Run the following commands to get started:"