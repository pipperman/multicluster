package data

import (
	"context"
	"multicluster/internal/biz"
	"multicluster/internal/data/ent/cluster"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
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
		log:  log.NewHelper(log.With(logger, "module", "repo/cluster")),
	}
}

func (r *clusterRepo) Create(ctx context.Context, c *biz.Cluster, option *biz.ClusterCreateOption) (*biz.Cluster, error) {
	item, err := r.data.db.Cluster.Create().
		SetName(c.Name).
		SetClusterSpec(c.ClusterSpec).
		SetClusterType("managed").
		SetVpcID("1c1101223").
		SetZoneID("232323").
		SetClusterID(uuid.New().String()).
		SetVersion(c.Version).
		SetRegionID("1c00033").
		SetEnableDeletionProtection(true).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	r.log.Info("Cluster was created")
	return &biz.Cluster{
		ClusterId: "",
		Name:      item.Name,
	}, nil
}

func (r *clusterRepo) Update(ctx context.Context, c *biz.Cluster, option *biz.ClusterUpdateOption) (*biz.Cluster, error) {
	return nil, nil
}

func (r *clusterRepo) Delete(ctx context.Context, option *biz.ClusterDeleteOption) error {
	return nil
}

func (r *clusterRepo) Get(ctx context.Context, option *biz.ClusterGetOption) (*biz.Cluster, error) {
	query := r.data.db.Cluster.Query()
	if option != nil {
		query = query.Where(
			cluster.ClusterID(option.ClusterId),
		)
	}

	item, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Cluster{
		Id:        item.ID,
		ClusterId: item.ClusterID,
		Name:      item.Name,
		Version:   item.Version,
	}, nil
}

func (r *clusterRepo) List(ctx context.Context, pageNum, pageSize int64, option *biz.ClusterListOption) ([]*biz.Cluster, error) {
	query := r.data.db.Cluster.Query().
		Offset(int((pageNum - 1) * pageSize)).
		Limit(int(pageSize))
	if option != nil {
		query = query.Where(
			cluster.NameHasPrefix(option.Name),
			cluster.ClusterType(option.ClusterType),
			cluster.ClusterSpec(option.ClusterSpec),
			cluster.Profile(option.Profile),
		)
	}

	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Cluster, 0)
	for _, item := range list {
		rv = append(rv, &biz.Cluster{
			Id:          item.ID,
			ClusterId:   item.ClusterID,
			Name:        item.Name,
			Version:     item.Version,
			ClusterType: item.ClusterType,
			ClusterSpec: item.ClusterSpec,
			Profile:     item.Profile,
		})
	}
	return rv, nil
}
