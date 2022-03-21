// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified zonal instance group. Get a list of available zonal instance groups by making a list() request. For managed instance groups, use the instanceGroupManagers or regionInstanceGroupManagers methods instead.
 */
export function getInstanceGroup(args: GetInstanceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/beta:getInstanceGroup", {
        "instanceGroup": args.instanceGroup,
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

export interface GetInstanceGroupArgs {
    instanceGroup: string;
    project?: string;
    zone: string;
}

export interface GetInstanceGroupResult {
    /**
     * The creation timestamp for this instance group in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * The fingerprint of the named ports. The system uses this fingerprint to detect conflicts when multiple users change the named ports concurrently.
     */
    readonly fingerprint: string;
    /**
     * The resource type, which is always compute#instanceGroup for instance groups.
     */
    readonly kind: string;
    /**
     * The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
     */
    readonly name: string;
    /**
     *  Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "app1", port: 8080}, {name: "app1", port: 8081}, {name: "app2", port: 8082}] Named ports apply to all instances in this instance group. 
     */
    readonly namedPorts: outputs.compute.beta.NamedPortResponse[];
    /**
     * The URL of the network to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
     */
    readonly network: string;
    /**
     * The URL of the region where the instance group is located (for regional resources).
     */
    readonly region: string;
    /**
     * The URL for this instance group. The server generates this URL.
     */
    readonly selfLink: string;
    /**
     * The total number of instances in the instance group.
     */
    readonly size: number;
    /**
     * The URL of the subnetwork to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
     */
    readonly subnetwork: string;
    /**
     * The URL of the zone where the instance group is located (for zonal resources).
     */
    readonly zone: string;
}

export function getInstanceGroupOutput(args: GetInstanceGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceGroupResult> {
    return pulumi.output(args).apply(a => getInstanceGroup(a, opts))
}

export interface GetInstanceGroupOutputArgs {
    instanceGroup: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
