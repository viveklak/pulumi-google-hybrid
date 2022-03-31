// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Outputs
{

    /// <summary>
    /// An image built by the pipeline.
    /// </summary>
    [OutputType]
    public sealed class BuiltImageResponse
    {
        /// <summary>
        /// Docker Registry 2.0 digest.
        /// </summary>
        public readonly string Digest;
        /// <summary>
        /// Name used to push the container image to Google Container Registry, as presented to `docker push`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Stores timing information for pushing the specified image.
        /// </summary>
        public readonly Outputs.TimeSpanResponse PushTiming;

        [OutputConstructor]
        private BuiltImageResponse(
            string digest,

            string name,

            Outputs.TimeSpanResponse pushTiming)
        {
            Digest = digest;
            Name = name;
            PushTiming = pushTiming;
        }
    }
}
