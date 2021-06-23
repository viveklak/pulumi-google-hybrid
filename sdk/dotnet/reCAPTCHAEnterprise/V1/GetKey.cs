// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.reCAPTCHAEnterprise.V1
{
    public static class GetKey
    {
        /// <summary>
        /// Returns the specified key.
        /// </summary>
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("google-native:recaptchaenterprise/v1:getKey", args ?? new GetKeyArgs(), options.WithVersion());
    }


    public sealed class GetKeyArgs : Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public string KeyId { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        /// <summary>
        /// Settings for keys that can be used by Android apps.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsResponse AndroidSettings;
        /// <summary>
        /// The timestamp corresponding to the creation of this Key.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Human-readable display name of this key. Modifiable by user.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Settings for keys that can be used by iOS apps.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse IosSettings;
        /// <summary>
        /// See Creating and managing labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name for the Key in the format "projects/{project}/keys/{key}".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Options for user acceptance testing.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1TestingOptionsResponse TestingOptions;
        /// <summary>
        /// Settings for keys that can be used by websites.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1WebKeySettingsResponse WebSettings;

        [OutputConstructor]
        private GetKeyResult(
            Outputs.GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsResponse androidSettings,

            string createTime,

            string displayName,

            Outputs.GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse iosSettings,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.GoogleCloudRecaptchaenterpriseV1TestingOptionsResponse testingOptions,

            Outputs.GoogleCloudRecaptchaenterpriseV1WebKeySettingsResponse webSettings)
        {
            AndroidSettings = androidSettings;
            CreateTime = createTime;
            DisplayName = displayName;
            IosSettings = iosSettings;
            Labels = labels;
            Name = name;
            TestingOptions = testingOptions;
            WebSettings = webSettings;
        }
    }
}
