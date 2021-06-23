// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates a network endpoint group in the specified project using the parameters that are included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:NetworkEndpointGroup")]
    public partial class NetworkEndpointGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Metadata defined as annotations on the network endpoint group.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>> Annotations { get; private set; } = null!;

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        [Output("appEngine")]
        public Output<Outputs.NetworkEndpointGroupAppEngineResponse> AppEngine { get; private set; } = null!;

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        [Output("cloudFunction")]
        public Output<Outputs.NetworkEndpointGroupCloudFunctionResponse> CloudFunction { get; private set; } = null!;

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        [Output("cloudRun")]
        public Output<Outputs.NetworkEndpointGroupCloudRunResponse> CloudRun { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// The default port used if the port number is not specified in the network endpoint.
        /// </summary>
        [Output("defaultPort")]
        public Output<int> DefaultPort { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
        /// </summary>
        [Output("networkEndpointType")]
        public Output<string> NetworkEndpointType { get; private set; } = null!;

        /// <summary>
        /// The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
        /// </summary>
        [Output("pscTargetService")]
        public Output<string> PscTargetService { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The URL of the region where the network endpoint group is located.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
        /// </summary>
        [Output("serverlessDeployment")]
        public Output<Outputs.NetworkEndpointGroupServerlessDeploymentResponse> ServerlessDeployment { get; private set; } = null!;

        /// <summary>
        /// [Output only] Number of network endpoints in the network endpoint group.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        /// </summary>
        [Output("subnetwork")]
        public Output<string> Subnetwork { get; private set; } = null!;

        /// <summary>
        /// Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The URL of the zone where the network endpoint group is located.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkEndpointGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkEndpointGroup(string name, NetworkEndpointGroupArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:NetworkEndpointGroup", name, args ?? new NetworkEndpointGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkEndpointGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:NetworkEndpointGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkEndpointGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkEndpointGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkEndpointGroup(name, id, options);
        }
    }

    public sealed class NetworkEndpointGroupArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Metadata defined as annotations on the network endpoint group.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        [Input("appEngine")]
        public Input<Inputs.NetworkEndpointGroupAppEngineArgs>? AppEngine { get; set; }

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        [Input("cloudFunction")]
        public Input<Inputs.NetworkEndpointGroupCloudFunctionArgs>? CloudFunction { get; set; }

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        /// </summary>
        [Input("cloudRun")]
        public Input<Inputs.NetworkEndpointGroupCloudRunArgs>? CloudRun { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// The default port used if the port number is not specified in the network endpoint.
        /// </summary>
        [Input("defaultPort")]
        public Input<int>? DefaultPort { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, or SERVERLESS.
        /// </summary>
        [Input("networkEndpointType")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.NetworkEndpointGroupNetworkEndpointType>? NetworkEndpointType { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
        /// </summary>
        [Input("pscTargetService")]
        public Input<string>? PscTargetService { get; set; }

        /// <summary>
        /// [Output Only] The URL of the region where the network endpoint group is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Input("selfLinkWithId")]
        public Input<string>? SelfLinkWithId { get; set; }

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
        /// </summary>
        [Input("serverlessDeployment")]
        public Input<Inputs.NetworkEndpointGroupServerlessDeploymentArgs>? ServerlessDeployment { get; set; }

        /// <summary>
        /// [Output only] Number of network endpoints in the network endpoint group.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        /// <summary>
        /// Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.NetworkEndpointGroupType>? Type { get; set; }

        /// <summary>
        /// [Output Only] The URL of the zone where the network endpoint group is located.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public NetworkEndpointGroupArgs()
        {
        }
    }
}
