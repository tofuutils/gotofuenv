/*
 *
 * Copyright 2024 gotofuenv authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package terraformretriever

import (
	_ "embed"
	"encoding/json"
	"slices"
	"testing"

	"github.com/tofuutils/gotofuenv/versionmanager/semantic"
)

//go:embed release.json
var releaseData []byte

var (
	releaseValue any
	releaseErr   error
)

//go:embed releases.json
var releasesData []byte

var (
	releasesValue any
	releasesErr   error
)

func init() {
	releaseErr = json.Unmarshal(releaseData, &releaseValue)
	releasesErr = json.Unmarshal(releasesData, &releasesValue)
}

func TestExtractAssetUrls(t *testing.T) {
	if releaseErr != nil {
		t.Fatal("Unexpected parsing error : ", releaseErr)
	}

	fileName, downloadUrl, downloadSumsUrl, downloadSumsSigUrl, err := extractAssetUrls("http://localhost:8080", "linux", "386", releaseValue)
	if err != nil {
		t.Fatal("Unexpected extract error : ", err)
	}

	if fileName != "terraform_1.7.0_linux_386.zip" {
		t.Error("Unexpected fileName, get :", fileName)
	}
	if downloadUrl != "https://releases.hashicorp.com/terraform/1.7.0/terraform_1.7.0_linux_386.zip" {
		t.Error("Unexpected downloadUrl, get :", downloadUrl)
	}
	if downloadSumsUrl != "http://localhost:8080/terraform_1.7.0_SHA256SUMS" {
		t.Error("Unexpected downloadSumsUrl, get :", downloadSumsUrl)
	}
	if downloadSumsSigUrl != "http://localhost:8080/terraform_1.7.0_SHA256SUMS.sig" {
		t.Error("Unexpected downloadSumsSigUrl, get :", downloadSumsSigUrl)
	}
}

func TestExtractReleases(t *testing.T) {
	if releasesErr != nil {
		t.Fatal("Unexpected parsing error : ", releasesErr)
	}

	releases, err := extractReleases(releasesValue)
	if err != nil {
		t.Fatal("Unexpected extract error : ", err)
	}

	slices.SortFunc(releases, semantic.CmpVersion)
	if !slices.Equal(releases, []string{"1.6.6", "1.7.0-rc1", "1.7.0-rc2", "1.7.0"}) {
		t.Error("Unmatching results, get :", releases)
	}
}
