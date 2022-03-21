// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1.Outputs
{

    /// <summary>
    /// Details required for a contact associated with a `Registration`.
    /// </summary>
    [OutputType]
    public sealed class ContactResponse
    {
        /// <summary>
        /// Email address of the contact.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Fax number of the contact in international format. For example, `"+1-800-555-0123"`.
        /// </summary>
        public readonly string FaxNumber;
        /// <summary>
        /// Phone number of the contact in international format. For example, `"+1-800-555-0123"`.
        /// </summary>
        public readonly string PhoneNumber;
        /// <summary>
        /// Postal address of the contact.
        /// </summary>
        public readonly Outputs.PostalAddressResponse PostalAddress;

        [OutputConstructor]
        private ContactResponse(
            string email,

            string faxNumber,

            string phoneNumber,

            Outputs.PostalAddressResponse postalAddress)
        {
            Email = email;
            FaxNumber = faxNumber;
            PhoneNumber = phoneNumber;
            PostalAddress = postalAddress;
        }
    }
}
