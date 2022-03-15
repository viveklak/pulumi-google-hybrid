// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified network firewall policy.
 */
export function getNetworkFirewallPolicy(args: GetNetworkFirewallPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkFirewallPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/alpha:getNetworkFirewallPolicy", {
        "firewallPolicy": args.firewallPolicy,
        "project": args.project,
    }, opts);
}

export interface GetNetworkFirewallPolicyArgs {
    firewallPolicy: string;
    project?: string;
}

export interface GetNetworkFirewallPolicyResult {
    /**
     * A list of associations that belong to this firewall policy.
     */
    readonly associations: outputs.compute.alpha.FirewallPolicyAssociationResponse[];
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     *
     * @deprecated Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly displayName: string;
    /**
     * Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the firewall policy.
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
     * URL of the region where the regional firewall policy resides. This field is not applicable to global firewall policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
     */
    readonly ruleTupleCount: number;
    /**
     * A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
     */
    readonly rules: outputs.compute.alpha.FirewallPolicyRuleResponse[];
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly shortName: string;
    /**
     * The scope of networks allowed to be associated with the firewall policy. This field can be either GLOBAL_VPC_NETWORK or REGIONAL_VPC_NETWORK. A firewall policy with the VPC scope set to GLOBAL_VPC_NETWORK is allowed to be attached only to global networks. When the VPC scope is set to REGIONAL_VPC_NETWORK the firewall policy is allowed to be attached only to regional networks in the same scope as the firewall policy. Note: if not specified then GLOBAL_VPC_NETWORK will be used.
     */
    readonly vpcNetworkScope: string;
}

export function getNetworkFirewallPolicyOutput(args: GetNetworkFirewallPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkFirewallPolicyResult> {
    return pulumi.output(args).apply(a => getNetworkFirewallPolicy(a, opts))
}

export interface GetNetworkFirewallPolicyOutputArgs {
    firewallPolicy: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
