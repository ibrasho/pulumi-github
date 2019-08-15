// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIPRanges(opts?: pulumi.InvokeOptions): Promise<GetIPRangesResult> & GetIPRangesResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetIPRangesResult> = pulumi.runtime.invoke("github:index/getIPRanges:getIPRanges", {
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of values returned by getIPRanges.
 */
export interface GetIPRangesResult {
    readonly gits: string[];
    readonly hooks: string[];
    readonly importers: string[];
    readonly pages: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
