package main

import (
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/cstdev/lambdahelpers/pkg/notification"
	"github.com/cstdev/notifierlambda/pkg/scraper"
	log "github.com/sirupsen/logrus"
)

func handleRequest() (string, error) {
	log.Info("Handling request")

	region := os.Getenv("REGION")
	if region == "" {
		return "", errors.New("REGION environment variable must be set")
	}

	_, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Error("Failed to create session")
		return "", err
	}

	notification.SendMessage("TEXT", "NUMBER")

	scraper.FindText("URL", "TEXT")

	return "", nil
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.Info("Notifier")
	//lambda.Start(handleRequest)
	handleRequest()
}
