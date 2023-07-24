package main

import (
	"log"

	pb "github.com/izanamiah/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (

	port = ":8080"

)


func main(){
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials((insecure.NewCredentials())))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Name:[]string{"Alice","Bob","Cathy"},
	}

	// // Unary call
	// callSayHello(client)
	// // Server Stream call
	// callSayHelloServerStream(client,names)
	// //Client Stream call
	// callSayHelloClientStream(client,names)
	
	callSayHelloBidirectionalStream(client, names)


}


