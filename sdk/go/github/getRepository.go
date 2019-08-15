// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a GitHub repository.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/d/repository.html.markdown.
func LookupRepository(ctx *pulumi.Context, args *GetRepositoryArgs) (*GetRepositoryResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["fullName"] = args.FullName
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("github:index/getRepository:getRepository", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRepositoryResult{
		AllowMergeCommit: outputs["allowMergeCommit"],
		AllowRebaseMerge: outputs["allowRebaseMerge"],
		AllowSquashMerge: outputs["allowSquashMerge"],
		Archived: outputs["archived"],
		DefaultBranch: outputs["defaultBranch"],
		Description: outputs["description"],
		FullName: outputs["fullName"],
		GitCloneUrl: outputs["gitCloneUrl"],
		HasDownloads: outputs["hasDownloads"],
		HasIssues: outputs["hasIssues"],
		HasProjects: outputs["hasProjects"],
		HasWiki: outputs["hasWiki"],
		HomepageUrl: outputs["homepageUrl"],
		HtmlUrl: outputs["htmlUrl"],
		HttpCloneUrl: outputs["httpCloneUrl"],
		Name: outputs["name"],
		Private: outputs["private"],
		SshCloneUrl: outputs["sshCloneUrl"],
		SvnUrl: outputs["svnUrl"],
		Topics: outputs["topics"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRepository.
type GetRepositoryArgs struct {
	// Full name of the repository (in `org/name` format).
	FullName interface{}
	// The name of the repository.
	Name interface{}
}

// A collection of values returned by getRepository.
type GetRepositoryResult struct {
	// Whether the repository allows merge commits.
	AllowMergeCommit interface{}
	// Whether the repository allows rebase merges.
	AllowRebaseMerge interface{}
	// Whether the repository allows squash merges.
	AllowSquashMerge interface{}
	// Whether the repository is archived.
	Archived interface{}
	// The name of the default branch of the repository.
	DefaultBranch interface{}
	// A description of the repository.
	Description interface{}
	FullName interface{}
	// URL that can be provided to `git clone` to clone the
	// repository anonymously via the git protocol.
	GitCloneUrl interface{}
	// Whether the repository has Downloads feature enabled.
	HasDownloads interface{}
	// Whether the repository has GitHub Issues enabled.
	HasIssues interface{}
	// Whether the repository has the GitHub Projects enabled.
	HasProjects interface{}
	// Whether the repository has the GitHub Wiki enabled.
	HasWiki interface{}
	// URL of a page describing the project.
	HomepageUrl interface{}
	// URL to the repository on the web.
	HtmlUrl interface{}
	// URL that can be provided to `git clone` to clone the
	// repository via HTTPS.
	HttpCloneUrl interface{}
	Name interface{}
	// Whether the repository is private.
	Private interface{}
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshCloneUrl interface{}
	// URL that can be provided to `svn checkout` to check out
	// the repository via GitHub's Subversion protocol emulation.
	SvnUrl interface{}
	// The list of topics of the repository.
	Topics interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
