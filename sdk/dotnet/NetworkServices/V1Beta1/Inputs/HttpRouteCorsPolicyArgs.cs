// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Inputs
{

    /// <summary>
    /// The Specification for allowing client side cross-origin requests.
    /// </summary>
    public sealed class HttpRouteCorsPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// In response to a preflight request, setting this to true indicates that the actual request can include user credentials. This translates to the Access-Control-Allow-Credentials header. Default value is false.
        /// </summary>
        [Input("allowCredentials")]
        public Input<bool>? AllowCredentials { get; set; }

        [Input("allowHeaders")]
        private InputList<string>? _allowHeaders;

        /// <summary>
        /// Specifies the content for Access-Control-Allow-Headers header.
        /// </summary>
        public InputList<string> AllowHeaders
        {
            get => _allowHeaders ?? (_allowHeaders = new InputList<string>());
            set => _allowHeaders = value;
        }

        [Input("allowMethods")]
        private InputList<string>? _allowMethods;

        /// <summary>
        /// Specifies the content for Access-Control-Allow-Methods header.
        /// </summary>
        public InputList<string> AllowMethods
        {
            get => _allowMethods ?? (_allowMethods = new InputList<string>());
            set => _allowMethods = value;
        }

        [Input("allowOriginRegexes")]
        private InputList<string>? _allowOriginRegexes;

        /// <summary>
        /// Specifies the regular expression patterns that match allowed origins. For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax.
        /// </summary>
        public InputList<string> AllowOriginRegexes
        {
            get => _allowOriginRegexes ?? (_allowOriginRegexes = new InputList<string>());
            set => _allowOriginRegexes = value;
        }

        [Input("allowOrigins")]
        private InputList<string>? _allowOrigins;

        /// <summary>
        /// Specifies the list of origins that will be allowed to do CORS requests. An origin is allowed if it matches either an item in allow_origins or an item in allow_origin_regexes.
        /// </summary>
        public InputList<string> AllowOrigins
        {
            get => _allowOrigins ?? (_allowOrigins = new InputList<string>());
            set => _allowOrigins = value;
        }

        /// <summary>
        /// If true, the CORS policy is disabled. The default value is false, which indicates that the CORS policy is in effect.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("exposeHeaders")]
        private InputList<string>? _exposeHeaders;

        /// <summary>
        /// Specifies the content for Access-Control-Expose-Headers header.
        /// </summary>
        public InputList<string> ExposeHeaders
        {
            get => _exposeHeaders ?? (_exposeHeaders = new InputList<string>());
            set => _exposeHeaders = value;
        }

        /// <summary>
        /// Specifies how long result of a preflight request can be cached in seconds. This translates to the Access-Control-Max-Age header.
        /// </summary>
        [Input("maxAge")]
        public Input<string>? MaxAge { get; set; }

        public HttpRouteCorsPolicyArgs()
        {
        }
    }
}
