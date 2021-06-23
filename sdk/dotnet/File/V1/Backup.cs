// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.File.V1
{
    /// <summary>
    /// Creates a backup.
    /// </summary>
    [GoogleNativeResourceType("google-native:file/v1:Backup")]
    public partial class Backup : Pulumi.CustomResource
    {
        /// <summary>
        /// Capacity of the source file share when the backup was created.
        /// </summary>
        [Output("capacityGb")]
        public Output<string> CapacityGb { get; private set; } = null!;

        /// <summary>
        /// The time when the backup was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Amount of bytes that will be downloaded if the backup is restored. This may be different than storage bytes, since sequential backups of the same disk will share storage.
        /// </summary>
        [Output("downloadBytes")]
        public Output<string> DownloadBytes { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name of the backup, in the format projects/{project_number}/locations/{location_id}/backups/{backup_id}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// Name of the file share in the source Cloud Filestore instance that the backup is created from.
        /// </summary>
        [Output("sourceFileShare")]
        public Output<string> SourceFileShare { get; private set; } = null!;

        /// <summary>
        /// The resource name of the source Cloud Filestore instance, in the format projects/{project_number}/locations/{location_id}/instances/{instance_id}, used to create this backup.
        /// </summary>
        [Output("sourceInstance")]
        public Output<string> SourceInstance { get; private set; } = null!;

        /// <summary>
        /// The service tier of the source Cloud Filestore instance that this backup is created from.
        /// </summary>
        [Output("sourceInstanceTier")]
        public Output<string> SourceInstanceTier { get; private set; } = null!;

        /// <summary>
        /// The backup state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
        /// </summary>
        [Output("storageBytes")]
        public Output<string> StorageBytes { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("google-native:file/v1:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:file/v1:Backup", name, null, MakeResourceOptions(options, id))
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
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        /// <summary>
        /// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Name of the file share in the source Cloud Filestore instance that the backup is created from.
        /// </summary>
        [Input("sourceFileShare")]
        public Input<string>? SourceFileShare { get; set; }

        /// <summary>
        /// The resource name of the source Cloud Filestore instance, in the format projects/{project_number}/locations/{location_id}/instances/{instance_id}, used to create this backup.
        /// </summary>
        [Input("sourceInstance")]
        public Input<string>? SourceInstance { get; set; }

        public BackupArgs()
        {
        }
    }
}
