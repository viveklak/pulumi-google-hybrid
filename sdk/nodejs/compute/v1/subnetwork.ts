// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a subnetwork in the specified project using the data included in the request.
 */
export class Subnetwork extends pulumi.CustomResource {
    /**
     * Get an existing Subnetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Subnetwork {
        return new Subnetwork(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/v1:Subnetwork';

    /**
     * Returns true if the given object is an instance of Subnetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnetwork.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
     */
    public readonly enableFlowLogs!: pulumi.Output<boolean>;
    /**
     * The range of external IPv6 addresses that are owned by this subnetwork.
     */
    public /*out*/ readonly externalIpv6Prefix!: pulumi.Output<string>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The gateway address for default routes to reach destination addresses outside this subnetwork.
     */
    public /*out*/ readonly gatewayAddress!: pulumi.Output<string>;
    /**
     * The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
     */
    public readonly ipCidrRange!: pulumi.Output<string>;
    /**
     * The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct path.
     */
    public readonly ipv6AccessType!: pulumi.Output<string>;
    /**
     * The range of internal IPv6 addresses that are owned by this subnetwork.
     */
    public /*out*/ readonly ipv6CidrRange!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#subnetwork for Subnetwork resources.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.v1.SubnetworkLogConfigResponse>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. This field can be set only at resource creation time.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
     */
    public readonly privateIpGoogleAccess!: pulumi.Output<boolean>;
    /**
     * The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority. This field can be both set at resource creation time and updated using patch.
     */
    public readonly privateIpv6GoogleAccess!: pulumi.Output<string>;
    /**
     * The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
     */
    public readonly purpose!: pulumi.Output<string>;
    /**
     * URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
     */
    public readonly secondaryIpRanges!: pulumi.Output<outputs.compute.v1.SubnetworkSecondaryRangeResponse[]>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will be used. This field can be both set at resource creation time and updated using patch.
     */
    public readonly stackType!: pulumi.Output<string>;
    /**
     * The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Subnetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetworkArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableFlowLogs"] = args ? args.enableFlowLogs : undefined;
            resourceInputs["ipCidrRange"] = args ? args.ipCidrRange : undefined;
            resourceInputs["ipv6AccessType"] = args ? args.ipv6AccessType : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["privateIpGoogleAccess"] = args ? args.privateIpGoogleAccess : undefined;
            resourceInputs["privateIpv6GoogleAccess"] = args ? args.privateIpv6GoogleAccess : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["purpose"] = args ? args.purpose : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["secondaryIpRanges"] = args ? args.secondaryIpRanges : undefined;
            resourceInputs["stackType"] = args ? args.stackType : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["externalIpv6Prefix"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["gatewayAddress"] = undefined /*out*/;
            resourceInputs["ipv6CidrRange"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["enableFlowLogs"] = undefined /*out*/;
            resourceInputs["externalIpv6Prefix"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["gatewayAddress"] = undefined /*out*/;
            resourceInputs["ipCidrRange"] = undefined /*out*/;
            resourceInputs["ipv6AccessType"] = undefined /*out*/;
            resourceInputs["ipv6CidrRange"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["logConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["privateIpGoogleAccess"] = undefined /*out*/;
            resourceInputs["privateIpv6GoogleAccess"] = undefined /*out*/;
            resourceInputs["purpose"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
            resourceInputs["secondaryIpRanges"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["stackType"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Subnetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Subnetwork resource.
 */
export interface SubnetworkArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
     */
    enableFlowLogs?: pulumi.Input<boolean>;
    /**
     * The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
     */
    ipCidrRange?: pulumi.Input<string>;
    /**
     * The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct path.
     */
    ipv6AccessType?: pulumi.Input<enums.compute.v1.SubnetworkIpv6AccessType>;
    /**
     * This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
     */
    logConfig?: pulumi.Input<inputs.compute.v1.SubnetworkLogConfigArgs>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. This field can be set only at resource creation time.
     */
    network?: pulumi.Input<string>;
    /**
     * Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
     */
    privateIpGoogleAccess?: pulumi.Input<boolean>;
    /**
     * The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority. This field can be both set at resource creation time and updated using patch.
     */
    privateIpv6GoogleAccess?: pulumi.Input<enums.compute.v1.SubnetworkPrivateIpv6GoogleAccess>;
    project?: pulumi.Input<string>;
    /**
     * The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
     */
    purpose?: pulumi.Input<enums.compute.v1.SubnetworkPurpose>;
    /**
     * URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
     */
    region: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
     */
    role?: pulumi.Input<enums.compute.v1.SubnetworkRole>;
    /**
     * An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
     */
    secondaryIpRanges?: pulumi.Input<pulumi.Input<inputs.compute.v1.SubnetworkSecondaryRangeArgs>[]>;
    /**
     * The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will be used. This field can be both set at resource creation time and updated using patch.
     */
    stackType?: pulumi.Input<enums.compute.v1.SubnetworkStackType>;
}
