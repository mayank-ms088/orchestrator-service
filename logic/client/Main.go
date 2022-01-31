package main

import (
	"context"
	"flag"
	"log"
	"orchestrator-service/orchestrator/stubs/orchestrator"

	"google.golang.org/grpc"
)

func main(){
	serverAddress := flag.String("address","","the Server Address");
	flag.Parse();
	log.Printf("dial Server %s",*serverAddress);

	conn, err := grpc.Dial(*serverAddress,grpc.WithInsecure());

	if err!=nil{
		log.Fatal("cannot dial server",err);
	}

	orchestratorClient := orchestrator.NewOrchestratorClient(conn);

	req := &orchestrator.RPCRequest{
		Name: "Mayank Sharma",
	}

	res, err := orchestratorClient.GetUserByName(context.Background(),req);

	if err!=nil{
		log.Fatal(err);
		return;
	}
	log.Println(res);



	

	
}