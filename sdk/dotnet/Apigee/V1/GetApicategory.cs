// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetApicategory
    {
        /// <summary>
        /// Gets a category on the portal.
        /// </summary>
        public static Task<GetApicategoryResult> InvokeAsync(GetApicategoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApicategoryResult>("google-native:apigee/v1:getApicategory", args ?? new GetApicategoryArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a category on the portal.
        /// </summary>
        public static Output<GetApicategoryResult> Invoke(GetApicategoryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApicategoryResult>("google-native:apigee/v1:getApicategory", args ?? new GetApicategoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApicategoryArgs : Pulumi.InvokeArgs
    {
        [Input("apicategoryId", required: true)]
        public string ApicategoryId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("siteId", required: true)]
        public string SiteId { get; set; } = null!;

        public GetApicategoryArgs()
        {
        }
    }

    public sealed class GetApicategoryInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("apicategoryId", required: true)]
        public Input<string> ApicategoryId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        public GetApicategoryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApicategoryResult
    {
        /// <summary>
        /// Details of category.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1ApiCategoryDataResponse Data;
        /// <summary>
        /// ID that can be used to find errors in the log files.
        /// </summary>
        public readonly string ErrorCode;
        /// <summary>
        /// Description of the operation.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// ID that can be used to find request details in the log files.
        /// </summary>
        public readonly string RequestId;
        /// <summary>
        /// Status of the operation.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetApicategoryResult(
            Outputs.GoogleCloudApigeeV1ApiCategoryDataResponse data,

            string errorCode,

            string message,

            string requestId,

            string status)
        {
            Data = data;
            ErrorCode = errorCode;
            Message = message;
            RequestId = requestId;
            Status = status;
        }
    }
}
