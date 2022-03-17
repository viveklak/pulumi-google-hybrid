// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2.Outputs
{

    /// <summary>
    /// Label object for Deployments
    /// </summary>
    [OutputType]
    public sealed class DeploymentLabelEntryResponse
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
        private DeploymentLabelEntryResponse(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
