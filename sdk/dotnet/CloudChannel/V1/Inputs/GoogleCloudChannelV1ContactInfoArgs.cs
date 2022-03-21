// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1.Inputs
{

    /// <summary>
    /// Contact information for a customer account.
    /// </summary>
    public sealed class GoogleCloudChannelV1ContactInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The customer account's contact email. Required for entitlements that create admin.google.com accounts, and serves as the customer's username for those accounts. Use this email to invite Team customers.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The customer account contact's first name. Optional for Team customers.
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// The customer account contact's last name. Optional for Team customers.
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        /// <summary>
        /// The customer account's contact phone number.
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        /// <summary>
        /// Optional. The customer account contact's job title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public GoogleCloudChannelV1ContactInfoArgs()
        {
        }
    }
}
