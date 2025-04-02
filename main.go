package main

import (
	customError "god/internal/common/error"
	"god/internal/router"
	"god/pkg/config"
	"god/pkg/logger"
	"log"
	"net/http"
	"time"
)

func init() {
	customError.InitErrMsg()
}
func main() {
	if err := config.LoadPathEnv(".env"); err != nil {
		logger.Error(err.Error())
	}
	routersInit := router.InitRouter()

	server := &http.Server{
		Addr:           ":8000",
		Handler:        routersInit,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening%s", ":8000")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
