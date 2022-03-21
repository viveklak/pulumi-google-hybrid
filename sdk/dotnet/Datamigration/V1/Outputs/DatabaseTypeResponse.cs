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
    /// A message defining the database engine and provider.
    /// </summary>
    [OutputType]
    public sealed class DatabaseTypeResponse
    {
        /// <summary>
        /// The database engine.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// The database provider.
        /// </summary>
        public readonly string Provider;

        [OutputConstructor]
        private DatabaseTypeResponse(
            string engine,

            string provider)
        {
            Engine = engine;
            Provider = provider;
        }
    }
}
