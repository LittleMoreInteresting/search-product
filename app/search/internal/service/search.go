package service

import (
	"context"

	pb "search-product/api/search/v1"
	"search-product/app/search/internal/biz"
)

type SearchService struct {
	pb.UnimplementedSearchServer

	uc *biz.SearcherUsecase
}

func NewSearchService(uc *biz.SearcherUsecase) *SearchService {
	return &SearchService{uc: uc}
}

func (s *SearchService) MustSearch(ctx context.Context, req *pb.MustSearchRequest) (*pb.MustSearchReply, error) {
	return &pb.MustSearchReply{}, nil
}
