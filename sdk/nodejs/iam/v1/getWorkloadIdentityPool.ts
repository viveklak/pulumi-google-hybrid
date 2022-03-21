// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets an individual WorkloadIdentityPool.
 */
export function getWorkloadIdentityPool(args: GetWorkloadIdentityPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkloadIdentityPoolResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:iam/v1:getWorkloadIdentityPool", {
        "location": args.location,
        "project": args.project,
        "workloadIdentityPoolId": args.workloadIdentityPoolId,
    }, opts);
}

export interface GetWorkloadIdentityPoolArgs {
    location: string;
    project?: string;
    workloadIdentityPoolId: string;
}

export interface GetWorkloadIdentityPoolResult {
    /**
     * A description of the pool. Cannot exceed 256 characters.
     */
    readonly description: string;
    /**
     * Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
     */
    readonly disabled: boolean;
    /**
     * A display name for the pool. Cannot exceed 32 characters.
     */
    readonly displayName: string;
    /**
     * The resource name of the pool.
     */
    readonly name: string;
    /**
     * The state of the pool.
     */
    readonly state: string;
}

export function getWorkloadIdentityPoolOutput(args: GetWorkloadIdentityPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkloadIdentityPoolResult> {
    return pulumi.output(args).apply(a => getWorkloadIdentityPool(a, opts))
}

export interface GetWorkloadIdentityPoolOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    workloadIdentityPoolId: pulumi.Input<string>;
}
