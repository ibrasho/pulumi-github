// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/r/issue_label.html.markdown.
type IssueLabel struct {
	s *pulumi.ResourceState
}

// NewIssueLabel registers a new resource with the given unique name, arguments, and options.
func NewIssueLabel(ctx *pulumi.Context,
	name string, args *IssueLabelArgs, opts ...pulumi.ResourceOpt) (*IssueLabel, error) {
	if args == nil || args.Color == nil {
		return nil, errors.New("missing required argument 'Color'")
	}
	if args == nil || args.Repository == nil {
		return nil, errors.New("missing required argument 'Repository'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["color"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["repository"] = nil
	} else {
		inputs["color"] = args.Color
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["repository"] = args.Repository
	}
	inputs["etag"] = nil
	inputs["url"] = nil
	s, err := ctx.RegisterResource("github:index/issueLabel:IssueLabel", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IssueLabel{s: s}, nil
}

// GetIssueLabel gets an existing IssueLabel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIssueLabel(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IssueLabelState, opts ...pulumi.ResourceOpt) (*IssueLabel, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["color"] = state.Color
		inputs["description"] = state.Description
		inputs["etag"] = state.Etag
		inputs["name"] = state.Name
		inputs["repository"] = state.Repository
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("github:index/issueLabel:IssueLabel", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IssueLabel{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IssueLabel) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IssueLabel) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A 6 character hex code, **without the leading #**, identifying the color of the label.
func (r *IssueLabel) Color() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["color"])
}

// A short description of the label.
func (r *IssueLabel) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *IssueLabel) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The name of the label.
func (r *IssueLabel) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The GitHub repository
func (r *IssueLabel) Repository() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["repository"])
}

// The URL to the issue label
func (r *IssueLabel) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering IssueLabel resources.
type IssueLabelState struct {
	// A 6 character hex code, **without the leading #**, identifying the color of the label.
	Color interface{}
	// A short description of the label.
	Description interface{}
	Etag interface{}
	// The name of the label.
	Name interface{}
	// The GitHub repository
	Repository interface{}
	// The URL to the issue label
	Url interface{}
}

// The set of arguments for constructing a IssueLabel resource.
type IssueLabelArgs struct {
	// A 6 character hex code, **without the leading #**, identifying the color of the label.
	Color interface{}
	// A short description of the label.
	Description interface{}
	// The name of the label.
	Name interface{}
	// The GitHub repository
	Repository interface{}
}
