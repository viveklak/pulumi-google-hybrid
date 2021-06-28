// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetTargetSslProxy
    {
        /// <summary>
        /// Returns the specified TargetSslProxy resource. Gets a list of available target SSL proxies by making a list() request.
        /// </summary>
        public static Task<GetTargetSslProxyResult> InvokeAsync(GetTargetSslProxyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetSslProxyResult>("google-native:compute/beta:getTargetSslProxy", args ?? new GetTargetSslProxyArgs(), options.WithVersion());
    }


    public sealed class GetTargetSslProxyArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("targetSslProxy", required: true)]
        public string TargetSslProxy { get; set; } = null!;

        public GetTargetSslProxyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetSslProxyResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Type of the resource. Always compute#targetSslProxy for target SSL proxies.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
        /// </summary>
        public readonly string ProxyHeader;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// URL to the BackendService resource.
        /// </summary>
        public readonly string Service;
        /// <summary>
        /// URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates.
        /// </summary>
        public readonly ImmutableArray<string> SslCertificates;
        /// <summary>
        /// URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
        /// </summary>
        public readonly string SslPolicy;

        [OutputConstructor]
        private GetTargetSslProxyResult(
            string creationTimestamp,

            string description,

            string kind,

            string name,

            string proxyHeader,

            string selfLink,

            string service,

            ImmutableArray<string> sslCertificates,

            string sslPolicy)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            Name = name;
            ProxyHeader = proxyHeader;
            SelfLink = selfLink;
            Service = service;
            SslCertificates = sslCertificates;
            SslPolicy = sslPolicy;
        }
    }
}
