package scraper

import (
	log "github.com/sirupsen/logrus"
)

func FindText(URL string, text string) bool {
	log.WithFields(log.Fields{
		"URL":  URL,
		"text": text,
	}).Debug("Looking for text")

	return false
}
