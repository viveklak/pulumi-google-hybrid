// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1.Inputs
{

    /// <summary>
    /// Oracle data source configuration
    /// </summary>
    public sealed class OracleSourceConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Oracle objects to include in the stream.
        /// </summary>
        [Input("allowlist")]
        public Input<Inputs.OracleRdbmsArgs>? Allowlist { get; set; }

        /// <summary>
        /// Drop large object values.
        /// </summary>
        [Input("dropLargeObjects")]
        public Input<Inputs.OracleDropLargeObjectsArgs>? DropLargeObjects { get; set; }

        /// <summary>
        /// Oracle objects to exclude from the stream.
        /// </summary>
        [Input("rejectlist")]
        public Input<Inputs.OracleRdbmsArgs>? Rejectlist { get; set; }

        public OracleSourceConfigArgs()
        {
        }
    }
}
