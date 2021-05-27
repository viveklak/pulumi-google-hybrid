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
    /// Creates a debug session for a deployed API Proxy revision.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:OrganizationEnvironmentApiRevisionDebugsession")]
    public partial class OrganizationEnvironmentApiRevisionDebugsession : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
        /// </summary>
        [Output("count")]
        public Output<int> Count { get; private set; } = null!;

        /// <summary>
        /// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        /// <summary>
        /// A unique ID for this DebugSession.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
        /// </summary>
        [Output("timeout")]
        public Output<string> Timeout { get; private set; } = null!;

        /// <summary>
        /// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
        /// </summary>
        [Output("tracesize")]
        public Output<int> Tracesize { get; private set; } = null!;

        /// <summary>
        /// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
        /// </summary>
        [Output("validity")]
        public Output<int> Validity { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationEnvironmentApiRevisionDebugsession resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationEnvironmentApiRevisionDebugsession(string name, OrganizationEnvironmentApiRevisionDebugsessionArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationEnvironmentApiRevisionDebugsession", name, args ?? new OrganizationEnvironmentApiRevisionDebugsessionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationEnvironmentApiRevisionDebugsession(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationEnvironmentApiRevisionDebugsession", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationEnvironmentApiRevisionDebugsession resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationEnvironmentApiRevisionDebugsession Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationEnvironmentApiRevisionDebugsession(name, id, options);
        }
    }

    public sealed class OrganizationEnvironmentApiRevisionDebugsessionArgs : Pulumi.ResourceArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// A unique ID for this DebugSession.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("revisionId", required: true)]
        public Input<string> RevisionId { get; set; } = null!;

        /// <summary>
        /// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        /// <summary>
        /// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
        /// </summary>
        [Input("tracesize")]
        public Input<int>? Tracesize { get; set; }

        /// <summary>
        /// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
        /// </summary>
        [Input("validity")]
        public Input<int>? Validity { get; set; }

        public OrganizationEnvironmentApiRevisionDebugsessionArgs()
        {
        }
    }
}
