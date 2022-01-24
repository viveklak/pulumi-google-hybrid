// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a VpnTunnel resource in the specified project and region using the data included in the request.
 */
export class VpnTunnel extends pulumi.CustomResource {
    /**
     * Get an existing VpnTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VpnTunnel {
        return new VpnTunnel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:VpnTunnel';

    /**
     * Returns true if the given object is an instance of VpnTunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnTunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnTunnel.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Detailed status message for the VPN tunnel.
     */
    public /*out*/ readonly detailedStatus!: pulumi.Output<string>;
    /**
     * IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
     */
    public readonly ikeVersion!: pulumi.Output<number>;
    /**
     * Type of resource. Always compute#vpnTunnel for VPN tunnels.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this VpnTunnel, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a VpnTunnel.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
     */
    public readonly localTrafficSelector!: pulumi.Output<string[]>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
     */
    public readonly peerExternalGateway!: pulumi.Output<string>;
    /**
     * The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
     */
    public readonly peerExternalGatewayInterface!: pulumi.Output<number>;
    /**
     * URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
     */
    public readonly peerGcpGateway!: pulumi.Output<string>;
    /**
     * IP address of the peer VPN gateway. Only IPv4 is supported.
     */
    public readonly peerIp!: pulumi.Output<string>;
    /**
     * URL of the region where the VPN tunnel resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
     */
    public readonly remoteTrafficSelector!: pulumi.Output<string[]>;
    /**
     * URL of the router resource to be used for dynamic routing.
     */
    public readonly router!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
     */
    public readonly sharedSecret!: pulumi.Output<string>;
    /**
     * Hash of the shared secret.
     */
    public readonly sharedSecretHash!: pulumi.Output<string>;
    /**
     * The status of the VPN tunnel, which can be one of the following: - PROVISIONING: Resource is being allocated for the VPN tunnel. - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. - ESTABLISHED: Secure session is successfully established with the peer VPN. - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). - NEGOTIATION_FAILURE: Handshake failed. - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. - NO_INCOMING_PACKETS: No incoming packets from peer. - REJECTED: Tunnel configuration was rejected, can be result of being denied access. - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel. 
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
     */
    public readonly targetVpnGateway!: pulumi.Output<string>;
    /**
     * URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
     */
    public readonly vpnGateway!: pulumi.Output<string>;
    /**
     * The interface ID of the VPN gateway with which this VPN tunnel is associated.
     */
    public readonly vpnGatewayInterface!: pulumi.Output<number>;

    /**
     * Create a VpnTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnTunnelArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ikeVersion"] = args ? args.ikeVersion : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["localTrafficSelector"] = args ? args.localTrafficSelector : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["peerExternalGateway"] = args ? args.peerExternalGateway : undefined;
            resourceInputs["peerExternalGatewayInterface"] = args ? args.peerExternalGatewayInterface : undefined;
            resourceInputs["peerGcpGateway"] = args ? args.peerGcpGateway : undefined;
            resourceInputs["peerIp"] = args ? args.peerIp : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["remoteTrafficSelector"] = args ? args.remoteTrafficSelector : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["router"] = args ? args.router : undefined;
            resourceInputs["sharedSecret"] = args ? args.sharedSecret : undefined;
            resourceInputs["sharedSecretHash"] = args ? args.sharedSecretHash : undefined;
            resourceInputs["targetVpnGateway"] = args ? args.targetVpnGateway : undefined;
            resourceInputs["vpnGateway"] = args ? args.vpnGateway : undefined;
            resourceInputs["vpnGatewayInterface"] = args ? args.vpnGatewayInterface : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["detailedStatus"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["detailedStatus"] = undefined /*out*/;
            resourceInputs["ikeVersion"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["localTrafficSelector"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["peerExternalGateway"] = undefined /*out*/;
            resourceInputs["peerExternalGatewayInterface"] = undefined /*out*/;
            resourceInputs["peerGcpGateway"] = undefined /*out*/;
            resourceInputs["peerIp"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["remoteTrafficSelector"] = undefined /*out*/;
            resourceInputs["router"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["sharedSecret"] = undefined /*out*/;
            resourceInputs["sharedSecretHash"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["targetVpnGateway"] = undefined /*out*/;
            resourceInputs["vpnGateway"] = undefined /*out*/;
            resourceInputs["vpnGatewayInterface"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpnTunnel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VpnTunnel resource.
 */
export interface VpnTunnelArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
     */
    ikeVersion?: pulumi.Input<number>;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
     */
    localTrafficSelector?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
     */
    peerExternalGateway?: pulumi.Input<string>;
    /**
     * The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
     */
    peerExternalGatewayInterface?: pulumi.Input<number>;
    /**
     * URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
     */
    peerGcpGateway?: pulumi.Input<string>;
    /**
     * IP address of the peer VPN gateway. Only IPv4 is supported.
     */
    peerIp?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    /**
     * Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
     */
    remoteTrafficSelector?: pulumi.Input<pulumi.Input<string>[]>;
    requestId?: pulumi.Input<string>;
    /**
     * URL of the router resource to be used for dynamic routing.
     */
    router?: pulumi.Input<string>;
    /**
     * Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
     */
    sharedSecret?: pulumi.Input<string>;
    /**
     * Hash of the shared secret.
     */
    sharedSecretHash?: pulumi.Input<string>;
    /**
     * URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
     */
    targetVpnGateway?: pulumi.Input<string>;
    /**
     * URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
     */
    vpnGateway?: pulumi.Input<string>;
    /**
     * The interface ID of the VPN gateway with which this VPN tunnel is associated.
     */
    vpnGatewayInterface?: pulumi.Input<number>;
}
