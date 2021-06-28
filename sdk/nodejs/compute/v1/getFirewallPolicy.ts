// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified firewall policy.
 */
export function getFirewallPolicy(args: GetFirewallPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/v1:getFirewallPolicy", {
        "firewallPolicy": args.firewallPolicy,
    }, opts);
}

export interface GetFirewallPolicyArgs {
    firewallPolicy: string;
}

export interface GetFirewallPolicyResult {
    /**
     * A list of associations that belong to this firewall policy.
     */
    readonly associations: outputs.compute.v1.FirewallPolicyAssociationResponse[];
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Depreacted, please use short name instead. User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly displayName: string;
    /**
     * Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make get() request to the firewall policy.
     */
    readonly fingerprint: string;
    /**
     * [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
     */
    readonly kind: string;
    /**
     * Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
     */
    readonly name: string;
    /**
     * The parent of the firewall policy.
     */
    readonly parent: string;
    /**
     * Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
     */
    readonly ruleTupleCount: number;
    /**
     * A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
     */
    readonly rules: outputs.compute.v1.FirewallPolicyRuleResponse[];
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly shortName: string;
}
