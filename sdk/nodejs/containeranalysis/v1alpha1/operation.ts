// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new `Operation`.
 */
export class Operation extends pulumi.CustomResource {
    /**
     * Get an existing Operation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Operation {
        return new Operation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:containeranalysis/v1alpha1:Operation';

    /**
     * Returns true if the given object is an instance of Operation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Operation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Operation.__pulumiType;
    }


    /**
     * Create a Operation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OperationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["operation"] = args ? args.operation : undefined;
            inputs["operationId"] = args ? args.operationId : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Operation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Operation resource.
 */
export interface OperationArgs {
    /**
     * The operation to create.
     */
    readonly operation?: pulumi.Input<inputs.containeranalysis.v1alpha1.Operation>;
    /**
     * The ID to use for this operation.
     */
    readonly operationId?: pulumi.Input<string>;
    readonly projectsId: pulumi.Input<string>;
}
