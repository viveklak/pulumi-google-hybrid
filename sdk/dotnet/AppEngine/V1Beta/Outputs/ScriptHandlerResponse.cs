// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Outputs
{

    /// <summary>
    /// Executes a script to handle the request that matches the URL pattern.
    /// </summary>
    [OutputType]
    public sealed class ScriptHandlerResponse
    {
        /// <summary>
        /// Path to the script from the application root directory.
        /// </summary>
        public readonly string ScriptPath;

        [OutputConstructor]
        private ScriptHandlerResponse(string scriptPath)
        {
            ScriptPath = scriptPath;
        }
    }
}
