package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "multicluster/api/cluster/v1"
)

var (
	// ErrClusterNotFound is user not found.
	ErrClusterNotFound = errors.NotFound(v1.ErrorReason_CLUSTER_NOT_FOUND.String(), "cluster not found")
)

// Cluster is a Cluster model.
type Cluster struct {
	ClusterId string
	Name      string
}

// ClusterRepo is a Greater repo.
type ClusterRepo interface {
	Create(context.Context, *Cluster) (*Cluster, error)
	Update(context.Context, *Cluster) (*Cluster, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (*Cluster, error)
	List(context.Context) ([]*Cluster, error)
}

// ClusterUsecase is a Cluster usecase.
type ClusterUsecase struct {
	repo ClusterRepo
	log  *log.Helper
}

// NewClusterUsecase new a Cluster usecase.
func NewClusterUsecase(repo ClusterRepo, logger log.Logger) *ClusterUsecase {
	return &ClusterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCluster creates a Cluster, and returns the new Cluster.
func (uc *ClusterUsecase) CreateCluster(ctx context.Context, c *Cluster) (*Cluster, error) {
	uc.log.WithContext(ctx).Infof("Create Cluster: %v", c.ClusterId)
	return uc.repo.Create(ctx, c)
}

// CreateCluster creates a Cluster, and returns the new Cluster.
func (uc *ClusterUsecase) UpdateCluster(ctx context.Context, c *Cluster) (*Cluster, error) {
	uc.log.WithContext(ctx).Infof("Update Cluster: %v", c.ClusterId)
	return uc.repo.Create(ctx, c)
}

// CreateCluster creates a Cluster, and returns the new Cluster.
func (uc *ClusterUsecase) DeleteCluster(ctx context.Context, clusterId string) error {
	uc.log.WithContext(ctx).Infof("Create Cluster: %v", clusterId)
	return uc.repo.Delete(ctx, clusterId)
}

func (uc *ClusterUsecase) GetCluster(ctx context.Context, clusterId string) (*Cluster, error) {
	uc.log.WithContext(ctx).Infof("Get Cluster: %v", clusterId)
	return uc.repo.Get(ctx, clusterId)
}

func (uc *ClusterUsecase) ListCluster(ctx context.Context, clusterId string) (*Cluster, error) {
	uc.log.WithContext(ctx).Infof("List Cluster: %v", clusterId)
	return uc.repo.Get(ctx, clusterId)
}
