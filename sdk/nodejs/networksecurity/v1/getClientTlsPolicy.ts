// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single ClientTlsPolicy.
 */
export function getClientTlsPolicy(args: GetClientTlsPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetClientTlsPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:networksecurity/v1:getClientTlsPolicy", {
        "clientTlsPolicyId": args.clientTlsPolicyId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetClientTlsPolicyArgs {
    clientTlsPolicyId: string;
    location: string;
    project?: string;
}

export interface GetClientTlsPolicyResult {
    /**
     * Optional. Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
     */
    readonly clientCertificate: outputs.networksecurity.v1.GoogleCloudNetworksecurityV1CertificateProviderResponse;
    /**
     * The timestamp when the resource was created.
     */
    readonly createTime: string;
    /**
     * Optional. Free-text description of the resource.
     */
    readonly description: string;
    /**
     * Optional. Set of label tags associated with the resource.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the ClientTlsPolicy resource. It matches the pattern `projects/*&#47;locations/{location}/clientTlsPolicies/{client_tls_policy}`
     */
    readonly name: string;
    /**
     * Optional. Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
     */
    readonly serverValidationCa: outputs.networksecurity.v1.ValidationCAResponse[];
    /**
     * Optional. Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
     */
    readonly sni: string;
    /**
     * The timestamp when the resource was updated.
     */
    readonly updateTime: string;
}

export function getClientTlsPolicyOutput(args: GetClientTlsPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClientTlsPolicyResult> {
    return pulumi.output(args).apply(a => getClientTlsPolicy(a, opts))
}

export interface GetClientTlsPolicyOutputArgs {
    clientTlsPolicyId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
