package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
	"todo/internal/config"
	"todo/internal/user"
	"todo/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()

	cnfg := config.GetConfig()

	//cfgMongo := cnfg.MongoDB
	//mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port,
	//	cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDb)
	//if err != nil {
	//	panic(err)
	//}
	//
	//storage := db.NewStorage(mongoDBClient, cnfg.MongoDB.Collection, logger)
	//
	//user1 := user.User{
	//	ID:           "",
	//	Email:        "myemail@example.com",
	//	Username:     "Maasa",
	//	PasswordHash: "hash",
	//}
	//
	//user1ID, err := storage.Create(context.Background(), user1)
	//if err != nil {
	//	panic(err)
	//}
	//logger.Info(user1ID)

	logger.Info("Register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cnfg)

}

func start(router *httprouter.Router, cnfg *config.Config) {
	logger := logging.GetLogger()
	logger.Print("Start application")

	var listener net.Listener
	var listenErr error

	if cnfg.Listen.Type == "sock" {
		logger.Info("Detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("Create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path %s", socketPath)

		logger.Info("Listen Unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("Server is listening Unix socket %s", socketPath)

	} else {
		logger.Info("Listen tcp ")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cnfg.Listen.BindIp, cnfg.Listen.Port))
		logger.Printf("Server is listening port %s:%s", cnfg.Listen.BindIp, cnfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
