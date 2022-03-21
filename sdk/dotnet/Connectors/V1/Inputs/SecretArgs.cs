// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1.Inputs
{

    /// <summary>
    /// Secret provides a reference to entries in Secret Manager.
    /// </summary>
    public sealed class SecretArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name of the secret version in the format, format as: `projects/*/secrets/*/versions/*`.
        /// </summary>
        [Input("secretVersion")]
        public Input<string>? SecretVersion { get; set; }

        public SecretArgs()
        {
        }
    }
}
