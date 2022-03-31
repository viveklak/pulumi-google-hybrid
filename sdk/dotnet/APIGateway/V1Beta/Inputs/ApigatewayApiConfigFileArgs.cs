// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIGateway.V1Beta.Inputs
{

    /// <summary>
    /// A lightweight description of a file.
    /// </summary>
    public sealed class ApigatewayApiConfigFileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bytes that constitute the file.
        /// </summary>
        [Input("contents")]
        public Input<string>? Contents { get; set; }

        /// <summary>
        /// The file path (full or relative path). This is typically the path of the file when it is uploaded.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public ApigatewayApiConfigFileArgs()
        {
        }
    }
}
