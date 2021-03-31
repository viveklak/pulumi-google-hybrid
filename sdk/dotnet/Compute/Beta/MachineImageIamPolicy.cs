// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Beta
{
    /// <summary>
    /// Sets the access control policy on the specified resource. Replaces any existing policy.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:compute/beta:MachineImageIamPolicy")]
    public partial class MachineImageIamPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a MachineImageIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MachineImageIamPolicy(string name, MachineImageIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:compute/beta:MachineImageIamPolicy", name, args ?? new MachineImageIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MachineImageIamPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:compute/beta:MachineImageIamPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MachineImageIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MachineImageIamPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MachineImageIamPolicy(name, id, options);
        }
    }

    public sealed class MachineImageIamPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("bindings")]
        private InputList<Inputs.BindingArgs>? _bindings;

        /// <summary>
        /// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify bindings.
        /// </summary>
        public InputList<Inputs.BindingArgs> Bindings
        {
            get => _bindings ?? (_bindings = new InputList<Inputs.BindingArgs>());
            set => _bindings = value;
        }

        /// <summary>
        /// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify the etag.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// REQUIRED: The complete policy to be applied to the 'resource'. The size of the policy is limited to a few 10s of KB. An empty policy is in general a valid policy but certain services (like Projects) might reject them.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.PolicyArgs>? Policy { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        public MachineImageIamPolicyArgs()
        {
        }
    }
}
