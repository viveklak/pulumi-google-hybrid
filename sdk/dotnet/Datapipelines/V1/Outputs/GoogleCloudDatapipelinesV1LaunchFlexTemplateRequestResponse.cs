// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datapipelines.V1.Outputs
{

    /// <summary>
    /// A request to launch a Dataflow job from a Flex Template.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatapipelinesV1LaunchFlexTemplateRequestResponse
    {
        /// <summary>
        /// Parameter to launch a job from a Flex Template.
        /// </summary>
        public readonly Outputs.GoogleCloudDatapipelinesV1LaunchFlexTemplateParameterResponse LaunchParameter;
        /// <summary>
        /// The [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints) to which to direct the request. For example, `us-central1`, `us-west1`.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The ID of the Cloud Platform project that the job belongs to.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// If true, the request is validated but not actually executed. Defaults to false.
        /// </summary>
        public readonly bool ValidateOnly;

        [OutputConstructor]
        private GoogleCloudDatapipelinesV1LaunchFlexTemplateRequestResponse(
            Outputs.GoogleCloudDatapipelinesV1LaunchFlexTemplateParameterResponse launchParameter,

            string location,

            string project,

            bool validateOnly)
        {
            LaunchParameter = launchParameter;
            Location = location;
            Project = project;
            ValidateOnly = validateOnly;
        }
    }
}
