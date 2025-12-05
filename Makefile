APP_NAME=passwot

all: build clean

build:
	echo "Building $(APP_NAME)..."
	$(MAKE) test
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

test:
	echo "Testing $(APP_NAME)..."
	go mod tidy
	go test ./...
	go vet ./...
	go test ./...

clean:
	go clean
	rm -f $(APP_NAME)
