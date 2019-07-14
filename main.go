package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//"net"

	//"qasir-supplier/inventory/pb"
	"qasir-supplier/merchant/database"
	"qasir-supplier/merchant/pkg/endpoint"
	"qasir-supplier/merchant/pkg/service"
	transport "qasir-supplier/merchant/pkg/transport"

	"github.com/joho/godotenv"
	//"google.golang.org/grpc"
)

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	db := database.DBInit()

	httpPort := os.Getenv("MERCHANT_HTTP_PORT")
	//grpcPort := os.Getenv("INVENTORY_GRPC_PORT")

	var (
		httpAddr = flag.String("http", ":"+httpPort, "http listen address")
		//grpcAddr = flag.String("grpc", ":"+grpcPort, "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// define our inventory services
	srvMerchant := service.NewMerchantService(db)
	srvOutlet := service.NewOutletService(db)
	srvLocation := service.NewLocationService(db)

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := endpoint.Endpoints{

		// merchant endpoint
		GetMerchantsCoverageEndpoint: endpoint.MakeGetMerchantsCoverageEndpoint(srvMerchant),
		ShowMerchantEndpoint:         endpoint.MakeShowMerchantEndpoint(srvMerchant),
		CreateMerchantEndpoint:       endpoint.MakeCreateMerchantEndpoint(srvMerchant),
		UpdateMerchantEndpoint:       endpoint.MakeUpdateMerchantEndpoint(srvMerchant),
		DeleteMerchantEndpoint:       endpoint.MakeDeleteMerchantEndpoint(srvMerchant),

		// outlet endpoint
		GetOutletsEndpoint:           endpoint.MakeGetOutletsEndpoint(srvOutlet),
		ShowOutletEndpoint:           endpoint.MakeShowOutletEndpoint(srvOutlet),
		CreateOutletEndpoint:         endpoint.MakeCreateOutletEndpoint(srvOutlet),
		CreateOutletLocationEndpoint: endpoint.MakeCreateOutletLocationEndpoint(srvOutlet),
		UpdateOutletEndpoint:         endpoint.MakeUpdateOutletEndpoint(srvOutlet),
		DeleteOutletEndpoint:         endpoint.MakeDeleteOutletEndpoint(srvOutlet),

		// location endpoint
		GetLocationsEndpoint:   endpoint.MakeGetLocationsEndpoint(srvLocation),
		ShowLocationEndpoint:   endpoint.MakeShowLocationEndpoint(srvLocation),
		CreateLocationEndpoint: endpoint.MakeCreateLocationEndpoint(srvLocation),
		UpdateLocationEndpoint: endpoint.MakeUpdateLocationEndpoint(srvLocation),
		DeleteLocationEndpoint: endpoint.MakeDeleteLocationEndpoint(srvLocation),
	}

	// Run HTTP Server
	go func() {
		log.Println("Merchant Service (http) is listening on port", *httpAddr)
		handler := transport.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// Run GRPC Server
	// go func() {
	// 	grpcListener, err := net.Listen("tcp", *grpcAddr)
	// 	if err != nil {
	// 		log.Println("Error connecting grpc server : ", err)
	// 	}

	// 	log.Println("Inventory Service (grpc) is listening on port", *grpcAddr)

	// 	defer grpcListener.Close()

	// 	handler := transport.NewGRPCServer(ctx, endpoints)
	// 	grpcServer := grpc.NewServer()

	// 	// register products server
	// 	pb.RegisterProductsServer(grpcServer, handler)

	// 	// register brands server
	// 	/*TODO*/

	// 	if err := grpcServer.Serve(grpcListener); err != nil {
	// 		log.Println("Failed To Serve", err)
	// 	}
	// }()

	log.Fatalln(<-errChan)
}
