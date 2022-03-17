// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Fetches the representation of an existing Response Policy.
 */
export function getResponsePolicy(args: GetResponsePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetResponsePolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dns/v1beta2:getResponsePolicy", {
        "clientOperationId": args.clientOperationId,
        "project": args.project,
        "responsePolicy": args.responsePolicy,
    }, opts);
}

export interface GetResponsePolicyArgs {
    clientOperationId?: string;
    project?: string;
    responsePolicy: string;
}

export interface GetResponsePolicyResult {
    /**
     * User-provided description for this Response Policy.
     */
    readonly description: string;
    /**
     * The list of Google Kubernetes Engine clusters to which this response policy is applied.
     */
    readonly gkeClusters: outputs.dns.v1beta2.ResponsePolicyGKEClusterResponse[];
    readonly kind: string;
    /**
     * List of network names specifying networks to which this policy is applied.
     */
    readonly networks: outputs.dns.v1beta2.ResponsePolicyNetworkResponse[];
    /**
     * User assigned name for this Response Policy.
     */
    readonly responsePolicyName: string;
}

export function getResponsePolicyOutput(args: GetResponsePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResponsePolicyResult> {
    return pulumi.output(args).apply(a => getResponsePolicy(a, opts))
}

export interface GetResponsePolicyOutputArgs {
    clientOperationId?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    responsePolicy: pulumi.Input<string>;
}
