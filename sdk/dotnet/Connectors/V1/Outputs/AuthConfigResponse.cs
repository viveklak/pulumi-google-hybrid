// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1.Outputs
{

    /// <summary>
    /// AuthConfig defines details of a authentication type.
    /// </summary>
    [OutputType]
    public sealed class AuthConfigResponse
    {
        /// <summary>
        /// List containing additional auth configs.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigVariableResponse> AdditionalVariables;
        /// <summary>
        /// The type of authentication configured.
        /// </summary>
        public readonly string AuthType;
        /// <summary>
        /// Oauth2ClientCredentials.
        /// </summary>
        public readonly Outputs.Oauth2ClientCredentialsResponse Oauth2ClientCredentials;
        /// <summary>
        /// Oauth2JwtBearer.
        /// </summary>
        public readonly Outputs.Oauth2JwtBearerResponse Oauth2JwtBearer;
        /// <summary>
        /// UserPassword.
        /// </summary>
        public readonly Outputs.UserPasswordResponse UserPassword;

        [OutputConstructor]
        private AuthConfigResponse(
            ImmutableArray<Outputs.ConfigVariableResponse> additionalVariables,

            string authType,

            Outputs.Oauth2ClientCredentialsResponse oauth2ClientCredentials,

            Outputs.Oauth2JwtBearerResponse oauth2JwtBearer,

            Outputs.UserPasswordResponse userPassword)
        {
            AdditionalVariables = additionalVariables;
            AuthType = authType;
            Oauth2ClientCredentials = oauth2ClientCredentials;
            Oauth2JwtBearer = oauth2JwtBearer;
            UserPassword = userPassword;
        }
    }
}