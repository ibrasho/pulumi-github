# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

__config__ = pulumi.Config('github')

base_url = __config__.get('baseUrl') or utilities.get_env('GITHUB_BASE_URL')
"""
The GitHub Base API URL
"""

individual = __config__.get('individual')
"""
Whether to run outside an organization.
"""

insecure = __config__.get('insecure')
"""
Whether server should be accessed without verifying the TLS certificate.
"""

organization = __config__.get('organization') or utilities.get_env('GITHUB_ORGANIZATION')
"""
The GitHub organization name to manage. If `individual` is false, organization is required.
"""

token = __config__.get('token') or utilities.get_env('GITHUB_TOKEN')
"""
The OAuth token used to connect to GitHub.
"""
