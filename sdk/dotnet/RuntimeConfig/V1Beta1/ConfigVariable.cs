// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RuntimeConfig.V1Beta1
{
    /// <summary>
    /// Creates a variable within the given configuration. You cannot create a variable with a name that is a prefix of an existing variable name, or a name that has an existing variable name as a prefix. To learn more about creating a variable, read the [Setting and Getting Data](/deployment-manager/runtime-configurator/set-and-get-variables) documentation.
    /// </summary>
    [GoogleNativeResourceType("google-native:runtimeconfig/v1beta1:ConfigVariable")]
    public partial class ConfigVariable : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the variable resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME] The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a valid RuntimeConfig resource and `[VARIABLE_NAME]` follows Unix file system file path naming. The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and dashes. Slashes are used as path element separators and are not part of the `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one non-slash character. Multiple slashes are coalesced into single slash character. Each path segment should match [0-9A-Za-z](?:[_.A-Za-z0-9-]{0,62}[_.A-Za-z0-9])? regular expression. The length of a `[VARIABLE_NAME]` must be less than 256 characters. Once you create a variable, you cannot change the variable name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current state of the variable. The variable state indicates the outcome of the `variables().watch` call and is visible through the `get` and `list` calls.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The string value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. For example, `text: "my text value"`. The string must be valid UTF-8.
        /// </summary>
        [Output("text")]
        public Output<string> Text { get; private set; } = null!;

        /// <summary>
        /// The time of the last variable update. Timestamp will be UTC timestamp.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The binary value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. The value must be base64 encoded, and must comply with IETF RFC4648 (https://www.ietf.org/rfc/rfc4648.txt). Only one of `value` or `text` can be set.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigVariable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigVariable(string name, ConfigVariableArgs args, CustomResourceOptions? options = null)
            : base("google-native:runtimeconfig/v1beta1:ConfigVariable", name, args ?? new ConfigVariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigVariable(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:runtimeconfig/v1beta1:ConfigVariable", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigVariable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigVariable Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigVariable(name, id, options);
        }
    }

    public sealed class ConfigVariableArgs : Pulumi.ResourceArgs
    {
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        /// <summary>
        /// The name of the variable resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME] The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a valid RuntimeConfig resource and `[VARIABLE_NAME]` follows Unix file system file path naming. The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and dashes. Slashes are used as path element separators and are not part of the `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one non-slash character. Multiple slashes are coalesced into single slash character. Each path segment should match [0-9A-Za-z](?:[_.A-Za-z0-9-]{0,62}[_.A-Za-z0-9])? regular expression. The length of a `[VARIABLE_NAME]` must be less than 256 characters. Once you create a variable, you cannot change the variable name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The current state of the variable. The variable state indicates the outcome of the `variables().watch` call and is visible through the `get` and `list` calls.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The string value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. For example, `text: "my text value"`. The string must be valid UTF-8.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        /// <summary>
        /// The time of the last variable update. Timestamp will be UTC timestamp.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// The binary value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. The value must be base64 encoded, and must comply with IETF RFC4648 (https://www.ietf.org/rfc/rfc4648.txt). Only one of `value` or `text` can be set.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ConfigVariableArgs()
        {
        }
    }
}
