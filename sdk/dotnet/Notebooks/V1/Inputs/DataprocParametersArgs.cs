// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1.Inputs
{

    /// <summary>
    /// Parameters used in Dataproc JobType executions.
    /// </summary>
    public sealed class DataprocParametersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// URI for cluster used to run Dataproc execution. Format: `projects/{PROJECT_ID}/regions/{REGION}/clusters/{CLUSTER_NAME}`
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        public DataprocParametersArgs()
        {
        }
    }
}
