// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Alpha.Inputs
{

    /// <summary>
    /// An RPM package file. RPM packages only support INSTALLED state.
    /// </summary>
    public sealed class OSPolicyResourcePackageResourceRPMArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether dependencies should also be installed. - install when false: `rpm --upgrade --replacepkgs package.rpm` - install when true: `yum -y install package.rpm` or `zypper -y install package.rpm`
        /// </summary>
        [Input("pullDeps")]
        public Input<bool>? PullDeps { get; set; }

        /// <summary>
        /// Required. An rpm package.
        /// </summary>
        [Input("source")]
        public Input<Inputs.OSPolicyResourceFileArgs>? Source { get; set; }

        public OSPolicyResourcePackageResourceRPMArgs()
        {
        }
    }
}
