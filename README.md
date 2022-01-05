  

# blog-grpc

#Generate grpc code
-------------------
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    blogpb/blog.proto

#install mongodb
--------------------
1)echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/5.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list   

2)sudo apt-get update

3)sudo apt-get install -y mongodb-org


#install grpui (to test the grpc services/apis)
-----------------------------------------------
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

#Run blog server
----------------
go run blog_server/server.go

#Test grpc services/apis through grpcui
---------------------------------------
grpcui -plaintext localhost:50051




