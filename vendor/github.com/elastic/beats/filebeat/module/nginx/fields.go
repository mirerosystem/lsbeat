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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package nginx

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "nginx", asset.ModuleFieldsPri, AssetNginx); err != nil {
		panic(err)
	}
}

// AssetNginx returns asset data.
// This is the base64 encoded gzipped contents of module/nginx.
func AssetNginx() string {
	return "eJy0l8Fu4zgMhu95CqLn1g+QwwKLAgV62MViT3tzFYl2iMqil6Q7k7cf2Enb1JETO+nolNrm/38UKVV6gFfcrSHVlH6uAIws4hru/u7/vlsBBFQv1BpxWsMfKwCAvzh0EaFigdaJUqrBtghDCESuoaKIWqwAdMtipedUUb0Gkw5XABVhDLoepB4guQY/7fthuxbXUAt37eFJhqEfT4MQVMLNFEA/jv2OPZ33qPrxOGd8xrwfj5zMUdKDxTAjnyB7/Z7nAyWHc4wk2LBhSW0ZSe3LJ+94TsTtRm/OIPbjz7SPAq4ODvD8D7gQBFVRC3g2IAUHvSls0LtOEWh46LlpOIExUPKxC3gPG1QKqEOmPhKmMSgcyd9/sdrXaosuoChEekV4+e/hieWHk4Ch//VSnKj9iy6Ccid+ACcFQTUWDD3Xy/5NQe1RaHZ2Nxx2pWKyYrMz1Pz0RnLjN62z7Rq2Zm0hqC0nxaLXyso0VIvbV+LQ76cgnaKU/c+FCH1ckYmb49mgbTlcl/P/HaoVWYVZ6UpcmqjEgoVqSm4cOsewxy7fUJQ4XZNxPnSO83t/lJ7D0up+bTA1Z53mdOZxVCiCcku9JzTm2Lv6dFeY09zlELi09NNrbJpjvMvDxM58LBnwjfy4GudTy6a318kt5HNZjmkmwhey3AjBOZOFCKxF1cWY2xCXoUwpLOe5FSXfv9ewTCstWg01Mo3b/fqF4DkZJUx2y4wf/m3XyMVFvbkT77lLJruSlHM751VoFxXnwkX2w2e3Q51RmgsjWBOnb6rfebHZxSPbfVdDnZFaOEPf10qXBafQ3pFQhOX3XVgG+UX3Fc8pobchrfzZMnKql91WHj80gQImo4pQLpzsI77h0jNm5LrIxc0537QT2U67tcL9bbA4jZzjZ1f72VbQheIq1wZVXb30GJuPGvv9CgAA//9Q+TWn"
}
