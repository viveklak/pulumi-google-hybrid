// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Sets the access control policy on the specified resource. Replaces any existing policy.
 */
export class InstanceTemplateIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing InstanceTemplateIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InstanceTemplateIamPolicy {
        return new InstanceTemplateIamPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:compute/beta:InstanceTemplateIamPolicy';

    /**
     * Returns true if the given object is an instance of InstanceTemplateIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceTemplateIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceTemplateIamPolicy.__pulumiType;
    }


    /**
     * Create a InstanceTemplateIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceTemplateIamPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            inputs["bindings"] = args ? args.bindings : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["resource"] = args ? args.resource : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InstanceTemplateIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a InstanceTemplateIamPolicy resource.
 */
export interface InstanceTemplateIamPolicyArgs {
    /**
     * Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify bindings.
     */
    readonly bindings?: pulumi.Input<pulumi.Input<inputs.compute.beta.Binding>[]>;
    /**
     * Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify the etag.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * REQUIRED: The complete policy to be applied to the 'resource'. The size of the policy is limited to a few 10s of KB. An empty policy is in general a valid policy but certain services (like Projects) might reject them.
     */
    readonly policy?: pulumi.Input<inputs.compute.beta.Policy>;
    readonly project: pulumi.Input<string>;
    readonly resource: pulumi.Input<string>;
}
