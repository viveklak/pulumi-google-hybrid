// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// Configuration for the Integration add-on.
    /// </summary>
    public sealed class GoogleCloudApigeeV1IntegrationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag that specifies whether the Integration add-on is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public GoogleCloudApigeeV1IntegrationConfigArgs()
        {
        }
    }
}
