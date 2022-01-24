// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Use this method to get details about a private connectivity configuration.
 */
export function getPrivateConnection(args: GetPrivateConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateConnectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:datastream/v1:getPrivateConnection", {
        "location": args.location,
        "privateConnectionId": args.privateConnectionId,
        "project": args.project,
    }, opts);
}

export interface GetPrivateConnectionArgs {
    location: string;
    privateConnectionId: string;
    project?: string;
}

export interface GetPrivateConnectionResult {
    /**
     * The create time of the resource.
     */
    readonly createTime: string;
    /**
     * Display name.
     */
    readonly displayName: string;
    /**
     * In case of error, the details of the error in a user-friendly format.
     */
    readonly error: outputs.datastream.v1.ErrorResponse;
    /**
     * Labels.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource's name.
     */
    readonly name: string;
    /**
     * The state of the Private Connection.
     */
    readonly state: string;
    /**
     * The update time of the resource.
     */
    readonly updateTime: string;
    /**
     * VPC Peering Config.
     */
    readonly vpcPeeringConfig: outputs.datastream.v1.VpcPeeringConfigResponse;
}

export function getPrivateConnectionOutput(args: GetPrivateConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrivateConnectionResult> {
    return pulumi.output(args).apply(a => getPrivateConnection(a, opts))
}

export interface GetPrivateConnectionOutputArgs {
    location: pulumi.Input<string>;
    privateConnectionId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
