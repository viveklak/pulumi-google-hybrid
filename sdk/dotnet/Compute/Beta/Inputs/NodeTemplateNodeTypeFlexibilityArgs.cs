// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class NodeTemplateNodeTypeFlexibilityArgs : Pulumi.ResourceArgs
    {
        [Input("cpus")]
        public Input<string>? Cpus { get; set; }

        [Input("localSsd")]
        public Input<string>? LocalSsd { get; set; }

        [Input("memory")]
        public Input<string>? Memory { get; set; }

        public NodeTemplateNodeTypeFlexibilityArgs()
        {
        }
    }
}
