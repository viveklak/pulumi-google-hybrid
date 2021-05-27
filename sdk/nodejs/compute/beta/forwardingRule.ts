// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a ForwardingRule resource in the specified project and region using the data included in the request.
 */
export class ForwardingRule extends pulumi.CustomResource {
    /**
     * Get an existing ForwardingRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ForwardingRule {
        return new ForwardingRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:ForwardingRule';

    /**
     * Returns true if the given object is an instance of ForwardingRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ForwardingRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ForwardingRule.__pulumiType;
    }

    /**
     * IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule.
     *
     * If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address:
     *
     * * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name * Partial URL or by name, as in:  
     * - projects/project_id/regions/region/addresses/address-name 
     * - regions/region/addresses/address-name 
     * - global/addresses/address-name 
     * - address-name  
     *
     * The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
     *
     * Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true.
     *
     * For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
     */
    public readonly IPAddress!: pulumi.Output<string>;
    /**
     * The IP protocol to which this rule applies.
     *
     * For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP.
     *
     * The valid IP protocols are different for different load balancing products:  
     * - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid. 
     * - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid.  
     * - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid. 
     * - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid. 
     * - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid.
     */
    public readonly IPProtocol!: pulumi.Output<string>;
    /**
     * This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. This field cannot be used with port or portRange fields.
     *
     * When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.
     */
    public readonly allPorts!: pulumi.Output<boolean>;
    /**
     * This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
     */
    public readonly allowGlobalAccess!: pulumi.Output<boolean>;
    /**
     * Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
     */
    public readonly backendService!: pulumi.Output<string>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request.
     *
     * To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
     */
    public readonly fingerprint!: pulumi.Output<string>;
    /**
     * The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
     */
    public readonly ipVersion!: pulumi.Output<string>;
    /**
     * Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
     */
    public readonly isMirroringCollector!: pulumi.Output<boolean>;
    /**
     * [Output Only] Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
     */
    public readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Specifies the forwarding rule type.
     *
     *  
     * - EXTERNAL is used for:  
     * - Classic Cloud VPN gateways 
     * - Protocol forwarding to VMs from an external IP address 
     * - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing    
     * - INTERNAL is used for:  
     * - Protocol forwarding to VMs from an internal IP address 
     * - Internal TCP/UDP Load Balancing   
     * - INTERNAL_MANAGED is used for:  
     * - Internal HTTP(S) Load Balancing   
     * - INTERNAL_SELF_MANAGED is used for:  
     * - Traffic Director    
     *
     * For more information about forwarding rules, refer to Forwarding rule concepts.
     */
    public readonly loadBalancingScheme!: pulumi.Output<string>;
    /**
     * Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to Loadbalancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule will not be visible to those proxies.
     * For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match.
     * metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references.
     * metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     */
    public readonly metadataFilters!: pulumi.Output<outputs.compute.beta.MetadataFilterResponse[]>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * This field is not used for external load balancing.
     *
     * For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.
     *
     * For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD.
     *
     * For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM.
     *
     * If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
     */
    public readonly networkTier!: pulumi.Output<string>;
    /**
     * This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP.
     *
     * Packets addressed to ports in the specified range will be forwarded to target or  backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges.
     *
     * Some types of forwarding target have constraints on the acceptable ports:  
     * - TargetHttpProxy: 80, 8080 
     * - TargetHttpsProxy: 443 
     * - TargetGrpcProxy: no constraints 
     * - TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222 
     * - TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222 
     * - TargetVpnGateway: 500, 4500
     */
    public readonly portRange!: pulumi.Output<string>;
    /**
     * The ports field is only supported when the forwarding rule references a backend_service directly. Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing. Only packets addressed to the specified list of ports are forwarded to backends.
     *
     * You can only use one of ports and port_range, or allPorts. The three are mutually exclusive.
     *
     * You can specify a list of up to five ports, which can be non-contiguous.
     *
     * For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports.
     *
     * For more information, see [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).
     */
    public readonly ports!: pulumi.Output<string[]>;
    /**
     * [Output Only] The PSC connection id of the PSC Forwarding Rule.
     */
    public readonly pscConnectionId!: pulumi.Output<string>;
    /**
     * [Output Only] URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
     *
     * It is only supported for Internal TCP/UDP Load Balancing and Internal HTTP(S) Load Balancing.
     */
    public readonly serviceDirectoryRegistrations!: pulumi.Output<outputs.compute.beta.ForwardingRuleServiceDirectoryRegistrationResponse[]>;
    /**
     * An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name.
     *
     * The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     *
     * This field is only used for internal load balancing.
     */
    public readonly serviceLabel!: pulumi.Output<string>;
    /**
     * [Output Only] The internal fully qualified service name for this Forwarding Rule.
     *
     * This field is only used for internal load balancing.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * This field is only used for internal load balancing.
     *
     * For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule.
     *
     * If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.
     */
    public readonly subnetwork!: pulumi.Output<string>;
    public readonly target!: pulumi.Output<string>;

