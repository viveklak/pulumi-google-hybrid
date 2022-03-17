// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Describes the auto-registration of the Forwarding Rule to Service Directory. The region and project of the Service Directory resource generated from this registration will be the same as this Forwarding Rule.
    /// </summary>
    [OutputType]
    public sealed class ForwardingRuleServiceDirectoryRegistrationResponse
    {
        /// <summary>
        /// Service Directory namespace to register the forwarding rule under.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// Service Directory service to register the forwarding rule under.
        /// </summary>
        public readonly string Service;
        /// <summary>
        /// [Optional] Service Directory region to register this global forwarding rule under. Default to "us-central1". Only used for PSC for Google APIs. All PSC for Google APIs Forwarding Rules on the same network should use the same Service Directory region.
        /// </summary>
        public readonly string ServiceDirectoryRegion;

        [OutputConstructor]
        private ForwardingRuleServiceDirectoryRegistrationResponse(
            string @namespace,

            string service,

            string serviceDirectoryRegion)
        {
            Namespace = @namespace;
            Service = service;
            ServiceDirectoryRegion = serviceDirectoryRegion;
        }
    }
}
