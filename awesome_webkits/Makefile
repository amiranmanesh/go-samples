build:
	go build
run:
	go run ./main.go
init:
	go mod init awesome_webkits 
get:
	init
	go get -u all
check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger:
	check_install
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models