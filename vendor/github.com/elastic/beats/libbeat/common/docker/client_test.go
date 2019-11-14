// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build integration

package docker

import (
	"os"
	"testing"

	"github.com/docker/docker/api"
	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestNewClient(t *testing.T) {
	host := "unix:///var/run/docker.sock"

	client, err := NewClient(host, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	_, err = client.ContainerList(context.Background(), types.ContainerListOptions{})
	assert.NoError(t, err)

	// This test only works on newer Docker versions (any supported one really)
	switch client.ClientVersion() {
	case "1.22":
		t.Skip("Docker version is too old for this test")
	case api.DefaultVersion:
		t.Logf("Using default API version: %s", api.DefaultVersion)
	default:
		t.Logf("Negotiated version: %s", client.ClientVersion())
	}

	// Test we can hardcode version
	os.Setenv("DOCKER_API_VERSION", "1.22")

	client, err = NewClient(host, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "1.22", client.ClientVersion())

	_, err = client.ContainerList(context.Background(), types.ContainerListOptions{})
	assert.NoError(t, err)
}
