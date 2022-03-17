// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2Beta.Outputs
{

    /// <summary>
    /// Label object for TypeProviders
    /// </summary>
    [OutputType]
    public sealed class TypeProviderLabelEntryResponse
    {
        /// <summary>
        /// Key of the label
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Value of the label
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private TypeProviderLabelEntryResponse(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
