// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1.Inputs
{

    /// <summary>
    /// Represents a target of an invocation over HTTP.
    /// </summary>
    public sealed class DestinationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Cloud Function resource name. Only Cloud Functions V2 is supported. Format: `projects/{project}/locations/{location}/functions/{function}`
        /// </summary>
        [Input("cloudFunction")]
        public Input<string>? CloudFunction { get; set; }

        /// <summary>
        /// Cloud Run fully-managed resource that receives the events. The resource should be in the same project as the trigger.
        /// </summary>
        [Input("cloudRun")]
        public Input<Inputs.CloudRunArgs>? CloudRun { get; set; }

        /// <summary>
        /// A GKE service capable of receiving events. The service should be running in the same project as the trigger.
        /// </summary>
        [Input("gke")]
        public Input<Inputs.GKEArgs>? Gke { get; set; }

        /// <summary>
        /// The resource name of the Workflow whose Executions are triggered by the events. The Workflow resource should be deployed in the same project as the trigger. Format: `projects/{project}/locations/{location}/workflows/{workflow}`
        /// </summary>
        [Input("workflow")]
        public Input<string>? Workflow { get; set; }

        public DestinationArgs()
        {
        }
    }
}
