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
    /// Create a License resource in the specified project. *Caution* This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:License")]
    public partial class License : Pulumi.CustomResource
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of resource. Always compute#license for licenses.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The unique code used to attach this license to images, snapshots, and disks.
        /// </summary>
        [Output("licenseCode")]
        public Output<string> LicenseCode { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("resourceRequirements")]
        public Output<Outputs.LicenseResourceRequirementsResponse> ResourceRequirements { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
        /// </summary>
        [Output("transferable")]
        public Output<bool> Transferable { get; private set; } = null!;


        /// <summary>
        /// Create a License resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public License(string name, LicenseArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:License", name, args ?? new LicenseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private License(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:License", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing License resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static License Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new License(name, id, options);
        }
    }

    public sealed class LicenseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] Type of resource. Always compute#license for licenses.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// [Output Only] The unique code used to attach this license to images, snapshots, and disks.
        /// </summary>
        [Input("licenseCode")]
        public Input<string>? LicenseCode { get; set; }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("resourceRequirements")]
        public Input<Inputs.LicenseResourceRequirementsArgs>? ResourceRequirements { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
        /// </summary>
        [Input("transferable")]
        public Input<bool>? Transferable { get; set; }

        public LicenseArgs()
        {
        }
    }
}
