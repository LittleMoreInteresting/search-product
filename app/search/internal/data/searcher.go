package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
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
func (r *searcherRepo) MustSearch(ctx context.Context, g *biz.Searcher) (*biz.Searcher, error) {
	return g, nil
}
