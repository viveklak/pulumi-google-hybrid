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
    /// Settings specific to keys that can be used by Android apps.
    /// </summary>
    public sealed class GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs : Pulumi.ResourceArgs
    {
        [Input("allowedPackageNames")]
        private InputList<string>? _allowedPackageNames;

        /// <summary>
        /// Android package names of apps allowed to use the key. Example: 'com.companyname.appname'
        /// </summary>
        public InputList<string> AllowedPackageNames
        {
            get => _allowedPackageNames ?? (_allowedPackageNames = new InputList<string>());
            set => _allowedPackageNames = value;
        }

        public GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs()
        {
        }
    }
}
