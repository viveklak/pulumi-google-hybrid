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
    public static class GetTagTemplate
    {
        /// <summary>
        /// Gets a tag template.
        /// </summary>
        public static Task<GetTagTemplateResult> InvokeAsync(GetTagTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagTemplateResult>("google-native:datacatalog/v1beta1:getTagTemplate", args ?? new GetTagTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a tag template.
        /// </summary>
        public static Output<GetTagTemplateResult> Invoke(GetTagTemplateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTagTemplateResult>("google-native:datacatalog/v1beta1:getTagTemplate", args ?? new GetTagTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagTemplateArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("tagTemplateId", required: true)]
        public string TagTemplateId { get; set; } = null!;

        public GetTagTemplateArgs()
        {
        }
    }

    public sealed class GetTagTemplateInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("tagTemplateId", required: true)]
        public Input<string> TagTemplateId { get; set; } = null!;

        public GetTagTemplateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagTemplateResult
    {
        /// <summary>
        /// The display name for this template. Defaults to an empty string.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. This map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. Field IDs can contain letters (both uppercase and lowercase), numbers (0-9) and underscores (_). Field IDs must be at least 1 character long and at most 64 characters long. Field IDs must start with a letter or underscore.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Fields;
        /// <summary>
        /// The resource name of the tag template in URL format. Example: * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id} Note that this TagTemplate and its child resources may not actually be stored in the location in this name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetTagTemplateResult(
            string displayName,

            ImmutableDictionary<string, string> fields,

            string name)
        {
            DisplayName = displayName;
            Fields = fields;
            Name = name;
        }
    }
}
