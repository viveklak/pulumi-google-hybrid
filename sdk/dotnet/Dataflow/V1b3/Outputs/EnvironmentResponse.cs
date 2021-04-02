// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dataflow.V1b3.Outputs
{

    [OutputType]
    public sealed class EnvironmentResponse
    {
        /// <summary>
        /// The type of cluster manager API to use. If unknown or unspecified, the service will attempt to choose a reasonable default. This should be in the form of the API service name, e.g. "compute.googleapis.com".
        /// </summary>
        public readonly string ClusterManagerApiService;
        /// <summary>
        /// The dataset for the current project where various workflow related tables are stored. The supported resource type is: Google BigQuery: bigquery.googleapis.com/{dataset}
        /// </summary>
        public readonly string Dataset;
        /// <summary>
        /// Any debugging options to be supplied to the job.
        /// </summary>
        public readonly Outputs.DebugOptionsResponse DebugOptions;
        /// <summary>
        /// The list of experiments to enable. This field should be used for SDK related experiments and not for service related experiments. The proper field for service related experiments is service_options. For more details see the rationale at go/user-specified-service-options.
        /// </summary>
        public readonly ImmutableArray<string> Experiments;
        /// <summary>
        /// Which Flexible Resource Scheduling mode to run in.
        /// </summary>
        public readonly string FlexResourceSchedulingGoal;
        /// <summary>
        /// Experimental settings.
        /// </summary>
        public readonly ImmutableDictionary<string, string> InternalExperiments;
        /// <summary>
        /// The Cloud Dataflow SDK pipeline options specified by the user. These options are passed through the service and are used to recreate the SDK pipeline options on the worker in a language agnostic and platform independent way.
        /// </summary>
        public readonly ImmutableDictionary<string, string> SdkPipelineOptions;
        /// <summary>
        /// Identity to run virtual machines as. Defaults to the default account.
        /// </summary>
        public readonly string ServiceAccountEmail;
        /// <summary>
        /// If set, contains the Cloud KMS key identifier used to encrypt data at rest, AKA a Customer Managed Encryption Key (CMEK). Format: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY
        /// </summary>
        public readonly string ServiceKmsKeyName;
        /// <summary>
        /// The list of service options to enable. This field should be used for service related experiments only. These experiments, when graduating to GA, should be replaced by dedicated fields or become default (i.e. always on). For more details see the rationale at go/user-specified-service-options.
        /// </summary>
        public readonly ImmutableArray<string> ServiceOptions;
        /// <summary>
        /// The shuffle mode used for the job.
        /// </summary>
        public readonly string ShuffleMode;
        /// <summary>
        /// The prefix of the resources the system should use for temporary storage. The system will append the suffix "/temp-{JOBNAME} to this resource prefix, where {JOBNAME} is the value of the job_name field. The resulting bucket and object prefix is used as the prefix of the resources used to store temporary data needed during the job execution. NOTE: This will override the value in taskrunner_settings. The supported resource type is: Google Cloud Storage: storage.googleapis.com/{bucket}/{object} bucket.storage.googleapis.com/{object}
        /// </summary>
        public readonly string TempStoragePrefix;
        /// <summary>
        /// A description of the process that generated the request.
        /// </summary>
        public readonly ImmutableDictionary<string, string> UserAgent;
        /// <summary>
        /// A structure describing which components and their versions of the service are required in order to run the job.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Version;
        /// <summary>
        /// The worker pools. At least one "harness" worker pool must be specified in order for the job to have workers.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkerPoolResponse> WorkerPools;
        /// <summary>
        /// The Compute Engine region (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1". Mutually exclusive with worker_zone. If neither worker_region nor worker_zone is specified, default to the control plane's region.
        /// </summary>
        public readonly string WorkerRegion;
        /// <summary>
        /// The Compute Engine zone (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1-a". Mutually exclusive with worker_region. If neither worker_region nor worker_zone is specified, a zone in the control plane's region is chosen based on available capacity.
        /// </summary>
        public readonly string WorkerZone;

        [OutputConstructor]
        private EnvironmentResponse(
            string clusterManagerApiService,

            string dataset,

            Outputs.DebugOptionsResponse debugOptions,

            ImmutableArray<string> experiments,

            string flexResourceSchedulingGoal,

            ImmutableDictionary<string, string> internalExperiments,

            ImmutableDictionary<string, string> sdkPipelineOptions,

            string serviceAccountEmail,

            string serviceKmsKeyName,

            ImmutableArray<string> serviceOptions,

            string shuffleMode,

            string tempStoragePrefix,

            ImmutableDictionary<string, string> userAgent,

            ImmutableDictionary<string, string> version,

            ImmutableArray<Outputs.WorkerPoolResponse> workerPools,

            string workerRegion,

            string workerZone)
        {
            ClusterManagerApiService = clusterManagerApiService;
            Dataset = dataset;
            DebugOptions = debugOptions;
            Experiments = experiments;
            FlexResourceSchedulingGoal = flexResourceSchedulingGoal;
            InternalExperiments = internalExperiments;
            SdkPipelineOptions = sdkPipelineOptions;
            ServiceAccountEmail = serviceAccountEmail;
            ServiceKmsKeyName = serviceKmsKeyName;
            ServiceOptions = serviceOptions;
            ShuffleMode = shuffleMode;
            TempStoragePrefix = tempStoragePrefix;
            UserAgent = userAgent;
            Version = version;
            WorkerPools = workerPools;
            WorkerRegion = workerRegion;
            WorkerZone = workerZone;
        }
    }
}