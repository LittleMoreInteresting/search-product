package data

import (
	"github.com/olivere/elastic/v7"
	"search-product/app/search/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSearcherRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	es *elastic.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	es := mustEsClient(c)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		es.Stop()
	}
	return &Data{es: es}, cleanup, nil
}

func mustEsClient(c *conf.Data) (esClient *elastic.Client) {
	//set Endpoints
	esOpts := []elastic.ClientOptionFunc{
		elastic.SetURL(c.Es.Endpoints...),
	}

	//set Auth
	if len(c.Es.Username) > 0 && len(c.Es.Password) > 0 {
		esOpts = append(esOpts, elastic.SetBasicAuth(c.Es.Username, c.Es.Password))
	}

	//set Sniff
	esOpts = append(esOpts, elastic.SetSniff(c.Es.Sniff))
	if c.Es.Sniff {
		if c.Es.SniffTimeout != nil {
			esOpts = append(esOpts, elastic.SetSnifferTimeout(c.Es.SniffTimeout.AsDuration()))
		}
		if c.Es.SniffInterval != nil {
			esOpts = append(esOpts, elastic.SetSnifferInterval(c.Es.SniffInterval.AsDuration()))
		}
	}
	esClient, _ = elastic.NewClient(esOpts...)
	return
}
