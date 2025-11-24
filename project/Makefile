
URL_SHORTENER=urlShortenerApp
CLIENT_SERVICE=clientApp
 
## up_build: build project and run contianer
up_build: build_url_shortener build_client
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Starting Docker images..."
	docker-compose up 
	@echo "Docker images started!"

## build_url_shortener: builds the url-shortener binary
build_url_shortener:
	@echo "Building url-shortener binary..."
	cd ../url-shortener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${URL_SHORTENER} ./cmd/api
	@echo "Done!"

build_client:
	@echo "Building client  binary..."
	cd ../client-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CLIENT_SERVICE} ./cmd/api
	@echo "Done!"


down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"