package main

import (
	"api_gateway/handler"
	"api_gateway/proto"
	"context"
	"net/http"

	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/client"
	//"go-micro.dev/v4/cmd/protoc-gen-micro/plugin/micro"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/profiling/service"
	//"go-micro.dev/v4/util/ctx"
)

func main() {
	r := gin.Default()

	addressServiceTransactionOpt := client.WithAddress(":9000")
	clientSrvTransaction := grpc.NewClient()
	// grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientSrvTransaction),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.POST("/balance", handler.NewAccount().BalanceAccount)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transfer-bank", handler.NewTransaction().TransferBank)
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction", srvTransaction.Client()).
			Login(context.Background(), &proto.LoginRequest{
				Username: "Ryan",
			}, addressServiceTransactionOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// import (
// 	"api_gateway/usecase"
// 	"fmt"
// )

// func main() {
// 	login := usecase.NewLogin()
// 	auth := login.Autentikasi("admin", "admin123")
// 	fmt.Println(auth)
// }
