// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ns1

import (
	// Allow metadata embedding
	_ "embed"
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/ns1-terraform/terraform-provider-ns1/ns1"
	"github.com/pulumi/pulumi-ns1/provider/v3/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/x"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "ns1"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

//go:embed cmd/pulumi-resource-ns1/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(ns1.Provider())

	prov := tfbridge.ProviderInfo{
		P:                p,
		Name:             "ns1",
		Version:          version.Version,
		Description:      "A Pulumi package for creating and managing ns1 cloud resources.",
		Keywords:         []string{"pulumi", "ns1"},
		License:          "Apache-2.0",
		Homepage:         "https://pulumi.io",
		GitHubOrg:        "ns1-terraform",
		Repository:       "https://github.com/pulumi/pulumi-ns1",
		Config:           map[string]*tfbridge.SchemaInfo{},
		UpstreamRepoPath: "./upstream",
		MetadataInfo:     tfbridge.NewProviderMetadata(metadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"ns1_zone": {
				Tok: makeResource(mainMod, "Zone"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"zone": {
						CSharpName: "ZoneName",
					},
				},
			},
			"ns1_monitoringjob": {Tok: makeResource(mainMod, "MonitoringJob")},
			"ns1_notifylist":    {Tok: makeResource(mainMod, "NotifyList")},
			"ns1_datasource":    {Tok: makeResource(mainMod, "DataSource")},
			"ns1_datafeed":      {Tok: makeResource(mainMod, "DataFeed")},
			"ns1_apikey":        {Tok: makeResource(mainMod, "APIKey")},
			"ns1_pulsarjob": {
				Tok: makeResource(mainMod, "PulsarJob"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"ns1_subnet": {
				Tok: makeResource(mainMod, "Subnet"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"ns1_tsigkey": {Tok: makeResource(mainMod, "Tsigkey")},
			"ns1_dnsview": {
				Tok: makeResource(mainMod, "Dnsview"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"ns1_dnssec": {Tok: makeDataSource(mainMod, "getDNSSec")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	err := x.ComputeDefaults(&prov, x.TokensSingleModule("ns1_", mainMod, x.MakeStandardToken(mainPkg)))
	contract.AssertNoErrorf(err, "failed to compute default tokens")

	prov.SetAutonaming(255, "-")

	err = x.AutoAliasing(&prov, prov.GetMetadata())
	contract.AssertNoErrorf(err, "failed to apply automatic aliases")

	return prov
}
