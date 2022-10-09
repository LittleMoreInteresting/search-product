package data

import (
	"context"
	"reflect"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/olivere/elastic/v7"
	"search-product/app/search/internal/biz"
)

type searcherRepo struct {
	data *Data
	log  *log.Helper
}

// NewSearcherRepo .
func NewSearcherRepo(data *Data, logger log.Logger) biz.SearcherRepo {
	return &searcherRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *searcherRepo) MustSearch(ctx context.Context, g *biz.Searcher) ([]*biz.Searcher, error) {
	var res []*biz.Searcher
	produceEs := NewESModel(r.data.es, "product")
	var filter []elastic.Query
	if len(g.Name) > 0 {
		filter = append(filter, elastic.NewTermQuery("name", g.Name))
	}
	if g.Price > 0 {
		filter = append(filter, elastic.NewTermQuery("price", g.Price))
	}

	query, err := produceEs.MustQuery(ctx, "name,price", filter, "price desc", 1, 20)
	if err != nil {
		return nil, err
	}
	var product Product
	for _, item := range query.Each(reflect.TypeOf(product)) {
		t := item.(Product)
		res = append(res, &biz.Searcher{
			Id:    t.Id,
			Name:  t.Name,
			Price: t.Price,
		})
	}

	return res, nil
}
