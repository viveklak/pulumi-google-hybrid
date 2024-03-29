// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1Alpha1.Outputs
{

    /// <summary>
    /// InstanceSpec is a description of an instance.
    /// </summary>
    [OutputType]
    public sealed class InstanceSpecResponse
    {
        /// <summary>
        /// Optional. Optional duration in seconds the instance may be active relative to StartTime before the system will actively try to mark it failed and kill associated containers. If set to zero, the system will never attempt to kill an instance based on time. Otherwise, value must be a positive integer. +optional
        /// </summary>
        public readonly string ActiveDeadlineSeconds;
        /// <summary>
        /// Optional. List of containers belonging to the instance. We disallow a number of fields on this Container. Only a single container may be provided.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerResponse> Containers;
        /// <summary>
        /// Optional. Restart policy for all containers within the instance. Allowed values are: - OnFailure: Instances will always be restarted on failure if the backoffLimit has not been reached. - Never: Instances are never restarted and all failures are permanent. Cannot be used if backoffLimit is set. +optional
        /// </summary>
        public readonly string RestartPolicy;
        /// <summary>
        /// Optional. Email address of the IAM service account associated with the instance of a Job. The service account represents the identity of the running instance, and determines what permissions the instance has. If not provided, the instance will use the project's default service account. +optional
        /// </summary>
        public readonly string ServiceAccountName;
        /// <summary>
        /// Optional. Optional duration in seconds the instance needs to terminate gracefully. Value must be non-negative integer. The value zero indicates delete immediately. The grace period is the duration in seconds after the processes running in the instance are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. +optional
        /// </summary>
        public readonly string TerminationGracePeriodSeconds;
        /// <summary>
        /// Optional. List of volumes that can be mounted by containers belonging to the instance. More info: https://kubernetes.io/docs/concepts/storage/volumes +optional
        /// </summary>
        public readonly ImmutableArray<Outputs.VolumeResponse> Volumes;

        [OutputConstructor]
        private InstanceSpecResponse(
            string activeDeadlineSeconds,

            ImmutableArray<Outputs.ContainerResponse> containers,

            string restartPolicy,

            string serviceAccountName,

            string terminationGracePeriodSeconds,

            ImmutableArray<Outputs.VolumeResponse> volumes)
        {
            ActiveDeadlineSeconds = activeDeadlineSeconds;
            Containers = containers;
            RestartPolicy = restartPolicy;
            ServiceAccountName = serviceAccountName;
            TerminationGracePeriodSeconds = terminationGracePeriodSeconds;
            Volumes = volumes;
        }
    }
}
