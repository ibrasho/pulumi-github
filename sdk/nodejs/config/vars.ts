// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("github");

/**
 * The GitHub Base API URL
 */
export let baseUrl: string | undefined = __config.get("baseUrl") || utilities.getEnv("GITHUB_BASE_URL");
/**
 * Whether to run outside an organization.
 */
export let individual: boolean | undefined = __config.getObject<boolean>("individual");
/**
 * Whether server should be accessed without verifying the TLS certificate.
 */
export let insecure: boolean | undefined = __config.getObject<boolean>("insecure");
/**
 * The GitHub organization name to manage. If `individual` is false, organization is required.
 */
export let organization: string | undefined = __config.get("organization") || utilities.getEnv("GITHUB_ORGANIZATION");
/**
 * The OAuth token used to connect to GitHub.
 */
export let token: string | undefined = __config.get("token") || utilities.getEnv("GITHUB_TOKEN");
