// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1.Inputs
{

    /// <summary>
    /// Defines the contact information associated with a `Registration`. [ICANN](https://icann.org/) requires all domain names to have associated contact information. The `registrant_contact` is considered the domain's legal owner, and often the other contacts are identical.
    /// </summary>
    public sealed class ContactSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative contact for the `Registration`.
        /// </summary>
        [Input("adminContact", required: true)]
        public Input<Inputs.ContactArgs> AdminContact { get; set; } = null!;

        /// <summary>
        /// Privacy setting for the contacts associated with the `Registration`.
        /// </summary>
        [Input("privacy", required: true)]
        public Input<Pulumi.GoogleNative.Domains.V1.ContactSettingsPrivacy> Privacy { get; set; } = null!;

        /// <summary>
        /// The registrant contact for the `Registration`. *Caution: Anyone with access to this email address, phone number, and/or postal address can take control of the domain.* *Warning: For new `Registration`s, the registrant receives an email confirmation that they must complete within 15 days to avoid domain suspension.*
        /// </summary>
        [Input("registrantContact", required: true)]
        public Input<Inputs.ContactArgs> RegistrantContact { get; set; } = null!;

        /// <summary>
        /// The technical contact for the `Registration`.
        /// </summary>
        [Input("technicalContact", required: true)]
        public Input<Inputs.ContactArgs> TechnicalContact { get; set; } = null!;

        public ContactSettingsArgs()
        {
        }
    }
}
