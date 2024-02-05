build-server:
	docker build -t health-check-deamon -f Dockerfile.server .

run-server:
	docker run -p 50051:50051 -it --rm --name health-check-deamon health-check-deamon

build-client:
	docker build -t health-check-client -f Dockerfile.client .

run-client:
	docker run -p 8443:8443 -e GRPC_ADDR=localhost:8443 -it --rm --name health-check-client health-check-client

gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative health_info/health_info.proto