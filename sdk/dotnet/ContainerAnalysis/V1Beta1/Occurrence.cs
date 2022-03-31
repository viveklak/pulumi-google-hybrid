// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1
{
    /// <summary>
    /// Creates a new occurrence.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:containeranalysis/v1beta1:Occurrence")]
    public partial class Occurrence : Pulumi.CustomResource
    {
        /// <summary>
        /// Describes an attestation of an artifact.
        /// </summary>
        [Output("attestation")]
        public Output<Outputs.DetailsResponse> Attestation { get; private set; } = null!;

        /// <summary>
        /// Describes a verifiable build.
        /// </summary>
        [Output("build")]
        public Output<Outputs.GrafeasV1beta1BuildDetailsResponse> Build { get; private set; } = null!;

        /// <summary>
        /// The time this occurrence was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Describes the deployment of an artifact on a runtime.
        /// </summary>
        [Output("deployment")]
        public Output<Outputs.GrafeasV1beta1DeploymentDetailsResponse> Deployment { get; private set; } = null!;

        /// <summary>
        /// Describes how this resource derives from the basis in the associated note.
        /// </summary>
        [Output("derivedImage")]
        public Output<Outputs.GrafeasV1beta1ImageDetailsResponse> DerivedImage { get; private set; } = null!;

        /// <summary>
        /// Describes when a resource was discovered.
        /// </summary>
        [Output("discovered")]
        public Output<Outputs.GrafeasV1beta1DiscoveryDetailsResponse> Discovered { get; private set; } = null!;

        /// <summary>
        /// Describes the installation of a package on the linked resource.
        /// </summary>
        [Output("installation")]
        public Output<Outputs.GrafeasV1beta1PackageDetailsResponse> Installation { get; private set; } = null!;

        /// <summary>
        /// Describes a specific in-toto link.
        /// </summary>
        [Output("intoto")]
        public Output<Outputs.GrafeasV1beta1IntotoDetailsResponse> Intoto { get; private set; } = null!;

        /// <summary>
        /// This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
        /// </summary>
        [Output("noteName")]
        public Output<string> NoteName { get; private set; } = null!;

        /// <summary>
        /// A description of actions that can be taken to remedy the note.
        /// </summary>
        [Output("remediation")]
        public Output<string> Remediation { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource for which the occurrence applies.
        /// </summary>
        [Output("resource")]
        public Output<Outputs.ResourceResponse> Resource { get; private set; } = null!;

        /// <summary>
        /// Describes a specific software bill of materials document.
        /// </summary>
        [Output("sbom")]
        public Output<Outputs.DocumentOccurrenceResponse> Sbom { get; private set; } = null!;

        /// <summary>
        /// Describes a specific SPDX File.
        /// </summary>
        [Output("spdxFile")]
        public Output<Outputs.FileOccurrenceResponse> SpdxFile { get; private set; } = null!;

        /// <summary>
        /// Describes a specific SPDX Package.
        /// </summary>
        [Output("spdxPackage")]
        public Output<Outputs.PackageInfoOccurrenceResponse> SpdxPackage { get; private set; } = null!;

        /// <summary>
        /// Describes a specific SPDX Relationship.
        /// </summary>
        [Output("spdxRelationship")]
        public Output<Outputs.RelationshipOccurrenceResponse> SpdxRelationship { get; private set; } = null!;

        /// <summary>
        /// The time this occurrence was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Describes a security vulnerability.
        /// </summary>
        [Output("vulnerability")]
        public Output<Outputs.GrafeasV1beta1VulnerabilityDetailsResponse> Vulnerability { get; private set; } = null!;


        /// <summary>
        /// Create a Occurrence resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Occurrence(string name, OccurrenceArgs args, CustomResourceOptions? options = null)
            : base("google-native:containeranalysis/v1beta1:Occurrence", name, args ?? new OccurrenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Occurrence(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:containeranalysis/v1beta1:Occurrence", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Occurrence resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Occurrence Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Occurrence(name, id, options);
        }
    }

    public sealed class OccurrenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes an attestation of an artifact.
        /// </summary>
        [Input("attestation")]
        public Input<Inputs.DetailsArgs>? Attestation { get; set; }

        /// <summary>
        /// Describes a verifiable build.
        /// </summary>
        [Input("build")]
        public Input<Inputs.GrafeasV1beta1BuildDetailsArgs>? Build { get; set; }

        /// <summary>
        /// Describes the deployment of an artifact on a runtime.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.GrafeasV1beta1DeploymentDetailsArgs>? Deployment { get; set; }

        /// <summary>
        /// Describes how this resource derives from the basis in the associated note.
        /// </summary>
        [Input("derivedImage")]
        public Input<Inputs.GrafeasV1beta1ImageDetailsArgs>? DerivedImage { get; set; }

        /// <summary>
        /// Describes when a resource was discovered.
        /// </summary>
        [Input("discovered")]
        public Input<Inputs.GrafeasV1beta1DiscoveryDetailsArgs>? Discovered { get; set; }

        /// <summary>
        /// Describes the installation of a package on the linked resource.
        /// </summary>
        [Input("installation")]
        public Input<Inputs.GrafeasV1beta1PackageDetailsArgs>? Installation { get; set; }

        /// <summary>
        /// Describes a specific in-toto link.
        /// </summary>
        [Input("intoto")]
        public Input<Inputs.GrafeasV1beta1IntotoDetailsArgs>? Intoto { get; set; }

        /// <summary>
        /// Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
        /// </summary>
        [Input("noteName", required: true)]
        public Input<string> NoteName { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A description of actions that can be taken to remedy the note.
        /// </summary>
        [Input("remediation")]
        public Input<string>? Remediation { get; set; }

        /// <summary>
        /// Immutable. The resource for which the occurrence applies.
        /// </summary>
        [Input("resource", required: true)]
        public Input<Inputs.ResourceArgs> Resource { get; set; } = null!;

        /// <summary>
        /// Describes a specific software bill of materials document.
        /// </summary>
        [Input("sbom")]
        public Input<Inputs.DocumentOccurrenceArgs>? Sbom { get; set; }

        /// <summary>
        /// Describes a specific SPDX File.
        /// </summary>
        [Input("spdxFile")]
        public Input<Inputs.FileOccurrenceArgs>? SpdxFile { get; set; }

        /// <summary>
        /// Describes a specific SPDX Package.
        /// </summary>
        [Input("spdxPackage")]
        public Input<Inputs.PackageInfoOccurrenceArgs>? SpdxPackage { get; set; }

        /// <summary>
        /// Describes a specific SPDX Relationship.
        /// </summary>
        [Input("spdxRelationship")]
        public Input<Inputs.RelationshipOccurrenceArgs>? SpdxRelationship { get; set; }

        /// <summary>
        /// Describes a security vulnerability.
        /// </summary>
        [Input("vulnerability")]
        public Input<Inputs.GrafeasV1beta1VulnerabilityDetailsArgs>? Vulnerability { get; set; }

        public OccurrenceArgs()
        {
        }
    }
}
