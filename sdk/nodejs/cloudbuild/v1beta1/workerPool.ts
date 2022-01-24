// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a `WorkerPool` to run the builds, and returns the new worker pool. NOTE: As of now, this method returns an `Operation` that is always complete.
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
    public static readonly __pulumiType = 'google-native:cloudbuild/v1beta1:WorkerPool';

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
     * User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * Time at which the request to create the `WorkerPool` was received.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Time at which the request to delete the `WorkerPool` was received.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The resource name of the `WorkerPool`, with format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. The value of `{worker_pool}` is provided by `worker_pool_id` in `CreateWorkerPool` request and the value of `{location}` is determined by the endpoint accessed.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Network configuration for the `WorkerPool`.
     */
    public readonly networkConfig!: pulumi.Output<outputs.cloudbuild.v1beta1.NetworkConfigResponse>;
    /**
     * `WorkerPool` state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * A unique identifier for the `WorkerPool`.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Time at which the request to update the `WorkerPool` was received.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Worker configuration for the `WorkerPool`.
     */
    public readonly workerConfig!: pulumi.Output<outputs.cloudbuild.v1beta1.WorkerConfigResponse>;

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
            if ((!args || args.workerPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workerPoolId'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["networkConfig"] = args ? args.networkConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["workerConfig"] = args ? args.workerConfig : undefined;
            resourceInputs["workerPoolId"] = args ? args.workerPoolId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["annotations"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkConfig"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["workerConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkerPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkerPool resource.
 */
export interface WorkerPoolArgs {
    /**
     * User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
     */
    displayName?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Network configuration for the `WorkerPool`.
     */
    networkConfig?: pulumi.Input<inputs.cloudbuild.v1beta1.NetworkConfigArgs>;
    project?: pulumi.Input<string>;
    /**
     * Worker configuration for the `WorkerPool`.
     */
    workerConfig?: pulumi.Input<inputs.cloudbuild.v1beta1.WorkerConfigArgs>;
    workerPoolId: pulumi.Input<string>;
}
