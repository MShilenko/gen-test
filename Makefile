.PHONY: gen-grpc-server gen-http-server

gen-grpc-server:
	sh ./scripts/gen-grpc-server.sh

gen-http-server:
	sh ./scripts/gen-http-server.sh