# blog-grpc

protoc blogpb/blog.proto --go_out=plugins=grpc:.
protoc blogpb/blog.proto --go-grpc_out=plugins=grpc:.
#Hello world

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    blogpb/blog.proto

protoc -I . --grpc-gateway_out blogpb/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    blogpb/blog.proto

curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto   


curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/http.proto