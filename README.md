# blog-grpc

protoc blogpb/blog.proto --go_out=plugins=grpc:.
protoc blogpb/blog.proto --go-grpc_out=plugins=grpc:.
#Hello world

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    blogpb/blog.proto

curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto >    