// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2
{
    public static class GetDeidentifyTemplate
    {
        /// <summary>
        /// Gets a DeidentifyTemplate. See https://cloud.google.com/dlp/docs/creating-templates-deid to learn more.
        /// </summary>
        public static Task<GetDeidentifyTemplateResult> InvokeAsync(GetDeidentifyTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeidentifyTemplateResult>("google-native:dlp/v2:getDeidentifyTemplate", args ?? new GetDeidentifyTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a DeidentifyTemplate. See https://cloud.google.com/dlp/docs/creating-templates-deid to learn more.
        /// </summary>
        public static Output<GetDeidentifyTemplateResult> Invoke(GetDeidentifyTemplateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDeidentifyTemplateResult>("google-native:dlp/v2:getDeidentifyTemplate", args ?? new GetDeidentifyTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeidentifyTemplateArgs : Pulumi.InvokeArgs
    {
        [Input("deidentifyTemplateId", required: true)]
        public string DeidentifyTemplateId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDeidentifyTemplateArgs()
        {
        }
    }

    public sealed class GetDeidentifyTemplateInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("deidentifyTemplateId", required: true)]
        public Input<string> DeidentifyTemplateId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDeidentifyTemplateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDeidentifyTemplateResult
    {
        /// <summary>
        /// The creation timestamp of an inspectTemplate.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The core content of the template.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2DeidentifyConfigResponse DeidentifyConfig;
        /// <summary>
        /// Short description (max 256 chars).
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Display name (max 256 chars).
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The template name. The template will have one of the following formats: `projects/PROJECT_ID/deidentifyTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/deidentifyTemplates/TEMPLATE_ID`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The last update timestamp of an inspectTemplate.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetDeidentifyTemplateResult(
            string createTime,

            Outputs.GooglePrivacyDlpV2DeidentifyConfigResponse deidentifyConfig,

            string description,

            string displayName,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            DeidentifyConfig = deidentifyConfig;
            Description = description;
            DisplayName = displayName;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
