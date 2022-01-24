// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Firestore.V1
{
    public static class GetIndex
    {
        /// <summary>
        /// Gets a composite index.
        /// </summary>
        public static Task<GetIndexResult> InvokeAsync(GetIndexArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIndexResult>("google-native:firestore/v1:getIndex", args ?? new GetIndexArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a composite index.
        /// </summary>
        public static Output<GetIndexResult> Invoke(GetIndexInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIndexResult>("google-native:firestore/v1:getIndex", args ?? new GetIndexInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIndexArgs : Pulumi.InvokeArgs
    {
        [Input("collectionGroupId", required: true)]
        public string CollectionGroupId { get; set; } = null!;

        [Input("databaseId", required: true)]
        public string DatabaseId { get; set; } = null!;

        [Input("indexId", required: true)]
        public string IndexId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetIndexArgs()
        {
        }
    }

    public sealed class GetIndexInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("collectionGroupId", required: true)]
        public Input<string> CollectionGroupId { get; set; } = null!;

        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        [Input("indexId", required: true)]
        public Input<string> IndexId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetIndexInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIndexResult
    {
        /// <summary>
        /// The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleFirestoreAdminV1IndexFieldResponse> Fields;
        /// <summary>
        /// A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        /// </summary>
        public readonly string QueryScope;
        /// <summary>
        /// The serving state of the index.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetIndexResult(
            ImmutableArray<Outputs.GoogleFirestoreAdminV1IndexFieldResponse> fields,

            string name,

            string queryScope,

            string state)
        {
            Fields = fields;
            Name = name;
            QueryScope = queryScope;
            State = state;
        }
    }
}
