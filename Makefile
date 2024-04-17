BINARY_NAME=myapp
DATABASE_HOST="aqsofzhcugpxbjtlnizp"
APIKEY="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFxc29memhjdWdweGJqdGxuaXpwIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTM0NjEyMDIsImV4cCI6MjAwOTAzNzIwMn0.bEowq4cL4nc3XkdxzCiMx7l8SjkDq_kURu79NSMdkiA"

## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -o ${BINARY_NAME} .
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env DATABASE_HOST=${DATABASE_HOST} APIKEY=${APIKEY} ./${BINARY_NAME} &
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	go test -v ./...