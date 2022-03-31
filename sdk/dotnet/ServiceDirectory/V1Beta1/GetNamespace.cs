// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceDirectory.V1Beta1
{
    public static class GetNamespace
    {
        /// <summary>
        /// Gets a namespace.
        /// </summary>
        public static Task<GetNamespaceResult> InvokeAsync(GetNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceResult>("google-native:servicedirectory/v1beta1:getNamespace", args ?? new GetNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a namespace.
        /// </summary>
        public static Output<GetNamespaceResult> Invoke(GetNamespaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNamespaceResult>("google-native:servicedirectory/v1beta1:getNamespace", args ?? new GetNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNamespaceArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("namespaceId", required: true)]
        public string NamespaceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetNamespaceArgs()
        {
        }
    }

    public sealed class GetNamespaceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetNamespaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespaceResult
    {
        /// <summary>
        /// The timestamp when the namespace was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Resource labels associated with this namespace. No more than 64 user labels can be associated with a given resource. Label keys and values can be no longer than 63 characters.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The resource name for the namespace in the format `projects/*/locations/*/namespaces/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The timestamp when the namespace was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetNamespaceResult(
            string createTime,

            ImmutableDictionary<string, string> labels,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            Labels = labels;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
