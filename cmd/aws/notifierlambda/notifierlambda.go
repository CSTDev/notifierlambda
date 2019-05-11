package main

import (
	"errors"
	"os"
	"strconv"

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

	alertIfPresent := getEnvBool("ALERT_IF_PRESENT")

	found, err := scraper.FindText(url, searchText)
	if err != nil {
		log.Error("Failed to search site")
		return err
	}

	log.WithFields(log.Fields{
		"result": found,
	}).Info("Search result")

	if !found && alertIfPresent {
		log.Info("Not found and should alert only if it found it")
		return nil
	}

	if found && !alertIfPresent {
		log.Info("Found it and should only alert if it doesn't")
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

func getEnvBool(key string) bool {
	s := os.Getenv(key)
	if s == "" {
		return false
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return v
}
