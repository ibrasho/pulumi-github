// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a GitHub user.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/d/user.html.markdown.
func LookupUser(ctx *pulumi.Context, args *GetUserArgs) (*GetUserResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["username"] = args.Username
	}
	outputs, err := ctx.Invoke("github:index/getUser:getUser", inputs)
	if err != nil {
		return nil, err
	}
	return &GetUserResult{
		AvatarUrl: outputs["avatarUrl"],
		Bio: outputs["bio"],
		Blog: outputs["blog"],
		Company: outputs["company"],
		CreatedAt: outputs["createdAt"],
		Email: outputs["email"],
		Followers: outputs["followers"],
		Following: outputs["following"],
		GpgKeys: outputs["gpgKeys"],
		GravatarId: outputs["gravatarId"],
		Location: outputs["location"],
		Login: outputs["login"],
		Name: outputs["name"],
		PublicGists: outputs["publicGists"],
		PublicRepos: outputs["publicRepos"],
		SiteAdmin: outputs["siteAdmin"],
		SshKeys: outputs["sshKeys"],
		UpdatedAt: outputs["updatedAt"],
		Username: outputs["username"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	// The username.
	Username interface{}
}

// A collection of values returned by getUser.
type GetUserResult struct {
	// the user's avatar URL.
	AvatarUrl interface{}
	// the user's bio.
	Bio interface{}
	// the user's blog location.
	Blog interface{}
	// the user's company name.
	Company interface{}
	// the creation date.
	CreatedAt interface{}
	// the user's email.
	Email interface{}
	// the number of followers.
	Followers interface{}
	// the number of following users.
	Following interface{}
	// list of user's GPG keys
	GpgKeys interface{}
	// the user's gravatar ID.
	GravatarId interface{}
	// the user's location.
	Location interface{}
	// the user's login.
	Login interface{}
	// the user's full name.
	Name interface{}
	// the number of public gists.
	PublicGists interface{}
	// the number of public repositories.
	PublicRepos interface{}
	// whether the user is a GitHub admin.
	SiteAdmin interface{}
	// list of user's SSH keys
	SshKeys interface{}
	// the update date.
	UpdatedAt interface{}
	Username interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}