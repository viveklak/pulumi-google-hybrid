// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1
{
    public static class GetDatacenterConnector
    {
        /// <summary>
        /// Gets details of a single DatacenterConnector.
        /// </summary>
        public static Task<GetDatacenterConnectorResult> InvokeAsync(GetDatacenterConnectorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatacenterConnectorResult>("google-native:vmmigration/v1alpha1:getDatacenterConnector", args ?? new GetDatacenterConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single DatacenterConnector.
        /// </summary>
        public static Output<GetDatacenterConnectorResult> Invoke(GetDatacenterConnectorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatacenterConnectorResult>("google-native:vmmigration/v1alpha1:getDatacenterConnector", args ?? new GetDatacenterConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatacenterConnectorArgs : Pulumi.InvokeArgs
    {
        [Input("datacenterConnectorId", required: true)]
        public string DatacenterConnectorId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetDatacenterConnectorArgs()
        {
        }
    }

    public sealed class GetDatacenterConnectorInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("datacenterConnectorId", required: true)]
        public Input<string> DatacenterConnectorId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetDatacenterConnectorInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatacenterConnectorResult
    {
        /// <summary>
        /// The communication channel between the datacenter connector and GCP.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The time the connector was created (as an API call, not when it was actually installed).
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Provides details on the state of the Datacenter Connector in case of an error.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// The connector's name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
        /// </summary>
        public readonly string RegistrationId;
        /// <summary>
        /// The service account to use in the connector when communicating with the cloud.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// State of the DatacenterConnector, as determined by the health checks.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The time the state was last set.
        /// </summary>
        public readonly string StateTime;
        /// <summary>
        /// The last time the connector was updated with an API call.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetDatacenterConnectorResult(
            string bucket,

            string createTime,

            Outputs.StatusResponse error,

            string name,

            string registrationId,

            string serviceAccount,

            string state,

            string stateTime,

            string updateTime,

            string version)
        {
            Bucket = bucket;
            CreateTime = createTime;
            Error = error;
            Name = name;
            RegistrationId = registrationId;
            ServiceAccount = serviceAccount;
            State = state;
            StateTime = stateTime;
            UpdateTime = updateTime;
            Version = version;
        }
    }
}
