// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3Beta1.Inputs
{

    /// <summary>
    /// The Google Cloud Storage location for the input content.
    /// </summary>
    public sealed class GcsSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Source data URI. For example, `gs://my_bucket/my_object`.
        /// </summary>
        [Input("inputUri", required: true)]
        public Input<string> InputUri { get; set; } = null!;

        public GcsSourceArgs()
        {
        }
    }
}
