// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class JobConfigurationResponse
    {
        /// <summary>
        /// [Pick one] Copies a table.
        /// </summary>
        public readonly Outputs.JobConfigurationTableCopyResponse Copy;
        /// <summary>
        /// [Optional] If set, don't actually run this job. A valid query will return a mostly empty response with some processing statistics, while an invalid query will return the same error it would if it wasn't a dry run. Behavior of non-query jobs is undefined.
        /// </summary>
        public readonly bool DryRun;
        /// <summary>
        /// [Pick one] Configures an extract job.
        /// </summary>
        public readonly Outputs.JobConfigurationExtractResponse Extract;
        /// <summary>
        /// [Optional] Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
        /// </summary>
        public readonly string JobTimeoutMs;
        /// <summary>
        /// The type of the job. Can be QUERY, LOAD, EXTRACT, COPY or UNKNOWN.
        /// </summary>
        public readonly string JobType;
        /// <summary>
        /// The labels associated with this job. You can use these to organize and group your jobs. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// [Pick one] Configures a load job.
        /// </summary>
        public readonly Outputs.JobConfigurationLoadResponse Load;
        /// <summary>
        /// [Pick one] Configures a query job.
        /// </summary>
        public readonly Outputs.JobConfigurationQueryResponse Query;

        [OutputConstructor]
        private JobConfigurationResponse(
            Outputs.JobConfigurationTableCopyResponse copy,

            bool dryRun,

            Outputs.JobConfigurationExtractResponse extract,

            string jobTimeoutMs,

            string jobType,

            ImmutableDictionary<string, string> labels,

            Outputs.JobConfigurationLoadResponse load,

            Outputs.JobConfigurationQueryResponse query)
        {
            Copy = copy;
            DryRun = dryRun;
            Extract = extract;
            JobTimeoutMs = jobTimeoutMs;
            JobType = jobType;
            Labels = labels;
            Load = load;
            Query = query;
        }
    }
}
