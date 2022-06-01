package example

import (
	"context"
	"fmt"
	"multicluster/api/cluster/v1"
	"multicluster/internal/biz/constant"
	"testing"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func TestHelloWorld(t *testing.T) {
	cli, err := http.NewClient(context.Background(), http.WithEndpoint("127.0.0.1:8000"), http.WithMiddleware(tracing.Client()))
	if err != nil {
		t.Fatal(err)
	}

	clientGreeter := v1.NewClusterHTTPClient(cli)
	reply, err := clientGreeter.CreateCluster(context.Background(),
		&v1.ClusterCreateRequest{
			Name:        "test-333",
			ClusterType: string(constant.Managed),
			ClusterSpec: "ha_k8s",
			Version:     string(constant.V_1_18_20),
			RegionId:    "cd10e",
			VpcId:       "vpc1",
			ZoneId:      "gx1",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reply.ClusterId)
}
