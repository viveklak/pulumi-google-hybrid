// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Jobs.V3.Inputs
{

    /// <summary>
    /// An object that represents a latitude/longitude pair. This is expressed as a pair of doubles to represent degrees latitude and degrees longitude. Unless specified otherwise, this object must conform to the WGS84 standard. Values must be within normalized ranges.
    /// </summary>
    public sealed class LatLngArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The latitude in degrees. It must be in the range [-90.0, +90.0].
        /// </summary>
        [Input("latitude")]
        public Input<double>? Latitude { get; set; }

        /// <summary>
        /// The longitude in degrees. It must be in the range [-180.0, +180.0].
        /// </summary>
        [Input("longitude")]
        public Input<double>? Longitude { get; set; }

        public LatLngArgs()
        {
        }
    }
}
