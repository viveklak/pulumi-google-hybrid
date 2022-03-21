// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// Common Vulnerability Scoring System version 3. For details, see https://www.first.org/cvss/specification-document
    /// </summary>
    public sealed class CVSSv3Args : Pulumi.ResourceArgs
    {
        [Input("attackComplexity")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.CVSSv3AttackComplexity>? AttackComplexity { get; set; }

        /// <summary>
        /// Base Metrics Represents the intrinsic characteristics of a vulnerability that are constant over time and across user environments.
        /// </summary>
        [Input("attackVector")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.CVSSv3AttackVector>? AttackVector { get; set; }

        [Input("availabilityImpact")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.CVSSv3AvailabilityImpact>? AvailabilityImpact { get; set; }

        /// <summary>
        /// The base score is a function of the base metric scores.
        /// </summary>
        [Input("baseScore")]
        public Input<double>? BaseScore { get; set; }

        [Input("confidentialityImpact")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.CVSSv3ConfidentialityImpact>? ConfidentialityImpact { get; set; }

        [Input("exploitabilityScore")]
        public Input<double>? ExploitabilityScore { get; set; }

        [Input("impactScore")]
        public Input<double>? ImpactScore { get; set; }

        [Input("integrityImpact")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.CVSSv3IntegrityImpact>? IntegrityImpact { get; set; }

        [Input("privilegesRequired")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.CVSSv3PrivilegesRequired>? PrivilegesRequired { get; set; }

        [Input("scope")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.CVSSv3Scope>? Scope { get; set; }

        [Input("userInteraction")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.CVSSv3UserInteraction>? UserInteraction { get; set; }

        public CVSSv3Args()
        {
        }
    }
}
