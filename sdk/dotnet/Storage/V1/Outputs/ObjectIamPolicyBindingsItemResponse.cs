// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Outputs
{

    [OutputType]
    public sealed class ObjectIamPolicyBindingsItemResponse
    {
        /// <summary>
        /// The condition that is associated with this binding. NOTE: an unsatisfied condition will not allow user access via current binding. Different bindings, including their conditions, are examined independently.
        /// </summary>
        public readonly Outputs.ExprResponse Condition;
        /// <summary>
        /// A collection of identifiers for members who may assume the provided role. Recognized identifiers are as follows:  
        /// - allUsers — A special identifier that represents anyone on the internet; with or without a Google account.  
        /// - allAuthenticatedUsers — A special identifier that represents anyone who is authenticated with a Google account or a service account.  
        /// - user:emailid — An email address that represents a specific account. For example, user:alice@gmail.com or user:joe@example.com.  
        /// - serviceAccount:emailid — An email address that represents a service account. For example,  serviceAccount:my-other-app@appspot.gserviceaccount.com .  
        /// - group:emailid — An email address that represents a Google group. For example, group:admins@example.com.  
        /// - domain:domain — A Google Apps domain name that represents all the users of that domain. For example, domain:google.com or domain:example.com.  
        /// - projectOwner:projectid — Owners of the given project. For example, projectOwner:my-example-project  
        /// - projectEditor:projectid — Editors of the given project. For example, projectEditor:my-example-project  
        /// - projectViewer:projectid — Viewers of the given project. For example, projectViewer:my-example-project
        /// </summary>
        public readonly ImmutableArray<string> Members;
        /// <summary>
        /// The role to which members belong. Two types of roles are supported: new IAM roles, which grant permissions that do not map directly to those provided by ACLs, and legacy IAM roles, which do map directly to ACL permissions. All roles are of the format roles/storage.specificRole.
        /// The new IAM roles are:  
        /// - roles/storage.admin — Full control of Google Cloud Storage resources.  
        /// - roles/storage.objectViewer — Read-Only access to Google Cloud Storage objects.  
        /// - roles/storage.objectCreator — Access to create objects in Google Cloud Storage.  
        /// - roles/storage.objectAdmin — Full control of Google Cloud Storage objects.   The legacy IAM roles are:  
        /// - roles/storage.legacyObjectReader — Read-only access to objects without listing. Equivalent to an ACL entry on an object with the READER role.  
        /// - roles/storage.legacyObjectOwner — Read/write access to existing objects without listing. Equivalent to an ACL entry on an object with the OWNER role.  
        /// - roles/storage.legacyBucketReader — Read access to buckets with object listing. Equivalent to an ACL entry on a bucket with the READER role.  
        /// - roles/storage.legacyBucketWriter — Read access to buckets with object listing/creation/deletion. Equivalent to an ACL entry on a bucket with the WRITER role.  
        /// - roles/storage.legacyBucketOwner — Read and write access to existing buckets with object listing/creation/deletion. Equivalent to an ACL entry on a bucket with the OWNER role.
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private ObjectIamPolicyBindingsItemResponse(
            Outputs.ExprResponse condition,

            ImmutableArray<string> members,

            string role)
        {
            Condition = condition;
            Members = members;
            Role = role;
        }
    }
}
