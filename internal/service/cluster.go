package service

import (
	"context"
	v1 "multicluster/api/cluster/v1"
	"multicluster/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/copier"
)

// CreateCluster implements multicluster.ClusterServer.
func (s *MultiClusterService) CreateCluster(ctx context.Context, in *v1.ClusterCreateRequest) (*v1.ClusterCreateReply, error) {
	cluster := &biz.Cluster{}
	copier.Copy(cluster, in)
	c, err := s.cls.CreateCluster(ctx, cluster, &biz.ClusterCreateOption{})
	if err != nil {
		return nil, errors.InternalServer(err.Error(), "cluster create error").WithMetadata(map[string]string{"a": "b"})

	}

	return &v1.ClusterCreateReply{ClusterId: "Hello " + c.Name}, nil
}

// GetCluster implements multicluster.ClusterServer.
func (s *MultiClusterService) GetCluster(ctx context.Context, in *v1.ClusterRequest) (*v1.ClusterReply, error) {
	c, err := s.cls.GetCluster(ctx, &biz.ClusterGetOption{
		ClusterId: in.ClusterId,
	})
	if err != nil {
		return nil, err
	}
	metadata := &v1.ClusterMetadata{}
	copier.Copy(metadata, c)

	return &v1.ClusterReply{Metadata: metadata}, nil
}
