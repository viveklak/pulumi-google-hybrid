// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Outputs
{

    /// <summary>
    /// Provides data to pass through to the worker harness.
    /// </summary>
    [OutputType]
    public sealed class WorkerSettingsResponse
    {
        /// <summary>
        /// The base URL for accessing Google Cloud APIs. When workers access Google Cloud APIs, they logically do so via relative URLs. If this field is specified, it supplies the base URL to use for resolving these relative URLs. The normative algorithm used is defined by RFC 1808, "Relative Uniform Resource Locators". If not specified, the default value is "http://www.googleapis.com/"
        /// </summary>
        public readonly string BaseUrl;
        /// <summary>
        /// Whether to send work progress updates to the service.
        /// </summary>
        public readonly bool ReportingEnabled;
        /// <summary>
        /// The Cloud Dataflow service path relative to the root URL, for example, "dataflow/v1b3/projects".
        /// </summary>
        public readonly string ServicePath;
        /// <summary>
        /// The Shuffle service path relative to the root URL, for example, "shuffle/v1beta1".
        /// </summary>
        public readonly string ShuffleServicePath;
        /// <summary>
        /// The prefix of the resources the system should use for temporary storage. The supported resource type is: Google Cloud Storage: storage.googleapis.com/{bucket}/{object} bucket.storage.googleapis.com/{object}
        /// </summary>
        public readonly string TempStoragePrefix;
        /// <summary>
        /// The ID of the worker running this pipeline.
        /// </summary>
        public readonly string WorkerId;

        [OutputConstructor]
        private WorkerSettingsResponse(
            string baseUrl,

            bool reportingEnabled,

            string servicePath,

            string shuffleServicePath,

            string tempStoragePrefix,

            string workerId)
        {
            BaseUrl = baseUrl;
            ReportingEnabled = reportingEnabled;
            ServicePath = servicePath;
            ShuffleServicePath = shuffleServicePath;
            TempStoragePrefix = tempStoragePrefix;
            WorkerId = workerId;
        }
    }
}
