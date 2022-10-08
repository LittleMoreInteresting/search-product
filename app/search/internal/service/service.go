package service

import (
	"github.com/google/wire"
	pb "search-product/api/search/v1"
	"search-product/app/search/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSearchService)

type SearchService struct {
	pb.UnimplementedSearchServer

	uc *biz.SearcherUsecase
}
