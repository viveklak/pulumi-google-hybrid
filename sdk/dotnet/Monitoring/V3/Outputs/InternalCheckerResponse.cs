// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Outputs
{

    /// <summary>
    /// An internal checker allows Uptime checks to run on private/internal GCP resources.
    /// </summary>
    [OutputType]
    public sealed class InternalCheckerResponse
    {
        /// <summary>
        /// The checker's human-readable name. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The GCP zone the Uptime check should egress from. Only respected for internal Uptime checks, where internal_network is specified.
        /// </summary>
        public readonly string GcpZone;
        /// <summary>
        /// A unique resource name for this InternalChecker. The format is: projects/[PROJECT_ID_OR_NUMBER]/internalCheckers/[INTERNAL_CHECKER_ID] [PROJECT_ID_OR_NUMBER] is the Stackdriver Workspace project for the Uptime check config associated with the internal checker.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The GCP VPC network (https://cloud.google.com/vpc/docs/vpc) where the internal resource lives (ex: "default").
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// The GCP project ID where the internal checker lives. Not necessary the same as the Workspace project.
        /// </summary>
        public readonly string PeerProjectId;
        /// <summary>
        /// The current operational state of the internal checker.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private InternalCheckerResponse(
            string displayName,

            string gcpZone,

            string name,

            string network,

            string peerProjectId,

            string state)
        {
            DisplayName = displayName;
            GcpZone = gcpZone;
            Name = name;
            Network = network;
            PeerProjectId = peerProjectId;
            State = state;
        }
    }
}
