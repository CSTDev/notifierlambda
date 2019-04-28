package scraper

import (
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func FindText(URL string, text string) (bool, error) {
	log.WithFields(log.Fields{
		"URL":  URL,
		"text": text,
	}).Debug("Looking for text")

	response, err := http.Get(URL)
	if err != nil {
		log.Error("Unable to get URL")
		return false, err
	}

	defer response.Body.Close()

	// Get response body
	dataInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("Unable to get response body")
		return false, err
	}
	// Find the text in content
	pageContent := string(dataInBytes)
	return strings.Contains(pageContent, text), nil
}
