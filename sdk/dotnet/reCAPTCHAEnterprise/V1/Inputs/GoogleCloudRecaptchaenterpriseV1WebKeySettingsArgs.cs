// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.Inputs
{

    /// <summary>
    /// Settings specific to keys that can be used by websites.
    /// </summary>
    public sealed class GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, it means allowed_domains will not be enforced.
        /// </summary>
        [Input("allowAllDomains")]
        public Input<bool>? AllowAllDomains { get; set; }

        /// <summary>
        /// Required. Whether this key can be used on AMP (Accelerated Mobile Pages) websites. This can only be set for the SCORE integration type.
        /// </summary>
        [Input("allowAmpTraffic")]
        public Input<bool>? AllowAmpTraffic { get; set; }

        [Input("allowedDomains")]
        private InputList<string>? _allowedDomains;

        /// <summary>
        /// Domains or subdomains of websites allowed to use the key. All subdomains of an allowed domain are automatically allowed. A valid domain requires a host and must not include any path, port, query or fragment. Examples: 'example.com' or 'subdomain.example.com'
        /// </summary>
        public InputList<string> AllowedDomains
        {
            get => _allowedDomains ?? (_allowedDomains = new InputList<string>());
            set => _allowedDomains = value;
        }

        /// <summary>
        /// Settings for the frequency and difficulty at which this key triggers captcha challenges. This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE.
        /// </summary>
        [Input("challengeSecurityPreference")]
        public Input<Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference>? ChallengeSecurityPreference { get; set; }

        /// <summary>
        /// Required. Describes how this key is integrated with the website.
        /// </summary>
        [Input("integrationType")]
        public Input<Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType>? IntegrationType { get; set; }

        public GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs()
        {
        }
    }
}
