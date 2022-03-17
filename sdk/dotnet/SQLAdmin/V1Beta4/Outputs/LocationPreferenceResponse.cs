// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Outputs
{

    /// <summary>
    /// Preferred location. This specifies where a Cloud SQL instance is located. Note that if the preferred location is not available, the instance will be located as close as possible within the region. Only one location may be specified.
    /// </summary>
    [OutputType]
    public sealed class LocationPreferenceResponse
    {
        /// <summary>
        /// The App Engine application to follow, it must be in the same region as the Cloud SQL instance. WARNING: Changing this might restart the instance.
        /// </summary>
        public readonly string FollowGaeApplication;
        /// <summary>
        /// This is always `sql#locationPreference`.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The preferred Compute Engine zone for the secondary/failover (for example: us-central1-a, us-central1-b, etc.).
        /// </summary>
        public readonly string SecondaryZone;
        /// <summary>
        /// The preferred Compute Engine zone (for example: us-central1-a, us-central1-b, etc.). WARNING: Changing this might restart the instance.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private LocationPreferenceResponse(
            string followGaeApplication,

            string kind,

            string secondaryZone,

            string zone)
        {
            FollowGaeApplication = followGaeApplication;
            Kind = kind;
            SecondaryZone = secondaryZone;
            Zone = zone;
        }
    }
}
