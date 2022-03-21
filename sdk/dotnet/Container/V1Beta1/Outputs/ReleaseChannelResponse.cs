// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// ReleaseChannel indicates which release channel a cluster is subscribed to. Release channels are arranged in order of risk. When a cluster is subscribed to a release channel, Google maintains both the master version and the node version. Node auto-upgrade defaults to true and cannot be disabled.
    /// </summary>
    [OutputType]
    public sealed class ReleaseChannelResponse
    {
        /// <summary>
        /// channel specifies which release channel the cluster is subscribed to.
        /// </summary>
        public readonly string Channel;

        [OutputConstructor]
        private ReleaseChannelResponse(string channel)
        {
            Channel = channel;
        }
    }
}
