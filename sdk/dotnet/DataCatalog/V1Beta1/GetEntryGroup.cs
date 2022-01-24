// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.DataCatalog.V1Beta1
{
    public static class GetEntryGroup
    {
        /// <summary>
        /// Gets an EntryGroup.
        /// </summary>
        public static Task<GetEntryGroupResult> InvokeAsync(GetEntryGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntryGroupResult>("google-native:datacatalog/v1beta1:getEntryGroup", args ?? new GetEntryGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an EntryGroup.
        /// </summary>
        public static Output<GetEntryGroupResult> Invoke(GetEntryGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEntryGroupResult>("google-native:datacatalog/v1beta1:getEntryGroup", args ?? new GetEntryGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEntryGroupArgs : Pulumi.InvokeArgs
    {
        [Input("entryGroupId", required: true)]
        public string EntryGroupId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("readMask")]
        public string? ReadMask { get; set; }

        public GetEntryGroupArgs()
        {
        }
    }

    public sealed class GetEntryGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("entryGroupId", required: true)]
        public Input<string> EntryGroupId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("readMask")]
        public Input<string>? ReadMask { get; set; }

        public GetEntryGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEntryGroupResult
    {
        /// <summary>
        /// Timestamps about this EntryGroup. Default value is empty timestamps.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse DataCatalogTimestamps;
        /// <summary>
        /// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents. Default value is an empty string.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A short name to identify the entry group, for example, "analytics data - jan 2011". Default value is an empty string.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The resource name of the entry group in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id} Note that this EntryGroup and its child resources may not actually be stored in the location in this name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetEntryGroupResult(
            Outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse dataCatalogTimestamps,

            string description,

            string displayName,

            string name)
        {
            DataCatalogTimestamps = dataCatalogTimestamps;
            Description = description;
            DisplayName = displayName;
            Name = name;
        }
    }
}
