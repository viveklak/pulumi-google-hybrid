// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Sets the access control policy on an Organization resource. Replaces any existing policy. The `resource` field should be the organization's resource name, e.g. "organizations/123".
 */
export class OrganizationIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationIamPolicy {
        return new OrganizationIamPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:cloudresourcemanager/v1beta1:OrganizationIamPolicy';

    /**
     * Returns true if the given object is an instance of OrganizationIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationIamPolicy.__pulumiType;
    }


    /**
     * Create a OrganizationIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationIamPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationsId'");
            }
            inputs["organizationsId"] = args ? args.organizationsId : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["updateMask"] = args ? args.updateMask : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationIamPolicy resource.
 */
export interface OrganizationIamPolicyArgs {
    readonly organizationsId: pulumi.Input<string>;
    /**
     * REQUIRED: The complete policy to be applied to the `resource`. The size of the policy is limited to a few 10s of KB. An empty policy is a valid policy but certain Cloud Platform services (such as Projects) might reject them.
     */
    readonly policy?: pulumi.Input<inputs.cloudresourcemanager.v1beta1.Policy>;
    /**
     * OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
     */
    readonly updateMask?: pulumi.Input<string>;
}
