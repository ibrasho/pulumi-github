// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource allows you to create and manage webhooks for GitHub organization.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/r/organization_webhook.html.markdown.
type OrganizationWebhook struct {
	s *pulumi.ResourceState
}

// NewOrganizationWebhook registers a new resource with the given unique name, arguments, and options.
func NewOrganizationWebhook(ctx *pulumi.Context,
	name string, args *OrganizationWebhookArgs, opts ...pulumi.ResourceOpt) (*OrganizationWebhook, error) {
	if args == nil || args.Events == nil {
		return nil, errors.New("missing required argument 'Events'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["active"] = nil
		inputs["configuration"] = nil
		inputs["events"] = nil
	} else {
		inputs["active"] = args.Active
		inputs["configuration"] = args.Configuration
		inputs["events"] = args.Events
	}
	inputs["etag"] = nil
	inputs["url"] = nil
	s, err := ctx.RegisterResource("github:index/organizationWebhook:OrganizationWebhook", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrganizationWebhook{s: s}, nil
}

// GetOrganizationWebhook gets an existing OrganizationWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationWebhook(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OrganizationWebhookState, opts ...pulumi.ResourceOpt) (*OrganizationWebhook, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["active"] = state.Active
		inputs["configuration"] = state.Configuration
		inputs["etag"] = state.Etag
		inputs["events"] = state.Events
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("github:index/organizationWebhook:OrganizationWebhook", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrganizationWebhook{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OrganizationWebhook) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OrganizationWebhook) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Indicate of the webhook should receive events. Defaults to `true`.
func (r *OrganizationWebhook) Active() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["active"])
}

// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
func (r *OrganizationWebhook) Configuration() *pulumi.Output {
	return r.s.State["configuration"]
}

func (r *OrganizationWebhook) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
func (r *OrganizationWebhook) Events() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["events"])
}

// URL of the webhook
func (r *OrganizationWebhook) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering OrganizationWebhook resources.
type OrganizationWebhookState struct {
	// Indicate of the webhook should receive events. Defaults to `true`.
	Active interface{}
	// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
	Configuration interface{}
	Etag interface{}
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
	Events interface{}
	// URL of the webhook
	Url interface{}
}

// The set of arguments for constructing a OrganizationWebhook resource.
type OrganizationWebhookArgs struct {
	// Indicate of the webhook should receive events. Defaults to `true`.
	Active interface{}
	// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
	Configuration interface{}
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
	Events interface{}
}
