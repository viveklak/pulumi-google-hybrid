// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetReport
    {
        /// <summary>
        /// Retrieve a custom report definition.
        /// </summary>
        public static Task<GetReportResult> InvokeAsync(GetReportArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReportResult>("google-native:apigee/v1:getReport", args ?? new GetReportArgs(), options.WithVersion());
    }


    public sealed class GetReportArgs : Pulumi.InvokeArgs
    {
        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("reportId", required: true)]
        public string ReportId { get; set; } = null!;

        public GetReportArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReportResult
    {
        /// <summary>
        /// This field contains the chart type for the report
        /// </summary>
        public readonly string ChartType;
        /// <summary>
        /// Legacy field: not used. This field contains a list of comments associated with custom report
        /// </summary>
        public readonly ImmutableArray<string> Comments;
        /// <summary>
        /// Unix time when the app was created json key: createdAt
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// This contains the list of dimensions for the report
        /// </summary>
        public readonly ImmutableArray<string> Dimensions;
        /// <summary>
        /// This is the display name for the report
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Environment name
        /// </summary>
        public readonly string Environment;
        /// <summary>
        /// This field contains the filter expression
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// Legacy field: not used. Contains the from time for the report
        /// </summary>
        public readonly string FromTime;
        /// <summary>
        /// Modified time of this entity as milliseconds since epoch. json key: lastModifiedAt
        /// </summary>
        public readonly string LastModifiedAt;
        /// <summary>
        /// Last viewed time of this entity as milliseconds since epoch
        /// </summary>
        public readonly string LastViewedAt;
        /// <summary>
        /// Legacy field: not used This field contains the limit for the result retrieved
        /// </summary>
        public readonly string Limit;
        /// <summary>
        /// This contains the list of metrics
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudApigeeV1CustomReportMetricResponse> Metrics;
        /// <summary>
        /// Unique identifier for the report T his is a legacy field used to encode custom report unique id
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Legacy field: not used. This field contains the offset for the data
        /// </summary>
        public readonly string Offset;
        /// <summary>
        /// Organization name
        /// </summary>
        public readonly string Organization;
        /// <summary>
        /// This field contains report properties such as ui metadata etc.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudApigeeV1ReportPropertyResponse> Properties;
        /// <summary>
        /// Legacy field: not used much. Contains the list of sort by columns
        /// </summary>
        public readonly ImmutableArray<string> SortByCols;
        /// <summary>
        /// Legacy field: not used much. Contains the sort order for the sort columns
        /// </summary>
        public readonly string SortOrder;
        /// <summary>
        /// Legacy field: not used. This field contains a list of tags associated with custom report
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// This field contains the time unit of aggregation for the report
        /// </summary>
        public readonly string TimeUnit;
        /// <summary>
        /// Legacy field: not used. Contains the end time for the report
        /// </summary>
        public readonly string ToTime;
        /// <summary>
        /// Legacy field: not used. This field contains the top k parameter value for restricting the result
        /// </summary>
        public readonly string Topk;

        [OutputConstructor]
        private GetReportResult(
            string chartType,

            ImmutableArray<string> comments,

            string createdAt,

            ImmutableArray<string> dimensions,

            string displayName,

            string environment,

            string filter,

            string fromTime,

            string lastModifiedAt,

            string lastViewedAt,

            string limit,

            ImmutableArray<Outputs.GoogleCloudApigeeV1CustomReportMetricResponse> metrics,

            string name,

            string offset,

            string organization,

            ImmutableArray<Outputs.GoogleCloudApigeeV1ReportPropertyResponse> properties,

            ImmutableArray<string> sortByCols,

            string sortOrder,

            ImmutableArray<string> tags,

            string timeUnit,

            string toTime,

            string topk)
        {
            ChartType = chartType;
            Comments = comments;
            CreatedAt = createdAt;
            Dimensions = dimensions;
            DisplayName = displayName;
            Environment = environment;
            Filter = filter;
            FromTime = fromTime;
            LastModifiedAt = lastModifiedAt;
            LastViewedAt = lastViewedAt;
            Limit = limit;
            Metrics = metrics;
            Name = name;
            Offset = offset;
            Organization = organization;
            Properties = properties;
            SortByCols = sortByCols;
            SortOrder = sortOrder;
            Tags = tags;
            TimeUnit = timeUnit;
            ToTime = toTime;
            Topk = topk;
        }
    }
}
