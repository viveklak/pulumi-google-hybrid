// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1.Outputs
{

    /// <summary>
    /// Represents a target of an invocation over HTTP.
    /// </summary>
    [OutputType]
    public sealed class DestinationResponse
    {
        /// <summary>
        /// The Cloud Function resource name. Only Cloud Functions V2 is supported. Format: `projects/{project}/locations/{location}/functions/{function}`
        /// </summary>
        public readonly string CloudFunction;
        /// <summary>
        /// Cloud Run fully-managed resource that receives the events. The resource should be in the same project as the trigger.
        /// </summary>
        public readonly Outputs.CloudRunResponse CloudRun;
        /// <summary>
        /// A GKE service capable of receiving events. The service should be running in the same project as the trigger.
        /// </summary>
        public readonly Outputs.GKEResponse Gke;

        [OutputConstructor]
        private DestinationResponse(
            string cloudFunction,

            Outputs.CloudRunResponse cloudRun,

            Outputs.GKEResponse gke)
        {
            CloudFunction = cloudFunction;
            CloudRun = cloudRun;
            Gke = gke;
        }
    }
}
