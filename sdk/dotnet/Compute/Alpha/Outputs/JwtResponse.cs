// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// [Deprecated] JWT configuration for origin authentication. JWT configuration for origin authentication.
    /// </summary>
    [OutputType]
    public sealed class JwtResponse
    {
        /// <summary>
        /// A JWT containing any of these audiences will be accepted. The service name will be accepted if audiences is empty. Examples: bookstore_android.apps.googleusercontent.com, bookstore_web.apps.googleusercontent.com
        /// </summary>
        public readonly ImmutableArray<string> Audiences;
        /// <summary>
        /// Identifies the issuer that issued the JWT, which is usually a URL or an email address. Examples: https://securetoken.google.com, 1234567-compute@developer.gserviceaccount.com
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// The provider's public key set to validate the signature of the JWT.
        /// </summary>
        public readonly string JwksPublicKeys;
        /// <summary>
        /// jwt_headers and jwt_params define where to extract the JWT from an HTTP request. If no explicit location is specified, the following default locations are tried in order: 1. The Authorization header using the Bearer schema. See `here `_. Example: Authorization: Bearer . 2. `access_token` query parameter. See `this `_ Multiple JWTs can be verified for a request. Each JWT has to be extracted from the locations its issuer specified or from the default locations. This field is set if JWT is sent in a request header. This field specifies the header name. For example, if `header=x-goog-iap-jwt-assertion`, the header format will be x-goog-iap-jwt-assertion: .
        /// </summary>
        public readonly ImmutableArray<Outputs.JwtHeaderResponse> JwtHeaders;
        /// <summary>
        /// This field is set if JWT is sent in a query parameter. This field specifies the query parameter name. For example, if jwt_params[0] is jwt_token, the JWT format in the query parameter is /path?jwt_token=.
        /// </summary>
        public readonly ImmutableArray<string> JwtParams;

        [OutputConstructor]
        private JwtResponse(
            ImmutableArray<string> audiences,

            string issuer,

            string jwksPublicKeys,

            ImmutableArray<Outputs.JwtHeaderResponse> jwtHeaders,

            ImmutableArray<string> jwtParams)
        {
            Audiences = audiences;
            Issuer = issuer;
            JwksPublicKeys = jwksPublicKeys;
            JwtHeaders = jwtHeaders;
            JwtParams = jwtParams;
        }
    }
}
