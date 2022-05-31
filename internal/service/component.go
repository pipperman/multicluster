package service

import (
	"context"

	v1 "multicluster/api/component/v1"
	"multicluster/internal/biz"
)

// Config implements multicluster.ClusterServer.
func (s *MultiClusterService) Config(ctx context.Context, in *v1.ComponentRequest) (*v1.ComponentConfigReply, error) {
	err := s.comp.Config(ctx, &biz.Component{ClusterId: in.ClusterId})
	if err != nil {
		return nil, err
	}
	return &v1.ComponentConfigReply{}, nil
}

// GetComponent implements multicluster.ComponentServer.
func (s *MultiClusterService) GetComponent(ctx context.Context, in *v1.ComponentRequest) (*v1.ComponentReply, error) {
	c, err := s.comp.GetComponent(ctx, in.ComponentId)
	if err != nil {
		return nil, err
	}
	return &v1.ComponentReply{Component: &v1.ComponentMetadata{Config: c.Config}}, nil
}

// ListComponent implements multicluster.ClusterServer.
func (s *MultiClusterService) ListComponent(ctx context.Context, in *v1.ComponentListRequest) (*v1.ComponentListReply, error) {
	_, err := s.comp.ListComponent(ctx, in.PageInfo.PageNum, in.PageInfo.PageSize, &biz.ListOptions{
		ComponentType: in.Option.ComponentType,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ComponentListReply{ClusterId: in.ClusterId, Components: []*v1.ComponentMetadata{}, PageInfo: in.PageInfo}, nil
}
