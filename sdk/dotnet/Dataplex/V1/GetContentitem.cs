// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1
{
    public static class GetContentitem
    {
        /// <summary>
        /// Get a content resource.
        /// </summary>
        public static Task<GetContentitemResult> InvokeAsync(GetContentitemArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetContentitemResult>("google-native:dataplex/v1:getContentitem", args ?? new GetContentitemArgs(), options.WithDefaults());

        /// <summary>
        /// Get a content resource.
        /// </summary>
        public static Output<GetContentitemResult> Invoke(GetContentitemInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetContentitemResult>("google-native:dataplex/v1:getContentitem", args ?? new GetContentitemInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContentitemArgs : Pulumi.InvokeArgs
    {
        [Input("contentitemId", required: true)]
        public string ContentitemId { get; set; } = null!;

        [Input("lakeId", required: true)]
        public string LakeId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        public GetContentitemArgs()
        {
        }
    }

    public sealed class GetContentitemInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("contentitemId", required: true)]
        public Input<string> ContentitemId { get; set; } = null!;

        [Input("lakeId", required: true)]
        public Input<string> LakeId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetContentitemInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetContentitemResult
    {
        /// <summary>
        /// Content creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Content data in string format.
        /// </summary>
        public readonly string DataText;
        /// <summary>
        /// Optional. Description of the content.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. User defined labels for the content.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The relative resource name of the content, of the form: projects/{project_id}/locations/{location_id}/lakes/{lake_id}/content/{content_id}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Notebook related configurations.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1ContentNotebookResponse Notebook;
        /// <summary>
        /// The path for the Content file, represented as directory structure. Unique within a lake. Limited to alphanumerics, hyphens, underscores, dots and slashes.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Sql Script related configurations.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1ContentSqlScriptResponse SqlScript;
        /// <summary>
        /// System generated globally unique ID for the content. This ID will be different if the content is deleted and re-created with the same name.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The time when the content was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetContentitemResult(
            string createTime,

            string dataText,

            string description,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.GoogleCloudDataplexV1ContentNotebookResponse notebook,

            string path,

            Outputs.GoogleCloudDataplexV1ContentSqlScriptResponse sqlScript,

            string uid,

            string updateTime)
        {
            CreateTime = createTime;
            DataText = dataText;
            Description = description;
            Labels = labels;
            Name = name;
            Notebook = notebook;
            Path = path;
            SqlScript = sqlScript;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}