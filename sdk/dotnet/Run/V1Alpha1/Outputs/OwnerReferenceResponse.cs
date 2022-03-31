// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1Alpha1.Outputs
{

    /// <summary>
    /// OwnerReference contains enough information to let you identify an owning object. Currently, an owning object must be in the same namespace, so there is no namespace field.
    /// </summary>
    [OutputType]
    public sealed class OwnerReferenceResponse
    {
        /// <summary>
        /// API version of the referent.
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned. +optional
        /// </summary>
        public readonly bool BlockOwnerDeletion;
        /// <summary>
        /// If true, this reference points to the managing controller. +optional
        /// </summary>
        public readonly bool Controller;
        /// <summary>
        /// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the referent. More info: https://kubernetes.io/docs/user-guide/identifiers#names
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// UID of the referent. More info: https://kubernetes.io/docs/user-guide/identifiers#uids
        /// </summary>
        public readonly string Uid;

        [OutputConstructor]
        private OwnerReferenceResponse(
            string apiVersion,

            bool blockOwnerDeletion,

            bool controller,

            string kind,

            string name,

            string uid)
        {
            ApiVersion = apiVersion;
            BlockOwnerDeletion = blockOwnerDeletion;
            Controller = controller;
            Kind = kind;
            Name = name;
            Uid = uid;
        }
    }
}
