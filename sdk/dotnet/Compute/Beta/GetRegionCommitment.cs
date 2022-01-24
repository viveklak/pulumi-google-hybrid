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
    public static class GetRegionCommitment
    {
        /// <summary>
        /// Returns the specified commitment resource. Gets a list of available commitments by making a list() request.
        /// </summary>
        public static Task<GetRegionCommitmentResult> InvokeAsync(GetRegionCommitmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionCommitmentResult>("google-native:compute/beta:getRegionCommitment", args ?? new GetRegionCommitmentArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified commitment resource. Gets a list of available commitments by making a list() request.
        /// </summary>
        public static Output<GetRegionCommitmentResult> Invoke(GetRegionCommitmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionCommitmentResult>("google-native:compute/beta:getRegionCommitment", args ?? new GetRegionCommitmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionCommitmentArgs : Pulumi.InvokeArgs
    {
        [Input("commitment", required: true)]
        public string Commitment { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetRegionCommitmentArgs()
        {
        }
    }

    public sealed class GetRegionCommitmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("commitment", required: true)]
        public Input<string> Commitment { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetRegionCommitmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionCommitmentResult
    {
        /// <summary>
        /// Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
        /// </summary>
        public readonly bool AutoRenew;
        /// <summary>
        /// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Commitment end time in RFC3339 text format.
        /// </summary>
        public readonly string EndTimestamp;
        /// <summary>
        /// Type of the resource. Always compute#commitment for commitments.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The license specification required as part of a license commitment.
        /// </summary>
        public readonly Outputs.LicenseResourceCommitmentResponse LicenseResource;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// URL of the region where this commitment may be used.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// List of reservations in this commitment.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReservationResponse> Reservations;
        /// <summary>
        /// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceCommitmentResponse> Resources;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Commitment start time in RFC3339 text format.
        /// </summary>
        public readonly string StartTimestamp;
        /// <summary>
        /// Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// An optional, human-readable explanation of the status.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRegionCommitmentResult(
            bool autoRenew,

            string category,

            string creationTimestamp,

            string description,

            string endTimestamp,

            string kind,

            Outputs.LicenseResourceCommitmentResponse licenseResource,

            string name,

            string plan,

            string region,

            ImmutableArray<Outputs.ReservationResponse> reservations,

            ImmutableArray<Outputs.ResourceCommitmentResponse> resources,

            string selfLink,

            string startTimestamp,

            string status,

            string statusMessage,

            string type)
        {
            AutoRenew = autoRenew;
            Category = category;
            CreationTimestamp = creationTimestamp;
            Description = description;
            EndTimestamp = endTimestamp;
            Kind = kind;
            LicenseResource = licenseResource;
            Name = name;
            Plan = plan;
            Region = region;
            Reservations = reservations;
            Resources = resources;
            SelfLink = selfLink;
            StartTimestamp = startTimestamp;
            Status = status;
            StatusMessage = statusMessage;
            Type = type;
        }
    }
}
