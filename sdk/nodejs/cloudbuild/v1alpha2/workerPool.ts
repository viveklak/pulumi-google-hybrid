// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a `WorkerPool` to run the builds, and returns the new worker pool.
 * Auto-naming is currently not supported for this resource.
 */
export class WorkerPool extends pulumi.CustomResource {
    /**
     * Get an existing WorkerPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkerPool {
        return new WorkerPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudbuild/v1alpha2:WorkerPool';

    /**
     * Returns true if the given object is an instance of WorkerPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkerPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkerPool.__pulumiType;
    }

    /**
     * Time at which the request to create the `WorkerPool` was received.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Time at which the request to delete the `WorkerPool` was received.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * The resource name of the `WorkerPool`. Format of the name is `projects/{project_id}/workerPools/{worker_pool_id}`, where the value of {worker_pool_id} is provided in the CreateWorkerPool request.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Network configuration for the `WorkerPool`.
     */
    public readonly networkConfig!: pulumi.Output<outputs.cloudbuild.v1alpha2.NetworkConfigResponse>;
    /**
     * Immutable. The region where the `WorkerPool` runs. Only "us-central1" is currently supported. Note that `region` cannot be changed once the `WorkerPool` is created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * WorkerPool state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Time at which the request to update the `WorkerPool` was received.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Worker configuration for the `WorkerPool`.
     */
    public readonly workerConfig!: pulumi.Output<outputs.cloudbuild.v1alpha2.WorkerConfigResponse>;

    /**
     * Create a WorkerPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkerPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.workerPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workerPoolId'");
            }
            resourceInputs["networkConfig"] = args ? args.networkConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["workerConfig"] = args ? args.workerConfig : undefined;
            resourceInputs["workerPoolId"] = args ? args.workerPoolId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkConfig"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["workerConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WorkerPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkerPool resource.
 */
export interface WorkerPoolArgs {
    /**
     * Network configuration for the `WorkerPool`.
     */
    networkConfig?: pulumi.Input<inputs.cloudbuild.v1alpha2.NetworkConfigArgs>;
    project?: pulumi.Input<string>;
    /**
     * Immutable. The region where the `WorkerPool` runs. Only "us-central1" is currently supported. Note that `region` cannot be changed once the `WorkerPool` is created.
     */
    region: pulumi.Input<string>;
    /**
     * Worker configuration for the `WorkerPool`.
     */
    workerConfig?: pulumi.Input<inputs.cloudbuild.v1alpha2.WorkerConfigArgs>;
    workerPoolId: pulumi.Input<string>;
}
