package service

import (
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
}

// NewMultiClusterService new a MultiCluster service.
func NewMultiClusterService(cls *biz.ClusterUsecase, comp *biz.ComponentUsecase) *MultiClusterService {
	return &MultiClusterService{cls: cls, comp: comp}
}
