// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1
{
    public static class GetConsentArtifact
    {
        /// <summary>
        /// Gets the specified Consent artifact.
        /// </summary>
        public static Task<GetConsentArtifactResult> InvokeAsync(GetConsentArtifactArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConsentArtifactResult>("google-native:healthcare/v1:getConsentArtifact", args ?? new GetConsentArtifactArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Consent artifact.
        /// </summary>
        public static Output<GetConsentArtifactResult> Invoke(GetConsentArtifactInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConsentArtifactResult>("google-native:healthcare/v1:getConsentArtifact", args ?? new GetConsentArtifactInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConsentArtifactArgs : Pulumi.InvokeArgs
    {
        [Input("consentArtifactId", required: true)]
        public string ConsentArtifactId { get; set; } = null!;

        [Input("consentStoreId", required: true)]
        public string ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetConsentArtifactArgs()
        {
        }
    }

    public sealed class GetConsentArtifactInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("consentArtifactId", required: true)]
        public Input<string> ConsentArtifactId { get; set; } = null!;

        [Input("consentStoreId", required: true)]
        public Input<string> ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetConsentArtifactInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConsentArtifactResult
    {
        /// <summary>
        /// Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
        /// </summary>
        public readonly ImmutableArray<Outputs.ImageResponse> ConsentContentScreenshots;
        /// <summary>
        /// Optional. An string indicating the version of the consent information shown to the user.
        /// </summary>
        public readonly string ConsentContentVersion;
        /// <summary>
        /// Optional. A signature from a guardian.
        /// </summary>
        public readonly Outputs.SignatureResponse GuardianSignature;
        /// <summary>
        /// Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// User's UUID provided by the client.
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// Optional. User's signature.
        /// </summary>
        public readonly Outputs.SignatureResponse UserSignature;
        /// <summary>
        /// Optional. A signature from a witness.
        /// </summary>
        public readonly Outputs.SignatureResponse WitnessSignature;

        [OutputConstructor]
        private GetConsentArtifactResult(
            ImmutableArray<Outputs.ImageResponse> consentContentScreenshots,

            string consentContentVersion,

            Outputs.SignatureResponse guardianSignature,

            ImmutableDictionary<string, string> metadata,

            string name,

            string userId,

            Outputs.SignatureResponse userSignature,

            Outputs.SignatureResponse witnessSignature)
        {
            ConsentContentScreenshots = consentContentScreenshots;
            ConsentContentVersion = consentContentVersion;
            GuardianSignature = guardianSignature;
            Metadata = metadata;
            Name = name;
            UserId = userId;
            UserSignature = userSignature;
            WitnessSignature = witnessSignature;
        }
    }
}
