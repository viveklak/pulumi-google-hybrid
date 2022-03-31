// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2
{
    /// <summary>
    /// Starts a new asynchronous job. Requires the Can View project role.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:bigquery/v2:Job")]
    public partial class Job : Pulumi.CustomResource
    {
        /// <summary>
        /// [Required] Describes the job configuration.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.JobConfigurationResponse> Configuration { get; private set; } = null!;

        /// <summary>
        /// A hash of this resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// [Optional] Reference describing the unique-per-user name of the job.
        /// </summary>
        [Output("jobReference")]
        public Output<Outputs.JobReferenceResponse> JobReference { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A URL that can be used to access this resource again.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Information about the job, including starting time and ending time of the job.
        /// </summary>
        [Output("statistics")]
        public Output<Outputs.JobStatisticsResponse> Statistics { get; private set; } = null!;

        /// <summary>
        /// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
        /// </summary>
        [Output("status")]
        public Output<Outputs.JobStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Email address of the user who ran the job.
        /// </summary>
        [Output("userEmail")]
        public Output<string> UserEmail { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:bigquery/v2:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:bigquery/v2:Job", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Job(name, id, options);
        }
    }

    public sealed class JobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Required] Describes the job configuration.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.JobConfigurationArgs>? Configuration { get; set; }

        /// <summary>
        /// [Optional] Reference describing the unique-per-user name of the job.
        /// </summary>
        [Input("jobReference")]
        public Input<Inputs.JobReferenceArgs>? JobReference { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("source")]
        public Input<AssetOrArchive>? Source { get; set; }

        public JobArgs()
        {
        }
    }
}
