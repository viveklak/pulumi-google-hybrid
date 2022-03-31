// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Submit a data export job to be processed in the background. If the request is successful, the API returns a 201 status, a URI that can be used to retrieve the status of the export job, and the `state` value of "enqueued".
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:Export")]
    public partial class Export : Pulumi.CustomResource
    {
        /// <summary>
        /// Time the export job was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Name of the datastore that is the destination of the export job [datastore]
        /// </summary>
        [Output("datastoreName")]
        public Output<string> DatastoreName { get; private set; } = null!;

        /// <summary>
        /// Description of the export job.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Error is set when export fails
        /// </summary>
        [Output("error")]
        public Output<string> Error { get; private set; } = null!;

        /// <summary>
        /// Execution time for this export job. If the job is still in progress, it will be set to the amount of time that has elapsed since`created`, in seconds. Else, it will set to (`updated` - `created`), in seconds.
        /// </summary>
        [Output("executionTime")]
        public Output<string> ExecutionTime { get; private set; } = null!;

        /// <summary>
        /// Display name of the export job.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Self link of the export job. A URI that can be used to retrieve the status of an export job. Example: `/organizations/myorg/environments/myenv/analytics/exports/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
        /// </summary>
        [Output("self")]
        public Output<string> Self { get; private set; } = null!;

        /// <summary>
        /// Status of the export job. Valid values include `enqueued`, `running`, `completed`, and `failed`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Time the export job was last updated.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a Export resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Export(string name, ExportArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:Export", name, args ?? new ExportArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Export(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:Export", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Export resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Export Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Export(name, id, options);
        }
    }

    public sealed class ExportArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
        /// </summary>
        [Input("csvDelimiter")]
        public Input<string>? CsvDelimiter { get; set; }

        /// <summary>
        /// Name of the preconfigured datastore.
        /// </summary>
        [Input("datastoreName", required: true)]
        public Input<string> DatastoreName { get; set; } = null!;

        /// <summary>
        /// Date range of the data to export.
        /// </summary>
        [Input("dateRange", required: true)]
        public Input<Inputs.GoogleCloudApigeeV1DateRangeArgs> DateRange { get; set; } = null!;

        /// <summary>
        /// Optional. Description of the export job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Display name of the export job.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Optional. Output format of the export. Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the `csvDelimiter` property.
        /// </summary>
        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        public ExportArgs()
        {
        }
    }
}
