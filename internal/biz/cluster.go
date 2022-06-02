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
	Id                       int64
	ClusterId                string
	Name                     string
	Version                  string
	ClusterType              string
	ClusterSpec              string
	UpgradeableVersion       string
	EnableDeletionProtection bool
	DockerVersion            string
	Eip                      string
	ElbId                    string
	MasterUrl                string
	Metadata                 string
	NetWorkMode              string
	Profile                  string
	RegionId                 string
	VpcId                    string
	ZoneId                   string
	SecurityGroupId          string
	ClusterSize              int
}

type ClusterGetOption struct {
	ClusterId string
}

type ClusterListOption struct {
	Name        string
	ClusterType string
	ClusterSpec string
	Profile     string
	RegionId    string
}

type ClusterDeleteOption struct {
	Id        int64
	ClusterId string
}

type ClusterCreateOption struct {
	Id        int64
	ClusterId string
}

type ClusterUpdateOption struct {
}

// ClusterRepo is a Greater repo.
type ClusterRepo interface {
	Create(context.Context, *Cluster, *ClusterCreateOption) (*Cluster, error)
	Update(context.Context, *Cluster, *ClusterUpdateOption) (*Cluster, error)
	Delete(context.Context, *ClusterDeleteOption) error
	Get(context.Context, *ClusterGetOption) (*Cluster, error)
	List(context.Context, int64, int64, *ClusterListOption) ([]*Cluster, error)
}

// ClusterUsecase is a Cluster usecase.
type ClusterUsecase struct {
	repo ClusterRepo
	log  *log.Helper
}

// NewClusterUsecase new a Cluster usecase.
func NewClusterUsecase(repo ClusterRepo, logger log.Logger) *ClusterUsecase {
	return &ClusterUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/cluster"))}
}

// CreateCluster creates a Cluster, and returns the new Cluster.
func (uc *ClusterUsecase) CreateCluster(ctx context.Context, c *Cluster, option *ClusterCreateOption) (*Cluster, error) {
	uc.log.WithContext(ctx).Infof("Create Cluster: %v", c.Name)
	return uc.repo.Create(ctx, c, option)
}

// CreateCluster creates a Cluster, and returns the new Cluster.
func (uc *ClusterUsecase) UpdateCluster(ctx context.Context, c *Cluster, option *ClusterUpdateOption) (*Cluster, error) {
	uc.log.WithContext(ctx).Infof("Update Cluster: %v", c.ClusterId)
	return uc.repo.Update(ctx, c, option)
}

// CreateCluster creates a Cluster, and returns the new Cluster.
func (uc *ClusterUsecase) DeleteCluster(ctx context.Context, option *ClusterDeleteOption) error {
	uc.log.WithContext(ctx).Infof("Delete Cluster: %v", option.ClusterId)
	return uc.repo.Delete(ctx, option)
}

func (uc *ClusterUsecase) GetCluster(ctx context.Context, option *ClusterGetOption) (*Cluster, error) {
	uc.log.WithContext(ctx).Infof("Get Cluster: %v", option.ClusterId)
	return uc.repo.Get(ctx, option)
}

func (uc *ClusterUsecase) ListCluster(ctx context.Context, pageNum, pageSize int64, options *ClusterListOption) ([]*Cluster, error) {
	uc.log.WithContext(ctx).Infof("List Cluster")
	return uc.repo.List(ctx, pageNum, pageSize, options)
}
