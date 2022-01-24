// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetLicense
    {
        /// <summary>
        /// Returns the specified License resource. *Caution* This resource is intended for use only by third-party partners who are creating Cloud Marketplace images. 
        /// </summary>
        public static Task<GetLicenseResult> InvokeAsync(GetLicenseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLicenseResult>("google-native:compute/beta:getLicense", args ?? new GetLicenseArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified License resource. *Caution* This resource is intended for use only by third-party partners who are creating Cloud Marketplace images. 
        /// </summary>
        public static Output<GetLicenseResult> Invoke(GetLicenseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLicenseResult>("google-native:compute/beta:getLicense", args ?? new GetLicenseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLicenseArgs : Pulumi.InvokeArgs
    {
        [Input("license", required: true)]
        public string License { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetLicenseArgs()
        {
        }
    }

    public sealed class GetLicenseInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("license", required: true)]
        public Input<string> License { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetLicenseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLicenseResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Type of resource. Always compute#license for licenses.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The unique code used to attach this license to images, snapshots, and disks.
        /// </summary>
        public readonly string LicenseCode;
        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.LicenseResourceRequirementsResponse ResourceRequirements;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
        /// </summary>
        public readonly bool Transferable;

        [OutputConstructor]
        private GetLicenseResult(
            string creationTimestamp,

            string description,

            string kind,

            string licenseCode,

            string name,

            Outputs.LicenseResourceRequirementsResponse resourceRequirements,

            string selfLink,

            bool transferable)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            LicenseCode = licenseCode;
            Name = name;
            ResourceRequirements = resourceRequirements;
            SelfLink = selfLink;
            Transferable = transferable;
        }
    }
}
