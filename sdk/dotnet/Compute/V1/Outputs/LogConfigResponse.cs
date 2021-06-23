// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class LogConfigResponse
    {
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly Outputs.LogConfigCloudAuditOptionsResponse CloudAudit;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly Outputs.LogConfigCounterOptionsResponse Counter;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly Outputs.LogConfigDataAccessOptionsResponse DataAccess;

        [OutputConstructor]
        private LogConfigResponse(
            Outputs.LogConfigCloudAuditOptionsResponse cloudAudit,

            Outputs.LogConfigCounterOptionsResponse counter,

            Outputs.LogConfigDataAccessOptionsResponse dataAccess)
        {
            CloudAudit = cloudAudit;
            Counter = counter;
            DataAccess = dataAccess;
        }
    }
}
