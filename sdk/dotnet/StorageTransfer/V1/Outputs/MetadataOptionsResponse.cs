// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Outputs
{

    /// <summary>
    /// Specifies the metadata options for running a transfer. These options only apply to transfers involving a POSIX filesystem and are ignored for other transfers.
    /// </summary>
    [OutputType]
    public sealed class MetadataOptionsResponse
    {
        /// <summary>
        /// Specifies how each object's ACLs should be preserved for transfers between Google Cloud Storage buckets. If unspecified, the default behavior is the same as ACL_DESTINATION_BUCKET_DEFAULT.
        /// </summary>
        public readonly string Acl;
        /// <summary>
        /// Specifies how each file's POSIX group ID (GID) attribute should be handled by the transfer. By default, GID is not preserved.
        /// </summary>
        public readonly string Gid;
        /// <summary>
        /// Specifies how each object's Cloud KMS customer-managed encryption key (CMEK) is preserved for transfers between Google Cloud Storage buckets. If unspecified, the default behavior is the same as KMS_KEY_DESTINATION_BUCKET_DEFAULT.
        /// </summary>
        public readonly string KmsKey;
        /// <summary>
        /// Specifies how each file's mode attribute should be handled by the transfer. By default, mode is not preserved.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// Specifies the storage class to set on objects being transferred to Google Cloud Storage buckets. If unspecified, the default behavior is the same as STORAGE_CLASS_DESTINATION_BUCKET_DEFAULT.
        /// </summary>
        public readonly string StorageClass;
        /// <summary>
        /// Specifies how symlinks should be handled by the transfer. By default, symlinks are not preserved.
        /// </summary>
        public readonly string Symlink;
        /// <summary>
        /// Specifies how each object's temporary hold status should be preserved for transfers between Google Cloud Storage buckets. If unspecified, the default behavior is the same as TEMPORARY_HOLD_PRESERVE.
        /// </summary>
        public readonly string TemporaryHold;
        /// <summary>
        /// Specifies how each object's `timeCreated` metadata is preserved for transfers between Google Cloud Storage buckets. If unspecified, the default behavior is the same as TIME_CREATED_SKIP.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// Specifies how each file's POSIX user ID (UID) attribute should be handled by the transfer. By default, UID is not preserved.
        /// </summary>
        public readonly string Uid;

        [OutputConstructor]
        private MetadataOptionsResponse(
            string acl,

            string gid,

            string kmsKey,

            string mode,

            string storageClass,

            string symlink,

            string temporaryHold,

            string timeCreated,

            string uid)
        {
            Acl = acl;
            Gid = gid;
            KmsKey = kmsKey;
            Mode = mode;
            StorageClass = storageClass;
            Symlink = symlink;
            TemporaryHold = temporaryHold;
            TimeCreated = timeCreated;
            Uid = uid;
        }
    }
}
