// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1
{
    /// <summary>
    /// Create security settings in the specified location.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3beta1:SecuritySetting")]
    public partial class SecuritySetting : Pulumi.CustomResource
    {
        /// <summary>
        /// The human-readable name of the security settings, unique within the location.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// DLP inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`
        /// </summary>
        [Output("inspectTemplate")]
        public Output<string> InspectTemplate { get; private set; } = null!;

        /// <summary>
        /// Resource name of the settings. Format: `projects//locations//securitySettings/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of types of data to remove when retention settings triggers purge.
        /// </summary>
        [Output("purgeDataTypes")]
        public Output<ImmutableArray<string>> PurgeDataTypes { get; private set; } = null!;

        /// <summary>
        /// Defines on what data we apply redaction. Note that we don't redact data to which we don't have access, e.g., Stackdriver logs.
        /// </summary>
        [Output("redactionScope")]
        public Output<string> RedactionScope { get; private set; } = null!;

        /// <summary>
        /// Strategy that defines how we do redaction.
        /// </summary>
        [Output("redactionStrategy")]
        public Output<string> RedactionStrategy { get; private set; } = null!;

        /// <summary>
        /// Retains the data for the specified number of days. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL.
        /// </summary>
        [Output("retentionWindowDays")]
        public Output<int> RetentionWindowDays { get; private set; } = null!;


        /// <summary>
        /// Create a SecuritySetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecuritySetting(string name, SecuritySettingArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:SecuritySetting", name, args ?? new SecuritySettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecuritySetting(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:SecuritySetting", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SecuritySetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecuritySetting Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecuritySetting(name, id, options);
        }
    }

    public sealed class SecuritySettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable name of the security settings, unique within the location.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// DLP inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`
        /// </summary>
        [Input("inspectTemplate")]
        public Input<string>? InspectTemplate { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Resource name of the settings. Format: `projects//locations//securitySettings/`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("purgeDataTypes")]
        private InputList<Pulumi.GoogleNative.Dialogflow.V3Beta1.SecuritySettingPurgeDataTypesItem>? _purgeDataTypes;

        /// <summary>
        /// List of types of data to remove when retention settings triggers purge.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Dialogflow.V3Beta1.SecuritySettingPurgeDataTypesItem> PurgeDataTypes
        {
            get => _purgeDataTypes ?? (_purgeDataTypes = new InputList<Pulumi.GoogleNative.Dialogflow.V3Beta1.SecuritySettingPurgeDataTypesItem>());
            set => _purgeDataTypes = value;
        }

        /// <summary>
        /// Defines on what data we apply redaction. Note that we don't redact data to which we don't have access, e.g., Stackdriver logs.
        /// </summary>
        [Input("redactionScope")]
        public Input<Pulumi.GoogleNative.Dialogflow.V3Beta1.SecuritySettingRedactionScope>? RedactionScope { get; set; }

        /// <summary>
        /// Strategy that defines how we do redaction.
        /// </summary>
        [Input("redactionStrategy")]
        public Input<Pulumi.GoogleNative.Dialogflow.V3Beta1.SecuritySettingRedactionStrategy>? RedactionStrategy { get; set; }

        /// <summary>
        /// Retains the data for the specified number of days. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL.
        /// </summary>
        [Input("retentionWindowDays")]
        public Input<int>? RetentionWindowDays { get; set; }

        public SecuritySettingArgs()
        {
        }
    }
}
