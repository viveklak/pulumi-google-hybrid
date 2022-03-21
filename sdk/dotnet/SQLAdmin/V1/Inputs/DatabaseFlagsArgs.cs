// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1.Inputs
{

    /// <summary>
    /// Database flags for Cloud SQL instances.
    /// </summary>
    public sealed class DatabaseFlagsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the flag. These flags are passed at instance startup, so include both server options and system variables. Flags are specified with underscores, not hyphens. For more information, see [Configuring Database Flags](https://cloud.google.com/sql/docs/mysql/flags) in the Cloud SQL documentation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The value of the flag. Boolean flags are set to `on` for true and `off` for false. This field must be omitted if the flag doesn't take a value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DatabaseFlagsArgs()
        {
        }
    }
}
