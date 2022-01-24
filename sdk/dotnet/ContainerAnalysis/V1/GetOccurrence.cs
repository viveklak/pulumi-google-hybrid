// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1
{
    public static class GetOccurrence
    {
        /// <summary>
        /// Gets the specified occurrence.
        /// </summary>
        public static Task<GetOccurrenceResult> InvokeAsync(GetOccurrenceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOccurrenceResult>("google-native:containeranalysis/v1:getOccurrence", args ?? new GetOccurrenceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified occurrence.
        /// </summary>
        public static Output<GetOccurrenceResult> Invoke(GetOccurrenceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOccurrenceResult>("google-native:containeranalysis/v1:getOccurrence", args ?? new GetOccurrenceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOccurrenceArgs : Pulumi.InvokeArgs
    {
        [Input("occurrenceId", required: true)]
        public string OccurrenceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetOccurrenceArgs()
        {
        }
    }

    public sealed class GetOccurrenceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("occurrenceId", required: true)]
        public Input<string> OccurrenceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetOccurrenceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOccurrenceResult
    {
        /// <summary>
        /// Describes an attestation of an artifact.
        /// </summary>
        public readonly Outputs.AttestationOccurrenceResponse Attestation;
        /// <summary>
        /// Describes a verifiable build.
        /// </summary>
        public readonly Outputs.BuildOccurrenceResponse Build;
        /// <summary>
        /// Describes a compliance violation on a linked resource.
        /// </summary>
        public readonly Outputs.ComplianceOccurrenceResponse Compliance;
        /// <summary>
        /// The time this occurrence was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Describes the deployment of an artifact on a runtime.
        /// </summary>
        public readonly Outputs.DeploymentOccurrenceResponse Deployment;
        /// <summary>
        /// Describes when a resource was discovered.
        /// </summary>
        public readonly Outputs.DiscoveryOccurrenceResponse Discovery;
        /// <summary>
        /// Describes an attestation of an artifact using dsse.
        /// </summary>
        public readonly Outputs.DSSEAttestationOccurrenceResponse DsseAttestation;
        /// <summary>
        /// https://github.com/secure-systems-lab/dsse
        /// </summary>
        public readonly Outputs.EnvelopeResponse Envelope;
        /// <summary>
        /// Describes how this resource derives from the basis in the associated note.
        /// </summary>
        public readonly Outputs.ImageOccurrenceResponse Image;
        /// <summary>
        /// This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
        /// </summary>
        public readonly string NoteName;
        /// <summary>
        /// Describes the installation of a package on the linked resource.
        /// </summary>
        public readonly Outputs.PackageOccurrenceResponse Package;
        /// <summary>
        /// A description of actions that can be taken to remedy the note.
        /// </summary>
        public readonly string Remediation;
        /// <summary>
        /// Immutable. A URI that represents the resource for which the occurrence applies. For example, `https://gcr.io/project/image@sha256:123abc` for a Docker image.
        /// </summary>
        public readonly string ResourceUri;
        /// <summary>
        /// The time this occurrence was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Describes an available package upgrade on the linked resource.
        /// </summary>
        public readonly Outputs.UpgradeOccurrenceResponse Upgrade;
        /// <summary>
        /// Describes a security vulnerability.
        /// </summary>
        public readonly Outputs.VulnerabilityOccurrenceResponse Vulnerability;

        [OutputConstructor]
        private GetOccurrenceResult(
            Outputs.AttestationOccurrenceResponse attestation,

            Outputs.BuildOccurrenceResponse build,

            Outputs.ComplianceOccurrenceResponse compliance,

            string createTime,

            Outputs.DeploymentOccurrenceResponse deployment,

            Outputs.DiscoveryOccurrenceResponse discovery,

            Outputs.DSSEAttestationOccurrenceResponse dsseAttestation,

            Outputs.EnvelopeResponse envelope,

            Outputs.ImageOccurrenceResponse image,

            string kind,

            string name,

            string noteName,

            Outputs.PackageOccurrenceResponse package,

            string remediation,

            string resourceUri,

            string updateTime,

            Outputs.UpgradeOccurrenceResponse upgrade,

            Outputs.VulnerabilityOccurrenceResponse vulnerability)
        {
            Attestation = attestation;
            Build = build;
            Compliance = compliance;
            CreateTime = createTime;
            Deployment = deployment;
            Discovery = discovery;
            DsseAttestation = dsseAttestation;
            Envelope = envelope;
            Image = image;
            Kind = kind;
            Name = name;
            NoteName = noteName;
            Package = package;
            Remediation = remediation;
            ResourceUri = resourceUri;
            UpdateTime = updateTime;
            Upgrade = upgrade;
            Vulnerability = vulnerability;
        }
    }
}
