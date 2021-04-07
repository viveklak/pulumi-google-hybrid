// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.DeploymentManager.V2Beta.Inputs
{

    /// <summary>
    /// Represents an Operation resource. Google Compute Engine has three Operation resources: * [Global](/compute/docs/reference/rest/{$api_version}/globalOperations) * [Regional](/compute/docs/reference/rest/{$api_version}/regionOperations) * [Zonal](/compute/docs/reference/rest/{$api_version}/zoneOperations) You can use an operation resource to manage asynchronous API requests. For more information, read Handling API responses. Operations can be global, regional or zonal. - For global operations, use the `globalOperations` resource. - For regional operations, use the `regionOperations` resource. - For zonal operations, use the `zonalOperations` resource. For more information, read Global, Regional, and Zonal Resources.
    /// </summary>
    public sealed class OperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] The value of `requestId` if you provided it in the request. Not present otherwise.
        /// </summary>
        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        /// <summary>
        /// [Deprecated] This field is deprecated.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// [Output Only] A textual description of the operation, which is set when the operation is created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// [Output Only] The time that this operation was completed. This value is in RFC3339 text format.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("error")]
        private InputMap<string>? _error;

        /// <summary>
        /// [Output Only] If errors are generated during processing of the operation, this field will be populated.
        /// </summary>
        public InputMap<string> Error
        {
            get => _error ?? (_error = new InputMap<string>());
            set => _error = value;
        }

        /// <summary>
        /// [Output Only] If the operation fails, this field contains the HTTP error message that was returned, such as `NOT FOUND`.
        /// </summary>
        [Input("httpErrorMessage")]
        public Input<string>? HttpErrorMessage { get; set; }

        /// <summary>
        /// [Output Only] If the operation fails, this field contains the HTTP error status code that was returned. For example, a `404` means the resource was not found.
        /// </summary>
        [Input("httpErrorStatusCode")]
        public Input<int>? HttpErrorStatusCode { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the operation. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] The time that this operation was requested. This value is in RFC3339 text format.
        /// </summary>
        [Input("insertTime")]
        public Input<string>? InsertTime { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always `compute#operation` for Operation resources.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// [Output Only] Name of the operation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [Output Only] An ID that represents a group of operations, such as when a group of operations results from a `bulkInsert` API request.
        /// </summary>
        [Input("operationGroupId")]
        public Input<string>? OperationGroupId { get; set; }

        /// <summary>
        /// [Output Only] The type of operation, such as `insert`, `update`, or `delete`, and so on.
        /// </summary>
        [Input("operationType")]
        public Input<string>? OperationType { get; set; }

        /// <summary>
        /// [Output Only] An optional progress indicator that ranges from 0 to 100. There is no requirement that this be linear or support any granularity of operations. This should not be used to guess when the operation will be complete. This number should monotonically increase as the operation progresses.
        /// </summary>
        [Input("progress")]
        public Input<int>? Progress { get; set; }

        /// <summary>
        /// [Output Only] The URL of the region where the operation resides. Only applicable when performing regional operations.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] The time that this operation was started by the server. This value is in RFC3339 text format.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// [Output Only] The status of the operation, which can be one of the following: `PENDING`, `RUNNING`, or `DONE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// [Output Only] An optional textual description of the current status of the operation.
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        /// <summary>
        /// [Output Only] The unique target ID, which identifies a specific incarnation of the target resource.
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        /// <summary>
        /// [Output Only] The URL of the resource that the operation modifies. For operations related to creating a snapshot, this points to the persistent disk that the snapshot was created from.
        /// </summary>
        [Input("targetLink")]
        public Input<string>? TargetLink { get; set; }

        /// <summary>
        /// [Output Only] User who requested the operation, for example: `user@example.com`.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        [Input("warnings")]
        private InputList<ImmutableDictionary<string, string>>? _warnings;

        /// <summary>
        /// [Output Only] If warning messages are generated during processing of the operation, this field will be populated.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Warnings
        {
            get => _warnings ?? (_warnings = new InputList<ImmutableDictionary<string, string>>());
            set => _warnings = value;
        }

        /// <summary>
        /// [Output Only] The URL of the zone where the operation resides. Only applicable when performing per-zone operations.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public OperationArgs()
        {
        }
    }
}