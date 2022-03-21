// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// License information: https://spdx.github.io/spdx-spec/3-package-information/#315-declared-license
    /// </summary>
    public sealed class LicenseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comments
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Expression: https://spdx.github.io/spdx-spec/appendix-IV-SPDX-license-expressions/
        /// </summary>
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        public LicenseArgs()
        {
        }
    }
}
