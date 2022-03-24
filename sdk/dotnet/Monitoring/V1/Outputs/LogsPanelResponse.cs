// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Outputs
{

    /// <summary>
    /// A widget that displays a stream of log.
    /// </summary>
    [OutputType]
    public sealed class LogsPanelResponse
    {
        /// <summary>
        /// A filter that chooses which log entries to return. See Advanced Logs Queries (https://cloud.google.com/logging/docs/view/advanced-queries). Only log entries that match the filter are returned. An empty filter matches all log entries.
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// The names of logging resources to collect logs for. Currently only projects are supported. If empty, the widget will default to the host project.
        /// </summary>
        public readonly ImmutableArray<string> ResourceNames;

        [OutputConstructor]
        private LogsPanelResponse(
            string filter,

            ImmutableArray<string> resourceNames)
        {
            Filter = filter;
            ResourceNames = resourceNames;
        }
    }
}