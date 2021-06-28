// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a ServiceAttachment in the specified project in the given scope using the parameters that are included in the request.
 */
export class ServiceAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceAttachment {
        return new ServiceAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:ServiceAttachment';

    /**
     * Returns true if the given object is an instance of ServiceAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAttachment.__pulumiType;
    }

    /**
     * An array of connections for all the consumers connected to this service attachment.
     */
    public /*out*/ readonly connectedEndpoints!: pulumi.Output<outputs.compute.alpha.ServiceAttachmentConnectedEndpointResponse[]>;
    /**
     * The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
     */
    public readonly connectionPreference!: pulumi.Output<string>;
    /**
     * Projects that are allowed to connect to this service attachment.
     */
    public readonly consumerAcceptLists!: pulumi.Output<outputs.compute.alpha.ServiceAttachmentConsumerProjectLimitResponse[]>;
    /**
     * An array of forwarding rules for all the consumers connected to this service attachment.
     */
    public /*out*/ readonly consumerForwardingRules!: pulumi.Output<outputs.compute.alpha.ServiceAttachmentConsumerForwardingRuleResponse[]>;
    /**
     * Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
     */
    public readonly consumerRejectLists!: pulumi.Output<string[]>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
     */
    public readonly enableProxyProtocol!: pulumi.Output<boolean>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ServiceAttachment. An up-to-date fingerprint must be provided in order to patch/update the ServiceAttachment; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the ServiceAttachment.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#serviceAttachment for service attachments.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
     */
    public readonly natSubnets!: pulumi.Output<string[]>;
    /**
     * The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
     */
    public readonly producerForwardingRule!: pulumi.Output<string>;
    /**
     * An 128-bit global unique ID of the PSC service attachment.
     */
    public /*out*/ readonly pscServiceAttachmentId!: pulumi.Output<outputs.compute.alpha.Uint128Response>;
    /**
     * URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The URL of a service serving the endpoint identified by this service attachment.
     */
    public readonly targetService!: pulumi.Output<string>;

    /**
     * Create a ServiceAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["connectionPreference"] = args ? args.connectionPreference : undefined;
            inputs["consumerAcceptLists"] = args ? args.consumerAcceptLists : undefined;
            inputs["consumerRejectLists"] = args ? args.consumerRejectLists : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableProxyProtocol"] = args ? args.enableProxyProtocol : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natSubnets"] = args ? args.natSubnets : undefined;
            inputs["producerForwardingRule"] = args ? args.producerForwardingRule : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["targetService"] = args ? args.targetService : undefined;
            inputs["connectedEndpoints"] = undefined /*out*/;
            inputs["consumerForwardingRules"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["pscServiceAttachmentId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        } else {
            inputs["connectedEndpoints"] = undefined /*out*/;
            inputs["connectionPreference"] = undefined /*out*/;
            inputs["consumerAcceptLists"] = undefined /*out*/;
            inputs["consumerForwardingRules"] = undefined /*out*/;
            inputs["consumerRejectLists"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["enableProxyProtocol"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["natSubnets"] = undefined /*out*/;
            inputs["producerForwardingRule"] = undefined /*out*/;
            inputs["pscServiceAttachmentId"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["targetService"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceAttachment resource.
 */
export interface ServiceAttachmentArgs {
    /**
     * The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
     */
    connectionPreference?: pulumi.Input<enums.compute.alpha.ServiceAttachmentConnectionPreference>;
    /**
     * Projects that are allowed to connect to this service attachment.
     */
    consumerAcceptLists?: pulumi.Input<pulumi.Input<inputs.compute.alpha.ServiceAttachmentConsumerProjectLimitArgs>[]>;
    /**
     * Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
     */
    consumerRejectLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
     */
    enableProxyProtocol?: pulumi.Input<boolean>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
     */
    natSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
     */
    producerForwardingRule?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    region: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * The URL of a service serving the endpoint identified by this service attachment.
     */
    targetService?: pulumi.Input<string>;
}
