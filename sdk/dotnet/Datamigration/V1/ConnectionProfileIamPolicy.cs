// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Datamigration.V1
{
    /// <summary>
    /// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:datamigration/v1:ConnectionProfileIamPolicy")]
    public partial class ConnectionProfileIamPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a ConnectionProfileIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectionProfileIamPolicy(string name, ConnectionProfileIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:datamigration/v1:ConnectionProfileIamPolicy", name, args ?? new ConnectionProfileIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectionProfileIamPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:datamigration/v1:ConnectionProfileIamPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectionProfileIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectionProfileIamPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConnectionProfileIamPolicy(name, id, options);
        }
    }

    public sealed class ConnectionProfileIamPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("connectionProfilesId", required: true)]
        public Input<string> ConnectionProfilesId { get; set; } = null!;

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        /// <summary>
        /// REQUIRED: The complete policy to be applied to the `resource`. The size of the policy is limited to a few 10s of KB. An empty policy is a valid policy but certain Cloud Platform services (such as Projects) might reject them.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.PolicyArgs>? Policy { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
        /// </summary>
        [Input("updateMask")]
        public Input<string>? UpdateMask { get; set; }

        public ConnectionProfileIamPolicyArgs()
        {
        }
    }
}
