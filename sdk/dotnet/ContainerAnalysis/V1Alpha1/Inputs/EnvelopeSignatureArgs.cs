// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// A DSSE signature
    /// </summary>
    public sealed class EnvelopeSignatureArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A reference id to the key being used for signing
        /// </summary>
        [Input("keyid")]
        public Input<string>? Keyid { get; set; }

        /// <summary>
        /// The signature itself
        /// </summary>
        [Input("sig")]
        public Input<string>? Sig { get; set; }

        public EnvelopeSignatureArgs()
        {
        }
    }
}
