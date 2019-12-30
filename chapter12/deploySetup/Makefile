PROJECT_NAME=bookServer
BINARY_NAME=app
GOCMD=go
GOBUILD=$(GOCMD) build

all:
	make build && make deploy

build:
	$(info Building the book server binary...)
	cd ${PROJECT_NAME} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GOBUILD} -o "$(BINARY_NAME)" -v

deploy:
	docker-compose rm -f
	docker-compose up --build -d
