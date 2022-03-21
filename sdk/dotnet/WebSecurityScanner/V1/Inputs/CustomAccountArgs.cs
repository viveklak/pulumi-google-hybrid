// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1.Inputs
{

    /// <summary>
    /// Describes authentication configuration that uses a custom account.
    /// </summary>
    public sealed class CustomAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The login form URL of the website.
        /// </summary>
        [Input("loginUrl", required: true)]
        public Input<string> LoginUrl { get; set; } = null!;

        /// <summary>
        /// Input only. The password of the custom account. The credential is stored encrypted and not returned in any response nor included in audit logs.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The user name of the custom account.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public CustomAccountArgs()
        {
        }
    }
}
