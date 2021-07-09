// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1Beta
{
    /// <summary>
    /// Create a Service Perimeter. The longrunning operation from this RPC will have a successful status once the Service Perimeter has propagated to long-lasting storage. Service Perimeters containing errors will result in an error response for the first error encountered.
    /// </summary>
    [GoogleNativeResourceType("google-native:accesscontextmanager/v1beta:ServicePerimeter")]
    public partial class ServicePerimeter : Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the `ServicePerimeter` and its use. Does not affect behavior.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
        /// </summary>
        [Output("perimeterType")]
        public Output<string> PerimeterType { get; private set; } = null!;

        /// <summary>
        /// Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
        /// </summary>
        [Output("status")]
        public Output<Outputs.ServicePerimeterConfigResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePerimeter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePerimeter(string name, ServicePerimeterArgs args, CustomResourceOptions? options = null)
            : base("google-native:accesscontextmanager/v1beta:ServicePerimeter", name, args ?? new ServicePerimeterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePerimeter(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:accesscontextmanager/v1beta:ServicePerimeter", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServicePerimeter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePerimeter Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServicePerimeter(name, id, options);
        }
    }

    public sealed class ServicePerimeterArgs : Pulumi.ResourceArgs
    {
        [Input("accessPolicyId", required: true)]
        public Input<string> AccessPolicyId { get; set; } = null!;

        /// <summary>
        /// Description of the `ServicePerimeter` and its use. Does not affect behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
        /// </summary>
        [Input("perimeterType")]
        public Input<Pulumi.GoogleNative.AccessContextManager.V1Beta.ServicePerimeterPerimeterType>? PerimeterType { get; set; }

        /// <summary>
        /// Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
        /// </summary>
        [Input("status")]
        public Input<Inputs.ServicePerimeterConfigArgs>? Status { get; set; }

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public ServicePerimeterArgs()
        {
        }
    }
}
