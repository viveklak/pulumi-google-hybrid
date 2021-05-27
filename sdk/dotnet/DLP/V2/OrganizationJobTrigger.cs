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
    /// Creates a job trigger to run DLP actions such as scanning storage for sensitive information on a set schedule. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.
    /// </summary>
    [GoogleNativeResourceType("google-native:dlp/v2:OrganizationJobTrigger")]
    public partial class OrganizationJobTrigger : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation timestamp of a triggeredJob.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// User provided description (max 256 chars)
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Display name (max 100 chars)
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
        /// </summary>
        [Output("errors")]
        public Output<ImmutableArray<Outputs.GooglePrivacyDlpV2ErrorResponse>> Errors { get; private set; } = null!;

        /// <summary>
        /// For inspect jobs, a snapshot of the configuration.
        /// </summary>
        [Output("inspectJob")]
        public Output<Outputs.GooglePrivacyDlpV2InspectJobConfigResponse> InspectJob { get; private set; } = null!;

        /// <summary>
        /// The timestamp of the last time this trigger executed.
        /// </summary>
        [Output("lastRunTime")]
        public Output<string> LastRunTime { get; private set; } = null!;

        /// <summary>
        /// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. A status for this trigger.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableArray<Outputs.GooglePrivacyDlpV2TriggerResponse>> Triggers { get; private set; } = null!;

        /// <summary>
        /// The last update timestamp of a triggeredJob.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationJobTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationJobTrigger(string name, OrganizationJobTriggerArgs args, CustomResourceOptions? options = null)
            : base("google-native:dlp/v2:OrganizationJobTrigger", name, args ?? new OrganizationJobTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationJobTrigger(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dlp/v2:OrganizationJobTrigger", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationJobTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationJobTrigger Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationJobTrigger(name, id, options);
        }
    }

    public sealed class OrganizationJobTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User provided description (max 256 chars)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name (max 100 chars)
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// For inspect jobs, a snapshot of the configuration.
        /// </summary>
        [Input("inspectJob")]
        public Input<Inputs.GooglePrivacyDlpV2InspectJobConfigArgs>? InspectJob { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Required. A status for this trigger.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The trigger id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        /// </summary>
        [Input("triggerId")]
        public Input<string>? TriggerId { get; set; }

        [Input("triggers")]
        private InputList<Inputs.GooglePrivacyDlpV2TriggerArgs>? _triggers;

        /// <summary>
        /// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2TriggerArgs> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<Inputs.GooglePrivacyDlpV2TriggerArgs>());
            set => _triggers = value;
        }

        public OrganizationJobTriggerArgs()
        {
        }
    }
}
