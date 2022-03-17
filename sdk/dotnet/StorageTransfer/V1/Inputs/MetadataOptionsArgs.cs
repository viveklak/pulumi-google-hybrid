// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Inputs
{

    /// <summary>
    /// Specifies the metadata options for running a transfer. These options only apply to transfers involving a POSIX filesystem and are ignored for other transfers.
    /// </summary>
    public sealed class MetadataOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how each object's ACLs should be preserved for transfers between Google Cloud Storage buckets. If unspecified, the default behavior is the same as ACL_DESTINATION_BUCKET_DEFAULT.
        /// </summary>
        [Input("acl")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsAcl>? Acl { get; set; }

        /// <summary>
        /// Specifies how each file's POSIX group ID (GID) attribute should be handled by the transfer. By default, GID is not preserved.
        /// </summary>
        [Input("gid")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsGid>? Gid { get; set; }

        /// <summary>
        /// Specifies how each object's Cloud KMS customer-managed encryption key (CMEK) is preserved for transfers between Google Cloud Storage buckets. If unspecified, the default behavior is the same as KMS_KEY_DESTINATION_BUCKET_DEFAULT.
        /// </summary>
        [Input("kmsKey")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsKmsKey>? KmsKey { get; set; }

        /// <summary>
        /// Specifies how each file's mode attribute should be handled by the transfer. By default, mode is not preserved.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsMode>? Mode { get; set; }

        /// <summary>
        /// Specifies the storage class to set on objects being transferred to Google Cloud Storage buckets. If unspecified, the default behavior is the same as STORAGE_CLASS_DESTINATION_BUCKET_DEFAULT.
        /// </summary>
        [Input("storageClass")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsStorageClass>? StorageClass { get; set; }

        /// <summary>
        /// Specifies how symlinks should be handled by the transfer. By default, symlinks are not preserved.
        /// </summary>
        [Input("symlink")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsSymlink>? Symlink { get; set; }

        /// <summary>
        /// Specifies how each object's temporary hold status should be preserved for transfers between Google Cloud Storage buckets. If unspecified, the default behavior is the same as TEMPORARY_HOLD_PRESERVE.
        /// </summary>
        [Input("temporaryHold")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsTemporaryHold>? TemporaryHold { get; set; }

        /// <summary>
        /// Specifies how each object's `timeCreated` metadata is preserved for transfers between Google Cloud Storage buckets. If unspecified, the default behavior is the same as TIME_CREATED_SKIP.
        /// </summary>
        [Input("timeCreated")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsTimeCreated>? TimeCreated { get; set; }

        /// <summary>
        /// Specifies how each file's POSIX user ID (UID) attribute should be handled by the transfer. By default, UID is not preserved.
        /// </summary>
        [Input("uid")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.MetadataOptionsUid>? Uid { get; set; }

        public MetadataOptionsArgs()
        {
        }
    }
}
