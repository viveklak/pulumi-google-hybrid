// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a machine image in the specified project using the data that is included in the request. If you are creating a new machine image to update an existing instance, your new machine image should use the same network or, if applicable, the same subnetwork as the original instance.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:MachineImage")]
    public partial class MachineImage : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation timestamp for this machine image in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process.
        /// </summary>
        [Output("guestFlush")]
        public Output<bool> GuestFlush { get; private set; } = null!;

        /// <summary>
        /// Properties of source instance
        /// </summary>
        [Output("instanceProperties")]
        public Output<Outputs.InstancePropertiesResponse> InstanceProperties { get; private set; } = null!;

        /// <summary>
        /// The resource type, which is always compute#machineImage for machine image.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
        /// </summary>
        [Output("machineImageEncryptionKey")]
        public Output<Outputs.CustomerEncryptionKeyResponse> MachineImageEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// An array of Machine Image specific properties for disks attached to the source instance
        /// </summary>
        [Output("savedDisks")]
        public Output<ImmutableArray<Outputs.SavedDiskResponse>> SavedDisks { get; private set; } = null!;

        /// <summary>
        /// The URL for this machine image. The server defines this URL.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        [Output("sourceDiskEncryptionKeys")]
        public Output<ImmutableArray<Outputs.SourceDiskEncryptionKeyResponse>> SourceDiskEncryptionKeys { get; private set; } = null!;

        /// <summary>
        /// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
        /// </summary>
        [Output("sourceInstance")]
        public Output<string> SourceInstance { get; private set; } = null!;

        /// <summary>
        /// DEPRECATED: Please use instance_properties instead for source instance related properties. New properties will not be added to this field.
        /// </summary>
        [Output("sourceInstanceProperties")]
        public Output<Outputs.SourceInstancePropertiesResponse> SourceInstanceProperties { get; private set; } = null!;

        /// <summary>
        /// The status of the machine image. One of the following values: INVALID, CREATING, READY, DELETING, and UPLOADING.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
        /// </summary>
        [Output("storageLocations")]
        public Output<ImmutableArray<string>> StorageLocations { get; private set; } = null!;

        /// <summary>
        /// Total size of the storage used by the machine image.
        /// </summary>
        [Output("totalStorageBytes")]
        public Output<string> TotalStorageBytes { get; private set; } = null!;


        /// <summary>
        /// Create a MachineImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MachineImage(string name, MachineImageArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:MachineImage", name, args ?? new MachineImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MachineImage(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:MachineImage", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MachineImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MachineImage Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MachineImage(name, id, options);
        }
    }

    public sealed class MachineImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process.
        /// </summary>
        [Input("guestFlush")]
        public Input<bool>? GuestFlush { get; set; }

        /// <summary>
        /// Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
        /// </summary>
        [Input("machineImageEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? MachineImageEncryptionKey { get; set; }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("savedDisks")]
        private InputList<Inputs.SavedDiskArgs>? _savedDisks;

        /// <summary>
        /// An array of Machine Image specific properties for disks attached to the source instance
        /// </summary>
        public InputList<Inputs.SavedDiskArgs> SavedDisks
        {
            get => _savedDisks ?? (_savedDisks = new InputList<Inputs.SavedDiskArgs>());
            set => _savedDisks = value;
        }

        [Input("sourceDiskEncryptionKeys")]
        private InputList<Inputs.SourceDiskEncryptionKeyArgs>? _sourceDiskEncryptionKeys;

        /// <summary>
        /// [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        public InputList<Inputs.SourceDiskEncryptionKeyArgs> SourceDiskEncryptionKeys
        {
            get => _sourceDiskEncryptionKeys ?? (_sourceDiskEncryptionKeys = new InputList<Inputs.SourceDiskEncryptionKeyArgs>());
            set => _sourceDiskEncryptionKeys = value;
        }

        /// <summary>
        /// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
        /// </summary>
        [Input("sourceInstance", required: true)]
        public Input<string> SourceInstance { get; set; } = null!;

        [Input("storageLocations")]
        private InputList<string>? _storageLocations;

        /// <summary>
        /// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
        /// </summary>
        public InputList<string> StorageLocations
        {
            get => _storageLocations ?? (_storageLocations = new InputList<string>());
            set => _storageLocations = value;
        }

        public MachineImageArgs()
        {
        }
    }
}
