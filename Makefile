gen_orchestrator:
	protoc --proto_path=orchestrator/proto orchestrator/proto/*.proto --go-grpc_out=orchestrator/stubs --go_out=orchestrator/stubs
clean_orchestrator:
	rm -r orchestrator/stubs/*
server-1:
	go run server/Main.go -port 9000
client-1:
	go run client/Main.go -address 0.0.0.0:9000
