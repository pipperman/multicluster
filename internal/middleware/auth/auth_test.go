package auth

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"

	"multicluster/internal/conf"
)

func TestGenerateToken(t *testing.T) {
	v := getConfig()
	tk := GenerateToken(v.Jwt.Secret, "robin")
	spew.Dump(tk)
}

func getConfig() *conf.Bootstrap {
	path := getConfigAbsolutePath()
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var v conf.Bootstrap
	err := c.Scan(&v)
	if err != nil {
		panic(err)
	}

	return &v
}

func getConfigAbsolutePath() string {
	workProject := "multicluster"
	configPath := path.Join(workProject, "configs/config.yaml")
	filePath, _ := os.Getwd()
	index := strings.LastIndex(filePath, workProject)
	absolutePath := strings.Replace(filePath[0:index], `\`, `/`, -1)
	return absolutePath + configPath
}
