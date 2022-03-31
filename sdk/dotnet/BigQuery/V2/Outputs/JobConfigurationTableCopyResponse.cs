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
    public sealed class JobConfigurationTableCopyResponse
    {
        /// <summary>
        /// [Optional] Specifies whether the job is allowed to create new tables. The following values are supported: CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table. CREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result. The default value is CREATE_IF_NEEDED. Creation, truncation and append actions occur as one atomic update upon job completion.
        /// </summary>
        public readonly string CreateDisposition;
        /// <summary>
        /// Custom encryption configuration (e.g., Cloud KMS keys).
        /// </summary>
        public readonly Outputs.EncryptionConfigurationResponse DestinationEncryptionConfiguration;
        /// <summary>
        /// [Optional] The time when the destination table expires. Expired tables will be deleted and their storage reclaimed.
        /// </summary>
        public readonly object DestinationExpirationTime;
        /// <summary>
        /// [Required] The destination table
        /// </summary>
        public readonly Outputs.TableReferenceResponse DestinationTable;
        /// <summary>
        /// [Optional] Supported operation types in table copy job.
        /// </summary>
        public readonly string OperationType;
        /// <summary>
        /// [Pick one] Source table to copy.
        /// </summary>
        public readonly Outputs.TableReferenceResponse SourceTable;
        /// <summary>
        /// [Pick one] Source tables to copy.
        /// </summary>
        public readonly ImmutableArray<Outputs.TableReferenceResponse> SourceTables;
        /// <summary>
        /// [Optional] Specifies the action that occurs if the destination table already exists. The following values are supported: WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data. WRITE_APPEND: If the table already exists, BigQuery appends the data to the table. WRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result. The default value is WRITE_EMPTY. Each action is atomic and only occurs if BigQuery is able to complete the job successfully. Creation, truncation and append actions occur as one atomic update upon job completion.
        /// </summary>
        public readonly string WriteDisposition;

        [OutputConstructor]
        private JobConfigurationTableCopyResponse(
            string createDisposition,

            Outputs.EncryptionConfigurationResponse destinationEncryptionConfiguration,

            object destinationExpirationTime,

            Outputs.TableReferenceResponse destinationTable,

            string operationType,

            Outputs.TableReferenceResponse sourceTable,

            ImmutableArray<Outputs.TableReferenceResponse> sourceTables,

            string writeDisposition)
        {
            CreateDisposition = createDisposition;
            DestinationEncryptionConfiguration = destinationEncryptionConfiguration;
            DestinationExpirationTime = destinationExpirationTime;
            DestinationTable = destinationTable;
            OperationType = operationType;
            SourceTable = sourceTable;
            SourceTables = sourceTables;
            WriteDisposition = writeDisposition;
        }
    }
}
