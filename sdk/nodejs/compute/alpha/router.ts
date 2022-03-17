// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
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
    public static readonly __pulumiType = 'google-native:compute/alpha:Router';

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
    public readonly bgp!: pulumi.Output<outputs.compute.alpha.RouterBgpResponse>;
    /**
     * BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
     */
    public readonly bgpPeers!: pulumi.Output<outputs.compute.alpha.RouterBgpPeerResponse[]>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments). Not currently available publicly. 
     */
    public readonly encryptedInterconnectRouter!: pulumi.Output<boolean>;
    /**
     * Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
     */
    public readonly interfaces!: pulumi.Output<outputs.compute.alpha.RouterInterfaceResponse[]>;
    /**
     * Type of resource. Always compute#router for routers.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Keys used for MD5 authentication.
     */
    public readonly md5AuthenticationKeys!: pulumi.Output<outputs.compute.alpha.RouterMd5AuthenticationKeyResponse[]>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of NAT services created in this router.
     */
    public readonly nats!: pulumi.Output<outputs.compute.alpha.RouterNatResponse[]>;
    /**
     * URI of the network to which this router belongs.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;

    /**
     * Create a Router resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["bgp"] = args ? args.bgp : undefined;
            resourceInputs["bgpPeers"] = args ? args.bgpPeers : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encryptedInterconnectRouter"] = args ? args.encryptedInterconnectRouter : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["md5AuthenticationKeys"] = args ? args.md5AuthenticationKeys : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nats"] = args ? args.nats : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
        } else {
            resourceInputs["bgp"] = undefined /*out*/;
            resourceInputs["bgpPeers"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["encryptedInterconnectRouter"] = undefined /*out*/;
            resourceInputs["interfaces"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["md5AuthenticationKeys"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nats"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Router.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Router resource.
 */
export interface RouterArgs {
    /**
     * BGP information specific to this router.
     */
    bgp?: pulumi.Input<inputs.compute.alpha.RouterBgpArgs>;
    /**
     * BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
     */
    bgpPeers?: pulumi.Input<pulumi.Input<inputs.compute.alpha.RouterBgpPeerArgs>[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments). Not currently available publicly. 
     */
    encryptedInterconnectRouter?: pulumi.Input<boolean>;
    /**
     * Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.compute.alpha.RouterInterfaceArgs>[]>;
    /**
     * Keys used for MD5 authentication.
     */
    md5AuthenticationKeys?: pulumi.Input<pulumi.Input<inputs.compute.alpha.RouterMd5AuthenticationKeyArgs>[]>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of NAT services created in this router.
     */
    nats?: pulumi.Input<pulumi.Input<inputs.compute.alpha.RouterNatArgs>[]>;
    /**
     * URI of the network to which this router belongs.
     */
    network?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
}
