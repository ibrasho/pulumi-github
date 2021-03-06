// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class RepositoryWebhook extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryWebhookState, opts?: pulumi.CustomResourceOptions): RepositoryWebhook {
        return new RepositoryWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryWebhook:RepositoryWebhook';

    /**
     * Returns true if the given object is an instance of RepositoryWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryWebhook.__pulumiType;
    }

    public readonly active!: pulumi.Output<boolean | undefined>;
    public readonly configuration!: pulumi.Output<{ contentType?: string, insecureSsl?: string, secret?: string, url: string } | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly events!: pulumi.Output<string[]>;
    public readonly repository!: pulumi.Output<string>;
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a RepositoryWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryWebhookArgs | RepositoryWebhookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RepositoryWebhookState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["events"] = state ? state.events : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as RepositoryWebhookArgs | undefined;
            if (!args || args.events === undefined) {
                throw new Error("Missing required property 'events'");
            }
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["events"] = args ? args.events : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RepositoryWebhook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryWebhook resources.
 */
export interface RepositoryWebhookState {
    readonly active?: pulumi.Input<boolean>;
    readonly configuration?: pulumi.Input<{ contentType?: pulumi.Input<string>, insecureSsl?: pulumi.Input<string>, secret?: pulumi.Input<string>, url: pulumi.Input<string> }>;
    readonly etag?: pulumi.Input<string>;
    readonly events?: pulumi.Input<pulumi.Input<string>[]>;
    readonly repository?: pulumi.Input<string>;
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryWebhook resource.
 */
export interface RepositoryWebhookArgs {
    readonly active?: pulumi.Input<boolean>;
    readonly configuration?: pulumi.Input<{ contentType?: pulumi.Input<string>, insecureSsl?: pulumi.Input<string>, secret?: pulumi.Input<string>, url: pulumi.Input<string> }>;
    readonly events: pulumi.Input<pulumi.Input<string>[]>;
    readonly repository: pulumi.Input<string>;
}
