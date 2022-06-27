package common

import (
	"github.com/meilisearch/meilisearch-go"
	log "github.com/sirupsen/logrus"
)

var (
	Client = &meilisearch.Client{}
)

type MeiliConfig struct {
	Url    string
	ApiKey string
}

func Init(c *MeiliConfig) {
	Client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   c.Url,
		APIKey: c.ApiKey,
	})
	_, err := Client.GetKeys()
	if err != nil {
		panic(err)
	}
	log.Info("MeiliSearch Api Success")
}
