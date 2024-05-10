# toll-calculator

// next 50

### some useful stuff
```
docker run --name kafka -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENERS=yes -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true bitnami/kafka:latest
```

## Installing protobuf compiler (protoc compiler)
For linux users or (WSL2)
```
sudo apt install -y protobuf-compiler
```

## Installing GRPC and Protobuffer plugins for Golang
1. Protobuffers
```
go get -u github.com/golang/protobuf/protoc-gen-go
```