// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1Alpha1.Inputs
{

    /// <summary>
    /// A single application container. This specifies both the container to run, the command to run in the container and the arguments to supply to it. Note that additional arguments may be supplied by the system to the container at runtime.
    /// </summary>
    public sealed class ContainerArgs : Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// (Optional) Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("command")]
        private InputList<string>? _command;
        public InputList<string> Command
        {
            get => _command ?? (_command = new InputList<string>());
            set => _command = value;
        }

        [Input("env")]
        private InputList<Inputs.EnvVarArgs>? _env;

        /// <summary>
        /// (Optional) List of environment variables to set in the container.
        /// </summary>
        public InputList<Inputs.EnvVarArgs> Env
        {
            get => _env ?? (_env = new InputList<Inputs.EnvVarArgs>());
            set => _env = value;
        }

        [Input("envFrom")]
        private InputList<Inputs.EnvFromSourceArgs>? _envFrom;

        /// <summary>
        /// (Optional) List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
        /// </summary>
        public InputList<Inputs.EnvFromSourceArgs> EnvFrom
        {
            get => _envFrom ?? (_envFrom = new InputList<Inputs.EnvFromSourceArgs>());
            set => _envFrom = value;
        }

        /// <summary>
        /// Only supports containers from Google Container Registry or Artifact Registry URL of the Container image. More info: https://kubernetes.io/docs/concepts/containers/images
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// (Optional) Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
        /// </summary>
        [Input("imagePullPolicy")]
        public Input<string>? ImagePullPolicy { get; set; }

        /// <summary>
        /// (Optional) Periodic probe of container liveness. Container will be restarted if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
        /// </summary>
        [Input("livenessProbe")]
        public Input<Inputs.ProbeArgs>? LivenessProbe { get; set; }

        /// <summary>
        /// (Optional) Name of the container specified as a DNS_LABEL. Currently unused in Cloud Run. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ports")]
        private InputList<Inputs.ContainerPortArgs>? _ports;

        /// <summary>
        /// (Optional) List of ports to expose from the container. Only a single port can be specified. The specified ports must be listening on all interfaces (0.0.0.0) within the container to be accessible. If omitted, a port number will be chosen and passed to the container through the PORT environment variable for the container to listen on.
        /// </summary>
        public InputList<Inputs.ContainerPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ContainerPortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// (Optional) Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
        /// </summary>
        [Input("readinessProbe")]
        public Input<Inputs.ProbeArgs>? ReadinessProbe { get; set; }

        /// <summary>
        /// (Optional) Compute Resources required by this container. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
        /// </summary>
        [Input("resources")]
        public Input<Inputs.ResourceRequirementsArgs>? Resources { get; set; }

        /// <summary>
        /// (Optional) Security options the pod should run with. More info: https://kubernetes.io/docs/concepts/policy/security-context/ More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
        /// </summary>
        [Input("securityContext")]
        public Input<Inputs.SecurityContextArgs>? SecurityContext { get; set; }

        /// <summary>
        /// (Optional) Startup probe of application within the container. All other probes are disabled if a startup probe is provided, until it succeeds. Container will not be added to service endpoints if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
        /// </summary>
        [Input("startupProbe")]
        public Input<Inputs.ProbeArgs>? StartupProbe { get; set; }

        /// <summary>
        /// (Optional) Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log.
        /// </summary>
        [Input("terminationMessagePath")]
        public Input<string>? TerminationMessagePath { get; set; }

        /// <summary>
        /// (Optional) Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
        /// </summary>
        [Input("terminationMessagePolicy")]
        public Input<string>? TerminationMessagePolicy { get; set; }

        [Input("volumeMounts")]
        private InputList<Inputs.VolumeMountArgs>? _volumeMounts;

        /// <summary>
        /// (Optional) Volume to mount into the container's filesystem. Only supports SecretVolumeSources. Pod volumes to mount into the container's filesystem.
        /// </summary>
        public InputList<Inputs.VolumeMountArgs> VolumeMounts
        {
            get => _volumeMounts ?? (_volumeMounts = new InputList<Inputs.VolumeMountArgs>());
            set => _volumeMounts = value;
        }

        /// <summary>
        /// (Optional) Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public ContainerArgs()
        {
        }
    }
}
