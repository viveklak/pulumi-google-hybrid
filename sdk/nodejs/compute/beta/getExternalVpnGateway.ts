// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified externalVpnGateway. Get a list of available externalVpnGateways by making a list() request.
 */
export function getExternalVpnGateway(args: GetExternalVpnGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetExternalVpnGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/beta:getExternalVpnGateway", {
        "externalVpnGateway": args.externalVpnGateway,
        "project": args.project,
    }, opts);
}

export interface GetExternalVpnGatewayArgs {
    externalVpnGateway: string;
    project: string;
}

export interface GetExternalVpnGatewayResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * List of interfaces for this external VPN gateway.
     */
    readonly interfaces: outputs.compute.beta.ExternalVpnGatewayInterfaceResponse[];
    /**
     * Type of the resource. Always compute#externalVpnGateway for externalVpnGateways.
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this ExternalVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an ExternalVpnGateway.
     */
    readonly labelFingerprint: string;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Indicates the user-supplied redundancy type of this external VPN gateway.
     */
    readonly redundancyType: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
}