    /**
     * Create a ForwardingRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ForwardingRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["IPAddress"] = args ? args.IPAddress : undefined;
            inputs["IPProtocol"] = args ? args.IPProtocol : undefined;
            inputs["allPorts"] = args ? args.allPorts : undefined;
            inputs["allowGlobalAccess"] = args ? args.allowGlobalAccess : undefined;
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["fingerprint"] = args ? args.fingerprint : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["isMirroringCollector"] = args ? args.isMirroringCollector : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["labelFingerprint"] = args ? args.labelFingerprint : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            inputs["metadataFilters"] = args ? args.metadataFilters : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["networkTier"] = args ? args.networkTier : undefined;
            inputs["portRange"] = args ? args.portRange : undefined;
            inputs["ports"] = args ? args.ports : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pscConnectionId"] = args ? args.pscConnectionId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["serviceDirectoryRegistrations"] = args ? args.serviceDirectoryRegistrations : undefined;
            inputs["serviceLabel"] = args ? args.serviceLabel : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["target"] = args ? args.target : undefined;
        } else {
            inputs["IPAddress"] = undefined /*out*/;
            inputs["IPProtocol"] = undefined /*out*/;
            inputs["allPorts"] = undefined /*out*/;
            inputs["allowGlobalAccess"] = undefined /*out*/;
            inputs["backendService"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["ipVersion"] = undefined /*out*/;
            inputs["isMirroringCollector"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["loadBalancingScheme"] = undefined /*out*/;
            inputs["metadataFilters"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["network"] = undefined /*out*/;
            inputs["networkTier"] = undefined /*out*/;
            inputs["portRange"] = undefined /*out*/;
            inputs["ports"] = undefined /*out*/;
            inputs["pscConnectionId"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["serviceDirectoryRegistrations"] = undefined /*out*/;
            inputs["serviceLabel"] = undefined /*out*/;
            inputs["serviceName"] = undefined /*out*/;
            inputs["subnetwork"] = undefined /*out*/;
            inputs["target"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ForwardingRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ForwardingRule resource.
 */
export interface ForwardingRuleArgs {
    /**
     * IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule.
     *
     * If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address:
     *
     * * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name * Partial URL or by name, as in:  
     * - projects/project_id/regions/region/addresses/address-name 
     * - regions/region/addresses/address-name 
     * - global/addresses/address-name 
     * - address-name  
     *
     * The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
     *
     * Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true.
     *
     * For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
     */
    readonly IPAddress?: pulumi.Input<string>;
    /**
     * The IP protocol to which this rule applies.
     *
     * For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP.
     *
     * The valid IP protocols are different for different load balancing products:  
     * - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid. 
     * - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid.  
     * - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid. 
     * - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid. 
     * - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid.
     */
    readonly IPProtocol?: pulumi.Input<string>;
    /**
     * This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. This field cannot be used with port or portRange fields.
     *
     * When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.
     */
    readonly allPorts?: pulumi.Input<boolean>;
    /**
     * This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
     */
    readonly allowGlobalAccess?: pulumi.Input<boolean>;
    /**
     * Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
     */
    readonly backendService?: pulumi.Input<string>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request.
     *
     * To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
     */
    readonly ipVersion?: pulumi.Input<string>;
    /**
     * Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
     */
    readonly isMirroringCollector?: pulumi.Input<boolean>;
    /**
     * [Output Only] Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the forwarding rule type.
     *
     *  
     * - EXTERNAL is used for:  
     * - Classic Cloud VPN gateways 
     * - Protocol forwarding to VMs from an external IP address 
     * - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing    
     * - INTERNAL is used for:  
     * - Protocol forwarding to VMs from an internal IP address 
     * - Internal TCP/UDP Load Balancing   
     * - INTERNAL_MANAGED is used for:  
     * - Internal HTTP(S) Load Balancing   
     * - INTERNAL_SELF_MANAGED is used for:  
     * - Traffic Director    
     *
     * For more information about forwarding rules, refer to Forwarding rule concepts.
     */
    readonly loadBalancingScheme?: pulumi.Input<string>;
    /**
     * Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to Loadbalancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule will not be visible to those proxies.
     * For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match.
     * metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references.
     * metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     */
    readonly metadataFilters?: pulumi.Input<pulumi.Input<inputs.compute.beta.MetadataFilterArgs>[]>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * This field is not used for external load balancing.
     *
     * For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.
     *
     * For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD.
     *
     * For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM.
     *
     * If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
     */
    readonly networkTier?: pulumi.Input<string>;
    /**
     * This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP.
     *
     * Packets addressed to ports in the specified range will be forwarded to target or  backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges.
     *
     * Some types of forwarding target have constraints on the acceptable ports:  
     * - TargetHttpProxy: 80, 8080 
     * - TargetHttpsProxy: 443 
     * - TargetGrpcProxy: no constraints 
     * - TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222 
     * - TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222 
     * - TargetVpnGateway: 500, 4500
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * The ports field is only supported when the forwarding rule references a backend_service directly. Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing. Only packets addressed to the specified list of ports are forwarded to backends.
     *
     * You can only use one of ports and port_range, or allPorts. The three are mutually exclusive.
     *
     * You can specify a list of up to five ports, which can be non-contiguous.
     *
     * For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports.
     *
     * For more information, see [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).
     */
    readonly ports?: pulumi.Input<pulumi.Input<string>[]>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] The PSC connection id of the PSC Forwarding Rule.
     */
    readonly pscConnectionId?: pulumi.Input<string>;
    /**
     * [Output Only] URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
     *
     * It is only supported for Internal TCP/UDP Load Balancing and Internal HTTP(S) Load Balancing.
     */
    readonly serviceDirectoryRegistrations?: pulumi.Input<pulumi.Input<inputs.compute.beta.ForwardingRuleServiceDirectoryRegistrationArgs>[]>;
    /**
     * An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name.
     *
     * The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     *
     * This field is only used for internal load balancing.
     */
    readonly serviceLabel?: pulumi.Input<string>;
    /**
     * [Output Only] The internal fully qualified service name for this Forwarding Rule.
     *
     * This field is only used for internal load balancing.
     */
    readonly serviceName?: pulumi.Input<string>;
    /**
     * This field is only used for internal load balancing.
     *
     * For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule.
     *
     * If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.
     */
    readonly subnetwork?: pulumi.Input<string>;
    readonly target?: pulumi.Input<string>;
}
