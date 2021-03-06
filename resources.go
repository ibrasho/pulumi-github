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

package github

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-github/github"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "github"
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

// // boolRef returns a reference to the bool argument.
// func boolRef(b bool) *bool {
// 	return &b
// }

// // stringValue gets a string value from a property map if present, else ""
// func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
// 	val, ok := vars[prop]
// 	if ok && val.IsString() {
// 		return val.StringValue()
// 	}
// 	return ""
// }

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := github.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "github",
		Description: "A Pulumi package for creating and managing github cloud resources.",
		Keywords:    []string{"pulumi", "github"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/ibrasho/pulumi-github",
		Config: map[string]*tfbridge.SchemaInfo{
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_TOKEN"},
				},
			},
			"organization": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_ORGANIZATION"},
				},
			},
			"base_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_BASE_URL"},
				},
			},
			"insecure":   {},
			"individual": {},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"github_branch_protection":        {Tok: makeResource(mainMod, "BranchProtection")},
			"github_issue_label":              {Tok: makeResource(mainMod, "IssueLabel")},
			"github_membership":               {Tok: makeResource(mainMod, "Membership")},
			"github_organization_block":       {Tok: makeResource(mainMod, "OrganizationBlock")},
			"github_organization_project":     {Tok: makeResource(mainMod, "OrganizationProject")},
			"github_organization_webhook":     {Tok: makeResource(mainMod, "OrganizationWebhook")},
			"github_project_column":           {Tok: makeResource(mainMod, "ProjectColumn")},
			"github_repository":               {Tok: makeResource(mainMod, "Repository")},
			"github_repository_collaborator":  {Tok: makeResource(mainMod, "RepositoryCollaborator")},
			"github_repository_deploy_key":    {Tok: makeResource(mainMod, "RepositoryDeployKey")},
			"github_repository_project":       {Tok: makeResource(mainMod, "RepositoryProject")},
			"github_repository_webhook":       {Tok: makeResource(mainMod, "RepositoryWebhook")},
			"github_team_membership":          {Tok: makeResource(mainMod, "TeamMembership")},
			"github_team_repository":          {Tok: makeResource(mainMod, "TeamRepository")},
			"github_user_gpg_key":             {Tok: makeResource(mainMod, "UserGPGKey")},
			"github_user_invitation_accepter": {Tok: makeResource(mainMod, "UserInvitationAccepter")},
			"github_user_ssh_key":             {Tok: makeResource(mainMod, "UserSSHKey")},
			"github_team":                     {Tok: makeResource(mainMod, "Team")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"github_collaborators": {Tok: makeDataSource(mainMod, "getCollaborators")},
			"github_ip_ranges":     {Tok: makeDataSource(mainMod, "getIPRanges")},
			"github_repository":    {Tok: makeDataSource(mainMod, "getRepository")},
			"github_repositories":  {Tok: makeDataSource(mainMod, "getRepositories")},
			"github_user":          {Tok: makeDataSource(mainMod, "getUser")},
			"github_team":          {Tok: makeDataSource(mainMod, "getTeam")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "latest",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=0.17.28,<1.0.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
