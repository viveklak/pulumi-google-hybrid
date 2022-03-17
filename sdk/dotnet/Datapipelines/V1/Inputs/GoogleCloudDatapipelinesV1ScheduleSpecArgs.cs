// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datapipelines.V1.Inputs
{

    /// <summary>
    /// Details of the schedule the pipeline runs on.
    /// </summary>
    public sealed class GoogleCloudDatapipelinesV1ScheduleSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unix-cron format of the schedule. This information is retrieved from the linked Cloud Scheduler.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Timezone ID. This matches the timezone IDs used by the Cloud Scheduler API. If empty, UTC time is assumed.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public GoogleCloudDatapipelinesV1ScheduleSpecArgs()
        {
        }
    }
}
