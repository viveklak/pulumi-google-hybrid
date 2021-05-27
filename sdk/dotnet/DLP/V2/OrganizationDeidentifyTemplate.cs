// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2
{
    /// <summary>
    /// Creates a DeidentifyTemplate for re-using frequently used configuration for de-identifying content, images, and storage. See https://cloud.google.com/dlp/docs/creating-templates-deid to learn more.
    /// </summary>
    [GoogleNativeResourceType("google-native:dlp/v2:OrganizationDeidentifyTemplate")]
    public partial class OrganizationDeidentifyTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation timestamp of an inspectTemplate.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The core content of the template.
        /// </summary>
        [Output("deidentifyConfig")]
        public Output<Outputs.GooglePrivacyDlpV2DeidentifyConfigResponse> DeidentifyConfig { get; private set; } = null!;

        /// <summary>
        /// Short description (max 256 chars).
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Display name (max 256 chars).
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The template name. The template will have one of the following formats: `projects/PROJECT_ID/deidentifyTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/deidentifyTemplates/TEMPLATE_ID`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The last update timestamp of an inspectTemplate.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationDeidentifyTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationDeidentifyTemplate(string name, OrganizationDeidentifyTemplateArgs args, CustomResourceOptions? options = null)
            : base("google-native:dlp/v2:OrganizationDeidentifyTemplate", name, args ?? new OrganizationDeidentifyTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationDeidentifyTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dlp/v2:OrganizationDeidentifyTemplate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationDeidentifyTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationDeidentifyTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationDeidentifyTemplate(name, id, options);
        }
    }

    public sealed class OrganizationDeidentifyTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The core content of the template.
        /// </summary>
        [Input("deidentifyConfig")]
        public Input<Inputs.GooglePrivacyDlpV2DeidentifyConfigArgs>? DeidentifyConfig { get; set; }

        /// <summary>
        /// Short description (max 256 chars).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name (max 256 chars).
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// The template id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public OrganizationDeidentifyTemplateArgs()
        {
        }
    }
}
