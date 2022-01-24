// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Spanner.V1
{
    public static class GetSession
    {
        /// <summary>
        /// Gets a session. Returns `NOT_FOUND` if the session does not exist. This is mainly useful for determining whether a session is still alive.
        /// </summary>
        public static Task<GetSessionResult> InvokeAsync(GetSessionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSessionResult>("google-native:spanner/v1:getSession", args ?? new GetSessionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a session. Returns `NOT_FOUND` if the session does not exist. This is mainly useful for determining whether a session is still alive.
        /// </summary>
        public static Output<GetSessionResult> Invoke(GetSessionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSessionResult>("google-native:spanner/v1:getSession", args ?? new GetSessionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSessionArgs : Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public string DatabaseId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sessionId", required: true)]
        public string SessionId { get; set; } = null!;

        public GetSessionArgs()
        {
        }
    }

    public sealed class GetSessionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sessionId", required: true)]
        public Input<string> SessionId { get; set; } = null!;

        public GetSessionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSessionResult
    {
        /// <summary>
        /// The approximate timestamp when the session is last used. It is typically earlier than the actual last use time.
        /// </summary>
        public readonly string ApproximateLastUseTime;
        /// <summary>
        /// The timestamp when the session is created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The labels for the session. * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`. * No more than 64 labels can be associated with a given session. See https://goo.gl/xmQnxf for more information on and examples of labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The name of the session. This is always system-assigned.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetSessionResult(
            string approximateLastUseTime,

            string createTime,

            ImmutableDictionary<string, string> labels,

            string name)
        {
            ApproximateLastUseTime = approximateLastUseTime;
            CreateTime = createTime;
            Labels = labels;
            Name = name;
        }
    }
}
