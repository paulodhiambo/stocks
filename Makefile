 # Variables
 IMAGE_NAME=dimingo/stocks:latest
 CONTAINER_NAME=stocks_container

go_install:
	go mod tidy
 # Generate wiring code
wire_gen: go_install
	wire gen stocks/config

 # Build the Go application
build: go_install
	go build -o ./stocks/build

# Build the Docker image
package: build
	docker build -t $(IMAGE_NAME) .

 # Run the Docker container
run: package
	docker run --name $(CONTAINER_NAME) -d $(IMAGE_NAME)

 # Stop and remove the Docker container
stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

.PHONY: wire_gen build package run stop
