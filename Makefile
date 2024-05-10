obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

agg:
	@go build -o bin/agg ./aggregator
	@./bin/agg

proto:
	protoc --go_out=. --go_opt=path=source_relative --go-grpc_out=. --go-grpc_opt=path=source_relative ctypes/ptypes.proto


.PHONY: obu invoicer