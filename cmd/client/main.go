package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	trainv1 "github.com/Shuv1Wolf/train_protos/gen/go"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := trainv1.NewAdderClient(conn)
	res, err := client.Add(context.Background(), &trainv1.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("result:", res.GetResult())
}
