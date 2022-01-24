// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Datastream.V1
{
    public static class GetPrivateConnection
    {
        /// <summary>
        /// Use this method to get details about a private connectivity configuration.
        /// </summary>
        public static Task<GetPrivateConnectionResult> InvokeAsync(GetPrivateConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateConnectionResult>("google-native:datastream/v1:getPrivateConnection", args ?? new GetPrivateConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this method to get details about a private connectivity configuration.
        /// </summary>
        public static Output<GetPrivateConnectionResult> Invoke(GetPrivateConnectionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPrivateConnectionResult>("google-native:datastream/v1:getPrivateConnection", args ?? new GetPrivateConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateConnectionArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("privateConnectionId", required: true)]
        public string PrivateConnectionId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetPrivateConnectionArgs()
        {
        }
    }

    public sealed class GetPrivateConnectionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("privateConnectionId", required: true)]
        public Input<string> PrivateConnectionId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetPrivateConnectionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateConnectionResult
    {
        /// <summary>
        /// The create time of the resource.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// In case of error, the details of the error in a user-friendly format.
        /// </summary>
        public readonly Outputs.ErrorResponse Error;
        /// <summary>
        /// Labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource's name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The state of the Private Connection.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The update time of the resource.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// VPC Peering Config.
        /// </summary>
        public readonly Outputs.VpcPeeringConfigResponse VpcPeeringConfig;

        [OutputConstructor]
        private GetPrivateConnectionResult(
            string createTime,

            string displayName,

            Outputs.ErrorResponse error,

            ImmutableDictionary<string, string> labels,

            string name,

            string state,

            string updateTime,

            Outputs.VpcPeeringConfigResponse vpcPeeringConfig)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            Error = error;
            Labels = labels;
            Name = name;
            State = state;
            UpdateTime = updateTime;
            VpcPeeringConfig = vpcPeeringConfig;
        }
    }
}
