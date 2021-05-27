// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1
{
    /// <summary>
    /// Creates a new custom dashboard. For examples on how you can use this API to create dashboards, see Managing dashboards by API. This method requires the monitoring.dashboards.create permission on the specified project. For more information about permissions, see Cloud Identity and Access Management.
    /// </summary>
    [GoogleNativeResourceType("google-native:monitoring/v1:Dashboard")]
    public partial class Dashboard : Pulumi.CustomResource
    {
        /// <summary>
        /// The content is divided into equally spaced columns and the widgets are arranged vertically.
        /// </summary>
        [Output("columnLayout")]
        public Output<Outputs.ColumnLayoutResponse> ColumnLayout { get; private set; } = null!;

        /// <summary>
        /// Required. The mutable, human-readable name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. An etag is returned in the response to GetDashboard, and users are expected to put that etag in the request to UpdateDashboard to ensure that their change will be applied to the same version of the Dashboard configuration. The field should not be passed during dashboard creation.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Content is arranged with a basic layout that re-flows a simple list of informational elements like widgets or tiles.
        /// </summary>
        [Output("gridLayout")]
        public Output<Outputs.GridLayoutResponse> GridLayout { get; private set; } = null!;

        /// <summary>
        /// The content is arranged as a grid of tiles, with each content widget occupying one or more grid blocks.
        /// </summary>
        [Output("mosaicLayout")]
        public Output<Outputs.MosaicLayoutResponse> MosaicLayout { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the dashboard.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The content is divided into equally spaced rows and the widgets are arranged horizontally.
        /// </summary>
        [Output("rowLayout")]
        public Output<Outputs.RowLayoutResponse> RowLayout { get; private set; } = null!;


        /// <summary>
        /// Create a Dashboard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dashboard(string name, DashboardArgs args, CustomResourceOptions? options = null)
            : base("google-native:monitoring/v1:Dashboard", name, args ?? new DashboardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dashboard(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:monitoring/v1:Dashboard", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Dashboard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dashboard Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Dashboard(name, id, options);
        }
    }

    public sealed class DashboardArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content is divided into equally spaced columns and the widgets are arranged vertically.
        /// </summary>
        [Input("columnLayout")]
        public Input<Inputs.ColumnLayoutArgs>? ColumnLayout { get; set; }

        /// <summary>
        /// Required. The mutable, human-readable name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. An etag is returned in the response to GetDashboard, and users are expected to put that etag in the request to UpdateDashboard to ensure that their change will be applied to the same version of the Dashboard configuration. The field should not be passed during dashboard creation.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Content is arranged with a basic layout that re-flows a simple list of informational elements like widgets or tiles.
        /// </summary>
        [Input("gridLayout")]
        public Input<Inputs.GridLayoutArgs>? GridLayout { get; set; }

        /// <summary>
        /// The content is arranged as a grid of tiles, with each content widget occupying one or more grid blocks.
        /// </summary>
        [Input("mosaicLayout")]
        public Input<Inputs.MosaicLayoutArgs>? MosaicLayout { get; set; }

        /// <summary>
        /// Immutable. The resource name of the dashboard.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The content is divided into equally spaced rows and the widgets are arranged horizontally.
        /// </summary>
        [Input("rowLayout")]
        public Input<Inputs.RowLayoutArgs>? RowLayout { get; set; }

        public DashboardArgs()
        {
        }
    }
}
