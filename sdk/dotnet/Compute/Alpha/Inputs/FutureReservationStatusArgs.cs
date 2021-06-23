// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// [Output only] Represents status related to the future reservation.
    /// </summary>
    public sealed class FutureReservationStatusArgs : Pulumi.ResourceArgs
    {
        [Input("autoCreatedReservations")]
        private InputList<string>? _autoCreatedReservations;

        /// <summary>
        /// Fully qualified urls of the automatically created reservations at start_time.
        /// </summary>
        public InputList<string> AutoCreatedReservations
        {
            get => _autoCreatedReservations ?? (_autoCreatedReservations = new InputList<string>());
            set => _autoCreatedReservations = value;
        }

        /// <summary>
        /// This count indicates the fulfilled capacity so far. This is set during "PROVISIONING" state. This count also includes capacity delivered as part of existing matching reservations.
        /// </summary>
        [Input("fulfilledCount")]
        public Input<string>? FulfilledCount { get; set; }

        /// <summary>
        /// Time when Future Reservation would become LOCKED, after which no modifications to Future Reservation will be allowed. Applicable only after the Future Reservation is in the APPROVED state. The lock_time is an RFC3339 string. The procurement_status will transition to PROCURING state at this time.
        /// </summary>
        [Input("lockTime")]
        public Input<string>? LockTime { get; set; }

        /// <summary>
        /// Current state of this Future Reservation
        /// </summary>
        [Input("procurementStatus")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.FutureReservationStatusProcurementStatus>? ProcurementStatus { get; set; }

        public FutureReservationStatusArgs()
        {
        }
    }
}
