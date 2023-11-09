package client

import (
	"fmt"
	"log"
	"time"

	"github.com/mas-wig/post-api-1/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "0.0.0.0:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatal("failed to connect : %w", err)
	}

	client := pb.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	newUser := &pb.SignUpUserInput{
		Name:            "Jhon doe",
		Email:           "jhondoe@gmail.com",
		Password:        "password1234",
		PasswordConfirm: "password1234",
	}
	result, err := client.SignUpUser(ctx, newUser)
	if err != nil {
		log.Fatalf("SignUpUser :%v", err)
	}
	fmt.Println(result)
}
