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
    public sealed class RowAccessPolicyReferenceResponse
    {
        /// <summary>
        /// [Required] The ID of the dataset containing this row access policy.
        /// </summary>
        public readonly string DatasetId;
        /// <summary>
        /// [Required] The ID of the row access policy. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
        /// </summary>
        public readonly string PolicyId;
        /// <summary>
        /// [Required] The ID of the project containing this row access policy.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// [Required] The ID of the table containing this row access policy.
        /// </summary>
        public readonly string TableId;

        [OutputConstructor]
        private RowAccessPolicyReferenceResponse(
            string datasetId,

            string policyId,

            string project,

            string tableId)
        {
            DatasetId = datasetId;
            PolicyId = policyId;
            Project = project;
            TableId = tableId;
        }
    }
}
