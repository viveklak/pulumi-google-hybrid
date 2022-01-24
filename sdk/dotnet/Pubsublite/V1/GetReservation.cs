// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Pubsublite.V1
{
    public static class GetReservation
    {
        /// <summary>
        /// Returns the reservation configuration.
        /// </summary>
        public static Task<GetReservationResult> InvokeAsync(GetReservationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReservationResult>("google-native:pubsublite/v1:getReservation", args ?? new GetReservationArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the reservation configuration.
        /// </summary>
        public static Output<GetReservationResult> Invoke(GetReservationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReservationResult>("google-native:pubsublite/v1:getReservation", args ?? new GetReservationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReservationArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("reservationId", required: true)]
        public string ReservationId { get; set; } = null!;

        public GetReservationArgs()
        {
        }
    }

    public sealed class GetReservationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("reservationId", required: true)]
        public Input<string> ReservationId { get; set; } = null!;

        public GetReservationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReservationResult
    {
        /// <summary>
        /// The name of the reservation. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The reserved throughput capacity. Every unit of throughput capacity is equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed messages. Any topics which are declared as using capacity from a Reservation will consume resources from this reservation instead of being charged individually.
        /// </summary>
        public readonly string ThroughputCapacity;

        [OutputConstructor]
        private GetReservationResult(
            string name,

            string throughputCapacity)
        {
            Name = name;
            ThroughputCapacity = throughputCapacity;
        }
    }
}
