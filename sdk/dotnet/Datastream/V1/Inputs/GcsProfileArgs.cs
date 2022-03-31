// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Inputs
{

    /// <summary>
    /// Cloud Storage bucket profile.
    /// </summary>
    public sealed class GcsProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Cloud Storage bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The root path inside the Cloud Storage bucket.
        /// </summary>
        [Input("rootPath")]
        public Input<string>? RootPath { get; set; }

        public GcsProfileArgs()
        {
        }
    }
}
