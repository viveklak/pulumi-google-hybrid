// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudApigeeV1ReportPropertyResponse
    {
        /// <summary>
        /// name of the property
        /// </summary>
        public readonly string Property;
        /// <summary>
        /// property values
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudApigeeV1AttributeResponse> Value;

        [OutputConstructor]
        private GoogleCloudApigeeV1ReportPropertyResponse(
            string property,

            ImmutableArray<Outputs.GoogleCloudApigeeV1AttributeResponse> value)
        {
            Property = property;
            Value = value;
        }
    }
}
