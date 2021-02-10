build:
	go build -o app .

deps: ## Installs dependencies using glide
	@glide --home /tmp/ install

build-docker-image:
	@docker build -t geekspace/gym-app:0.0.1 .

test:
	@go test

vet:
	@go vet

fmt:
	@go fmt

push: ## Push the images to local and remote registry
	@docker push geekspace/gym-app:0.0.1
