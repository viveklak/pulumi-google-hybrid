// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Create an `AccessPolicy`. Fails if this organization already has a `AccessPolicy`. The longrunning Operation will have a successful status once the `AccessPolicy` has propagated to long-lasting storage. Syntactic and basic semantic errors will be returned in `metadata` as a BadRequest proto.
 */
export class AccessPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccessPolicy {
        return new AccessPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:accesscontextmanager/v1:AccessPolicy';

    /**
     * Returns true if the given object is an instance of AccessPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPolicy.__pulumiType;
    }


    /**
     * Create a AccessPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accessPoliciesId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessPoliciesId'");
            }
            inputs["accessPoliciesId"] = args ? args.accessPoliciesId : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["title"] = args ? args.title : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AccessPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccessPolicy resource.
 */
export interface AccessPolicyArgs {
    readonly accessPoliciesId: pulumi.Input<string>;
    /**
     * Output only. An opaque identifier for the current version of the `AccessPolicy`. This will always be a strongly validated etag, meaning that two Access Polices will be identical if and only if their etags are identical. Clients should not expect this to be in any specific format.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Output only. Resource name of the `AccessPolicy`. Format: `accessPolicies/{policy_id}`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Required. The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * Required. Human readable title. Does not affect behavior.
     */
    readonly title?: pulumi.Input<string>;
}
