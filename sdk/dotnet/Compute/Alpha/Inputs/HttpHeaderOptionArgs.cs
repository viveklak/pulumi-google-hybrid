// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Specification determining how headers are added to requests or responses.
    /// </summary>
    public sealed class HttpHeaderOptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the header.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// The value of the header to add.
        /// </summary>
        [Input("headerValue")]
        public Input<string>? HeaderValue { get; set; }

        /// <summary>
        /// If false, headerValue is appended to any values that already exist for the header. If true, headerValue is set for the header, discarding any values that were set for that header. The default value is false. 
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        public HttpHeaderOptionArgs()
        {
        }
    }
}
