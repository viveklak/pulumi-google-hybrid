// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class FutureReservationStatusResponse
    {
        /// <summary>
        /// Fully qualified urls of the automatically created reservations at start_time.
        /// </summary>
        public readonly ImmutableArray<string> AutoCreatedReservations;
        /// <summary>
        /// This count indicates the fulfilled capacity so far. This is set during "PROVISIONING" state. This count also includes capacity delivered as part of existing matching reservations.
        /// </summary>
        public readonly string FulfilledCount;
        /// <summary>
        /// Time when Future Reservation would become LOCKED, after which no modifications to Future Reservation will be allowed. Applicable only after the Future Reservation is in the APPROVED state. The lock_time is an RFC3339 string. The procurement_status will transition to PROCURING state at this time.
        /// </summary>
        public readonly string LockTime;
        /// <summary>
        /// Current state of this Future Reservation
        /// </summary>
        public readonly string ProcurementStatus;

        [OutputConstructor]
        private FutureReservationStatusResponse(
            ImmutableArray<string> autoCreatedReservations,

            string fulfilledCount,

            string lockTime,

            string procurementStatus)
        {
            AutoCreatedReservations = autoCreatedReservations;
            FulfilledCount = fulfilledCount;
            LockTime = lockTime;
            ProcurementStatus = procurementStatus;
        }
    }
}
