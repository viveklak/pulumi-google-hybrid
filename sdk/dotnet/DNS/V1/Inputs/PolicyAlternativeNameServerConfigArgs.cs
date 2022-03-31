// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Inputs
{

    public sealed class PolicyAlternativeNameServerConfigArgs : Pulumi.ResourceArgs
    {
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("targetNameServers")]
        private InputList<Inputs.PolicyAlternativeNameServerConfigTargetNameServerArgs>? _targetNameServers;

        /// <summary>
        /// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        /// </summary>
        public InputList<Inputs.PolicyAlternativeNameServerConfigTargetNameServerArgs> TargetNameServers
        {
            get => _targetNameServers ?? (_targetNameServers = new InputList<Inputs.PolicyAlternativeNameServerConfigTargetNameServerArgs>());
            set => _targetNameServers = value;
        }

        public PolicyAlternativeNameServerConfigArgs()
        {
        }
    }
}
