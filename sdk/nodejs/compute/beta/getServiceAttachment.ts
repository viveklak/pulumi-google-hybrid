// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified ServiceAttachment resource in the given scope.
 */
export function getServiceAttachment(args: GetServiceAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceAttachmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/beta:getServiceAttachment", {
        "project": args.project,
        "region": args.region,
        "serviceAttachment": args.serviceAttachment,
    }, opts);
}

export interface GetServiceAttachmentArgs {
    project: string;
    region: string;
    serviceAttachment: string;
}

export interface GetServiceAttachmentResult {
    /**
     * An array of connections for all the consumers connected to this service attachment.
     */
    readonly connectedEndpoints: outputs.compute.beta.ServiceAttachmentConnectedEndpointResponse[];
    /**
     * The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
     */
    readonly connectionPreference: string;
    /**
     * An array of forwarding rules for all the consumers connected to this service attachment.
     */
    readonly consumerForwardingRules: outputs.compute.beta.ServiceAttachmentConsumerForwardingRuleResponse[];
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
     */
    readonly enableProxyProtocol: boolean;
    /**
     * Type of the resource. Always compute#serviceAttachment for service attachments.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
     */
    readonly natSubnets: string[];
    /**
     * The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
     */
    readonly producerForwardingRule: string;
    /**
     * An 128-bit global unique ID of the PSC service attachment.
     */
    readonly pscServiceAttachmentId: outputs.compute.beta.Uint128Response;
    /**
     * URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * The URL of a service serving the endpoint identified by this service attachment.
     */
    readonly targetService: string;
}
