// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1
{
    public static class GetTagTemplate
    {
        /// <summary>
        /// Gets a tag template.
        /// </summary>
        public static Task<GetTagTemplateResult> InvokeAsync(GetTagTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagTemplateResult>("google-native:datacatalog/v1:getTagTemplate", args ?? new GetTagTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a tag template.
        /// </summary>
        public static Output<GetTagTemplateResult> Invoke(GetTagTemplateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTagTemplateResult>("google-native:datacatalog/v1:getTagTemplate", args ?? new GetTagTemplateInvokeArgs(), options.WithDefaults());
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
        /// Display name for this template. Defaults to an empty string. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum length is 200 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. The map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. The IDs have the following limitations: * Can contain uppercase and lowercase letters, numbers (0-9) and underscores (_). * Must be at least 1 character and at most 64 characters long. * Must start with a letter or underscore.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Fields;
        /// <summary>
        /// Indicates whether tags created with this template are public. Public tags do not require tag template access to appear in ListTags API response. Additionally, you can search for a public tag by value with a simple search query in addition to using a ``tag:`` predicate.
        /// </summary>
        public readonly bool IsPubliclyReadable;
        /// <summary>
        /// The resource name of the tag template in URL format. Note: The tag template itself and its child resources might not be stored in the location specified in its name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetTagTemplateResult(
            string displayName,

            ImmutableDictionary<string, string> fields,

            bool isPubliclyReadable,

            string name)
        {
            DisplayName = displayName;
            Fields = fields;
            IsPubliclyReadable = isPubliclyReadable;
            Name = name;
        }
    }
}
