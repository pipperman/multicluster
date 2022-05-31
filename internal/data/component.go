package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"multicluster/internal/biz"
)

var _ biz.ComponentRepo = (*componentRepo)(nil)

type componentRepo struct {
	data *Data
	log  *log.Helper
}

// NewComponentRepo .
func NewComponentRepo(data *Data, logger log.Logger) biz.ComponentRepo {
	return &componentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *componentRepo) Config(ctx context.Context, c *biz.Component) error {
	return nil
}

func (r *componentRepo) Get(context.Context, string) (*biz.Component, error) {
	return nil, nil
}

func (r *componentRepo) List(ctx context.Context, pageNum, pageSize int64, options *biz.ComponentListOptions) ([]*biz.Component, error) {
	return nil, nil
}
