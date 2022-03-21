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
    /// The period during which some deployable was active in a runtime.
    /// </summary>
    [OutputType]
    public sealed class DeploymentOccurrenceResponse
    {
        /// <summary>
        /// Address of the runtime element hosting this deployment.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// Configuration used to create this deployment.
        /// </summary>
        public readonly string Config;
        /// <summary>
        /// Beginning of the lifetime of this deployment.
        /// </summary>
        public readonly string DeployTime;
        /// <summary>
        /// Platform hosting this deployment.
        /// </summary>
        public readonly string Platform;
        /// <summary>
        /// Resource URI for the artifact being deployed taken from the deployable field with the same name.
        /// </summary>
        public readonly ImmutableArray<string> ResourceUri;
        /// <summary>
        /// End of the lifetime of this deployment.
        /// </summary>
        public readonly string UndeployTime;
        /// <summary>
        /// Identity of the user that triggered this deployment.
        /// </summary>
        public readonly string UserEmail;

        [OutputConstructor]
        private DeploymentOccurrenceResponse(
            string address,

            string config,

            string deployTime,

            string platform,

            ImmutableArray<string> resourceUri,

            string undeployTime,

            string userEmail)
        {
            Address = address;
            Config = config;
            DeployTime = deployTime;
            Platform = platform;
            ResourceUri = resourceUri;
            UndeployTime = undeployTime;
            UserEmail = userEmail;
        }
    }
}
