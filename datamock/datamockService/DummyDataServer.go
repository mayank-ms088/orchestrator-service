package datamockService

import (
	"context"
	"errors"
	"orchestrator-service/datamock/stubs/datamock"
	"strconv"
)

type DummyDataServer struct {
	datamock.UnimplementedDummyDataServiceServer
}

func NewDummyDataServer() *DummyDataServer {
	return &DummyDataServer{}
}

func (server *DummyDataServer) GetMockUserData(
	ctx context.Context,
	req *datamock.RPCRequest,
)(*datamock.User,error){
	
	if(len(req.GetName()) < 6){
		return nil,errors.New("name should be > 6 chars")
	}
	name:=req.GetName();
	class:=strconv.Itoa(len(req.GetName()));
	roll:=int64(len(name))*10;
	return &datamock.User{
		Name: name,
		Class: class,
		Roll: roll,
	},nil;

}