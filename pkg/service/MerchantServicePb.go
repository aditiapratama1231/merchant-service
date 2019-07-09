package service

// import (
// 	"context"
// 	"log"
// 	pbinventory "qasir-supplier/merchant/pb/inventory"
// 	"time"

// 	"google.golang.org/grpc"
// )

// type MerchantService interface {
// 	GetMerchants(ctx context.Context) (string, error)
// }

// type merchantService struct{}

// func NewMerchantService() MerchantService {
// 	return merchantService{}
// }

// // GetProduct from Inventory Service
// func GetProduct(client pbinventory.ProductsClient, point *pbinventory.ProductRequest) string {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	feature, err := client.GetProducts(ctx, point)
// 	if err != nil {
// 		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
// 	}
// 	return feature.Message
// }

// // Get will return today's date
// func (merchantService) GetMerchants(ctx context.Context) (string, error) {
// 	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
// 	if err != nil {
// 		panic(err)
// 	}

// 	// client := pb.NewAddClient(conn)
// 	client := pbinventory.NewProductsClient(conn)
// 	product := GetProduct(client, &pbinventory.ProductRequest{Id: 1})
// 	return "Merchants Successfully Loaded" + " and " + product, nil
// }
