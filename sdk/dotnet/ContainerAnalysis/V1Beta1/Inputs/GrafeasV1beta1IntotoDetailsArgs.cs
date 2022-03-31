// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// This corresponds to a signed in-toto link - it is made up of one or more signatures and the in-toto link itself. This is used for occurrences of a Grafeas in-toto note.
    /// </summary>
    public sealed class GrafeasV1beta1IntotoDetailsArgs : Pulumi.ResourceArgs
    {
        [Input("signatures")]
        private InputList<Inputs.GrafeasV1beta1IntotoSignatureArgs>? _signatures;
        public InputList<Inputs.GrafeasV1beta1IntotoSignatureArgs> Signatures
        {
            get => _signatures ?? (_signatures = new InputList<Inputs.GrafeasV1beta1IntotoSignatureArgs>());
            set => _signatures = value;
        }

        [Input("signed")]
        public Input<Inputs.LinkArgs>? Signed { get; set; }

        public GrafeasV1beta1IntotoDetailsArgs()
        {
        }
    }
}
