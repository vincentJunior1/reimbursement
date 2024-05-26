package main

import (
	"os"
	"os/signal"
	"reimbursement/controllers"
	"reimbursement/helper"
	"reimbursement/middleware"
	"reimbursement/repository"
	"reimbursement/routers"
	"reimbursement/usecase"
	"syscall"
	"google.golang.org/grpc"
	pb "reimbursement/proto-go"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load(".env")
	logging := helper.InitializeNewLogs()
	mid := middleware.InitMiddleware(logging)
	db := repository.InitializeMysqlDatabase(logging)
	uc := usecase.InitializeV1Usecase(db, logging)
	controller := controllers.InitializeV1Controller(uc, logging)
	route := routers.InitializeRouter(controller, mid, logging)

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- route.StartServer()
	}()

	grpc := grpc.NewServer()

	pb.

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		log.Println("got an interrupt, exiting...")
	case err := <-serverErr:
		if err != nil {
			log.Println("error while running api, exiting...", err)
		}
	}
}
