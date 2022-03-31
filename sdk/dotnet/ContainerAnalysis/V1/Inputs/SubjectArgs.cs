// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    public sealed class SubjectArgs : Pulumi.ResourceArgs
    {
        [Input("digest")]
        private InputMap<string>? _digest;

        /// <summary>
        /// `"": ""` Algorithms can be e.g. sha256, sha512 See https://github.com/in-toto/attestation/blob/main/spec/field_types.md#DigestSet
        /// </summary>
        public InputMap<string> Digest
        {
            get => _digest ?? (_digest = new InputMap<string>());
            set => _digest = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public SubjectArgs()
        {
        }
    }
}
