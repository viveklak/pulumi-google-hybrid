// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1Beta1.Inputs
{

    /// <summary>
    /// The Kubernetes workloads configuration for GKE cluster associated with the Cloud Composer environment. Supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
    /// </summary>
    public sealed class WorkloadsConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Resources used by Airflow schedulers.
        /// </summary>
        [Input("scheduler")]
        public Input<Inputs.SchedulerResourceArgs>? Scheduler { get; set; }

        /// <summary>
        /// Optional. Resources used by Airflow web server.
        /// </summary>
        [Input("webServer")]
        public Input<Inputs.WebServerResourceArgs>? WebServer { get; set; }

        /// <summary>
        /// Optional. Resources used by Airflow workers.
        /// </summary>
        [Input("worker")]
        public Input<Inputs.WorkerResourceArgs>? Worker { get; set; }

        public WorkloadsConfigArgs()
        {
        }
    }
}
