package server

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"multicluster/internal/errors"
)

func TestHTTPStruct(t *testing.T) {
	a := &errors.HTTPError{
		Errors: make(map[string][]string),
	}
	a.Errors["body"] = []string{"can't be empty"}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Printf("%s", string(b))
}
