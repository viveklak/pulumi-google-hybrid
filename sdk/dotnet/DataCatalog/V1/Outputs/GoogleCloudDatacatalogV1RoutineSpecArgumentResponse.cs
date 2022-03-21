// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Outputs
{

    /// <summary>
    /// Input or output argument of a function or stored procedure.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatacatalogV1RoutineSpecArgumentResponse
    {
        /// <summary>
        /// Specifies whether the argument is input or output.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The name of the argument. A return argument of a function might not have a name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of the argument. The exact value depends on the source system and the language.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1RoutineSpecArgumentResponse(
            string mode,

            string name,

            string type)
        {
            Mode = mode;
            Name = name;
            Type = type;
        }
    }
}
