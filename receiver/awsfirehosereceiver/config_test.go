// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awsfirehosereceiver

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestLoadConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	sub, err := cm.Sub(config.NewComponentIDWithName(typeStr, "").String())
	require.NoError(t, err)
	require.NoError(t, config.UnmarshalReceiver(sub, cfg))

	assert.NoError(t, cfg.Validate())

	require.Equal(t, &Config{
		ReceiverSettings: config.NewReceiverSettings(config.NewComponentID(typeStr)),
		RecordType:       "cwmetrics",
		AccessKey:        "some_access_key",
		HTTPServerSettings: confighttp.HTTPServerSettings{
			Endpoint: "0.0.0.0:4433",
			TLSSetting: &configtls.TLSServerSetting{
				TLSSetting: configtls.TLSSetting{
					CertFile: "server.crt",
					KeyFile:  "server.key",
				},
			},
		},
	}, cfg)
}