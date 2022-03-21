// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1.Outputs
{

    /// <summary>
    /// Config for image bounding poly (and bounding box) human labeling task.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatalabelingV1beta1BoundingPolyConfigResponse
    {
        /// <summary>
        /// Annotation spec set resource name.
        /// </summary>
        public readonly string AnnotationSpecSet;
        /// <summary>
        /// Optional. Instruction message showed on contributors UI.
        /// </summary>
        public readonly string InstructionMessage;

        [OutputConstructor]
        private GoogleCloudDatalabelingV1beta1BoundingPolyConfigResponse(
            string annotationSpecSet,

            string instructionMessage)
        {
            AnnotationSpecSet = annotationSpecSet;
            InstructionMessage = instructionMessage;
        }
    }
}
