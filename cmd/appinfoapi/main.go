package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/thomas-kbleung/appinfo/pkg/appinfo"
)

func main() {

	logrus.Info("Starting appinfo api")

	// read endpoint addresses
	address := os.Getenv("API_ENDPOINT")
	if len(address) == 0 {
		logrus.Fatal("API_ENDPOINT has not defined")
	}

	mockRepo := appinfo.NewMockRepository()
	api := appinfo.NewApi(
		appinfo.ApiConfig {
			Repo: mockRepo,
		})

	// create echo server
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())

	apiBaseURI := e.Group("/v1/app_info")

	// binding resources path to API
	apiBaseURI.GET("/:system", api.GetInfo)

	// Running echo server
	logrus.Info("Running echo server")
	go func() {
		if err := e.Start(address); err != nil && err.Error() != "http: Server closed" {
			logrus.WithError(err).WithField("address", address).Fatal("Start API echo server failure")
		}
	}()

	logrus.Info("Service is ready")

	// wait for exit signal
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	// running shutdown routine
	logrus.Printf("Waiting for gracefully shutdown echo server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	logrus.Info("Stopped appinfo api")

}
