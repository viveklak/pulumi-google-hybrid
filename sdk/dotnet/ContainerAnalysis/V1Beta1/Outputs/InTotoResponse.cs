// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.ContainerAnalysis.V1Beta1.Outputs
{

    [OutputType]
    public sealed class InTotoResponse
    {
        /// <summary>
        /// This field contains the expected command used to perform the step.
        /// </summary>
        public readonly ImmutableArray<string> ExpectedCommand;
        /// <summary>
        /// The following fields contain in-toto artifact rules identifying the artifacts that enter this supply chain step, and exit the supply chain step, i.e. materials and products of the step.
        /// </summary>
        public readonly ImmutableArray<Outputs.ArtifactRuleResponse> ExpectedMaterials;
        public readonly ImmutableArray<Outputs.ArtifactRuleResponse> ExpectedProducts;
        /// <summary>
        /// This field contains the public keys that can be used to verify the signatures on the step metadata.
        /// </summary>
        public readonly ImmutableArray<Outputs.SigningKeyResponse> SigningKeys;
        /// <summary>
        /// This field identifies the name of the step in the supply chain.
        /// </summary>
        public readonly string StepName;
        /// <summary>
        /// This field contains a value that indicates the minimum number of keys that need to be used to sign the step's in-toto link.
        /// </summary>
        public readonly string Threshold;

        [OutputConstructor]
        private InTotoResponse(
            ImmutableArray<string> expectedCommand,

            ImmutableArray<Outputs.ArtifactRuleResponse> expectedMaterials,

            ImmutableArray<Outputs.ArtifactRuleResponse> expectedProducts,

            ImmutableArray<Outputs.SigningKeyResponse> signingKeys,

            string stepName,

            string threshold)
        {
            ExpectedCommand = expectedCommand;
            ExpectedMaterials = expectedMaterials;
            ExpectedProducts = expectedProducts;
            SigningKeys = signingKeys;
            StepName = stepName;
            Threshold = threshold;
        }
    }
}