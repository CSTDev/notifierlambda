package main

import (
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/cstdev/lambdahelpers/pkg/notification"
	"github.com/cstdev/notifierlambda/pkg/scraper"
	log "github.com/sirupsen/logrus"
)

func handleRequest() error {
	log.Info("Handling request")

	region := os.Getenv("REGION")
	if region == "" {
		return errors.New("REGION environment variable must be set")
	}

	url := os.Getenv("URL")
	if url == "" {
		return errors.New("URL environment variable must be set")
	}

	searchText := os.Getenv("SEARCH_TEXT")
	if searchText == "" {
		return errors.New("SEARCH_TEXT environment variable must be set")
	}

	phoneNumber := os.Getenv("PHONE_NUMBER")
	if phoneNumber == "" {
		return errors.New("PRHONE_NUMBER environment variable must be set")
	}

	result, _ := scraper.FindText(url, searchText)
	log.WithFields(log.Fields{
		"result": result,
	}).Info("Search result")

	if result {
		log.Info("Found text, returning.")
		return nil
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Error("Failed to create session")
		return err
	}

	sms := notification.SMS{
		Client: sns.New(sess),
	}

	err = sms.SendMessage(searchText+" Not Found", phoneNumber)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed to send text")
		return err
	}

	return nil
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
	lambda.Start(handleRequest)
	// handleRequest()
}
