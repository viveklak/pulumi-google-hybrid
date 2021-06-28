// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates an instant snapshot in the specified zone.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:ZoneInstantSnapshot")]
    public partial class ZoneInstantSnapshot : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Size of the source disk, specified in GB.
        /// </summary>
        [Output("diskSizeGb")]
        public Output<string> DiskSizeGb { get; private set; } = null!;

        /// <summary>
        /// Whether to attempt an application consistent instant snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
        /// </summary>
        [Output("guestFlush")]
        public Output<bool> GuestFlush { get; private set; } = null!;

        /// <summary>
        /// Type of the resource. Always compute#instantSnapshot for InstantSnapshot resources.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for the labels being applied to this InstantSnapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a InstantSnapshot.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this InstantSnapshot. These can be later modified by the setLabels method. Label values may be empty.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for this resource's resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// URL of the source disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:  
        /// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk  
        /// - https://www.googleapis.com/compute/v1/projects/project/regions/region/disks/disk  
        /// - projects/project/zones/zone/disks/disk  
        /// - projects/project/regions/region/disks/disk  
        /// - zones/zone/disks/disk  
        /// - regions/region/disks/disk
        /// </summary>
        [Output("sourceDisk")]
        public Output<string> SourceDisk { get; private set; } = null!;

        /// <summary>
        /// The ID value of the disk used to create this InstantSnapshot. This value may be used to determine whether the InstantSnapshot was taken from the current or a previous instance of a given disk name.
        /// </summary>
        [Output("sourceDiskId")]
        public Output<string> SourceDiskId { get; private set; } = null!;

        /// <summary>
        /// The status of the instantSnapshot. This can be CREATING, DELETING, FAILED, or READY.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// URL of the zone where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a ZoneInstantSnapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZoneInstantSnapshot(string name, ZoneInstantSnapshotArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:ZoneInstantSnapshot", name, args ?? new ZoneInstantSnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZoneInstantSnapshot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:ZoneInstantSnapshot", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ZoneInstantSnapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZoneInstantSnapshot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ZoneInstantSnapshot(name, id, options);
        }
    }

    public sealed class ZoneInstantSnapshotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to attempt an application consistent instant snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
        /// </summary>
        [Input("guestFlush")]
        public Input<bool>? GuestFlush { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this InstantSnapshot. These can be later modified by the setLabels method. Label values may be empty.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

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
        /// URL of the source disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:  
        /// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk  
        /// - https://www.googleapis.com/compute/v1/projects/project/regions/region/disks/disk  
        /// - projects/project/zones/zone/disks/disk  
        /// - projects/project/regions/region/disks/disk  
        /// - zones/zone/disks/disk  
        /// - regions/region/disks/disk
        /// </summary>
        [Input("sourceDisk")]
        public Input<string>? SourceDisk { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ZoneInstantSnapshotArgs()
        {
        }
    }
}
