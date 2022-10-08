package service

import (
	"context"

	pb "search-product/api/search/v1"
	"search-product/app/search/internal/biz"
)

func NewSearchService(uc *biz.SearcherUsecase) *SearchService {
	return &SearchService{uc: uc}
}

func (s *SearchService) MustSearch(ctx context.Context, req *pb.MustSearchRequest) (*pb.MustSearchReply, error) {
	searcher, err := s.uc.MustSearcher(ctx, &biz.Searcher{
		Name:  req.Product.Name,
		Price: float64(req.Product.Price),
		Group: req.Product.Group,
	})
	if err != nil {
		return nil, err
	}
	var Product []*pb.Product
	for _, s := range searcher {
		Product = append(Product, &pb.Product{
			Name:  s.Name,
			Price: float32(s.Price),
		})
	}

	return &pb.MustSearchReply{
		Product: Product,
	}, nil
}
