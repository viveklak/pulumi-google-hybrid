// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.IAM.V1
{
    public static class GetServiceAccount
    {
        /// <summary>
        /// Gets a ServiceAccount.
        /// </summary>
        public static Task<GetServiceAccountResult> InvokeAsync(GetServiceAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceAccountResult>("google-native:iam/v1:getServiceAccount", args ?? new GetServiceAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a ServiceAccount.
        /// </summary>
        public static Output<GetServiceAccountResult> Invoke(GetServiceAccountInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServiceAccountResult>("google-native:iam/v1:getServiceAccount", args ?? new GetServiceAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceAccountArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceAccountId", required: true)]
        public string ServiceAccountId { get; set; } = null!;

        public GetServiceAccountArgs()
        {
        }
    }

    public sealed class GetServiceAccountInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceAccountId", required: true)]
        public Input<string> ServiceAccountId { get; set; } = null!;

        public GetServiceAccountInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceAccountResult
    {
        /// <summary>
        /// Optional. A user-specified, human-readable description of the service account. The maximum length is 256 UTF-8 bytes.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether the service account is disabled.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// Optional. A user-specified, human-readable name for the service account. The maximum length is 100 UTF-8 bytes.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The email address of the service account.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The resource name of the service account. Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes. For example, if you try to get the service account `projects/-/serviceAccounts/fake@example.com`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The OAuth 2.0 client ID for the service account.
        /// </summary>
        public readonly string Oauth2ClientId;
        /// <summary>
        /// The ID of the project that owns the service account.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The unique, stable numeric ID for the service account. Each service account retains its unique ID even if you delete the service account. For example, if you delete a service account, then create a new service account with the same name, the new service account has a different unique ID than the deleted service account.
        /// </summary>
        public readonly string UniqueId;

        [OutputConstructor]
        private GetServiceAccountResult(
            string description,

            bool disabled,

            string displayName,

            string email,

            string name,

            string oauth2ClientId,

            string project,

            string uniqueId)
        {
            Description = description;
            Disabled = disabled;
            DisplayName = displayName;
            Email = email;
            Name = name;
            Oauth2ClientId = oauth2ClientId;
            Project = project;
            UniqueId = uniqueId;
        }
    }
}
