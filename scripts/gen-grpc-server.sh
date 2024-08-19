#!/bin/bash
buf generate --template docs/proto/grpc/server/gen.yaml
echo "gRPC server successfully generated in gen/grpc/server"