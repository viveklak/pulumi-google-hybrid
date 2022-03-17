// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class UDPHealthCheckResponse
    {
        /// <summary>
        /// The UDP port number for the health check request. Valid values are 1 through 65535.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence.
        /// </summary>
        public readonly string PortName;
        /// <summary>
        /// Raw data of request to send in payload of UDP packet. It is an error if this is empty. The request data can only be ASCII.
        /// </summary>
        public readonly string Request;
        /// <summary>
        /// The bytes to match against the beginning of the response data. It is an error if this is empty. The response data can only be ASCII.
        /// </summary>
        public readonly string Response;

        [OutputConstructor]
        private UDPHealthCheckResponse(
            int port,

            string portName,

            string request,

            string response)
        {
            Port = port;
            PortName = portName;
            Request = request;
            Response = response;
        }
    }
}
