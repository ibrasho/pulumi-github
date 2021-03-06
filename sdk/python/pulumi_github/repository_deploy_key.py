# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class RepositoryDeployKey(pulumi.CustomResource):
    etag: pulumi.Output[str]
    key: pulumi.Output[str]
    """
    A ssh key.
    """
    read_only: pulumi.Output[bool]
    """
    A boolean qualifying the key to be either read only or read/write.
    """
    repository: pulumi.Output[str]
    """
    Name of the GitHub repository.
    """
    title: pulumi.Output[str]
    """
    A title.
    """
    def __init__(__self__, resource_name, opts=None, key=None, read_only=None, repository=None, title=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a GitHub repository deploy key resource.
        
        A deploy key is an SSH key that is stored on your server and grants
        access to a single GitHub repository. This key is attached directly to the repository instead of to a personal user
        account.
        
        This resource allows you to add/remove repository deploy keys.
        
        Further documentation on GitHub repository deploy keys:
        - [About deploy keys](https://developer.github.com/guides/managing-deploy-keys/#deploy-keys)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: A ssh key.
        :param pulumi.Input[bool] read_only: A boolean qualifying the key to be either read only or read/write.
        :param pulumi.Input[str] repository: Name of the GitHub repository.
        :param pulumi.Input[str] title: A title.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/r/repository_deploy_key.html.markdown.
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

            if key is None:
                raise TypeError("Missing required property 'key'")
            __props__['key'] = key
            __props__['read_only'] = read_only
            if repository is None:
                raise TypeError("Missing required property 'repository'")
            __props__['repository'] = repository
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            __props__['etag'] = None
        super(RepositoryDeployKey, __self__).__init__(
            'github:index/repositoryDeployKey:RepositoryDeployKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, etag=None, key=None, read_only=None, repository=None, title=None):
        """
        Get an existing RepositoryDeployKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: A ssh key.
        :param pulumi.Input[bool] read_only: A boolean qualifying the key to be either read only or read/write.
        :param pulumi.Input[str] repository: Name of the GitHub repository.
        :param pulumi.Input[str] title: A title.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/r/repository_deploy_key.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["etag"] = etag
        __props__["key"] = key
        __props__["read_only"] = read_only
        __props__["repository"] = repository
        __props__["title"] = title
        return RepositoryDeployKey(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

