// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RuntimeConfig.V1Beta1
{
    public static class GetVariable
    {
        /// <summary>
        /// Gets information about a single variable.
        /// </summary>
        public static Task<GetVariableResult> InvokeAsync(GetVariableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVariableResult>("google-native:runtimeconfig/v1beta1:getVariable", args ?? new GetVariableArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a single variable.
        /// </summary>
        public static Output<GetVariableResult> Invoke(GetVariableInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVariableResult>("google-native:runtimeconfig/v1beta1:getVariable", args ?? new GetVariableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVariableArgs : Pulumi.InvokeArgs
    {
        [Input("configId", required: true)]
        public string ConfigId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("variableId", required: true)]
        public string VariableId { get; set; } = null!;

        public GetVariableArgs()
        {
        }
    }

    public sealed class GetVariableInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("variableId", required: true)]
        public Input<string> VariableId { get; set; } = null!;

        public GetVariableInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVariableResult
    {
        /// <summary>
        /// The name of the variable resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME] The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a valid RuntimeConfig resource and `[VARIABLE_NAME]` follows Unix file system file path naming. The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and dashes. Slashes are used as path element separators and are not part of the `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one non-slash character. Multiple slashes are coalesced into single slash character. Each path segment should match [0-9A-Za-z](?:[_.A-Za-z0-9-]{0,62}[_.A-Za-z0-9])? regular expression. The length of a `[VARIABLE_NAME]` must be less than 256 characters. Once you create a variable, you cannot change the variable name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current state of the variable. The variable state indicates the outcome of the `variables().watch` call and is visible through the `get` and `list` calls.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The string value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. For example, `text: "my text value"`. The string must be valid UTF-8.
        /// </summary>
        public readonly string Text;
        /// <summary>
        /// The time of the last variable update. Timestamp will be UTC timestamp.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The binary value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. The value must be base64 encoded, and must comply with IETF RFC4648 (https://www.ietf.org/rfc/rfc4648.txt). Only one of `value` or `text` can be set.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetVariableResult(
            string name,

            string state,

            string text,

            string updateTime,

            string value)
        {
            Name = name;
            State = state;
            Text = text;
            UpdateTime = updateTime;
            Value = value;
        }
    }
}
