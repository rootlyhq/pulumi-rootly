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

package rootly

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/rootlyhq/pulumi-rootly/provider/pkg/version"
	rootly "github.com/rootlyhq/terraform-provider-rootly/provider"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "rootly"
	// modules:
	mainMod = "index" // the rootly module
)

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(rootly.New("pulumi")())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "rootly",
		DisplayName:       "Rootly",
		Publisher:         "Rootly",
		LogoURL:           "https://assets.rootly.com/assets/logo/rootly-7d4aa42752841c6da862630427150431a13aadee3d9e528b92bd0fded5dbca1e.svg",
		PluginDownloadURL: "https://github.com/rootlyhq/pulumi-rootly/releases/v${VERSION}",
		Description:       "A Pulumi package for creating and managing Rootly cloud resources.",
		Keywords:          []string{"pulumi", "rootly", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://rootly.com",
		Repository:        "https://github.com/rootlyhq/pulumi-rootly",
		GitHubOrg:         "rootlyhq",
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"rootly_cause": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Cause"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"cause": {
						CSharpName: "rootly_cause",
					},
				},
			},
			"rootly_functionality": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Functionality"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"functionality": {
						CSharpName: "rootly_functionality",
					},
				},
			},
			"rootly_incident_role": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IncidentRole"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"incident_role": {
						CSharpName: "rootly_incident_role",
					},
				},
			},
			"rootly_incident_type": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IncidentType"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"incident_type": {
						CSharpName: "rootly_incident_type",
					},
				},
			},
			"rootly_service": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Service"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"service": {
						CSharpName: "rootly_service",
					},
				},
			},
			"rootly_severity": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Severity"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"severity": {
						CSharpName: "rootly_severity",
					},
				},
			},
			"rootly_team": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Team"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"team": {
						CSharpName: "rootly_team",
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
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
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
