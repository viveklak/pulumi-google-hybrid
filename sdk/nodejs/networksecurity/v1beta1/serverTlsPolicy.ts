// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new ServerTlsPolicy in a given project and location.
 */
export class ServerTlsPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ServerTlsPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerTlsPolicy {
        return new ServerTlsPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networksecurity/v1beta1:ServerTlsPolicy';

    /**
     * Returns true if the given object is an instance of ServerTlsPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerTlsPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerTlsPolicy.__pulumiType;
    }

    /**
     *  Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
     */
    public readonly allowOpen!: pulumi.Output<boolean>;
    /**
     * The timestamp when the resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Free-text description of the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Set of label tags associated with the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     *  Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
     */
    public readonly mtlsPolicy!: pulumi.Output<outputs.networksecurity.v1beta1.MTLSPolicyResponse>;
    /**
     * Name of the ServerTlsPolicy resource. It matches the pattern `projects/*&#47;locations/{location}/serverTlsPolicies/{server_tls_policy}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     *  Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
     */
    public readonly serverCertificate!: pulumi.Output<outputs.networksecurity.v1beta1.GoogleCloudNetworksecurityV1beta1CertificateProviderResponse>;
    /**
     * The timestamp when the resource was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ServerTlsPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerTlsPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.serverTlsPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverTlsPolicyId'");
            }
            resourceInputs["allowOpen"] = args ? args.allowOpen : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["mtlsPolicy"] = args ? args.mtlsPolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serverCertificate"] = args ? args.serverCertificate : undefined;
            resourceInputs["serverTlsPolicyId"] = args ? args.serverTlsPolicyId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["allowOpen"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["mtlsPolicy"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serverCertificate"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerTlsPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerTlsPolicy resource.
 */
export interface ServerTlsPolicyArgs {
    /**
     *  Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
     */
    allowOpen?: pulumi.Input<boolean>;
    /**
     * Free-text description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     *  Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
     */
    mtlsPolicy?: pulumi.Input<inputs.networksecurity.v1beta1.MTLSPolicyArgs>;
    /**
     * Name of the ServerTlsPolicy resource. It matches the pattern `projects/*&#47;locations/{location}/serverTlsPolicies/{server_tls_policy}`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     *  Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
     */
    serverCertificate?: pulumi.Input<inputs.networksecurity.v1beta1.GoogleCloudNetworksecurityV1beta1CertificateProviderArgs>;
    serverTlsPolicyId: pulumi.Input<string>;
}
