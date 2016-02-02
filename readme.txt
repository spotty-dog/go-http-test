To build for alpine linux:

CGO_ENABLED=0 go build -installsuffix cgo hello.go

docker build -t go-http .
docker run -p 80:8080 --name go-http go-http

http://localhost:8080/json?pretty=true

