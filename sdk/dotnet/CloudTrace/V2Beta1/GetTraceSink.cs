// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTrace.V2Beta1
{
    public static class GetTraceSink
    {
        /// <summary>
        /// Get a trace sink by name under the parent resource (GCP project).
        /// </summary>
        public static Task<GetTraceSinkResult> InvokeAsync(GetTraceSinkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTraceSinkResult>("google-native:cloudtrace/v2beta1:getTraceSink", args ?? new GetTraceSinkArgs(), options.WithVersion());
    }


    public sealed class GetTraceSinkArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("traceSinkId", required: true)]
        public string TraceSinkId { get; set; } = null!;

        public GetTraceSinkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTraceSinkResult
    {
        /// <summary>
        /// The canonical sink resource name, unique within the project. Must be of the form: project/[PROJECT_NUMBER]/traceSinks/[SINK_ID]. E.g.: `"projects/12345/traceSinks/my-project-trace-sink"`. Sink identifiers are limited to 256 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The export destination.
        /// </summary>
        public readonly Outputs.OutputConfigResponse OutputConfig;
        /// <summary>
        /// A service account name for exporting the data. This field is set by sinks.create and sinks.update. The service account will need to be granted write access to the destination specified in the output configuration, see [Granting access for a resource](/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource). To create tables and write data this account will need the dataEditor role. Read more about roles in the [BigQuery documentation](https://cloud.google.com/bigquery/docs/access-control). E.g.: "service-00000001@00000002.iam.gserviceaccount.com"
        /// </summary>
        public readonly string WriterIdentity;

        [OutputConstructor]
        private GetTraceSinkResult(
            string name,

            Outputs.OutputConfigResponse outputConfig,

            string writerIdentity)
        {
            Name = name;
            OutputConfig = outputConfig;
            WriterIdentity = writerIdentity;
        }
    }
}
