docker-image: 
	@rm -rf ./gen/
	@cp -r ../back_service/gen/ ./gen/
	@docker build -t ameena3/front_service .
	@rm -rf ./gen/