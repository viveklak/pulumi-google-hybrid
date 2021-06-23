// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetRegionInstantSnapshot
    {
        /// <summary>
        /// Returns the specified InstantSnapshot resource in the specified region.
        /// </summary>
        public static Task<GetRegionInstantSnapshotResult> InvokeAsync(GetRegionInstantSnapshotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionInstantSnapshotResult>("google-native:compute/alpha:getRegionInstantSnapshot", args ?? new GetRegionInstantSnapshotArgs(), options.WithVersion());
    }


    public sealed class GetRegionInstantSnapshotArgs : Pulumi.InvokeArgs
    {
        [Input("instantSnapshot", required: true)]
        public string InstantSnapshot { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetRegionInstantSnapshotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionInstantSnapshotResult
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// [Output Only] Size of the source disk, specified in GB.
        /// </summary>
        public readonly string DiskSizeGb;
        /// <summary>
        /// Whether to attempt an application consistent instant snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
        /// </summary>
        public readonly bool GuestFlush;
        /// <summary>
        /// [Output Only] Type of the resource. Always compute#instantSnapshot for InstantSnapshot resources.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for the labels being applied to this InstantSnapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a InstantSnapshot.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels to apply to this InstantSnapshot. These can be later modified by the setLabels method. Label values may be empty.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// [Output Only] URL of the region where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// [Output Only] Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Output Only] Server-defined URL for this resource's resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// URL of the source disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk 
        /// </summary>
        public readonly string SourceDisk;
        /// <summary>
        /// [Output Only] The ID value of the disk used to create this InstantSnapshot. This value may be used to determine whether the InstantSnapshot was taken from the current or a previous instance of a given disk name.
        /// </summary>
        public readonly string SourceDiskId;
        /// <summary>
        /// [Output Only] The status of the instantSnapshot. This can be CREATING, DELETING, FAILED, or READY.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// [Output Only] URL of the zone where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetRegionInstantSnapshotResult(
            string creationTimestamp,

            string description,

            string diskSizeGb,

            bool guestFlush,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string name,

            string region,

            bool satisfiesPzs,

            string selfLink,

            string selfLinkWithId,

            string sourceDisk,

            string sourceDiskId,

            string status,

            string zone)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            DiskSizeGb = diskSizeGb;
            GuestFlush = guestFlush;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            Name = name;
            Region = region;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            SourceDisk = sourceDisk;
            SourceDiskId = sourceDiskId;
            Status = status;
            Zone = zone;
        }
    }
}
