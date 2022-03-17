// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.Alpha.Outputs
{

    /// <summary>
    /// Options allows customized resource handling by Deployment Manager.
    /// </summary>
    [OutputType]
    public sealed class OptionsResponse
    {
        /// <summary>
        /// Options regarding how to thread async requests.
        /// </summary>
        public readonly ImmutableArray<Outputs.AsyncOptionsResponse> AsyncOptions;
        /// <summary>
        /// The mappings that apply for requests.
        /// </summary>
        public readonly ImmutableArray<Outputs.InputMappingResponse> InputMappings;
        /// <summary>
        /// The json path to the field in the resource JSON body into which the resource name should be mapped. Leaving this empty indicates that there should be no mapping performed.
        /// </summary>
        public readonly string NameProperty;
        /// <summary>
        /// Options for how to validate and process properties on a resource.
        /// </summary>
        public readonly Outputs.ValidationOptionsResponse ValidationOptions;

        [OutputConstructor]
        private OptionsResponse(
            ImmutableArray<Outputs.AsyncOptionsResponse> asyncOptions,

            ImmutableArray<Outputs.InputMappingResponse> inputMappings,

            string nameProperty,

            Outputs.ValidationOptionsResponse validationOptions)
        {
            AsyncOptions = asyncOptions;
            InputMappings = inputMappings;
            NameProperty = nameProperty;
            ValidationOptions = validationOptions;
        }
    }
}
