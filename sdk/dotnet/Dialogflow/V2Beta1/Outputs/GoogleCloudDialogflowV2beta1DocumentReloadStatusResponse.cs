// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Outputs
{

    /// <summary>
    /// The status of a reload attempt.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1DocumentReloadStatusResponse
    {
        /// <summary>
        /// The status of a reload attempt or the initial load.
        /// </summary>
        public readonly Outputs.GoogleRpcStatusResponse Status;
        /// <summary>
        /// The time of a reload attempt. This reload may have been triggered automatically or manually and may not have succeeded.
        /// </summary>
        public readonly string Time;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1DocumentReloadStatusResponse(
            Outputs.GoogleRpcStatusResponse status,

            string time)
        {
            Status = status;
            Time = time;
        }
    }
}
