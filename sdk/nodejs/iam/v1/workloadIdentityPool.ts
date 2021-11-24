// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new WorkloadIdentityPool. You cannot reuse the name of a deleted pool until 30 days after deletion.
 * Auto-naming is currently not supported for this resource.
 */
export class WorkloadIdentityPool extends pulumi.CustomResource {
    /**
     * Get an existing WorkloadIdentityPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkloadIdentityPool {
        return new WorkloadIdentityPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:iam/v1:WorkloadIdentityPool';

    /**
     * Returns true if the given object is an instance of WorkloadIdentityPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkloadIdentityPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkloadIdentityPool.__pulumiType;
    }

    /**
     * A description of the pool. Cannot exceed 256 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * A display name for the pool. Cannot exceed 32 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The resource name of the pool.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The state of the pool.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a WorkloadIdentityPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkloadIdentityPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.workloadIdentityPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workloadIdentityPoolId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["workloadIdentityPoolId"] = args ? args.workloadIdentityPoolId : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["disabled"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WorkloadIdentityPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkloadIdentityPool resource.
 */
export interface WorkloadIdentityPoolArgs {
    /**
     * A description of the pool. Cannot exceed 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * A display name for the pool. Cannot exceed 32 characters.
     */
    displayName?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    workloadIdentityPoolId: pulumi.Input<string>;
}
