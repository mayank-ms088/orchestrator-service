package orchestratorService

import (
	"context"
	"errors"
	"orchestrator-service/orchestrator/stubs/orchestrator"
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
	return nil,errors.New("not implemented yet. <yourname> will implement me");

}

