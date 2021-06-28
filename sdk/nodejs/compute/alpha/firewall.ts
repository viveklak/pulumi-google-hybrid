// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a firewall rule in the specified project using the data included in the request.
 */
export class Firewall extends pulumi.CustomResource {
    /**
     * Get an existing Firewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Firewall {
        return new Firewall(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:Firewall';

    /**
     * Returns true if the given object is an instance of Firewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Firewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Firewall.__pulumiType;
    }

    /**
     * The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a permitted connection.
     */
    public readonly allowed!: pulumi.Output<outputs.compute.alpha.FirewallAllowedItemResponse[]>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection.
     */
    public readonly denied!: pulumi.Output<outputs.compute.alpha.FirewallDeniedItemResponse[]>;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
     */
    public readonly destinationRanges!: pulumi.Output<string[]>;
    /**
     * Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS`. The default is `INGRESS`. For `INGRESS` traffic, you cannot specify the destinationRanges field, and for `EGRESS` traffic, you cannot specify the sourceRanges or sourceTags fields.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * Denotes whether the firewall rule is disabled. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * Type of the resource. Always compute#firewall for firewall rules.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * This field denotes the logging options for a particular firewall rule. If logging is enabled, logs will be exported to Cloud Logging.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.alpha.FirewallLogConfigResponse>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * URL of the network resource for this firewall rule. If not specified when creating a firewall rule, the default network is used:
     * global/networks/default
     * If you choose to specify this field, you can specify the network as a full or partial URL. For example, the following are all valid URLs:  
     * - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network 
     * - projects/myproject/global/networks/my-network 
     * - global/networks/default
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Priority for this rule. This is an integer between `0` and `65535`, both inclusive. The default value is `1000`. Relative priorities determine which rule takes effect if multiple rules apply. Lower values indicate higher priority. For example, a rule with priority `0` has higher precedence than a rule with priority `1`. DENY rules take precedence over ALLOW rules if they have equal priority. Note that VPC networks have implied rules with a priority of `65535`. To avoid conflicts with the implied rules, use a priority number less than `65535`.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the rule to apply. Only IPv4 is supported.
     */
    public readonly sourceRanges!: pulumi.Output<string[]>;
    /**
     * If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both fields for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.
     */
    public readonly sourceServiceAccounts!: pulumi.Output<string[]>;
    /**
     * If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network. Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network. Because tags are associated with instances, not IP addresses. One or both of sourceRanges and sourceTags may be set. If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the firewall to apply.
     */
    public readonly sourceTags!: pulumi.Output<string[]>;
    /**
     * A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
     */
    public readonly targetServiceAccounts!: pulumi.Output<string[]>;
    /**
     * A list of tags that controls which instances the firewall rule applies to. If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags. If no targetTags are specified, the firewall rule applies to all instances on the specified network.
     */
    public readonly targetTags!: pulumi.Output<string[]>;

    /**
     * Create a Firewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["allowed"] = args ? args.allowed : undefined;
            inputs["denied"] = args ? args.denied : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationRanges"] = args ? args.destinationRanges : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["sourceRanges"] = args ? args.sourceRanges : undefined;
            inputs["sourceServiceAccounts"] = args ? args.sourceServiceAccounts : undefined;
            inputs["sourceTags"] = args ? args.sourceTags : undefined;
            inputs["targetServiceAccounts"] = args ? args.targetServiceAccounts : undefined;
            inputs["targetTags"] = args ? args.targetTags : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
        } else {
            inputs["allowed"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["denied"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["destinationRanges"] = undefined /*out*/;
            inputs["direction"] = undefined /*out*/;
            inputs["disabled"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["logConfig"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["network"] = undefined /*out*/;
            inputs["priority"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["sourceRanges"] = undefined /*out*/;
            inputs["sourceServiceAccounts"] = undefined /*out*/;
            inputs["sourceTags"] = undefined /*out*/;
            inputs["targetServiceAccounts"] = undefined /*out*/;
            inputs["targetTags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Firewall.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Firewall resource.
 */
export interface FirewallArgs {
    /**
     * The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a permitted connection.
     */
    allowed?: pulumi.Input<pulumi.Input<inputs.compute.alpha.FirewallAllowedItemArgs>[]>;
    /**
     * The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection.
     */
    denied?: pulumi.Input<pulumi.Input<inputs.compute.alpha.FirewallDeniedItemArgs>[]>;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
     */
    destinationRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS`. The default is `INGRESS`. For `INGRESS` traffic, you cannot specify the destinationRanges field, and for `EGRESS` traffic, you cannot specify the sourceRanges or sourceTags fields.
     */
    direction?: pulumi.Input<enums.compute.alpha.FirewallDirection>;
    /**
     * Denotes whether the firewall rule is disabled. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * This field denotes the logging options for a particular firewall rule. If logging is enabled, logs will be exported to Cloud Logging.
     */
    logConfig?: pulumi.Input<inputs.compute.alpha.FirewallLogConfigArgs>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    name?: pulumi.Input<string>;
    /**
     * URL of the network resource for this firewall rule. If not specified when creating a firewall rule, the default network is used:
     * global/networks/default
     * If you choose to specify this field, you can specify the network as a full or partial URL. For example, the following are all valid URLs:  
     * - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network 
     * - projects/myproject/global/networks/my-network 
     * - global/networks/default
     */
    network?: pulumi.Input<string>;
    /**
     * Priority for this rule. This is an integer between `0` and `65535`, both inclusive. The default value is `1000`. Relative priorities determine which rule takes effect if multiple rules apply. Lower values indicate higher priority. For example, a rule with priority `0` has higher precedence than a rule with priority `1`. DENY rules take precedence over ALLOW rules if they have equal priority. Note that VPC networks have implied rules with a priority of `65535`. To avoid conflicts with the implied rules, use a priority number less than `65535`.
     */
    priority?: pulumi.Input<number>;
    project: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the rule to apply. Only IPv4 is supported.
     */
    sourceRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both fields for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.
     */
    sourceServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network. Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network. Because tags are associated with instances, not IP addresses. One or both of sourceRanges and sourceTags may be set. If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the firewall to apply.
     */
    sourceTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
     */
    targetServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of tags that controls which instances the firewall rule applies to. If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags. If no targetTags are specified, the firewall rule applies to all instances on the specified network.
     */
    targetTags?: pulumi.Input<pulumi.Input<string>[]>;
}
