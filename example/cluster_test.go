package example

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"multicluster/api/cluster/v1"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	cli, err := http.NewClient(context.Background(), http.WithEndpoint("127.0.0.1:8000"))
	if err != nil {
		t.Fatal(err)
	}
	clientGreeter := v1.NewClusterHTTPClient(cli)
	reply, err := clientGreeter.CreateCluster(context.Background(), &v1.ClusterCreateRequest{Name: "test-333",ClusterSpec:"ha_k8s",Version:"333"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reply.ClusterId)
}
