// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Outputs
{

    /// <summary>
    /// Dump flag definition.
    /// </summary>
    [OutputType]
    public sealed class DumpFlagResponse
    {
        /// <summary>
        /// The name of the flag
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the flag.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private DumpFlagResponse(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
