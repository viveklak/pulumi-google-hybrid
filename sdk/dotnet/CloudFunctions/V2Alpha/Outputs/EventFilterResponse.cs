// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2Alpha.Outputs
{

    /// <summary>
    /// Filters events based on exact matches on the CloudEvents attributes.
    /// </summary>
    [OutputType]
    public sealed class EventFilterResponse
    {
        /// <summary>
        /// The name of a CloudEvents attribute.
        /// </summary>
        public readonly string Attribute;
        /// <summary>
        /// The value for the attribute.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private EventFilterResponse(
            string attribute,

            string value)
        {
            Attribute = attribute;
            Value = value;
        }
    }
}
