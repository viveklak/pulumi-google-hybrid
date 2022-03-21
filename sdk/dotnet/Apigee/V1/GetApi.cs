// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetApi
    {
        /// <summary>
        /// Gets an API proxy including a list of existing revisions.
        /// </summary>
        public static Task<GetApiResult> InvokeAsync(GetApiArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiResult>("google-native:apigee/v1:getApi", args ?? new GetApiArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an API proxy including a list of existing revisions.
        /// </summary>
        public static Output<GetApiResult> Invoke(GetApiInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApiResult>("google-native:apigee/v1:getApi", args ?? new GetApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiArgs : Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetApiArgs()
        {
        }
    }

    public sealed class GetApiInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetApiInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiResult
    {
        /// <summary>
        /// The type of the API proxy.
        /// </summary>
        public readonly string ApiProxyType;
        /// <summary>
        /// User labels applied to this API Proxy.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The id of the most recently created revision for this api proxy.
        /// </summary>
        public readonly string LatestRevisionId;
        /// <summary>
        /// Metadata describing the API proxy.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1EntityMetadataResponse MetaData;
        /// <summary>
        /// Name of the API proxy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether this proxy is read-only. A read-only proxy cannot have new revisions created through calls to CreateApiProxyRevision. A proxy is read-only if it was generated by an archive.
        /// </summary>
        public readonly bool ReadOnly;
        /// <summary>
        /// List of revisons defined for the API proxy.
        /// </summary>
        public readonly ImmutableArray<string> Revision;

        [OutputConstructor]
        private GetApiResult(
            string apiProxyType,

            ImmutableDictionary<string, string> labels,

            string latestRevisionId,

            Outputs.GoogleCloudApigeeV1EntityMetadataResponse metaData,

            string name,

            bool readOnly,

            ImmutableArray<string> revision)
        {
            ApiProxyType = apiProxyType;
            Labels = labels;
            LatestRevisionId = latestRevisionId;
            MetaData = metaData;
            Name = name;
            ReadOnly = readOnly;
            Revision = revision;
        }
    }
}
