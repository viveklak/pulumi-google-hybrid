// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3
{
    /// <summary>
    /// Creates a new alerting policy.
    /// </summary>
    [GoogleNativeResourceType("google-native:monitoring/v3:AlertPolicy")]
    public partial class AlertPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
        /// </summary>
        [Output("combiner")]
        public Output<string> Combiner { get; private set; } = null!;

        /// <summary>
        /// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. If condition_time_series_query_language is present, it must be the only condition.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.ConditionResponse>> Conditions { get; private set; } = null!;

        /// <summary>
        /// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored.
        /// </summary>
        [Output("creationRecord")]
        public Output<Outputs.MutationRecordResponse> CreationRecord { get; private set; } = null!;

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Documentation that is included with notifications and incidents related to this policy. Best practice is for the documentation to include information to help responders understand, mitigate, escalate, and correct the underlying problems detected by the alerting policy. Notification channels that have limited capacity might not show this documentation.
        /// </summary>
        [Output("documentation")]
        public Output<Outputs.DocumentationResponse> Documentation { get; private set; } = null!;

        /// <summary>
        /// Whether or not the policy is enabled. On write, the default interpretation if unset is that the policy is enabled. On read, clients should not make any assumption about the state if it has not been populated. The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// A read-only record of the most recent change to the alerting policy. If provided in a call to create or update, this field will be ignored.
        /// </summary>
        [Output("mutationRecord")]
        public Output<Outputs.MutationRecordResponse> MutationRecord { get; private set; } = null!;

        /// <summary>
        /// Required if the policy exists. The resource name for this policy. The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID] [ALERT_POLICY_ID] is assigned by Stackdriver Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method. The format of the entries in this field is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] 
        /// </summary>
        [Output("notificationChannels")]
        public Output<ImmutableArray<string>> NotificationChannels { get; private set; } = null!;

        /// <summary>
        /// User-supplied key/value data to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        /// </summary>
        [Output("userLabels")]
        public Output<ImmutableDictionary<string, string>> UserLabels { get; private set; } = null!;

        /// <summary>
        /// Read-only description of how the alert policy is invalid. OK if the alert policy is valid. If not OK, the alert policy will not generate incidents.
        /// </summary>
        [Output("validity")]
        public Output<Outputs.StatusResponse> Validity { get; private set; } = null!;


        /// <summary>
        /// Create a AlertPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertPolicy(string name, AlertPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:monitoring/v3:AlertPolicy", name, args ?? new AlertPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:monitoring/v3:AlertPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AlertPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AlertPolicy(name, id, options);
        }
    }

    public sealed class AlertPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
        /// </summary>
        [Input("combiner")]
        public Input<string>? Combiner { get; set; }

        [Input("conditions")]
        private InputList<Inputs.ConditionArgs>? _conditions;

        /// <summary>
        /// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. If condition_time_series_query_language is present, it must be the only condition.
        /// </summary>
        public InputList<Inputs.ConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored.
        /// </summary>
        [Input("creationRecord")]
        public Input<Inputs.MutationRecordArgs>? CreationRecord { get; set; }

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Documentation that is included with notifications and incidents related to this policy. Best practice is for the documentation to include information to help responders understand, mitigate, escalate, and correct the underlying problems detected by the alerting policy. Notification channels that have limited capacity might not show this documentation.
        /// </summary>
        [Input("documentation")]
        public Input<Inputs.DocumentationArgs>? Documentation { get; set; }

        /// <summary>
        /// Whether or not the policy is enabled. On write, the default interpretation if unset is that the policy is enabled. On read, clients should not make any assumption about the state if it has not been populated. The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// A read-only record of the most recent change to the alerting policy. If provided in a call to create or update, this field will be ignored.
        /// </summary>
        [Input("mutationRecord")]
        public Input<Inputs.MutationRecordArgs>? MutationRecord { get; set; }

        /// <summary>
        /// Required if the policy exists. The resource name for this policy. The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID] [ALERT_POLICY_ID] is assigned by Stackdriver Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<string>? _notificationChannels;

        /// <summary>
        /// Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method. The format of the entries in this field is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] 
        /// </summary>
        public InputList<string> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<string>());
            set => _notificationChannels = value;
        }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// User-supplied key/value data to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        /// <summary>
        /// Read-only description of how the alert policy is invalid. OK if the alert policy is valid. If not OK, the alert policy will not generate incidents.
        /// </summary>
        [Input("validity")]
        public Input<Inputs.StatusArgs>? Validity { get; set; }

        public AlertPolicyArgs()
        {
        }
    }
}
