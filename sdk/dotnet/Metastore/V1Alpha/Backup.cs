// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Alpha
{
    /// <summary>
    /// Creates a new backup in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:metastore/v1alpha:Backup")]
    public partial class Backup : Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the backup was started.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the backup.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The time when the backup finished creating.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Immutable. The relative resource name of the backup, in the following form:projects/{project_number}/locations/{location_id}/services/{service_id}/backups/{backup_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Services that are restoring from the backup.
        /// </summary>
        [Output("restoringServices")]
        public Output<ImmutableArray<string>> RestoringServices { get; private set; } = null!;

        /// <summary>
        /// The revision of the service at the time of backup.
        /// </summary>
        [Output("serviceRevision")]
        public Output<Outputs.ServiceResponse> ServiceRevision { get; private set; } = null!;

        /// <summary>
        /// The current state of the backup.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("google-native:metastore/v1alpha:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:metastore/v1alpha:Backup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, options);
        }
    }

    public sealed class BackupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The ID of the backup, which is used as the final component of the backup's name.This value must be between 1 and 64 characters long, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
        /// </summary>
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        /// <summary>
        /// The description of the backup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The relative resource name of the backup, in the following form:projects/{project_number}/locations/{location_id}/services/{service_id}/backups/{backup_id}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public BackupArgs()
        {
        }
    }
}
