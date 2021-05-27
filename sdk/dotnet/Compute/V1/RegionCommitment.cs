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
    /// Creates a commitment in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:RegionCommitment")]
    public partial class RegionCommitment : Pulumi.CustomResource
    {
        /// <summary>
        /// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Commitment end time in RFC3339 text format.
        /// </summary>
        [Output("endTimestamp")]
        public Output<string> EndTimestamp { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#commitment for commitments.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The license specification required as part of a license commitment.
        /// </summary>
        [Output("licenseResource")]
        public Output<Outputs.LicenseResourceCommitmentResponse> LicenseResource { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where this commitment may be used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// List of reservations in this commitment.
        /// </summary>
        [Output("reservations")]
        public Output<ImmutableArray<Outputs.ReservationResponse>> Reservations { get; private set; } = null!;

        /// <summary>
        /// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.ResourceCommitmentResponse>> Resources { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Commitment start time in RFC3339 text format.
        /// </summary>
        [Output("startTimestamp")]
        public Output<string> StartTimestamp { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// [Output Only] An optional, human-readable explanation of the status.
        /// </summary>
        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;


        /// <summary>
        /// Create a RegionCommitment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionCommitment(string name, RegionCommitmentArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:RegionCommitment", name, args ?? new RegionCommitmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionCommitment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:RegionCommitment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RegionCommitment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionCommitment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionCommitment(name, id, options);
        }
    }

    public sealed class RegionCommitmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// [Output Only] Commitment end time in RFC3339 text format.
        /// </summary>
        [Input("endTimestamp")]
        public Input<string>? EndTimestamp { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#commitment for commitments.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The license specification required as part of a license commitment.
        /// </summary>
        [Input("licenseResource")]
        public Input<Inputs.LicenseResourceCommitmentArgs>? LicenseResource { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where this commitment may be used.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("reservations")]
        private InputList<Inputs.ReservationArgs>? _reservations;

        /// <summary>
        /// List of reservations in this commitment.
        /// </summary>
        public InputList<Inputs.ReservationArgs> Reservations
        {
            get => _reservations ?? (_reservations = new InputList<Inputs.ReservationArgs>());
            set => _reservations = value;
        }

        [Input("resources")]
        private InputList<Inputs.ResourceCommitmentArgs>? _resources;

        /// <summary>
        /// A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        /// </summary>
        public InputList<Inputs.ResourceCommitmentArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.ResourceCommitmentArgs>());
            set => _resources = value;
        }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] Commitment start time in RFC3339 text format.
        /// </summary>
        [Input("startTimestamp")]
        public Input<string>? StartTimestamp { get; set; }

        /// <summary>
        /// [Output Only] Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// [Output Only] An optional, human-readable explanation of the status.
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        public RegionCommitmentArgs()
        {
        }
    }
}
