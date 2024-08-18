#!/bin/bash
SERVER_FOLDER=gen/http-server
SWAGGER_FOLDER=docs/swagger/http/server

mkdir -p $SERVER_FOLDER
for dir in $(find $SWAGGER_FOLDER -mindepth 1 -type d -exec basename {} \;); do
    mkdir -p $SERVER_FOLDER/$dir
    oapi-codegen -package $dir -config $SWAGGER_FOLDER/gen.yaml $SWAGGER_FOLDER/$dir/swagger.yml > $SERVER_FOLDER/$dir/$dir.go
done

echo "HTTP server successfully generated in $SERVER_FOLDER"
