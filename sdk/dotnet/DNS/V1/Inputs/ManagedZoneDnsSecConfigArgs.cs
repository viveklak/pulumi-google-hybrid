// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Inputs
{

    public sealed class ManagedZoneDnsSecConfigArgs : Pulumi.ResourceArgs
    {
        [Input("defaultKeySpecs")]
        private InputList<Inputs.DnsKeySpecArgs>? _defaultKeySpecs;

        /// <summary>
        /// Specifies parameters for generating initial DnsKeys for this ManagedZone. Can only be changed while the state is OFF.
        /// </summary>
        public InputList<Inputs.DnsKeySpecArgs> DefaultKeySpecs
        {
            get => _defaultKeySpecs ?? (_defaultKeySpecs = new InputList<Inputs.DnsKeySpecArgs>());
            set => _defaultKeySpecs = value;
        }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Specifies the mechanism for authenticated denial-of-existence responses. Can only be changed while the state is OFF.
        /// </summary>
        [Input("nonExistence")]
        public Input<Pulumi.GoogleNative.DNS.V1.ManagedZoneDnsSecConfigNonExistence>? NonExistence { get; set; }

        /// <summary>
        /// Specifies whether DNSSEC is enabled, and what mode it is in.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.DNS.V1.ManagedZoneDnsSecConfigState>? State { get; set; }

        public ManagedZoneDnsSecConfigArgs()
        {
        }
    }
}
