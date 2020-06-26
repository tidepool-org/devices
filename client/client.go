package main

import (
	"context"
	"fmt"
	"github.com/tidepool-org/devices/api"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewDevicesClient(conn)

	ctx := context.Background()
	omnipod, err  := c.GetPumpById(ctx, &api.GetPumpByIdRequest{Id: "6678c377-928c-49b3-84c1-19e2dafaff8d"})
	fmt.Printf("%v %v\n", omnipod, err)
	pumpList, err  := c.ListPumps(ctx, &api.ListPumpsRequest{})
	fmt.Printf("%v %v\n", pumpList, err)
	g6, err  := c.GetCgmById(ctx, &api.GetCgmByIdRequest{Id: "d25c3f1b-a2e8-44e2-b3a3-fd07806fc245"})
	fmt.Printf("%v %v\n", g6, err)
	cgmList, err  := c.ListCgms(ctx, &api.ListCgmsRequest{})
	fmt.Printf("%v %v\n", cgmList, err)
}
