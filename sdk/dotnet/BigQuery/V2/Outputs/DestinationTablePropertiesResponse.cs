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
    public sealed class DestinationTablePropertiesResponse
    {
        /// <summary>
        /// [Optional] The description for the destination table. This will only be used if the destination table is newly created. If the table already exists and a value different than the current description is provided, the job will fail.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// [Internal] This field is for Google internal use only.
        /// </summary>
        public readonly string ExpirationTime;
        /// <summary>
        /// [Optional] The friendly name for the destination table. This will only be used if the destination table is newly created. If the table already exists and a value different than the current friendly name is provided, the job will fail.
        /// </summary>
        public readonly string FriendlyName;
        /// <summary>
        /// [Optional] The labels associated with this table. You can use these to organize and group your tables. This will only be used if the destination table is newly created. If the table already exists and labels are different than the current labels are provided, the job will fail.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;

        [OutputConstructor]
        private DestinationTablePropertiesResponse(
            string description,

            string expirationTime,

            string friendlyName,

            ImmutableDictionary<string, string> labels)
        {
            Description = description;
            ExpirationTime = expirationTime;
            FriendlyName = friendlyName;
            Labels = labels;
        }
    }
}
