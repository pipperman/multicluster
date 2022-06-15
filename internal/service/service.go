package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	cluster "multicluster/api/cluster/v1"
	component "multicluster/api/component/v1"
	"multicluster/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMultiClusterService)

// MultiClusterService is a MultiCluster service.
type MultiClusterService struct {
	cluster.UnimplementedClusterServer
	component.UnimplementedComponentServer

	cls  *biz.ClusterUsecase
	comp *biz.ComponentUsecase
	log  *log.Helper
}

// NewMultiClusterService new a MultiCluster service.
func NewMultiClusterService(cls *biz.ClusterUsecase, comp *biz.ComponentUsecase, logger log.Logger) *MultiClusterService {
	return &MultiClusterService{
		cls:  cls,
		comp: comp,
		log:  log.NewHelper(log.With(logger, "multicluster", "service/multicluster")),
	}
}
