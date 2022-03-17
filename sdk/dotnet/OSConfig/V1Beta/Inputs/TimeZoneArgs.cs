// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Inputs
{

    /// <summary>
    /// Represents a time zone from the [IANA Time Zone Database](https://www.iana.org/time-zones).
    /// </summary>
    public sealed class TimeZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IANA Time Zone Database time zone, e.g. "America/New_York".
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Optional. IANA Time Zone Database version number, e.g. "2019a".
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public TimeZoneArgs()
        {
        }
    }
}
