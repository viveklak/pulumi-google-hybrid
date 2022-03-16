// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetEndpointAttachment
    {
        /// <summary>
        /// Gets the endpoint attachment.
        /// </summary>
        public static Task<GetEndpointAttachmentResult> InvokeAsync(GetEndpointAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEndpointAttachmentResult>("google-native:apigee/v1:getEndpointAttachment", args ?? new GetEndpointAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the endpoint attachment.
        /// </summary>
        public static Output<GetEndpointAttachmentResult> Invoke(GetEndpointAttachmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEndpointAttachmentResult>("google-native:apigee/v1:getEndpointAttachment", args ?? new GetEndpointAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEndpointAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("endpointAttachmentId", required: true)]
        public string EndpointAttachmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetEndpointAttachmentArgs()
        {
        }
    }

    public sealed class GetEndpointAttachmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("endpointAttachmentId", required: true)]
        public Input<string> EndpointAttachmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetEndpointAttachmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEndpointAttachmentResult
    {
        /// <summary>
        /// Host that can be used in either the HTTP target endpoint directly or as the host in target server.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Location of the endpoint attachment.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Format: projects/*/regions/*/serviceAttachments/*
        /// </summary>
        public readonly string ServiceAttachment;

        [OutputConstructor]
        private GetEndpointAttachmentResult(
            string host,

            string location,

            string name,

            string serviceAttachment)
        {
            Host = host;
            Location = location;
            Name = name;
            ServiceAttachment = serviceAttachment;
        }
    }
}