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
    /// An OS policy resource is used to define the desired state configuration and provides a specific functionality like installing/removing packages, executing a script etc. The system ensures that resources are always in their desired state by taking necessary actions if they have drifted from their desired state.
    /// </summary>
    public sealed class OSPolicyResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Exec resource
        /// </summary>
        [Input("exec")]
        public Input<Inputs.OSPolicyResourceExecResourceArgs>? Exec { get; set; }

        /// <summary>
        /// File resource
        /// </summary>
        [Input("file")]
        public Input<Inputs.OSPolicyResourceFileResourceArgs>? File { get; set; }

        /// <summary>
        /// The id of the resource with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the OS policy.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Package resource
        /// </summary>
        [Input("pkg")]
        public Input<Inputs.OSPolicyResourcePackageResourceArgs>? Pkg { get; set; }

        /// <summary>
        /// Package repository resource
        /// </summary>
        [Input("repository")]
        public Input<Inputs.OSPolicyResourceRepositoryResourceArgs>? Repository { get; set; }

        public OSPolicyResourceArgs()
        {
        }
    }
}
