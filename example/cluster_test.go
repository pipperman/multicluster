package example

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"

	"multicluster/api/cluster/v1"
	"multicluster/internal/biz/constant"
	//"multicluster/internal/middleware/auth"
)

func TestClusterClient(t *testing.T) {
	ctx := context.Background()
	cli, err := http.NewClient(ctx, http.WithEndpoint("127.0.0.1:8000"),
		http.WithMiddleware(
			tracing.Client(),
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte("ctyun"), nil
			}, jwt.WithClaims(func() jwtv4.Claims {
				return jwtv4.MapClaims{
					"username": "robin",
					"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
				}
			}),
			),
		))
	if err != nil {
		t.Fatal(err)
	}

	clientCluster := v1.NewClusterHTTPClient(cli)
	reply, err := clientCluster.CreateCluster(context.Background(),
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
