// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Managedidentities.V1Alpha1
{
    public static class GetPeering
    {
        /// <summary>
        /// Gets details of a single Peering.
        /// </summary>
        public static Task<GetPeeringResult> InvokeAsync(GetPeeringArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPeeringResult>("google-native:managedidentities/v1alpha1:getPeering", args ?? new GetPeeringArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Peering.
        /// </summary>
        public static Output<GetPeeringResult> Invoke(GetPeeringInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPeeringResult>("google-native:managedidentities/v1alpha1:getPeering", args ?? new GetPeeringInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPeeringArgs : Pulumi.InvokeArgs
    {
        [Input("peeringId", required: true)]
        public string PeeringId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetPeeringArgs()
        {
        }
    }

    public sealed class GetPeeringInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("peeringId", required: true)]
        public Input<string> PeeringId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetPeeringInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPeeringResult
    {
        /// <summary>
        /// The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
        /// </summary>
        public readonly string AuthorizedNetwork;
        /// <summary>
        /// The time the instance was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
        /// </summary>
        public readonly string DomainResource;
        /// <summary>
        /// Optional. Resource labels to represent user provided metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Unique name of the peering in this scope including projects and location using the form: `projects/{project_id}/locations/global/peerings/{peering_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current state of this Peering.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Additional information about the current status of this peering, if available.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// Last update time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetPeeringResult(
            string authorizedNetwork,

            string createTime,

            string domainResource,

            ImmutableDictionary<string, string> labels,

            string name,

            string state,

            string statusMessage,

            string updateTime)
        {
            AuthorizedNetwork = authorizedNetwork;
            CreateTime = createTime;
            DomainResource = domainResource;
            Labels = labels;
            Name = name;
            State = state;
            StatusMessage = statusMessage;
            UpdateTime = updateTime;
        }
    }
}
