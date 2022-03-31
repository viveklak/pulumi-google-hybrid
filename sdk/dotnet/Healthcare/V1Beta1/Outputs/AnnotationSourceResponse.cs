// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Outputs
{

    /// <summary>
    /// AnnotationSource holds the source information of the annotation.
    /// </summary>
    [OutputType]
    public sealed class AnnotationSourceResponse
    {
        /// <summary>
        /// Cloud Healthcare API resource.
        /// </summary>
        public readonly Outputs.CloudHealthcareSourceResponse CloudHealthcareSource;

        [OutputConstructor]
        private AnnotationSourceResponse(Outputs.CloudHealthcareSourceResponse cloudHealthcareSource)
        {
            CloudHealthcareSource = cloudHealthcareSource;
        }
    }
}
