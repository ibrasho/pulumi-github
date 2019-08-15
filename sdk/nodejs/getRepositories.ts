// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getRepositories(args: GetRepositoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoriesResult> & GetRepositoriesResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetRepositoriesResult> = pulumi.runtime.invoke("github:index/getRepositories:getRepositories", {
        "query": args.query,
        "sort": args.sort,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getRepositories.
 */
export interface GetRepositoriesArgs {
    readonly query: string;
    readonly sort?: string;
}

/**
 * A collection of values returned by getRepositories.
 */
export interface GetRepositoriesResult {
    readonly fullNames: string[];
    readonly names: string[];
    readonly query: string;
    readonly sort?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}