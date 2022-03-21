// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Outputs
{

    /// <summary>
    /// Defines the layout properties and content for a row.
    /// </summary>
    [OutputType]
    public sealed class RowResponse
    {
        /// <summary>
        /// The relative weight of this row. The row weight is used to adjust the height of rows on the screen (relative to peers). Greater the weight, greater the height of the row on the screen. If omitted, a value of 1 is used while rendering.
        /// </summary>
        public readonly string Weight;
        /// <summary>
        /// The display widgets arranged horizontally in this row.
        /// </summary>
        public readonly ImmutableArray<Outputs.WidgetResponse> Widgets;

        [OutputConstructor]
        private RowResponse(
            string weight,

            ImmutableArray<Outputs.WidgetResponse> widgets)
        {
            Weight = weight;
            Widgets = widgets;
        }
    }
}
