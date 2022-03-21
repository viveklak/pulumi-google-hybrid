// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Outputs
{

    /// <summary>
    /// The status of a cluster and its instances.
    /// </summary>
    [OutputType]
    public sealed class ClusterStatusResponse
    {
        /// <summary>
        /// Optional details of cluster's state.
        /// </summary>
        public readonly string Detail;
        /// <summary>
        /// The cluster's state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Time when this state was entered (see JSON representation of Timestamp (https://developers.google.com/protocol-buffers/docs/proto3#json)).
        /// </summary>
        public readonly string StateStartTime;
        /// <summary>
        /// Additional state information that includes status reported by the agent.
        /// </summary>
        public readonly string Substate;

        [OutputConstructor]
        private ClusterStatusResponse(
            string detail,

            string state,

            string stateStartTime,

            string substate)
        {
            Detail = detail;
            State = state;
            StateStartTime = stateStartTime;
            Substate = substate;
        }
    }
}
