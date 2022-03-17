// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
 */
export function getNetworkEndpointGroup(args: GetNetworkEndpointGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkEndpointGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/beta:getNetworkEndpointGroup", {
        "networkEndpointGroup": args.networkEndpointGroup,
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

export interface GetNetworkEndpointGroupArgs {
    networkEndpointGroup: string;
    project?: string;
    zone: string;
}

export interface GetNetworkEndpointGroupResult {
    /**
     * Metadata defined as annotations on the network endpoint group.
     */
    readonly annotations: {[key: string]: string};
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    readonly appEngine: outputs.compute.beta.NetworkEndpointGroupAppEngineResponse;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    readonly cloudFunction: outputs.compute.beta.NetworkEndpointGroupCloudFunctionResponse;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    readonly cloudRun: outputs.compute.beta.NetworkEndpointGroupCloudRunResponse;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * The default port used if the port number is not specified in the network endpoint.
     */
    readonly defaultPort: number;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
     */
    readonly kind: string;
    /**
     * This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
     *
     * @deprecated This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
     */
    readonly loadBalancer: outputs.compute.beta.NetworkEndpointGroupLbNetworkEndpointGroupResponse;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
     */
    readonly network: string;
    /**
     * Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
     */
    readonly networkEndpointType: string;
    /**
     * The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
     */
    readonly pscTargetService: string;
    /**
     * The URL of the region where the network endpoint group is located.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
     */
    readonly serverlessDeployment: outputs.compute.beta.NetworkEndpointGroupServerlessDeploymentResponse;
    /**
     * [Output only] Number of network endpoints in the network endpoint group.
     */
    readonly size: number;
    /**
     * Optional URL of the subnetwork to which all network endpoints in the NEG belong.
     */
    readonly subnetwork: string;
    /**
     * The URL of the zone where the network endpoint group is located.
     */
    readonly zone: string;
}

export function getNetworkEndpointGroupOutput(args: GetNetworkEndpointGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkEndpointGroupResult> {
    return pulumi.output(args).apply(a => getNetworkEndpointGroup(a, opts))
}

export interface GetNetworkEndpointGroupOutputArgs {
    networkEndpointGroup: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
