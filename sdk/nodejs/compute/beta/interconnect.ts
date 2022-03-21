// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Interconnect in the specified project using the data included in the request.
 */
export class Interconnect extends pulumi.CustomResource {
    /**
     * Get an existing Interconnect resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Interconnect {
        return new Interconnect(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:Interconnect';

    /**
     * Returns true if the given object is an instance of Interconnect.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Interconnect {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Interconnect.__pulumiType;
    }

    /**
     * Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
     */
    public readonly adminEnabled!: pulumi.Output<boolean>;
    /**
     * A list of CircuitInfo objects, that describe the individual circuits in this LAG.
     */
    public /*out*/ readonly circuitInfos!: pulumi.Output<outputs.compute.beta.InterconnectCircuitInfoResponse[]>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
     */
    public readonly customerName!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * A list of outages expected for this Interconnect.
     */
    public /*out*/ readonly expectedOutages!: pulumi.Output<outputs.compute.beta.InterconnectOutageNotificationResponse[]>;
    /**
     * IP address configured on the Google side of the Interconnect link. This can be used only for ping tests.
     */
    public /*out*/ readonly googleIpAddress!: pulumi.Output<string>;
    /**
     * Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.
     */
    public /*out*/ readonly googleReferenceId!: pulumi.Output<string>;
    /**
     * A list of the URLs of all InterconnectAttachments configured to use this Interconnect.
     */
    public /*out*/ readonly interconnectAttachments!: pulumi.Output<string[]>;
    /**
     * Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
     */
    public readonly interconnectType!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#interconnect for interconnects.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this Interconnect, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Interconnect.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
     */
    public readonly linkType!: pulumi.Output<string>;
    /**
     * URL of the InterconnectLocation object that represents where this connection is to be provisioned.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Stackdriver logs alerting and Cloud Notifications.
     */
    public readonly nocContactEmail!: pulumi.Output<string>;
    /**
     * The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use. Attachments may be provisioned on this Interconnect. - OS_UNPROVISIONED: An Interconnect that has not completed turnup. No attachments may be provisioned on this Interconnect. - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
     */
    public /*out*/ readonly operationalStatus!: pulumi.Output<string>;
    /**
     * IP address configured on the customer side of the Interconnect link. The customer should configure this IP address during turnup when prompted by Google NOC. This can be used only for ping tests.
     */
    public /*out*/ readonly peerIpAddress!: pulumi.Output<string>;
    /**
     * Number of links actually provisioned in this interconnect.
     */
    public /*out*/ readonly provisionedLinkCount!: pulumi.Output<number>;
    /**
     * Target number of physical links in the link bundle, as requested by the customer.
     */
    public readonly requestedLinkCount!: pulumi.Output<number>;
    /**
     * Set to true if the resource satisfies the zone separation organization policy constraints and false otherwise. Defaults to false if the field is not present.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use. Attachments may be provisioned on this Interconnect. - UNPROVISIONED: The Interconnect has not completed turnup. No attachments may be provisioned on this Interconnect. - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance. No attachments may be provisioned or updated on this Interconnect. 
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Interconnect resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InterconnectArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["adminEnabled"] = args ? args.adminEnabled : undefined;
            resourceInputs["customerName"] = args ? args.customerName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["interconnectType"] = args ? args.interconnectType : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["linkType"] = args ? args.linkType : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nocContactEmail"] = args ? args.nocContactEmail : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["requestedLinkCount"] = args ? args.requestedLinkCount : undefined;
            resourceInputs["circuitInfos"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["expectedOutages"] = undefined /*out*/;
            resourceInputs["googleIpAddress"] = undefined /*out*/;
            resourceInputs["googleReferenceId"] = undefined /*out*/;
            resourceInputs["interconnectAttachments"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["operationalStatus"] = undefined /*out*/;
            resourceInputs["peerIpAddress"] = undefined /*out*/;
            resourceInputs["provisionedLinkCount"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["adminEnabled"] = undefined /*out*/;
            resourceInputs["circuitInfos"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["customerName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["expectedOutages"] = undefined /*out*/;
            resourceInputs["googleIpAddress"] = undefined /*out*/;
            resourceInputs["googleReferenceId"] = undefined /*out*/;
            resourceInputs["interconnectAttachments"] = undefined /*out*/;
            resourceInputs["interconnectType"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["linkType"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nocContactEmail"] = undefined /*out*/;
            resourceInputs["operationalStatus"] = undefined /*out*/;
            resourceInputs["peerIpAddress"] = undefined /*out*/;
            resourceInputs["provisionedLinkCount"] = undefined /*out*/;
            resourceInputs["requestedLinkCount"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Interconnect.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Interconnect resource.
 */
export interface InterconnectArgs {
    /**
     * Administrative status of the interconnect. When this is set to true, the Interconnect is functional and can carry traffic. When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
     */
    adminEnabled?: pulumi.Input<boolean>;
    /**
     * Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
     */
    customerName?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner. - DEDICATED: A dedicated physical interconnection with the customer. Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
     */
    interconnectType?: pulumi.Input<enums.compute.beta.InterconnectInterconnectType>;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle.
     */
    linkType?: pulumi.Input<enums.compute.beta.InterconnectLinkType>;
    /**
     * URL of the InterconnectLocation object that represents where this connection is to be provisioned.
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect. If specified, this will be used for notifications in addition to all other forms described, such as Stackdriver logs alerting and Cloud Notifications.
     */
    nocContactEmail?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * Target number of physical links in the link bundle, as requested by the customer.
     */
    requestedLinkCount?: pulumi.Input<number>;
}
