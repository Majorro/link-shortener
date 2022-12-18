package helpers

import (
	"os"
)

const ContextLoggerKey = "logger"
const ContextLinkKey = "link"
const ContextInputKey = "input"
const ContextAuthorIpKey = "authorIP"

type config struct {
	ShortLinkBase string `json:"shortLinkBase"`
	DbConnString  string `json:"dbConnString"`
	CorsAllowFrom string `json:"corsAllowFrom"`
}

type Input struct {
	Link string `json:"link"`
}

type Output struct {
	InputLink   string `json:"inputLink"`
	OutputLink  string `json:"outputLink"`
	Error       bool   `json:"error"`
	ErrorString string `json:"errorString"`
}

var configInstance = config{}

func GetConfig() config {
	if configInstance.ShortLinkBase == "" {
		shortLinkBase := os.Getenv("SHORT_LINK_BASE")
		if shortLinkBase == "" {
			configInstance.ShortLinkBase = "http://localhost:15001/"
		} else {
			configInstance.ShortLinkBase = shortLinkBase
		}
	}
	if configInstance.DbConnString == "" {
		postgresConnString := os.Getenv("POSTGRES_CONN_STRING")
		if postgresConnString == "" {
			configInstance.DbConnString = "postgres://user:password@localhost:5432/linkShortener"
		} else {
			configInstance.DbConnString = postgresConnString
		}
	}
	if configInstance.CorsAllowFrom == "" {
		corsAllowString := os.Getenv("ALLOW_CORS_FROM")
		if corsAllowString == "" {
			configInstance.CorsAllowFrom = "http://localhost:3000"
		} else {
			configInstance.CorsAllowFrom = corsAllowString
		}
	}
	return configInstance
}
