package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"orchestrator-service/datamock/datamockService"
	"orchestrator-service/datamock/stubs/datamock"

	"google.golang.org/grpc"
)





func main(){
	//fetching the server port
	port := flag.Int("port",0,"the server port");
	flag.Parse();
	log.Printf("start the first orchestrator on port %d",*port);

	//instantiating new orchestrator server
	dummyDataServer := datamockService.NewDummyDataServer();
	grpcServer := grpc.NewServer();
	datamock.RegisterDummyDataServiceServer(grpcServer,dummyDataServer);
	address := fmt.Sprintf("0.0.0.0:%d",*port);
	listener,err := net.Listen("tcp",address);

	if err!=nil{
		log.Fatal("cannot start first orchestrator:", err)
	}

	err = grpcServer.Serve(listener);

	if err != nil{
		log.Fatal("cannot start the server:", err);
	}

}