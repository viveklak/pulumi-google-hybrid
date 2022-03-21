// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Node kubelet configs.
    /// </summary>
    public sealed class NodeKubeletConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable CPU CFS quota enforcement for containers that specify CPU limits. This option is enabled by default which makes kubelet use CFS quota (https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt) to enforce container CPU limits. Otherwise, CPU limits will not be enforced at all. Disable this option to mitigate CPU throttling problems while still having your pods to be in Guaranteed QoS class by specifying the CPU limits. The default value is 'true' if unspecified.
        /// </summary>
        [Input("cpuCfsQuota")]
        public Input<bool>? CpuCfsQuota { get; set; }

        /// <summary>
        /// Set the CPU CFS quota period value 'cpu.cfs_period_us'. The string must be a sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". The value must be a positive duration.
        /// </summary>
        [Input("cpuCfsQuotaPeriod")]
        public Input<string>? CpuCfsQuotaPeriod { get; set; }

        /// <summary>
        /// Control the CPU management policy on the node. See https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/ The following values are allowed. * "none": the default, which represents the existing scheduling behavior. * "static": allows pods with certain resource characteristics to be granted increased CPU affinity and exclusivity on the node. The default value is 'none' if unspecified.
        /// </summary>
        [Input("cpuManagerPolicy")]
        public Input<string>? CpuManagerPolicy { get; set; }

        public NodeKubeletConfigArgs()
        {
        }
    }
}
