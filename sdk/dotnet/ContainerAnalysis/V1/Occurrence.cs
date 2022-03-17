// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1
{
    /// <summary>
    /// Creates a new occurrence.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:containeranalysis/v1:Occurrence")]
    public partial class Occurrence : Pulumi.CustomResource
    {
        /// <summary>
        /// Describes an attestation of an artifact.
        /// </summary>
        [Output("attestation")]
        public Output<Outputs.AttestationOccurrenceResponse> Attestation { get; private set; } = null!;

        /// <summary>
        /// Describes a verifiable build.
        /// </summary>
        [Output("build")]
        public Output<Outputs.BuildOccurrenceResponse> Build { get; private set; } = null!;

        /// <summary>
        /// Describes a compliance violation on a linked resource.
        /// </summary>
        [Output("compliance")]
        public Output<Outputs.ComplianceOccurrenceResponse> Compliance { get; private set; } = null!;

        /// <summary>
        /// The time this occurrence was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Describes the deployment of an artifact on a runtime.
        /// </summary>
        [Output("deployment")]
        public Output<Outputs.DeploymentOccurrenceResponse> Deployment { get; private set; } = null!;

        /// <summary>
        /// Describes when a resource was discovered.
        /// </summary>
        [Output("discovery")]
        public Output<Outputs.DiscoveryOccurrenceResponse> Discovery { get; private set; } = null!;

        /// <summary>
        /// Describes an attestation of an artifact using dsse.
        /// </summary>
        [Output("dsseAttestation")]
        public Output<Outputs.DSSEAttestationOccurrenceResponse> DsseAttestation { get; private set; } = null!;

        /// <summary>
        /// https://github.com/secure-systems-lab/dsse
        /// </summary>
        [Output("envelope")]
        public Output<Outputs.EnvelopeResponse> Envelope { get; private set; } = null!;

        /// <summary>
        /// Describes how this resource derives from the basis in the associated note.
        /// </summary>
        [Output("image")]
        public Output<Outputs.ImageOccurrenceResponse> Image { get; private set; } = null!;

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
        /// Describes the installation of a package on the linked resource.
        /// </summary>
        [Output("package")]
        public Output<Outputs.PackageOccurrenceResponse> Package { get; private set; } = null!;

        /// <summary>
        /// A description of actions that can be taken to remedy the note.
        /// </summary>
        [Output("remediation")]
        public Output<string> Remediation { get; private set; } = null!;

        /// <summary>
        /// Immutable. A URI that represents the resource for which the occurrence applies. For example, `https://gcr.io/project/image@sha256:123abc` for a Docker image.
        /// </summary>
        [Output("resourceUri")]
        public Output<string> ResourceUri { get; private set; } = null!;

        /// <summary>
        /// The time this occurrence was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Describes an available package upgrade on the linked resource.
        /// </summary>
        [Output("upgrade")]
        public Output<Outputs.UpgradeOccurrenceResponse> Upgrade { get; private set; } = null!;

        /// <summary>
        /// Describes a security vulnerability.
        /// </summary>
        [Output("vulnerability")]
        public Output<Outputs.VulnerabilityOccurrenceResponse> Vulnerability { get; private set; } = null!;


        /// <summary>
        /// Create a Occurrence resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Occurrence(string name, OccurrenceArgs args, CustomResourceOptions? options = null)
            : base("google-native:containeranalysis/v1:Occurrence", name, args ?? new OccurrenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Occurrence(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:containeranalysis/v1:Occurrence", name, null, MakeResourceOptions(options, id))
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
        public Input<Inputs.AttestationOccurrenceArgs>? Attestation { get; set; }

        /// <summary>
        /// Describes a verifiable build.
        /// </summary>
        [Input("build")]
        public Input<Inputs.BuildOccurrenceArgs>? Build { get; set; }

        /// <summary>
        /// Describes a compliance violation on a linked resource.
        /// </summary>
        [Input("compliance")]
        public Input<Inputs.ComplianceOccurrenceArgs>? Compliance { get; set; }

        /// <summary>
        /// Describes the deployment of an artifact on a runtime.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.DeploymentOccurrenceArgs>? Deployment { get; set; }

        /// <summary>
        /// Describes when a resource was discovered.
        /// </summary>
        [Input("discovery")]
        public Input<Inputs.DiscoveryOccurrenceArgs>? Discovery { get; set; }

        /// <summary>
        /// Describes an attestation of an artifact using dsse.
        /// </summary>
        [Input("dsseAttestation")]
        public Input<Inputs.DSSEAttestationOccurrenceArgs>? DsseAttestation { get; set; }

        /// <summary>
        /// https://github.com/secure-systems-lab/dsse
        /// </summary>
        [Input("envelope")]
        public Input<Inputs.EnvelopeArgs>? Envelope { get; set; }

        /// <summary>
        /// Describes how this resource derives from the basis in the associated note.
        /// </summary>
        [Input("image")]
        public Input<Inputs.ImageOccurrenceArgs>? Image { get; set; }

        /// <summary>
        /// Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
        /// </summary>
        [Input("noteName", required: true)]
        public Input<string> NoteName { get; set; } = null!;

        /// <summary>
        /// Describes the installation of a package on the linked resource.
        /// </summary>
        [Input("package")]
        public Input<Inputs.PackageOccurrenceArgs>? Package { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A description of actions that can be taken to remedy the note.
        /// </summary>
        [Input("remediation")]
        public Input<string>? Remediation { get; set; }

        /// <summary>
        /// Immutable. A URI that represents the resource for which the occurrence applies. For example, `https://gcr.io/project/image@sha256:123abc` for a Docker image.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        /// <summary>
        /// Describes an available package upgrade on the linked resource.
        /// </summary>
        [Input("upgrade")]
        public Input<Inputs.UpgradeOccurrenceArgs>? Upgrade { get; set; }

        /// <summary>
        /// Describes a security vulnerability.
        /// </summary>
        [Input("vulnerability")]
        public Input<Inputs.VulnerabilityOccurrenceArgs>? Vulnerability { get; set; }

        public OccurrenceArgs()
        {
        }
    }
}
