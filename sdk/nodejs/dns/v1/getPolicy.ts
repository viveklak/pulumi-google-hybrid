// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Fetches the representation of an existing Policy.
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dns/v1:getPolicy", {
        "clientOperationId": args.clientOperationId,
        "policy": args.policy,
        "project": args.project,
    }, opts);
}

export interface GetPolicyArgs {
    clientOperationId?: string;
    policy: string;
    project?: string;
}

export interface GetPolicyResult {
    /**
     * Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
     */
    readonly alternativeNameServerConfig: outputs.dns.v1.PolicyAlternativeNameServerConfigResponse;
    /**
     * A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
     */
    readonly description: string;
    /**
     * Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
     */
    readonly enableInboundForwarding: boolean;
    /**
     * Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
     */
    readonly enableLogging: boolean;
    readonly kind: string;
    /**
     * User-assigned name for this policy.
     */
    readonly name: string;
    /**
     * List of network names specifying networks to which this policy is applied.
     */
    readonly networks: outputs.dns.v1.PolicyNetworkResponse[];
}

export function getPolicyOutput(args: GetPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyResult> {
    return pulumi.output(args).apply(a => getPolicy(a, opts))
}

export interface GetPolicyOutputArgs {
    clientOperationId?: pulumi.Input<string>;
    policy: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
