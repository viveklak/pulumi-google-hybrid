// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Router resource in the specified project and region using the data included in the request.
 */
export class Router extends pulumi.CustomResource {
    /**
     * Get an existing Router resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Router {
        return new Router(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/v1:Router';

    /**
     * Returns true if the given object is an instance of Router.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Router {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Router.__pulumiType;
    }

    /**
     * BGP information specific to this router.
     */
    public readonly bgp!: pulumi.Output<outputs.compute.v1.RouterBgpResponse>;
    /**
     * BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
     */
    public readonly bgpPeers!: pulumi.Output<outputs.compute.v1.RouterBgpPeerResponse[]>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Field to indicate if a router is dedicated to use with encrypted Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).
     * Not currently available in all Interconnect locations.
     */
    public readonly encryptedInterconnectRouter!: pulumi.Output<boolean>;
    /**
     * Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
     */
    public readonly interfaces!: pulumi.Output<outputs.compute.v1.RouterInterfaceResponse[]>;
    /**
     * [Output Only] Type of resource. Always compute#router for routers.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of NAT services created in this router.
     */
    public readonly nats!: pulumi.Output<outputs.compute.v1.RouterNatResponse[]>;
    /**
     * URI of the network to which this router belongs.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * [Output Only] URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a Router resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["bgp"] = args ? args.bgp : undefined;
            inputs["bgpPeers"] = args ? args.bgpPeers : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encryptedInterconnectRouter"] = args ? args.encryptedInterconnectRouter : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["interfaces"] = args ? args.interfaces : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nats"] = args ? args.nats : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
        } else {
            inputs["bgp"] = undefined /*out*/;
            inputs["bgpPeers"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["encryptedInterconnectRouter"] = undefined /*out*/;
            inputs["interfaces"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nats"] = undefined /*out*/;
            inputs["network"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Router.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Router resource.
 */
export interface RouterArgs {
    /**
     * BGP information specific to this router.
     */
    readonly bgp?: pulumi.Input<inputs.compute.v1.RouterBgpArgs>;
    /**
     * BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
     */
    readonly bgpPeers?: pulumi.Input<pulumi.Input<inputs.compute.v1.RouterBgpPeerArgs>[]>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Field to indicate if a router is dedicated to use with encrypted Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).
     * Not currently available in all Interconnect locations.
     */
    readonly encryptedInterconnectRouter?: pulumi.Input<boolean>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
     */
    readonly interfaces?: pulumi.Input<pulumi.Input<inputs.compute.v1.RouterInterfaceArgs>[]>;
    /**
     * [Output Only] Type of resource. Always compute#router for routers.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of NAT services created in this router.
     */
    readonly nats?: pulumi.Input<pulumi.Input<inputs.compute.v1.RouterNatArgs>[]>;
    /**
     * URI of the network to which this router belongs.
     */
    readonly network?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}
