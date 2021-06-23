// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// MetadataFilter label name value pairs that are expected to match corresponding labels presented as metadata to the loadbalancer.
    /// </summary>
    public sealed class MetadataFilterLabelMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of metadata label. The name can have a maximum length of 1024 characters and must be at least 1 character long.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The value of the label must match the specified value. value can have a maximum length of 1024 characters.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public MetadataFilterLabelMatchArgs()
        {
        }
    }
}
