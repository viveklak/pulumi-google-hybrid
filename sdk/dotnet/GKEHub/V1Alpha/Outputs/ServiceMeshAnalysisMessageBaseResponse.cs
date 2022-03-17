// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// AnalysisMessageBase describes some common information that is needed for all messages.
    /// </summary>
    [OutputType]
    public sealed class ServiceMeshAnalysisMessageBaseResponse
    {
        /// <summary>
        /// A url pointing to the Service Mesh or Istio documentation for this specific error type.
        /// </summary>
        public readonly string DocumentationUrl;
        /// <summary>
        /// Represents how severe a message is.
        /// </summary>
        public readonly string Level;
        /// <summary>
        /// Represents the specific type of a message.
        /// </summary>
        public readonly Outputs.ServiceMeshTypeResponse Type;

        [OutputConstructor]
        private ServiceMeshAnalysisMessageBaseResponse(
            string documentationUrl,

            string level,

            Outputs.ServiceMeshTypeResponse type)
        {
            DocumentationUrl = documentationUrl;
            Level = level;
            Type = type;
        }
    }
}
