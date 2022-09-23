package biz

import (
	"context"

	v1 "search-product/api/search/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Searcher struct {
	Id    int64
	Name  string
	Price float64
	Desc  string
}

// GreeterRepo is a Greater repo.
type SearcherRepo interface {
	MustSearch(context.Context, *Searcher) (*Searcher, error)
}

// SearcherUsecase is a SearcherRepo usecase.
type SearcherUsecase struct {
	repo SearcherRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewSearcherUsecase(repo SearcherRepo, logger log.Logger) *SearcherUsecase {
	return &SearcherUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *SearcherUsecase) MustSearcher(ctx context.Context, g *Searcher) (*Searcher, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Id)
	return uc.repo.MustSearch(ctx, g)
}
