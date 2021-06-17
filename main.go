package main

import (
	"Blob/config"
	"Blob/global"
	"Blob/internal/router"
	"Blob/utils/logger"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	if err := setup(); err != nil {
		log.Fatalf("init error: %+v", err)
	}
}

func main() {
	// DB
	var err error
	if global.DB, err = global.NewDB(global.Database); err != nil {
		panic(err)
	}

	// Logger
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.App.LogPath + "/" + global.App.LogFileName + global.App.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

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

	if err = set.ReadConfig("Server", &global.Server); err != nil {
		return err
	}
	if err = set.ReadConfig("App", &global.App); err != nil {
		return err
	}
	if err = set.ReadConfig("Database", &global.Database); err != nil {
		return err
	}

	return nil
}
