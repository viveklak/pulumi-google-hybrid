// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1
{
    public static class GetRegistration
    {
        /// <summary>
        /// Gets the details of a `Registration` resource.
        /// </summary>
        public static Task<GetRegistrationResult> InvokeAsync(GetRegistrationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistrationResult>("google-native:domains/v1:getRegistration", args ?? new GetRegistrationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of a `Registration` resource.
        /// </summary>
        public static Output<GetRegistrationResult> Invoke(GetRegistrationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegistrationResult>("google-native:domains/v1:getRegistration", args ?? new GetRegistrationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistrationArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("registrationId", required: true)]
        public string RegistrationId { get; set; } = null!;

        public GetRegistrationArgs()
        {
        }
    }

    public sealed class GetRegistrationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("registrationId", required: true)]
        public Input<string> RegistrationId { get; set; } = null!;

        public GetRegistrationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistrationResult
    {
        /// <summary>
        /// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
        /// </summary>
        public readonly Outputs.ContactSettingsResponse ContactSettings;
        /// <summary>
        /// The creation timestamp of the `Registration` resource.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
        /// </summary>
        public readonly Outputs.DnsSettingsResponse DnsSettings;
        /// <summary>
        /// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The expiration timestamp of the `Registration`.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// The set of issues with the `Registration` that require attention.
        /// </summary>
        public readonly ImmutableArray<string> Issues;
        /// <summary>
        /// Set of labels associated with the `Registration`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
        /// </summary>
        public readonly Outputs.ManagementSettingsResponse ManagementSettings;
        /// <summary>
        /// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
        /// </summary>
        public readonly Outputs.ContactSettingsResponse PendingContactSettings;
        /// <summary>
        /// The state of the `Registration`
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
        /// </summary>
        public readonly ImmutableArray<string> SupportedPrivacy;

        [OutputConstructor]
        private GetRegistrationResult(
            Outputs.ContactSettingsResponse contactSettings,

            string createTime,

            Outputs.DnsSettingsResponse dnsSettings,

            string domainName,

            string expireTime,

            ImmutableArray<string> issues,

            ImmutableDictionary<string, string> labels,

            Outputs.ManagementSettingsResponse managementSettings,

            string name,

            Outputs.ContactSettingsResponse pendingContactSettings,

            string state,

            ImmutableArray<string> supportedPrivacy)
        {
            ContactSettings = contactSettings;
            CreateTime = createTime;
            DnsSettings = dnsSettings;
            DomainName = domainName;
            ExpireTime = expireTime;
            Issues = issues;
            Labels = labels;
            ManagementSettings = managementSettings;
            Name = name;
            PendingContactSettings = pendingContactSettings;
            State = state;
            SupportedPrivacy = supportedPrivacy;
        }
    }
}
