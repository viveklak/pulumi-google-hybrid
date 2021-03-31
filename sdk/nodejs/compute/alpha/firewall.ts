// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
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
    public static readonly __pulumiType = 'google-cloud:compute/alpha:Firewall';

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
            if ((!args || args.firewall === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firewall'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["allowed"] = args ? args.allowed : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["denied"] = args ? args.denied : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationRanges"] = args ? args.destinationRanges : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["enableLogging"] = args ? args.enableLogging : undefined;
            inputs["firewall"] = args ? args.firewall : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["selfLinkWithId"] = args ? args.selfLinkWithId : undefined;
            inputs["sourceRanges"] = args ? args.sourceRanges : undefined;
            inputs["sourceServiceAccounts"] = args ? args.sourceServiceAccounts : undefined;
            inputs["sourceTags"] = args ? args.sourceTags : undefined;
            inputs["targetServiceAccounts"] = args ? args.targetServiceAccounts : undefined;
            inputs["targetTags"] = args ? args.targetTags : undefined;
        } else {
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
    readonly allowed?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection.
     */
    readonly denied?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
     */
    readonly destinationRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS`. The default is `INGRESS`. For `INGRESS` traffic, you cannot specify the destinationRanges field, and for `EGRESS` traffic, you cannot specify the sourceRanges or sourceTags fields.
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * Denotes whether the firewall rule is disabled. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Deprecated in favor of enable in LogConfig. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported t Cloud Logging.
     */
    readonly enableLogging?: pulumi.Input<boolean>;
    readonly firewall: pulumi.Input<string>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output Only] Type of the resource. Always compute#firewall for firewall rules.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * This field denotes the logging options for a particular firewall rule. If logging is enabled, logs will be exported to Cloud Logging.
     */
    readonly logConfig?: pulumi.Input<inputs.compute.alpha.FirewallLogConfig>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * URL of the network resource for this firewall rule. If not specified when creating a firewall rule, the default network is used:
     * global/networks/default
     * If you choose to specify this field, you can specify the network as a full or partial URL. For example, the following are all valid URLs:  
     * - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network 
     * - projects/myproject/global/networks/my-network 
     * - global/networks/default
     */
    readonly network?: pulumi.Input<string>;
    /**
     * Priority for this rule. This is an integer between `0` and `65535`, both inclusive. The default value is `1000`. Relative priorities determine which rule takes effect if multiple rules apply. Lower values indicate higher priority. For example, a rule with priority `0` has higher precedence than a rule with priority `1`. DENY rules take precedence over ALLOW rules if they have equal priority. Note that VPC networks have implied rules with a priority of `65535`. To avoid conflicts with the implied rules, use a priority number less than `65535`.
     */
    readonly priority?: pulumi.Input<number>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId?: pulumi.Input<string>;
    /**
     * If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the rule to apply. Only IPv4 is supported.
     */
    readonly sourceRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both fields for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.
     */
    readonly sourceServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network. Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network. Because tags are associated with instances, not IP addresses. One or both of sourceRanges and sourceTags may be set. If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the firewall to apply.
     */
    readonly sourceTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
     */
    readonly targetServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of tags that controls which instances the firewall rule applies to. If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags. If no targetTags are specified, the firewall rule applies to all instances on the specified network.
     */
    readonly targetTags?: pulumi.Input<pulumi.Input<string>[]>;
}
