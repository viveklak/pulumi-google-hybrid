// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// A Dataproc job for running Presto (https://prestosql.io/) queries. IMPORTANT: The Dataproc Presto Optional Component (https://cloud.google.com/dataproc/docs/concepts/components/presto) must be enabled when the cluster is created to submit a Presto job to the cluster.
    /// </summary>
    [OutputType]
    public sealed class PrestoJobResponse
    {
        /// <summary>
        /// Optional. Presto client tags to attach to this query
        /// </summary>
        public readonly ImmutableArray<string> ClientTags;
        /// <summary>
        /// Optional. Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries.
        /// </summary>
        public readonly bool ContinueOnFailure;
        /// <summary>
        /// Optional. The runtime log config for job execution.
        /// </summary>
        public readonly Outputs.LoggingConfigResponse LoggingConfig;
        /// <summary>
        /// Optional. The format in which query output will be displayed. See the Presto documentation for supported output formats
        /// </summary>
        public readonly string OutputFormat;
        /// <summary>
        /// Optional. A mapping of property names to values. Used to set Presto session properties (https://prestodb.io/docs/current/sql/set-session.html) Equivalent to using the --session flag in the Presto CLI
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        /// <summary>
        /// The HCFS URI of the script that contains SQL queries.
        /// </summary>
        public readonly string QueryFileUri;
        /// <summary>
        /// A list of queries.
        /// </summary>
        public readonly Outputs.QueryListResponse QueryList;

        [OutputConstructor]
        private PrestoJobResponse(
            ImmutableArray<string> clientTags,

            bool continueOnFailure,

            Outputs.LoggingConfigResponse loggingConfig,

            string outputFormat,

            ImmutableDictionary<string, string> properties,

            string queryFileUri,

            Outputs.QueryListResponse queryList)
        {
            ClientTags = clientTags;
            ContinueOnFailure = continueOnFailure;
            LoggingConfig = loggingConfig;
            OutputFormat = outputFormat;
            Properties = properties;
            QueryFileUri = queryFileUri;
            QueryList = queryList;
        }
    }
}
