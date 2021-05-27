// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates a firewall rule in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:Firewall")]
    public partial class Firewall : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a permitted connection.
        /// </summary>
        [Output("allowed")]
        public Output<ImmutableArray<Outputs.FirewallAllowedItemResponse>> Allowed { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection.
        /// </summary>
        [Output("denied")]
        public Output<ImmutableArray<Outputs.FirewallDeniedItemResponse>> Denied { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
        /// </summary>
        [Output("destinationRanges")]
        public Output<ImmutableArray<string>> DestinationRanges { get; private set; } = null!;

        /// <summary>
        /// Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS`. The default is `INGRESS`. For `INGRESS` traffic, you cannot specify the destinationRanges field, and for `EGRESS` traffic, you cannot specify the sourceRanges or sourceTags fields.
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// Denotes whether the firewall rule is disabled. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.
        /// </summary>
        [Output("disabled")]
        public Output<bool> Disabled { get; private set; } = null!;

        /// <summary>
        /// Deprecated in favor of enable in LogConfig. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported t Cloud Logging.
        /// </summary>
        [Output("enableLogging")]
        public Output<bool> EnableLogging { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#firewall for firewall rules.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// This field denotes the logging options for a particular firewall rule. If logging is enabled, logs will be exported to Cloud Logging.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.FirewallLogConfigResponse> LogConfig { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// URL of the network resource for this firewall rule. If not specified when creating a firewall rule, the default network is used:
        /// global/networks/default
        /// If you choose to specify this field, you can specify the network as a full or partial URL. For example, the following are all valid URLs:  
        /// - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network 
        /// - projects/myproject/global/networks/my-network 
        /// - global/networks/default
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Priority for this rule. This is an integer between `0` and `65535`, both inclusive. The default value is `1000`. Relative priorities determine which rule takes effect if multiple rules apply. Lower values indicate higher priority. For example, a rule with priority `0` has higher precedence than a rule with priority `1`. DENY rules take precedence over ALLOW rules if they have equal priority. Note that VPC networks have implied rules with a priority of `65535`. To avoid conflicts with the implied rules, use a priority number less than `65535`.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the rule to apply. Only IPv4 is supported.
        /// </summary>
        [Output("sourceRanges")]
        public Output<ImmutableArray<string>> SourceRanges { get; private set; } = null!;

        /// <summary>
        /// If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both fields for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.
        /// </summary>
        [Output("sourceServiceAccounts")]
        public Output<ImmutableArray<string>> SourceServiceAccounts { get; private set; } = null!;

        /// <summary>
        /// If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network. Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network. Because tags are associated with instances, not IP addresses. One or both of sourceRanges and sourceTags may be set. If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the firewall to apply.
        /// </summary>
        [Output("sourceTags")]
        public Output<ImmutableArray<string>> SourceTags { get; private set; } = null!;

        /// <summary>
        /// A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
        /// </summary>
        [Output("targetServiceAccounts")]
        public Output<ImmutableArray<string>> TargetServiceAccounts { get; private set; } = null!;

        /// <summary>
        /// A list of tags that controls which instances the firewall rule applies to. If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags. If no targetTags are specified, the firewall rule applies to all instances on the specified network.
        /// </summary>
        [Output("targetTags")]
        public Output<ImmutableArray<string>> TargetTags { get; private set; } = null!;


        /// <summary>
        /// Create a Firewall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Firewall(string name, FirewallArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:Firewall", name, args ?? new FirewallArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Firewall(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:Firewall", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Firewall resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Firewall Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Firewall(name, id, options);
        }
    }

    public sealed class FirewallArgs : Pulumi.ResourceArgs
    {
        [Input("allowed")]
        private InputList<Inputs.FirewallAllowedItemArgs>? _allowed;

        /// <summary>
        /// The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a permitted connection.
        /// </summary>
        public InputList<Inputs.FirewallAllowedItemArgs> Allowed
        {
            get => _allowed ?? (_allowed = new InputList<Inputs.FirewallAllowedItemArgs>());
            set => _allowed = value;
        }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("denied")]
        private InputList<Inputs.FirewallDeniedItemArgs>? _denied;

        /// <summary>
        /// The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection.
        /// </summary>
        public InputList<Inputs.FirewallDeniedItemArgs> Denied
        {
            get => _denied ?? (_denied = new InputList<Inputs.FirewallDeniedItemArgs>());
            set => _denied = value;
        }

        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationRanges")]
        private InputList<string>? _destinationRanges;

        /// <summary>
        /// If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
        /// </summary>
        public InputList<string> DestinationRanges
        {
            get => _destinationRanges ?? (_destinationRanges = new InputList<string>());
            set => _destinationRanges = value;
        }

        /// <summary>
        /// Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS`. The default is `INGRESS`. For `INGRESS` traffic, you cannot specify the destinationRanges field, and for `EGRESS` traffic, you cannot specify the sourceRanges or sourceTags fields.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Denotes whether the firewall rule is disabled. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Deprecated in favor of enable in LogConfig. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported t Cloud Logging.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#firewall for firewall rules.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// This field denotes the logging options for a particular firewall rule. If logging is enabled, logs will be exported to Cloud Logging.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.FirewallLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL of the network resource for this firewall rule. If not specified when creating a firewall rule, the default network is used:
        /// global/networks/default
        /// If you choose to specify this field, you can specify the network as a full or partial URL. For example, the following are all valid URLs:  
        /// - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network 
        /// - projects/myproject/global/networks/my-network 
        /// - global/networks/default
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Priority for this rule. This is an integer between `0` and `65535`, both inclusive. The default value is `1000`. Relative priorities determine which rule takes effect if multiple rules apply. Lower values indicate higher priority. For example, a rule with priority `0` has higher precedence than a rule with priority `1`. DENY rules take precedence over ALLOW rules if they have equal priority. Note that VPC networks have implied rules with a priority of `65535`. To avoid conflicts with the implied rules, use a priority number less than `65535`.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Input("selfLinkWithId")]
        public Input<string>? SelfLinkWithId { get; set; }

        [Input("sourceRanges")]
        private InputList<string>? _sourceRanges;

        /// <summary>
        /// If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the rule to apply. Only IPv4 is supported.
        /// </summary>
        public InputList<string> SourceRanges
        {
            get => _sourceRanges ?? (_sourceRanges = new InputList<string>());
            set => _sourceRanges = value;
        }

        [Input("sourceServiceAccounts")]
        private InputList<string>? _sourceServiceAccounts;

        /// <summary>
        /// If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both fields for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.
        /// </summary>
        public InputList<string> SourceServiceAccounts
        {
            get => _sourceServiceAccounts ?? (_sourceServiceAccounts = new InputList<string>());
            set => _sourceServiceAccounts = value;
        }

        [Input("sourceTags")]
        private InputList<string>? _sourceTags;

        /// <summary>
        /// If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network. Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network. Because tags are associated with instances, not IP addresses. One or both of sourceRanges and sourceTags may be set. If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the firewall to apply.
        /// </summary>
        public InputList<string> SourceTags
        {
            get => _sourceTags ?? (_sourceTags = new InputList<string>());
            set => _sourceTags = value;
        }

        [Input("targetServiceAccounts")]
        private InputList<string>? _targetServiceAccounts;

        /// <summary>
        /// A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
        /// </summary>
        public InputList<string> TargetServiceAccounts
        {
            get => _targetServiceAccounts ?? (_targetServiceAccounts = new InputList<string>());
            set => _targetServiceAccounts = value;
        }

        [Input("targetTags")]
        private InputList<string>? _targetTags;

        /// <summary>
        /// A list of tags that controls which instances the firewall rule applies to. If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags. If no targetTags are specified, the firewall rule applies to all instances on the specified network.
        /// </summary>
        public InputList<string> TargetTags
        {
            get => _targetTags ?? (_targetTags = new InputList<string>());
            set => _targetTags = value;
        }

        public FirewallArgs()
        {
        }
    }
}
