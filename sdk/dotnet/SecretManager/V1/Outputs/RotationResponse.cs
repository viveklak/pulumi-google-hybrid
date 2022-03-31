// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SecretManager.V1.Outputs
{

    /// <summary>
    /// The rotation time and period for a Secret. At next_rotation_time, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. Secret.topics must be set to configure rotation.
    /// </summary>
    [OutputType]
    public sealed class RotationResponse
    {
        /// <summary>
        /// Optional. Timestamp in UTC at which the Secret is scheduled to rotate. Cannot be set to less than 300s (5 min) in the future and at most 3153600000s (100 years). next_rotation_time MUST be set if rotation_period is set.
        /// </summary>
        public readonly string NextRotationTime;
        /// <summary>
        /// Input only. The Duration between rotation notifications. Must be in seconds and at least 3600s (1h) and at most 3153600000s (100 years). If rotation_period is set, next_rotation_time must be set. next_rotation_time will be advanced by this period when the service automatically sends rotation notifications.
        /// </summary>
        public readonly string RotationPeriod;

        [OutputConstructor]
        private RotationResponse(
            string nextRotationTime,

            string rotationPeriod)
        {
            NextRotationTime = nextRotationTime;
            RotationPeriod = rotationPeriod;
        }
    }
}
