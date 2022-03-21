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
    /// Multiplexing settings for output stream.
    /// </summary>
    public sealed class MuxStreamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The container format. The default is `mp4` Supported container formats: - `ts` - `fmp4`- the corresponding file extension is `.m4s` - `mp4` - `vtt`
        /// </summary>
        [Input("container")]
        public Input<string>? Container { get; set; }

        [Input("elementaryStreams")]
        private InputList<string>? _elementaryStreams;

        /// <summary>
        /// List of `ElementaryStream.key`s multiplexed in this stream.
        /// </summary>
        public InputList<string> ElementaryStreams
        {
            get => _elementaryStreams ?? (_elementaryStreams = new InputList<string>());
            set => _elementaryStreams = value;
        }

        /// <summary>
        /// The name of the generated file. The default is `MuxStream.key` with the extension suffix corresponding to the `MuxStream.container`. Individual segments also have an incremental 10-digit zero-padded suffix starting from 0 before the extension, such as `mux_stream0000000123.ts`.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        /// <summary>
        /// A unique key for this multiplexed stream. HLS media manifests will be named `MuxStream.key` with the `.m3u8` extension suffix.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Segment settings for `ts`, `fmp4` and `vtt`.
        /// </summary>
        [Input("segmentSettings")]
        public Input<Inputs.SegmentSettingsArgs>? SegmentSettings { get; set; }

        public MuxStreamArgs()
        {
        }
    }
}
