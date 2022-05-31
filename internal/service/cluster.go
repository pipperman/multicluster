package service

import (
	"context"

	v1 "multicluster/api/cluster/v1"
	"multicluster/internal/biz"
)

// CreateCluster implements multicluster.ClusterServer.
func (s *MultiClusterService) CreateCluster(ctx context.Context, in *v1.ClusterCreateRequest) (*v1.ClusterCreateReply, error) {
	c, err := s.cls.CreateCluster(ctx, &biz.Cluster{Name: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.ClusterCreateReply{ClusterId: "Hello " + c.Name}, nil
}

// GetCluster implements multicluster.ClusterServer.
func (s *MultiClusterService) GetCluster(ctx context.Context, in *v1.ClusterRequest) (*v1.ClusterReply, error) {
	_, err := s.cls.GetCluster(ctx, in.ClusterId)
	if err != nil {
		return nil, err
	}
	return &v1.ClusterReply{}, nil
}
