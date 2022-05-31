package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"multicluster/internal/biz"
)

var _ biz.ClusterRepo = (*clusterRepo)(nil)

type clusterRepo struct {
	data *Data
	log  *log.Helper
}

// NewClusterRepo .
func NewClusterRepo(data *Data, logger log.Logger) biz.ClusterRepo {
	return &clusterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *clusterRepo) Create(ctx context.Context, g *biz.Cluster) (*biz.Cluster, error) {
	return g, nil
}

func (r *clusterRepo) Update(ctx context.Context, g *biz.Cluster) (*biz.Cluster, error) {
	return g, nil
}

func (r *clusterRepo) Delete(context.Context, string) error {
	return nil
}

func (r *clusterRepo) Get(context.Context, string) (*biz.Cluster, error) {
	return nil, nil
}

func (r *clusterRepo) List(context.Context) ([]*biz.Cluster, error) {
	return nil, nil
}
