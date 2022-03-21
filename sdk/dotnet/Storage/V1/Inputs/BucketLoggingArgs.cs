// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Inputs
{

    /// <summary>
    /// The bucket's logging configuration, which defines the destination bucket and optional name prefix for the current bucket's logs.
    /// </summary>
    public sealed class BucketLoggingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination bucket where the current bucket's logs should be placed.
        /// </summary>
        [Input("logBucket")]
        public Input<string>? LogBucket { get; set; }

        /// <summary>
        /// A prefix for log object names.
        /// </summary>
        [Input("logObjectPrefix")]
        public Input<string>? LogObjectPrefix { get; set; }

        public BucketLoggingArgs()
        {
        }
    }
}
