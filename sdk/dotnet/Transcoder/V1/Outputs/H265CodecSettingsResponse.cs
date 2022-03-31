// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Outputs
{

    /// <summary>
    /// H265 codec settings.
    /// </summary>
    [OutputType]
    public sealed class H265CodecSettingsResponse
    {
        /// <summary>
        /// Specifies whether an open Group of Pictures (GOP) structure should be allowed or not. The default is `false`.
        /// </summary>
        public readonly bool AllowOpenGop;
        /// <summary>
        /// Specify the intensity of the adaptive quantizer (AQ). Must be between 0 and 1, where 0 disables the quantizer and 1 maximizes the quantizer. A higher value equals a lower bitrate but smoother image. The default is 0.
        /// </summary>
        public readonly double AqStrength;
        /// <summary>
        /// The number of consecutive B-frames. Must be greater than or equal to zero. Must be less than `VideoStream.gop_frame_count` if set. The default is 0.
        /// </summary>
        public readonly int BFrameCount;
        /// <summary>
        /// Allow B-pyramid for reference frame selection. This may not be supported on all decoders. The default is `false`.
        /// </summary>
        public readonly bool BPyramid;
        /// <summary>
        /// The video bitrate in bits per second. The minimum value is 1,000. The maximum value is 800,000,000.
        /// </summary>
        public readonly int BitrateBps;
        /// <summary>
        /// Target CRF level. Must be between 10 and 36, where 10 is the highest quality and 36 is the most efficient compression. The default is 21.
        /// </summary>
        public readonly int CrfLevel;
        /// <summary>
        /// Use two-pass encoding strategy to achieve better video quality. `VideoStream.rate_control_mode` must be `vbr`. The default is `false`.
        /// </summary>
        public readonly bool EnableTwoPass;
        /// <summary>
        /// The target video frame rate in frames per second (FPS). Must be less than or equal to 120. Will default to the input frame rate if larger than the input frame rate. The API will generate an output FPS that is divisible by the input FPS, and smaller or equal to the target FPS. See [Calculating frame rate](https://cloud.google.com/transcoder/docs/concepts/frame-rate) for more information.
        /// </summary>
        public readonly double FrameRate;
        /// <summary>
        /// Select the GOP size based on the specified duration. The default is `3s`. Note that `gopDuration` must be less than or equal to [`segmentDuration`](#SegmentSettings), and [`segmentDuration`](#SegmentSettings) must be divisible by `gopDuration`.
        /// </summary>
        public readonly string GopDuration;
        /// <summary>
        /// Select the GOP size based on the specified frame count. Must be greater than zero.
        /// </summary>
        public readonly int GopFrameCount;
        /// <summary>
        /// The height of the video in pixels. Must be an even integer. When not specified, the height is adjusted to match the specified width and input aspect ratio. If both are omitted, the input height is used.
        /// </summary>
        public readonly int HeightPixels;
        /// <summary>
        /// Pixel format to use. The default is `yuv420p`. Supported pixel formats: - `yuv420p` pixel format - `yuv422p` pixel format - `yuv444p` pixel format - `yuv420p10` 10-bit HDR pixel format - `yuv422p10` 10-bit HDR pixel format - `yuv444p10` 10-bit HDR pixel format - `yuv420p12` 12-bit HDR pixel format - `yuv422p12` 12-bit HDR pixel format - `yuv444p12` 12-bit HDR pixel format
        /// </summary>
        public readonly string PixelFormat;
        /// <summary>
        /// Enforces the specified codec preset. The default is `veryfast`. The available options are [FFmpeg-compatible](https://trac.ffmpeg.org/wiki/Encode/H.265). Note that certain values for this field may cause the transcoder to override other fields you set in the `H265CodecSettings` message.
        /// </summary>
        public readonly string Preset;
        /// <summary>
        /// Enforces the specified codec profile. The following profiles are supported: * 8-bit profiles * `main` (default) * `main-intra` * `mainstillpicture` * 10-bit profiles * `main10` (default) * `main10-intra` * `main422-10` * `main422-10-intra` * `main444-10` * `main444-10-intra` * 12-bit profiles * `main12` (default) * `main12-intra` * `main422-12` * `main422-12-intra` * `main444-12` * `main444-12-intra` The available options are [FFmpeg-compatible](https://x265.readthedocs.io/). Note that certain values for this field may cause the transcoder to override other fields you set in the `H265CodecSettings` message.
        /// </summary>
        public readonly string Profile;
        /// <summary>
        /// Specify the `rate_control_mode`. The default is `vbr`. Supported rate control modes: - `vbr` - variable bitrate - `crf` - constant rate factor
        /// </summary>
        public readonly string RateControlMode;
        /// <summary>
        /// Enforces the specified codec tune. The available options are [FFmpeg-compatible](https://trac.ffmpeg.org/wiki/Encode/H.265). Note that certain values for this field may cause the transcoder to override other fields you set in the `H265CodecSettings` message.
        /// </summary>
        public readonly string Tune;
        /// <summary>
        /// Initial fullness of the Video Buffering Verifier (VBV) buffer in bits. Must be greater than zero. The default is equal to 90% of `VideoStream.vbv_size_bits`.
        /// </summary>
        public readonly int VbvFullnessBits;
        /// <summary>
        /// Size of the Video Buffering Verifier (VBV) buffer in bits. Must be greater than zero. The default is equal to `VideoStream.bitrate_bps`.
        /// </summary>
        public readonly int VbvSizeBits;
        /// <summary>
        /// The width of the video in pixels. Must be an even integer. When not specified, the width is adjusted to match the specified height and input aspect ratio. If both are omitted, the input width is used.
        /// </summary>
        public readonly int WidthPixels;

        [OutputConstructor]
        private H265CodecSettingsResponse(
            bool allowOpenGop,

            double aqStrength,

            int bFrameCount,

            bool bPyramid,

            int bitrateBps,

            int crfLevel,

            bool enableTwoPass,

            double frameRate,

            string gopDuration,

            int gopFrameCount,

            int heightPixels,

            string pixelFormat,

            string preset,

            string profile,

            string rateControlMode,

            string tune,

            int vbvFullnessBits,

            int vbvSizeBits,

            int widthPixels)
        {
            AllowOpenGop = allowOpenGop;
            AqStrength = aqStrength;
            BFrameCount = bFrameCount;
            BPyramid = bPyramid;
            BitrateBps = bitrateBps;
            CrfLevel = crfLevel;
            EnableTwoPass = enableTwoPass;
            FrameRate = frameRate;
            GopDuration = gopDuration;
            GopFrameCount = gopFrameCount;
            HeightPixels = heightPixels;
            PixelFormat = pixelFormat;
            Preset = preset;
            Profile = profile;
            RateControlMode = rateControlMode;
            Tune = tune;
            VbvFullnessBits = vbvFullnessBits;
            VbvSizeBits = vbvSizeBits;
            WidthPixels = widthPixels;
        }
    }
}
