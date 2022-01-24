// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetFutureReservation
    {
        /// <summary>
        /// Retrieves information about the specified future reservation.
        /// </summary>
        public static Task<GetFutureReservationResult> InvokeAsync(GetFutureReservationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFutureReservationResult>("google-native:compute/alpha:getFutureReservation", args ?? new GetFutureReservationArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about the specified future reservation.
        /// </summary>
        public static Output<GetFutureReservationResult> Invoke(GetFutureReservationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFutureReservationResult>("google-native:compute/alpha:getFutureReservation", args ?? new GetFutureReservationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFutureReservationArgs : Pulumi.InvokeArgs
    {
        [Input("futureReservation", required: true)]
        public string FutureReservation { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetFutureReservationArgs()
        {
        }
    }

    public sealed class GetFutureReservationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("futureReservation", required: true)]
        public Input<string> FutureReservation { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetFutureReservationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFutureReservationResult
    {
        /// <summary>
        /// The creation timestamp for this future reservation in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the future reservation.
        /// </summary>
        public readonly string Description;
        public readonly string Kind;
        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
        /// </summary>
        public readonly string NamePrefix;
        /// <summary>
        /// Server-defined fully-qualified URL for this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// List of Projects/Folders to share with.
        /// </summary>
        public readonly Outputs.ShareSettingsResponse ShareSettings;
        /// <summary>
        /// Future Reservation configuration to indicate instance properties and total count.
        /// </summary>
        public readonly Outputs.FutureReservationSpecificSKUPropertiesResponse SpecificSkuProperties;
        /// <summary>
        /// [Output only] Status of the Future Reservation
        /// </summary>
        public readonly Outputs.FutureReservationStatusResponse Status;
        /// <summary>
        /// Time window for this Future Reservation.
        /// </summary>
        public readonly Outputs.FutureReservationTimeWindowResponse TimeWindow;
        /// <summary>
        /// URL of the Zone where this future reservation resides.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetFutureReservationResult(
            string creationTimestamp,

            string description,

            string kind,

            string name,

            string namePrefix,

            string selfLink,

            string selfLinkWithId,

            Outputs.ShareSettingsResponse shareSettings,

            Outputs.FutureReservationSpecificSKUPropertiesResponse specificSkuProperties,

            Outputs.FutureReservationStatusResponse status,

            Outputs.FutureReservationTimeWindowResponse timeWindow,

            string zone)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            Name = name;
            NamePrefix = namePrefix;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            ShareSettings = shareSettings;
            SpecificSkuProperties = specificSkuProperties;
            Status = status;
            TimeWindow = timeWindow;
            Zone = zone;
        }
    }
}
