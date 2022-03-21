// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1
{
    public static class GetUserDataMapping
    {
        /// <summary>
        /// Gets the specified User data mapping.
        /// </summary>
        public static Task<GetUserDataMappingResult> InvokeAsync(GetUserDataMappingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserDataMappingResult>("google-native:healthcare/v1:getUserDataMapping", args ?? new GetUserDataMappingArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified User data mapping.
        /// </summary>
        public static Output<GetUserDataMappingResult> Invoke(GetUserDataMappingInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUserDataMappingResult>("google-native:healthcare/v1:getUserDataMapping", args ?? new GetUserDataMappingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserDataMappingArgs : Pulumi.InvokeArgs
    {
        [Input("consentStoreId", required: true)]
        public string ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("userDataMappingId", required: true)]
        public string UserDataMappingId { get; set; } = null!;

        public GetUserDataMappingArgs()
        {
        }
    }

    public sealed class GetUserDataMappingInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("consentStoreId", required: true)]
        public Input<string> ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("userDataMappingId", required: true)]
        public Input<string> UserDataMappingId { get; set; } = null!;

        public GetUserDataMappingInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserDataMappingResult
    {
        /// <summary>
        /// Indicates the time when this mapping was archived.
        /// </summary>
        public readonly string ArchiveTime;
        /// <summary>
        /// Indicates whether this mapping is archived.
        /// </summary>
        public readonly bool Archived;
        /// <summary>
        /// A unique identifier for the mapped resource.
        /// </summary>
        public readonly string DataId;
        /// <summary>
        /// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
        /// </summary>
        public readonly ImmutableArray<Outputs.AttributeResponse> ResourceAttributes;
        /// <summary>
        /// User's UUID provided by the client.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private GetUserDataMappingResult(
            string archiveTime,

            bool archived,

            string dataId,

            string name,

            ImmutableArray<Outputs.AttributeResponse> resourceAttributes,

            string userId)
        {
            ArchiveTime = archiveTime;
            Archived = archived;
            DataId = dataId;
            Name = name;
            ResourceAttributes = resourceAttributes;
            UserId = userId;
        }
    }
}
