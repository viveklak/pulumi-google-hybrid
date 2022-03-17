// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Outputs
{

    /// <summary>
    /// A condition type that checks whether a log message in the scoping project (https://cloud.google.com/monitoring/api/v3#project_name) satisfies the given filter. Logs from other projects in the metrics scope are not evaluated.
    /// </summary>
    [OutputType]
    public sealed class LogMatchResponse
    {
        /// <summary>
        /// A logs-based filter. See Advanced Logs Queries (https://cloud.google.com/logging/docs/view/advanced-queries) for how this filter should be constructed.
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// Optional. A map from a label key to an extractor expression, which is used to extract the value for this label key. Each entry in this map is a specification for how data should be extracted from log entries that match filter. Each combination of extracted values is treated as a separate rule for the purposes of triggering notifications. Label keys and corresponding values can be used in notifications generated by this condition.Please see the documentation on logs-based metric valueExtractors (https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics#LogMetric.FIELDS.value_extractor) for syntax and examples.
        /// </summary>
        public readonly ImmutableDictionary<string, string> LabelExtractors;

        [OutputConstructor]
        private LogMatchResponse(
            string filter,

            ImmutableDictionary<string, string> labelExtractors)
        {
            Filter = filter;
            LabelExtractors = labelExtractors;
        }
    }
}
