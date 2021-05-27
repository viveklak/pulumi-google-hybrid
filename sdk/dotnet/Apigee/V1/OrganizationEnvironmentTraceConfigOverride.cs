// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates a trace configuration override. The response contains a system-generated UUID, that can be used to view, update, or delete the configuration override. Use the List API to view the existing trace configuration overrides.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:OrganizationEnvironmentTraceConfigOverride")]
    public partial class OrganizationEnvironmentTraceConfigOverride : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the API proxy that will have its trace configuration overridden.
        /// </summary>
        [Output("apiProxy")]
        public Output<string> ApiProxy { get; private set; } = null!;

        /// <summary>
        /// ID of the trace configuration override specified as a system-generated UUID.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Trace configuration to override.
        /// </summary>
        [Output("samplingConfig")]
        public Output<Outputs.GoogleCloudApigeeV1TraceSamplingConfigResponse> SamplingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationEnvironmentTraceConfigOverride resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationEnvironmentTraceConfigOverride(string name, OrganizationEnvironmentTraceConfigOverrideArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationEnvironmentTraceConfigOverride", name, args ?? new OrganizationEnvironmentTraceConfigOverrideArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationEnvironmentTraceConfigOverride(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationEnvironmentTraceConfigOverride", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrganizationEnvironmentTraceConfigOverride resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationEnvironmentTraceConfigOverride Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationEnvironmentTraceConfigOverride(name, id, options);
        }
    }

    public sealed class OrganizationEnvironmentTraceConfigOverrideArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the API proxy that will have its trace configuration overridden.
        /// </summary>
        [Input("apiProxy")]
        public Input<string>? ApiProxy { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// ID of the trace configuration override specified as a system-generated UUID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Trace configuration to override.
        /// </summary>
        [Input("samplingConfig")]
        public Input<Inputs.GoogleCloudApigeeV1TraceSamplingConfigArgs>? SamplingConfig { get; set; }

        public OrganizationEnvironmentTraceConfigOverrideArgs()
        {
        }
    }
}
