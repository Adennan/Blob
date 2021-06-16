package main

import (
	"Blob/config"
	"Blob/global"
	"Blob/internal/router"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := setup(); err != nil {
		log.Fatalf("init error: %+v", err)
	}
}

func main() {
	gin.SetMode(global.Server.Mode)
	router := router.NewRouter()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", global.Server.HttpPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}

func setup() error {
	set, err := config.NewSetup()
	if err != nil {
		return err
	}

	if err = set.ReadConfig("Server", global.Server); err != nil {
		return err
	}
	if err = set.ReadConfig("App", global.App); err != nil {
		return err
	}
	if err = set.ReadConfig("Database", global.Database); err != nil {
		return err
	}

	return nil
}
