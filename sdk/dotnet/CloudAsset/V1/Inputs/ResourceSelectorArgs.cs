// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudAsset.V1.Inputs
{

    /// <summary>
    /// Specifies the resource to analyze for access policies, which may be set directly on the resource, or on ancestors such as organizations, folders or projects.
    /// </summary>
    public sealed class ResourceSelectorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of a resource of [supported resource types](https://cloud.google.com/asset-inventory/docs/supported-asset-types#analyzable_asset_types).
        /// </summary>
        [Input("fullResourceName", required: true)]
        public Input<string> FullResourceName { get; set; } = null!;

        public ResourceSelectorArgs()
        {
        }
    }
}
