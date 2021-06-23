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
    public sealed class ReservationResponse
    {
        /// <summary>
        /// [Output Only] Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        /// </summary>
        public readonly string Commitment;
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// [Output Only] Type of the resource. Always compute#reservations for reservations.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// [Output Only] Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// [Output Only] Server-defined fully-qualified URL for this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// Share-settings for shared-reservation
        /// </summary>
        public readonly Outputs.ShareSettingsResponse ShareSettings;
        /// <summary>
        /// Reservation for instances with specific machine shapes.
        /// </summary>
        public readonly Outputs.AllocationSpecificSKUReservationResponse SpecificReservation;
        /// <summary>
        /// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
        /// </summary>
        public readonly bool SpecificReservationRequired;
        /// <summary>
        /// [Output Only] The status of the reservation.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private ReservationResponse(
            string commitment,

            string creationTimestamp,

            string description,

            string kind,

            string name,

            bool satisfiesPzs,

            string selfLink,

            string selfLinkWithId,

            Outputs.ShareSettingsResponse shareSettings,

            Outputs.AllocationSpecificSKUReservationResponse specificReservation,

            bool specificReservationRequired,

            string status,

            string zone)
        {
            Commitment = commitment;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            Name = name;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            ShareSettings = shareSettings;
            SpecificReservation = specificReservation;
            SpecificReservationRequired = specificReservationRequired;
            Status = status;
            Zone = zone;
        }
    }
}
