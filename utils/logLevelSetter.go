package utils

import (
	log "github.com/sirupsen/logrus"
)

func SetLogLevel(env string) {
	if env == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	if env == "development" {
		log.SetFormatter(&log.TextFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
	}
}
