# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Membership(pulumi.CustomResource):
    etag: pulumi.Output[str]
    role: pulumi.Output[str]
    """
    The role of the user within the organization.
    Must be one of `member` or `admin`. Defaults to `member`.
    """
    username: pulumi.Output[str]
    """
    The user to add to the organization.
    """
    def __init__(__self__, resource_name, opts=None, role=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a GitHub membership resource.
        
        This resource allows you to add/remove users from your organization. When applied,
        an invitation will be sent to the user to become part of the organization. When
        destroyed, either the invitation will be cancelled or the user will be removed.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] role: The role of the user within the organization.
               Must be one of `member` or `admin`. Defaults to `member`.
        :param pulumi.Input[str] username: The user to add to the organization.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/r/membership.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['role'] = role
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
            __props__['etag'] = None
        super(Membership, __self__).__init__(
            'github:index/membership:Membership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, etag=None, role=None, username=None):
        """
        Get an existing Membership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] role: The role of the user within the organization.
               Must be one of `member` or `admin`. Defaults to `member`.
        :param pulumi.Input[str] username: The user to add to the organization.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/r/membership.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["etag"] = etag
        __props__["role"] = role
        __props__["username"] = username
        return Membership(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

