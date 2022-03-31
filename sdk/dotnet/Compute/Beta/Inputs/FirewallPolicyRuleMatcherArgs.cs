// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Represents a match condition that incoming traffic is evaluated against. Exactly one field must be specified.
    /// </summary>
    public sealed class FirewallPolicyRuleMatcherArgs : Pulumi.ResourceArgs
    {
        [Input("destIpRanges")]
        private InputList<string>? _destIpRanges;

        /// <summary>
        /// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
        /// </summary>
        public InputList<string> DestIpRanges
        {
            get => _destIpRanges ?? (_destIpRanges = new InputList<string>());
            set => _destIpRanges = value;
        }

        [Input("destRegionCodes")]
        private InputList<string>? _destRegionCodes;

        /// <summary>
        /// Region codes whose IP addresses will be used to match for destination of traffic. Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of dest region codes allowed is 5000.
        /// </summary>
        public InputList<string> DestRegionCodes
        {
            get => _destRegionCodes ?? (_destRegionCodes = new InputList<string>());
            set => _destRegionCodes = value;
        }

        [Input("destThreatIntelligences")]
        private InputList<string>? _destThreatIntelligences;

        /// <summary>
        /// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic destination.
        /// </summary>
        public InputList<string> DestThreatIntelligences
        {
            get => _destThreatIntelligences ?? (_destThreatIntelligences = new InputList<string>());
            set => _destThreatIntelligences = value;
        }

        [Input("layer4Configs")]
        private InputList<Inputs.FirewallPolicyRuleMatcherLayer4ConfigArgs>? _layer4Configs;

        /// <summary>
        /// Pairs of IP protocols and ports that the rule should match.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleMatcherLayer4ConfigArgs> Layer4Configs
        {
            get => _layer4Configs ?? (_layer4Configs = new InputList<Inputs.FirewallPolicyRuleMatcherLayer4ConfigArgs>());
            set => _layer4Configs = value;
        }

        [Input("srcIpRanges")]
        private InputList<string>? _srcIpRanges;

        /// <summary>
        /// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
        /// </summary>
        public InputList<string> SrcIpRanges
        {
            get => _srcIpRanges ?? (_srcIpRanges = new InputList<string>());
            set => _srcIpRanges = value;
        }

        [Input("srcRegionCodes")]
        private InputList<string>? _srcRegionCodes;

        /// <summary>
        /// Region codes whose IP addresses will be used to match for source of traffic. Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of source region codes allowed is 5000.
        /// </summary>
        public InputList<string> SrcRegionCodes
        {
            get => _srcRegionCodes ?? (_srcRegionCodes = new InputList<string>());
            set => _srcRegionCodes = value;
        }

        [Input("srcSecureTags")]
        private InputList<Inputs.FirewallPolicyRuleSecureTagArgs>? _srcSecureTags;

        /// <summary>
        /// List of secure tag values, which should be matched at the source of the traffic. For INGRESS rule, if all the srcSecureTag are INEFFECTIVE, and there is no srcIpRange, this rule will be ignored. Maximum number of source tag values allowed is 256.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleSecureTagArgs> SrcSecureTags
        {
            get => _srcSecureTags ?? (_srcSecureTags = new InputList<Inputs.FirewallPolicyRuleSecureTagArgs>());
            set => _srcSecureTags = value;
        }

        [Input("srcThreatIntelligences")]
        private InputList<string>? _srcThreatIntelligences;

        /// <summary>
        /// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic source.
        /// </summary>
        public InputList<string> SrcThreatIntelligences
        {
            get => _srcThreatIntelligences ?? (_srcThreatIntelligences = new InputList<string>());
            set => _srcThreatIntelligences = value;
        }

        public FirewallPolicyRuleMatcherArgs()
        {
        }
    }
}
