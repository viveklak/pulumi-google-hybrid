// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new ClientTlsPolicy in a given project and location.
 */
export class ClientTlsPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ClientTlsPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ClientTlsPolicy {
        return new ClientTlsPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networksecurity/v1:ClientTlsPolicy';

    /**
     * Returns true if the given object is an instance of ClientTlsPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientTlsPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientTlsPolicy.__pulumiType;
    }

    /**
     * Optional. Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
     */
    public readonly clientCertificate!: pulumi.Output<outputs.networksecurity.v1.GoogleCloudNetworksecurityV1CertificateProviderResponse>;
    /**
     * The timestamp when the resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Free-text description of the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Set of label tags associated with the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of the ClientTlsPolicy resource. It matches the pattern `projects/*&#47;locations/{location}/clientTlsPolicies/{client_tls_policy}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
     */
    public readonly serverValidationCa!: pulumi.Output<outputs.networksecurity.v1.ValidationCAResponse[]>;
    /**
     * Optional. Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
     */
    public readonly sni!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ClientTlsPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientTlsPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clientTlsPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientTlsPolicyId'");
            }
            resourceInputs["clientCertificate"] = args ? args.clientCertificate : undefined;
            resourceInputs["clientTlsPolicyId"] = args ? args.clientTlsPolicyId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serverValidationCa"] = args ? args.serverValidationCa : undefined;
            resourceInputs["sni"] = args ? args.sni : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["clientCertificate"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serverValidationCa"] = undefined /*out*/;
            resourceInputs["sni"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClientTlsPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ClientTlsPolicy resource.
 */
export interface ClientTlsPolicyArgs {
    /**
     * Optional. Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
     */
    clientCertificate?: pulumi.Input<inputs.networksecurity.v1.GoogleCloudNetworksecurityV1CertificateProviderArgs>;
    /**
     * Required. Short name of the ClientTlsPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "client_mtls_policy".
     */
    clientTlsPolicyId: pulumi.Input<string>;
    /**
     * Optional. Free-text description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Set of label tags associated with the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Name of the ClientTlsPolicy resource. It matches the pattern `projects/*&#47;locations/{location}/clientTlsPolicies/{client_tls_policy}`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
     */
    serverValidationCa?: pulumi.Input<pulumi.Input<inputs.networksecurity.v1.ValidationCAArgs>[]>;
    /**
     * Optional. Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
     */
    sni?: pulumi.Input<string>;
}
