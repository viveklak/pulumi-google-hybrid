// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    /// <summary>
    /// An artifact that can be deployed in some runtime.
    /// </summary>
    [OutputType]
    public sealed class DeploymentNoteResponse
    {
        /// <summary>
        /// Resource URI for the artifact being deployed.
        /// </summary>
        public readonly ImmutableArray<string> ResourceUri;

        [OutputConstructor]
        private DeploymentNoteResponse(ImmutableArray<string> resourceUri)
        {
            ResourceUri = resourceUri;
        }
    }
}
