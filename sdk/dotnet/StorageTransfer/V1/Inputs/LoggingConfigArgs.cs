// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Inputs
{

    /// <summary>
    /// Specifies the logging behavior for transfer operations. For cloud-to-cloud transfers, logs are sent to Cloud Logging. See [Read transfer logs](https://cloud.google.com/storage-transfer/docs/read-transfer-logs) for details. For transfers to or from a POSIX file system, logs are stored in the Cloud Storage bucket that is the source or sink of the transfer. See [Managing Transfer for on-premises jobs] (https://cloud.google.com/storage-transfer/docs/managing-on-prem-jobs#viewing-logs) for details.
    /// </summary>
    public sealed class LoggingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For transfers with a PosixFilesystem source, this option enables the Cloud Storage transfer logs for this transfer.
        /// </summary>
        [Input("enableOnpremGcsTransferLogs")]
        public Input<bool>? EnableOnpremGcsTransferLogs { get; set; }

        [Input("logActionStates")]
        private InputList<Pulumi.GoogleNative.StorageTransfer.V1.LoggingConfigLogActionStatesItem>? _logActionStates;

        /// <summary>
        /// States in which `log_actions` are logged. If empty, no logs are generated. Not supported for transfers with PosixFilesystem data sources; use enable_onprem_gcs_transfer_logs instead.
        /// </summary>
        public InputList<Pulumi.GoogleNative.StorageTransfer.V1.LoggingConfigLogActionStatesItem> LogActionStates
        {
            get => _logActionStates ?? (_logActionStates = new InputList<Pulumi.GoogleNative.StorageTransfer.V1.LoggingConfigLogActionStatesItem>());
            set => _logActionStates = value;
        }

        [Input("logActions")]
        private InputList<Pulumi.GoogleNative.StorageTransfer.V1.LoggingConfigLogActionsItem>? _logActions;

        /// <summary>
        /// Specifies the actions to be logged. If empty, no logs are generated. Not supported for transfers with PosixFilesystem data sources; use enable_onprem_gcs_transfer_logs instead.
        /// </summary>
        public InputList<Pulumi.GoogleNative.StorageTransfer.V1.LoggingConfigLogActionsItem> LogActions
        {
            get => _logActions ?? (_logActions = new InputList<Pulumi.GoogleNative.StorageTransfer.V1.LoggingConfigLogActionsItem>());
            set => _logActions = value;
        }

        public LoggingConfigArgs()
        {
        }
    }
}
