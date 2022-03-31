// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Inputs
{

    /// <summary>
    /// k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
    /// </summary>
    public sealed class ObjectMetaArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// (Optional) Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/user-guide/annotations
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// (Optional) Not supported by Cloud Run The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// (Optional) CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// (Optional) Not supported by Cloud Run Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.
        /// </summary>
        [Input("deletionGracePeriodSeconds")]
        public Input<int>? DeletionGracePeriodSeconds { get; set; }

        /// <summary>
        /// (Optional) Not supported by Cloud Run DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested. Populated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
        /// </summary>
        [Input("deletionTimestamp")]
        public Input<string>? DeletionTimestamp { get; set; }

        [Input("finalizers")]
        private InputList<string>? _finalizers;

        /// <summary>
        /// (Optional) Not supported by Cloud Run Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. +patchStrategy=merge
        /// </summary>
        public InputList<string> Finalizers
        {
            get => _finalizers ?? (_finalizers = new InputList<string>());
            set => _finalizers = value;
        }

        /// <summary>
        /// (Optional) Not supported by Cloud Run GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server. If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header). Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency string generateName = 2;
        /// </summary>
        [Input("generateName")]
        public Input<string>? GenerateName { get; set; }

        /// <summary>
        /// (Optional) A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
        /// </summary>
        [Input("generation")]
        public Input<int>? Generation { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// (Optional) Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and routes. More info: https://kubernetes.io/docs/user-guide/labels
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name must be unique within a namespace, within a Cloud Run region. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/user-guide/identifiers#names +optional
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Namespace defines the space within each name must be unique, within a Cloud Run region. In Cloud Run the namespace must be equal to either the project ID or project number.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("ownerReferences")]
        private InputList<Inputs.OwnerReferenceArgs>? _ownerReferences;

        /// <summary>
        /// (Optional) Not supported by Cloud Run List of objects that own this object. If ALL objects in the list have been deleted, this object will be garbage collected.
        /// </summary>
        public InputList<Inputs.OwnerReferenceArgs> OwnerReferences
        {
            get => _ownerReferences ?? (_ownerReferences = new InputList<Inputs.OwnerReferenceArgs>());
            set => _ownerReferences = value;
        }

        /// <summary>
        /// Optional. An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server or omit the value to disable conflict-detection. They may only be valid for a particular resource or set of resources. Populated by the system. Read-only. Value must be treated as opaque by clients or omitted. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
        /// </summary>
        [Input("resourceVersion")]
        public Input<string>? ResourceVersion { get; set; }

        /// <summary>
        /// (Optional) SelfLink is a URL representing this object. Populated by the system. Read-only. string selfLink = 4;
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// (Optional) UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations. Populated by the system. Read-only. More info: https://kubernetes.io/docs/user-guide/identifiers#uids
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public ObjectMetaArgs()
        {
        }
    }
}
