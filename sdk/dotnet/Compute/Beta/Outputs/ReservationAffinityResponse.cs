// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Specifies the reservations that this instance can consume from.
    /// </summary>
    [OutputType]
    public sealed class ReservationAffinityResponse
    {
        /// <summary>
        /// Specifies the type of reservation from which this instance can consume resources: ANY_RESERVATION (default), SPECIFIC_RESERVATION, or NO_RESERVATION. See Consuming reserved instances for examples.
        /// </summary>
        public readonly string ConsumeReservationType;
        /// <summary>
        /// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify googleapis.com/reservation-name as the key and specify the name of your reservation as its value.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Corresponds to the label values of a reservation resource. This can be either a name to a reservation in the same project or "projects/different-project/reservations/some-reservation-name" to target a shared reservation in the same zone but in a different project.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ReservationAffinityResponse(
            string consumeReservationType,

            string key,

            ImmutableArray<string> values)
        {
            ConsumeReservationType = consumeReservationType;
            Key = key;
            Values = values;
        }
    }
}
