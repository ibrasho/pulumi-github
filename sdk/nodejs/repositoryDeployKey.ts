// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class RepositoryDeployKey extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryDeployKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryDeployKeyState, opts?: pulumi.CustomResourceOptions): RepositoryDeployKey {
        return new RepositoryDeployKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryDeployKey:RepositoryDeployKey';

    /**
     * Returns true if the given object is an instance of RepositoryDeployKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryDeployKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryDeployKey.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly key!: pulumi.Output<string>;
    public readonly readOnly!: pulumi.Output<boolean | undefined>;
    public readonly repository!: pulumi.Output<string>;
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a RepositoryDeployKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryDeployKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryDeployKeyArgs | RepositoryDeployKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RepositoryDeployKeyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["readOnly"] = state ? state.readOnly : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as RepositoryDeployKeyArgs | undefined;
            if (!args || args.key === undefined) {
                throw new Error("Missing required property 'key'");
            }
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["key"] = args ? args.key : undefined;
            inputs["readOnly"] = args ? args.readOnly : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RepositoryDeployKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryDeployKey resources.
 */
export interface RepositoryDeployKeyState {
    readonly etag?: pulumi.Input<string>;
    readonly key?: pulumi.Input<string>;
    readonly readOnly?: pulumi.Input<boolean>;
    readonly repository?: pulumi.Input<string>;
    readonly title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryDeployKey resource.
 */
export interface RepositoryDeployKeyArgs {
    readonly key: pulumi.Input<string>;
    readonly readOnly?: pulumi.Input<boolean>;
    readonly repository: pulumi.Input<string>;
    readonly title: pulumi.Input<string>;
}