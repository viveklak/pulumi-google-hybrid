// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Outputs
{

    /// <summary>
    /// Logging configuration of the service. The following example shows how to configure logs to be sent to the producer and consumer projects. In the example, the `activity_history` log is sent to both the producer and consumer projects, whereas the `purchase_history` log is only sent to the producer project. monitored_resources: - type: library.googleapis.com/branch labels: - key: /city description: The city where the library branch is located in. - key: /name description: The name of the branch. logs: - name: activity_history labels: - key: /customer_id - name: purchase_history logging: producer_destinations: - monitored_resource: library.googleapis.com/branch logs: - activity_history - purchase_history consumer_destinations: - monitored_resource: library.googleapis.com/branch logs: - activity_history
    /// </summary>
    [OutputType]
    public sealed class LoggingResponse
    {
        /// <summary>
        /// Logging configurations for sending logs to the consumer project. There can be multiple consumer destinations, each one must have a different monitored resource type. A log can be used in at most one consumer destination.
        /// </summary>
        public readonly ImmutableArray<Outputs.LoggingDestinationResponse> ConsumerDestinations;
        /// <summary>
        /// Logging configurations for sending logs to the producer project. There can be multiple producer destinations, each one must have a different monitored resource type. A log can be used in at most one producer destination.
        /// </summary>
        public readonly ImmutableArray<Outputs.LoggingDestinationResponse> ProducerDestinations;

        [OutputConstructor]
        private LoggingResponse(
            ImmutableArray<Outputs.LoggingDestinationResponse> consumerDestinations,

            ImmutableArray<Outputs.LoggingDestinationResponse> producerDestinations)
        {
            ConsumerDestinations = consumerDestinations;
            ProducerDestinations = producerDestinations;
        }
    }
}
