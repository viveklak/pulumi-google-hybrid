// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Memcache.V1Beta2.Outputs
{

    [OutputType]
    public sealed class InstanceMessageResponse
    {
        /// <summary>
        /// A code that correspond to one type of user-facing message.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Message on memcached instance which will be exposed to users.
        /// </summary>
        public readonly string Message;

        [OutputConstructor]
        private InstanceMessageResponse(
            string code,

            string message)
        {
            Code = code;
            Message = message;
        }
    }
}
