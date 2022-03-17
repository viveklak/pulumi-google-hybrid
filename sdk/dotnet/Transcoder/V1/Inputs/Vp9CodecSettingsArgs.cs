// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Inputs
{

    /// <summary>
    /// VP9 codec settings.
    /// </summary>
    public sealed class Vp9CodecSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The video bitrate in bits per second. The minimum value is 1,000. The maximum value is 480,000,000.
        /// </summary>
        [Input("bitrateBps", required: true)]
        public Input<int> BitrateBps { get; set; } = null!;

        /// <summary>
        /// Target CRF level. Must be between 10 and 36, where 10 is the highest quality and 36 is the most efficient compression. The default is 21. *Note*: This field is not supported. 
        /// </summary>
        [Input("crfLevel")]
        public Input<int>? CrfLevel { get; set; }

        /// <summary>
        /// The target video frame rate in frames per second (FPS). Must be less than or equal to 120. Will default to the input frame rate if larger than the input frame rate. The API will generate an output FPS that is divisible by the input FPS, and smaller or equal to the target FPS. See [Calculating frame rate](https://cloud.google.com/transcoder/docs/concepts/frame-rate) for more information.
        /// </summary>
        [Input("frameRate", required: true)]
        public Input<double> FrameRate { get; set; } = null!;

        /// <summary>
        /// Select the GOP size based on the specified duration. The default is `3s`. Note that `gopDuration` must be less than or equal to [`segmentDuration`](#SegmentSettings), and [`segmentDuration`](#SegmentSettings) must be divisible by `gopDuration`.
        /// </summary>
        [Input("gopDuration")]
        public Input<string>? GopDuration { get; set; }

        /// <summary>
        /// Select the GOP size based on the specified frame count. Must be greater than zero.
        /// </summary>
        [Input("gopFrameCount")]
        public Input<int>? GopFrameCount { get; set; }

        /// <summary>
        /// The height of the video in pixels. Must be an even integer. When not specified, the height is adjusted to match the specified width and input aspect ratio. If both are omitted, the input height is used.
        /// </summary>
        [Input("heightPixels")]
        public Input<int>? HeightPixels { get; set; }

        /// <summary>
        /// Pixel format to use. The default is `yuv420p`. Supported pixel formats: - `yuv420p` pixel format - `yuv422p` pixel format - `yuv444p` pixel format - `yuv420p10` 10-bit HDR pixel format - `yuv422p10` 10-bit HDR pixel format - `yuv444p10` 10-bit HDR pixel format - `yuv420p12` 12-bit HDR pixel format - `yuv422p12` 12-bit HDR pixel format - `yuv444p12` 12-bit HDR pixel format
        /// </summary>
        [Input("pixelFormat")]
        public Input<string>? PixelFormat { get; set; }

        /// <summary>
        /// Enforces the specified codec profile. The following profiles are supported: * `profile0` (default) * `profile1` * `profile2` * `profile3` The available options are [WebM-compatible](https://www.webmproject.org/vp9/profiles/). Note that certain values for this field may cause the transcoder to override other fields you set in the `Vp9CodecSettings` message.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// Specify the `rate_control_mode`. The default is `vbr`. Supported rate control modes: - `vbr` - variable bitrate
        /// </summary>
        [Input("rateControlMode")]
        public Input<string>? RateControlMode { get; set; }

        /// <summary>
        /// The width of the video in pixels. Must be an even integer. When not specified, the width is adjusted to match the specified height and input aspect ratio. If both are omitted, the input width is used.
        /// </summary>
        [Input("widthPixels")]
        public Input<int>? WidthPixels { get; set; }

        public Vp9CodecSettingsArgs()
        {
        }
    }
}
