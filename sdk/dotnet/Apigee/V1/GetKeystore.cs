// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetKeystore
    {
        /// <summary>
        /// Gets a keystore or truststore.
        /// </summary>
        public static Task<GetKeystoreResult> InvokeAsync(GetKeystoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeystoreResult>("google-native:apigee/v1:getKeystore", args ?? new GetKeystoreArgs(), options.WithVersion());
    }


    public sealed class GetKeystoreArgs : Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("keystoreId", required: true)]
        public string KeystoreId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetKeystoreArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeystoreResult
    {
        /// <summary>
        /// Aliases in this keystore.
        /// </summary>
        public readonly ImmutableArray<string> Aliases;
        /// <summary>
        /// Resource ID for this keystore. Values must match the regular expression `[\w[:space:]-.]{1,255}`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetKeystoreResult(
            ImmutableArray<string> aliases,

            string name)
        {
            Aliases = aliases;
            Name = name;
        }
    }
}
