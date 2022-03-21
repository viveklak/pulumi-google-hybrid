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
    /// Create an Access Level. The longrunning operation from this RPC will have a successful status once the Access Level has propagated to long-lasting storage. Access Levels containing errors will result in an error response for the first error encountered.
    /// </summary>
    [GoogleNativeResourceType("google-native:accesscontextmanager/v1beta:AccessLevel")]
    public partial class AccessLevel : Pulumi.CustomResource
    {
        /// <summary>
        /// A `BasicLevel` composed of `Conditions`.
        /// </summary>
        [Output("basic")]
        public Output<Outputs.BasicLevelResponse> Basic { get; private set; } = null!;

        /// <summary>
        /// A `CustomLevel` written in the Common Expression Language.
        /// </summary>
        [Output("custom")]
        public Output<Outputs.CustomLevelResponse> Custom { get; private set; } = null!;

        /// <summary>
        /// Description of the `AccessLevel` and its use. Does not affect behavior.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Resource name for the Access Level. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/accessLevels/{short_name}`. The maximum length // of the `short_name` component is 50 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a AccessLevel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessLevel(string name, AccessLevelArgs args, CustomResourceOptions? options = null)
            : base("google-native:accesscontextmanager/v1beta:AccessLevel", name, args ?? new AccessLevelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessLevel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:accesscontextmanager/v1beta:AccessLevel", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccessLevel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessLevel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessLevel(name, id, options);
        }
    }

    public sealed class AccessLevelArgs : Pulumi.ResourceArgs
    {
        [Input("accessPolicyId", required: true)]
        public Input<string> AccessPolicyId { get; set; } = null!;

        /// <summary>
        /// A `BasicLevel` composed of `Conditions`.
        /// </summary>
        [Input("basic")]
        public Input<Inputs.BasicLevelArgs>? Basic { get; set; }

        /// <summary>
        /// A `CustomLevel` written in the Common Expression Language.
        /// </summary>
        [Input("custom")]
        public Input<Inputs.CustomLevelArgs>? Custom { get; set; }

        /// <summary>
        /// Description of the `AccessLevel` and its use. Does not affect behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource name for the Access Level. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/accessLevels/{short_name}`. The maximum length // of the `short_name` component is 50 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public AccessLevelArgs()
        {
        }
    }
}
