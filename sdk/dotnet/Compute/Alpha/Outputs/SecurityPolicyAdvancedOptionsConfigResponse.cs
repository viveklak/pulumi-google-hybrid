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
    public sealed class SecurityPolicyAdvancedOptionsConfigResponse
    {
        public readonly string JsonParsing;
        public readonly string LogLevel;

        [OutputConstructor]
        private SecurityPolicyAdvancedOptionsConfigResponse(
            string jsonParsing,

            string logLevel)
        {
            JsonParsing = jsonParsing;
            LogLevel = logLevel;
        }
    }
}
