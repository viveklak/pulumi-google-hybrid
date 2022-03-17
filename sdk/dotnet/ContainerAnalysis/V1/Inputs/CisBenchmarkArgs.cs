// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// A compliance check that is a CIS benchmark.
    /// </summary>
    public sealed class CisBenchmarkArgs : Pulumi.ResourceArgs
    {
        [Input("profileLevel")]
        public Input<int>? ProfileLevel { get; set; }

        [Input("severity")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1.CisBenchmarkSeverity>? Severity { get; set; }

        public CisBenchmarkArgs()
        {
        }
    }
}
