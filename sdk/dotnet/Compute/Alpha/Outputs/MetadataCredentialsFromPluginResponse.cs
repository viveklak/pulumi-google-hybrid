// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// [Deprecated] Custom authenticator credentials. Custom authenticator credentials.
    /// </summary>
    [OutputType]
    public sealed class MetadataCredentialsFromPluginResponse
    {
        /// <summary>
        /// Plugin name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A text proto that conforms to a Struct type definition interpreted by the plugin.
        /// </summary>
        public readonly string StructConfig;

        [OutputConstructor]
        private MetadataCredentialsFromPluginResponse(
            string name,

            string structConfig)
        {
            Name = name;
            StructConfig = structConfig;
        }
    }
}
