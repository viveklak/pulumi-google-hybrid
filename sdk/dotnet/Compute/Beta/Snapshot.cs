// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    /// <summary>
    /// Creates a snapshot in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:Snapshot")]
    public partial class Snapshot : Pulumi.CustomResource
    {
        /// <summary>
        /// [Output Only] Set to true if snapshots are automatically created by applying resource policy on the target disk.
        /// </summary>
        [Output("autoCreated")]
        public Output<bool> AutoCreated { get; private set; } = null!;

        /// <summary>
        /// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
        /// </summary>
        [Output("chainName")]
        public Output<string> ChainName { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Size of the source disk, specified in GB.
        /// </summary>
        [Output("diskSizeGb")]
        public Output<string> DiskSizeGb { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Number of bytes downloaded to restore a snapshot to a disk.
        /// </summary>
        [Output("downloadBytes")]
        public Output<string> DownloadBytes { get; private set; } = null!;

        /// <summary>
        /// [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
        /// </summary>
        [Output("guestFlush")]
        public Output<bool> GuestFlush { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#snapshot for Snapshot resources.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a snapshot.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Integer license codes indicating which licenses are attached to this snapshot.
        /// </summary>
        [Output("licenseCodes")]
        public Output<ImmutableArray<string>> LicenseCodes { get; private set; } = null!;

        /// <summary>
        /// [Output Only] A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image).
        /// </summary>
        [Output("licenses")]
        public Output<ImmutableArray<string>> Licenses { get; private set; } = null!;

        /// <summary>
        /// An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        [Output("locationHint")]
        public Output<string> LocationHint { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Encrypts the snapshot using a customer-supplied encryption key.
        /// 
        /// After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request.
        /// 
        /// Customer-supplied encryption keys do not protect access to metadata of the snapshot.
        /// 
        /// If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
        /// </summary>
        [Output("snapshotEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> SnapshotEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The source disk used to create this snapshot.
        /// </summary>
        [Output("sourceDisk")]
        public Output<string> SourceDisk { get; private set; } = null!;

        /// <summary>
        /// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        [Output("sourceDiskEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> SourceDiskEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The ID value of the disk used to create this snapshot. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given disk name.
        /// </summary>
        [Output("sourceDiskId")]
        public Output<string> SourceDiskId { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The status of the snapshot. This can be CREATING, DELETING, FAILED, READY, or UPLOADING.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// [Output Only] A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
        /// </summary>
        [Output("storageBytes")]
        public Output<string> StorageBytes { get; private set; } = null!;

        /// <summary>
        /// [Output Only] An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation. This status can either be UPDATING, meaning the size of the snapshot is being updated, or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
        /// </summary>
        [Output("storageBytesStatus")]
        public Output<string> StorageBytesStatus { get; private set; } = null!;

        /// <summary>
        /// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
        /// </summary>
        [Output("storageLocations")]
        public Output<ImmutableArray<string>> StorageLocations { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:Snapshot", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, options);
        }
    }

    public sealed class SnapshotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] Set to true if snapshots are automatically created by applying resource policy on the target disk.
        /// </summary>
        [Input("autoCreated")]
        public Input<bool>? AutoCreated { get; set; }

        /// <summary>
        /// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
        /// </summary>
        [Input("chainName")]
        public Input<string>? ChainName { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// [Output Only] Size of the source disk, specified in GB.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<string>? DiskSizeGb { get; set; }

        /// <summary>
        /// [Output Only] Number of bytes downloaded to restore a snapshot to a disk.
        /// </summary>
        [Input("downloadBytes")]
        public Input<string>? DownloadBytes { get; set; }

        /// <summary>
        /// [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
        /// </summary>
        [Input("guestFlush")]
        public Input<bool>? GuestFlush { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#snapshot for Snapshot resources.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a snapshot.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("licenseCodes")]
        private InputList<string>? _licenseCodes;

        /// <summary>
        /// [Output Only] Integer license codes indicating which licenses are attached to this snapshot.
        /// </summary>
        public InputList<string> LicenseCodes
        {
            get => _licenseCodes ?? (_licenseCodes = new InputList<string>());
            set => _licenseCodes = value;
        }

        [Input("licenses")]
        private InputList<string>? _licenses;

        /// <summary>
        /// [Output Only] A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image).
        /// </summary>
        public InputList<string> Licenses
        {
            get => _licenses ?? (_licenses = new InputList<string>());
            set => _licenses = value;
        }

        /// <summary>
        /// An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        [Input("locationHint")]
        public Input<string>? LocationHint { get; set; }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Reserved for future use.
        /// </summary>
        [Input("satisfiesPzs")]
        public Input<bool>? SatisfiesPzs { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Encrypts the snapshot using a customer-supplied encryption key.
        /// 
        /// After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request.
        /// 
        /// Customer-supplied encryption keys do not protect access to metadata of the snapshot.
        /// 
        /// If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
        /// </summary>
        [Input("snapshotEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SnapshotEncryptionKey { get; set; }

        /// <summary>
        /// The source disk used to create this snapshot.
        /// </summary>
        [Input("sourceDisk")]
        public Input<string>? SourceDisk { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        [Input("sourceDiskEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceDiskEncryptionKey { get; set; }

        /// <summary>
        /// [Output Only] The ID value of the disk used to create this snapshot. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given disk name.
        /// </summary>
        [Input("sourceDiskId")]
        public Input<string>? SourceDiskId { get; set; }

        /// <summary>
        /// [Output Only] The status of the snapshot. This can be CREATING, DELETING, FAILED, READY, or UPLOADING.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// [Output Only] A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
        /// </summary>
        [Input("storageBytes")]
        public Input<string>? StorageBytes { get; set; }

        /// <summary>
        /// [Output Only] An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation. This status can either be UPDATING, meaning the size of the snapshot is being updated, or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
        /// </summary>
        [Input("storageBytesStatus")]
        public Input<string>? StorageBytesStatus { get; set; }

        [Input("storageLocations")]
        private InputList<string>? _storageLocations;

        /// <summary>
        /// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
        /// </summary>
        public InputList<string> StorageLocations
        {
            get => _storageLocations ?? (_storageLocations = new InputList<string>());
            set => _storageLocations = value;
        }

        public SnapshotArgs()
        {
        }
    }
}
