// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Memcache.V1.Outputs
{

    [OutputType]
    public sealed class NodeResponse
    {
        /// <summary>
        /// Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// User defined parameters currently applied to the node.
        /// </summary>
        public readonly Outputs.MemcacheParametersResponse Parameters;
        /// <summary>
        /// The port number of the Memcached server on this node.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Current state of the Memcached node.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Location (GCP Zone) for the Memcached node.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private NodeResponse(
            string host,

            string nodeId,

            Outputs.MemcacheParametersResponse parameters,

            int port,

            string state,

            string zone)
        {
            Host = host;
            NodeId = nodeId;
            Parameters = parameters;
            Port = port;
            State = state;
            Zone = zone;
        }
    }
}
