gen_orchestrator:
	protoc --proto_path=orchestrator/proto orchestrator/proto/*.proto --go-grpc_out=orchestrator/stubs --go_out=orchestrator/stubs
clean_orchestrator:
	rm -r orchestrator/stubs/*

gen_data:
	protoc --proto_path=datamock/proto datamock/proto/*.proto --go-grpc_out=datamock/stubs --go_out=datamock/stubs
clean_data:
	rm -r datamock/stubs/*
data-server:
	go run logic/server/dummyDataServer/Main.go -port 10000

server-1:
	go run logic/server/orchestratorServer/Main.go -port 9000
server-2:
	go run logic/server/orchestratorServer/Main.go -port 9001
client:
	go run logic/client/Main.go -address 0.0.0.0:9000