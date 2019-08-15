// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> & GetUserResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetUserResult> = pulumi.runtime.invoke("github:index/getUser:getUser", {
        "username": args.username,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    readonly username: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    readonly avatarUrl: string;
    readonly bio: string;
    readonly blog: string;
    readonly company: string;
    readonly createdAt: string;
    readonly email: string;
    readonly followers: number;
    readonly following: number;
    readonly gpgKeys: string[];
    readonly gravatarId: string;
    readonly location: string;
    readonly login: string;
    readonly name: string;
    readonly publicGists: number;
    readonly publicRepos: number;
    readonly siteAdmin: boolean;
    readonly sshKeys: string[];
    readonly updatedAt: string;
    readonly username: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}