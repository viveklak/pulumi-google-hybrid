// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details about the specified hub.
 */
export function getHub(args: GetHubArgs, opts?: pulumi.InvokeOptions): Promise<GetHubResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:networkconnectivity/v1:getHub", {
        "hubId": args.hubId,
        "project": args.project,
    }, opts);
}

export interface GetHubArgs {
    hubId: string;
    project?: string;
}

export interface GetHubResult {
    /**
     * The time the hub was created.
     */
    readonly createTime: string;
    /**
     * An optional description of the hub.
     */
    readonly description: string;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    readonly labels: {[key: string]: string};
    /**
     * Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`
     */
    readonly name: string;
    /**
     * The VPC networks associated with this hub's spokes. This field is read-only. Network Connectivity Center automatically populates it based on the set of spokes attached to the hub.
     */
    readonly routingVpcs: outputs.networkconnectivity.v1.RoutingVPCResponse[];
    /**
     * The current lifecycle state of this hub.
     */
    readonly state: string;
    /**
     * The Google-generated UUID for the hub. This value is unique across all hub resources. If a hub is deleted and another with the same name is created, the new hub is assigned a different unique_id.
     */
    readonly uniqueId: string;
    /**
     * The time the hub was last updated.
     */
    readonly updateTime: string;
}

export function getHubOutput(args: GetHubOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHubResult> {
    return pulumi.output(args).apply(a => getHub(a, opts))
}

export interface GetHubOutputArgs {
    hubId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
