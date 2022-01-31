package orchestratorService

import (
	"context"
	"log"
	"orchestrator-service/datamock/stubs/datamock"
	"orchestrator-service/orchestrator/stubs/orchestrator"

	"google.golang.org/grpc"
)

type OrchestratorServer struct {
	orchestrator.UnimplementedOrchestratorServer
}

func NewOrchestratorServer() *OrchestratorServer {
	return &OrchestratorServer{}
}

func (server *OrchestratorServer) GetUserByName(
	ctx context.Context,
	req *orchestrator.RPCRequest,
)(*orchestrator.User,error){
	serverAddress := "0.0.0.0:9001"
	conn, err := grpc.Dial(*&serverAddress,grpc.WithInsecure());

	if err!=nil{
		log.Fatal("cannot dial server",err);
	}

	orchestratorClient := orchestrator.NewOrchestratorClient(conn);


	return orchestratorClient.GetUser(context.Background(),req);

}

func (server *OrchestratorServer) GetUser(
	ctx context.Context,
	req *orchestrator.RPCRequest,
)(*orchestrator.User,error){
		serverAddress := "0.0.0.0:10000"
	conn, err := grpc.Dial(*&serverAddress,grpc.WithInsecure());

	if err!=nil{
		log.Fatal("cannot dial server",err);
	}

	dummyDataClient := datamock.NewDummyDataServiceClient(conn);

	req_data := &datamock.RPCRequest{
		Name: req.GetName(),
	}

	res, err := dummyDataClient.GetMockUserData(context.Background(),req_data);

	if err!=nil{
		return nil,err;
	}
	res_orch := orchestrator.User{
		Name: res.Name,
		Class: res.Class,
		Roll: res.Roll,	
	}
	return &res_orch,nil;

}

