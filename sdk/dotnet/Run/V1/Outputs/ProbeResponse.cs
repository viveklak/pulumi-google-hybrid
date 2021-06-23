// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Outputs
{

    [OutputType]
    public sealed class ProbeResponse
    {
        /// <summary>
        /// (Optional) One and only one of the following should be specified. Exec specifies the action to take. A field inlined from the Handler message.
        /// </summary>
        public readonly Outputs.ExecActionResponse Exec;
        /// <summary>
        /// (Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
        /// </summary>
        public readonly int FailureThreshold;
        /// <summary>
        /// (Optional) HTTPGet specifies the http request to perform. A field inlined from the Handler message.
        /// </summary>
        public readonly Outputs.HTTPGetActionResponse HttpGet;
        /// <summary>
        /// (Optional) Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
        /// </summary>
        public readonly int InitialDelaySeconds;
        /// <summary>
        /// (Optional) How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
        /// </summary>
        public readonly int PeriodSeconds;
        /// <summary>
        /// (Optional) Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.
        /// </summary>
        public readonly int SuccessThreshold;
        /// <summary>
        /// (Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported A field inlined from the Handler message.
        /// </summary>
        public readonly Outputs.TCPSocketActionResponse TcpSocket;
        /// <summary>
        /// (Optional) Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
        /// </summary>
        public readonly int TimeoutSeconds;

        [OutputConstructor]
        private ProbeResponse(
            Outputs.ExecActionResponse exec,

            int failureThreshold,

            Outputs.HTTPGetActionResponse httpGet,

            int initialDelaySeconds,

            int periodSeconds,

            int successThreshold,

            Outputs.TCPSocketActionResponse tcpSocket,

            int timeoutSeconds)
        {
            Exec = exec;
            FailureThreshold = failureThreshold;
            HttpGet = httpGet;
            InitialDelaySeconds = initialDelaySeconds;
            PeriodSeconds = periodSeconds;
            SuccessThreshold = successThreshold;
            TcpSocket = tcpSocket;
            TimeoutSeconds = timeoutSeconds;
        }
    }
}
