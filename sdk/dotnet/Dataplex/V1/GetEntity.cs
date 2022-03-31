// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1
{
    public static class GetEntity
    {
        /// <summary>
        /// Get a metadata entity.
        /// </summary>
        public static Task<GetEntityResult> InvokeAsync(GetEntityArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntityResult>("google-native:dataplex/v1:getEntity", args ?? new GetEntityArgs(), options.WithDefaults());

        /// <summary>
        /// Get a metadata entity.
        /// </summary>
        public static Output<GetEntityResult> Invoke(GetEntityInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEntityResult>("google-native:dataplex/v1:getEntity", args ?? new GetEntityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEntityArgs : Pulumi.InvokeArgs
    {
        [Input("entityId", required: true)]
        public string EntityId { get; set; } = null!;

        [Input("lakeId", required: true)]
        public string LakeId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetEntityArgs()
        {
        }
    }

    public sealed class GetEntityInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        [Input("lakeId", required: true)]
        public Input<string> LakeId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetEntityInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEntityResult
    {
        /// <summary>
        /// Immutable. The ID of the asset associated with the storage location containing the entity data. The entity must be with in the same zone with the asset.
        /// </summary>
        public readonly string Asset;
        /// <summary>
        /// The name of the associated Data Catalog entry.
        /// </summary>
        public readonly string CatalogEntry;
        /// <summary>
        /// Metadata stores that the entity is compatible with.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1EntityCompatibilityStatusResponse Compatibility;
        /// <summary>
        /// The time when the entity was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Immutable. The storage path of the entity data. For Cloud Storage data, this is the fully-qualified path to the entity, such as gs://bucket/path/to/data. For BigQuery data, this is the name of the table resource, such as projects/project_id/datasets/dataset_id/tables/table_id.
        /// </summary>
        public readonly string DataPath;
        /// <summary>
        /// Optional. The set of items within the data path constituting the data in the entity, represented as a glob path. Example: gs://bucket/path/to/data/**/*.csv.
        /// </summary>
        public readonly string DataPathPattern;
        /// <summary>
        /// Optional. User friendly longer description text. Must be shorter than or equal to 1024 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Display name must be shorter than or equal to 256 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. The etag associated with the entity, which can be retrieved with a GetEntity request. Required for update and delete requests.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Identifies the storage format of the entity data. It does not apply to entities with data stored in BigQuery.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1StorageFormatResponse Format;
        /// <summary>
        /// The resource name of the entity, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/entities/{id}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The description of the data structure and layout. The schema is not included in list responses. It is only included in SCHEMA and FULL entity views of a GetEntity response.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1SchemaResponse Schema;
        /// <summary>
        /// Immutable. Identifies the storage system of the entity data.
        /// </summary>
        public readonly string System;
        /// <summary>
        /// Immutable. The type of entity.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The time when the entity was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetEntityResult(
            string asset,

            string catalogEntry,

            Outputs.GoogleCloudDataplexV1EntityCompatibilityStatusResponse compatibility,

            string createTime,

            string dataPath,

            string dataPathPattern,

            string description,

            string displayName,

            string etag,

            Outputs.GoogleCloudDataplexV1StorageFormatResponse format,

            string name,

            Outputs.GoogleCloudDataplexV1SchemaResponse schema,

            string system,

            string type,

            string updateTime)
        {
            Asset = asset;
            CatalogEntry = catalogEntry;
            Compatibility = compatibility;
            CreateTime = createTime;
            DataPath = dataPath;
            DataPathPattern = dataPathPattern;
            Description = description;
            DisplayName = displayName;
            Etag = etag;
            Format = format;
            Name = name;
            Schema = schema;
            System = system;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
