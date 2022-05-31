package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "multicluster/api/component/v1"
)

var (
	// ErrComponentNotFound is component not found.
	ErrComponentNotFound = errors.NotFound(v1.ErrorReason_COMPONENT_NOT_FOUND.String(), "component not found")
)

// Component is a Component model.
type Component struct {
	ClusterId string
	Config    string
}

type ComponentListOptions struct {
	ComponentType string
}

// ComponentRepo is a Component repo.
type ComponentRepo interface {
	Config(context.Context, *Component) error
	Get(context.Context, string) (*Component, error)
	List(context.Context, int64, int64, *ComponentListOptions) ([]*Component, error)
}

// ClusterUsecase is a Cluster usecase.
type ComponentUsecase struct {
	repo ComponentRepo
	log  *log.Helper
}

// NewComponentUsecase new a Component usecase.
func NewComponentUsecase(repo ComponentRepo, logger log.Logger) *ComponentUsecase {
	return &ComponentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCluster creates a Cluster, and returns the new Cluster.
func (comp *ComponentUsecase) Config(ctx context.Context, c *Component) error {
	comp.log.WithContext(ctx).Infof("Config Component: %v", c.ClusterId)
	return comp.repo.Config(ctx, c)
}

func (comp *ComponentUsecase) GetComponent(ctx context.Context, componentId string) (*Component, error) {
	comp.log.WithContext(ctx).Infof("Get Component: %v", componentId)
	return comp.repo.Get(ctx, componentId)
}

func (comp *ComponentUsecase) ListComponent(ctx context.Context, pageNum, pageSize int64, option *ComponentListOptions) ([]*Component, error) {
	comp.log.WithContext(ctx).Infof("List Component")
	return comp.repo.List(ctx, pageNum, pageSize, option)
}
