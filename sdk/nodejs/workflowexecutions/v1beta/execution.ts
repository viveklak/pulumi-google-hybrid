// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new execution using the latest revision of the given workflow.
 * Auto-naming is currently not supported for this resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class Execution extends pulumi.CustomResource {
    /**
     * Get an existing Execution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Execution {
        return new Execution(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:workflowexecutions/v1beta:Execution';

    /**
     * Returns true if the given object is an instance of Execution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Execution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Execution.__pulumiType;
    }

    /**
     * Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
     */
    public readonly argument!: pulumi.Output<string>;
    /**
     * The call logging level associated to this execution.
     */
    public readonly callLogLevel!: pulumi.Output<string>;
    /**
     * Marks the end of execution, successful or not.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.workflowexecutions.v1beta.ErrorResponse>;
    /**
     * The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
     */
    public /*out*/ readonly result!: pulumi.Output<string>;
    /**
     * Marks the beginning of execution.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * Current state of the execution.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Revision of the workflow this execution is using.
     */
    public /*out*/ readonly workflowRevisionId!: pulumi.Output<string>;

    /**
     * Create a Execution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExecutionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.workflowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workflowId'");
            }
            resourceInputs["argument"] = args ? args.argument : undefined;
            resourceInputs["callLogLevel"] = args ? args.callLogLevel : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["workflowId"] = args ? args.workflowId : undefined;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["result"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["workflowRevisionId"] = undefined /*out*/;
        } else {
            resourceInputs["argument"] = undefined /*out*/;
            resourceInputs["callLogLevel"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["result"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["workflowRevisionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Execution.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Execution resource.
 */
export interface ExecutionArgs {
    /**
     * Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
     */
    argument?: pulumi.Input<string>;
    /**
     * The call logging level associated to this execution.
     */
    callLogLevel?: pulumi.Input<enums.workflowexecutions.v1beta.ExecutionCallLogLevel>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    workflowId: pulumi.Input<string>;
}
