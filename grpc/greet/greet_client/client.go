package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/afasari/go-workspace/GRPC/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a Client")

	tls := true
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate authority trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

	// doServerStreaming(c)

	// doClientStreaming(c)

	// doBiDiStreaming(c)

	// doUnaryWithDeadline(c, 5*time.Second) // should complete
	// doUnaryWithDeadline(c, 1*time.Second) // should timeout

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Afas",
			LastName:  "Ari",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling unary Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a server streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Afas",
			LastName:  "Ari",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling server streaming Greet RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from Greet Many Times: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a client streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Afas",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ari",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Batiar",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling client streaming Greet RPC: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending a req %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when receiving response from 	LongGreet: %v\n", err)
	}
	fmt.Printf("LongGreet REsponse: %v\n", res)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
		return
	}
	requests := []*greetpb.GreetEveryOneRequest{
		&greetpb.GreetEveryOneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Afas",
			},
		},
		&greetpb.GreetEveryOneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ari",
			},
		},
		&greetpb.GreetEveryOneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Batiar",
			},
		},
	}
	waitc := make(chan struct{})
	// we send abuch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Starting to do a unarywithdeadline RPC...")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Afas",
			LastName:  "Ari",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("unexpected error: %v\n", statusErr.Message())
			}
		} else {
			log.Fatalf("Error while calling unary Greet RPC: %v", err)
		}
		return
	}
	log.Printf("Response from Greet: %v", res.Result)
}
