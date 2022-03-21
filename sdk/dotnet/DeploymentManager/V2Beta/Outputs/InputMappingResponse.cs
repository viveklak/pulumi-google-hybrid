// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2Beta.Outputs
{

    /// <summary>
    /// InputMapping creates a 'virtual' property that will be injected into the properties before sending the request to the underlying API.
    /// </summary>
    [OutputType]
    public sealed class InputMappingResponse
    {
        /// <summary>
        /// The name of the field that is going to be injected.
        /// </summary>
        public readonly string FieldName;
        /// <summary>
        /// The location where this mapping applies.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Regex to evaluate on method to decide if input applies.
        /// </summary>
        public readonly string MethodMatch;
        /// <summary>
        /// A jsonPath expression to select an element.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private InputMappingResponse(
            string fieldName,

            string location,

            string methodMatch,

            string value)
        {
            FieldName = fieldName;
            Location = location;
            MethodMatch = methodMatch;
            Value = value;
        }
    }
}
