// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetDataCollector
    {
        /// <summary>
        /// Gets a data collector.
        /// </summary>
        public static Task<GetDataCollectorResult> InvokeAsync(GetDataCollectorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataCollectorResult>("google-native:apigee/v1:getDataCollector", args ?? new GetDataCollectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a data collector.
        /// </summary>
        public static Output<GetDataCollectorResult> Invoke(GetDataCollectorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataCollectorResult>("google-native:apigee/v1:getDataCollector", args ?? new GetDataCollectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataCollectorArgs : Pulumi.InvokeArgs
    {
        [Input("datacollectorId", required: true)]
        public string DatacollectorId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetDataCollectorArgs()
        {
        }
    }

    public sealed class GetDataCollectorInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("datacollectorId", required: true)]
        public Input<string> DatacollectorId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetDataCollectorInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataCollectorResult
    {
        /// <summary>
        /// The time at which the data collector was created in milliseconds since the epoch.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A description of the data collector.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The time at which the Data Collector was last updated in milliseconds since the epoch.
        /// </summary>
        public readonly string LastModifiedAt;
        /// <summary>
        /// ID of the data collector. Must begin with `dc_`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The type of data this data collector will collect.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDataCollectorResult(
            string createdAt,

            string description,

            string lastModifiedAt,

            string name,

            string type)
        {
            CreatedAt = createdAt;
            Description = description;
            LastModifiedAt = lastModifiedAt;
            Name = name;
            Type = type;
        }
    }
}
