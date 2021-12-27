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
    /// Manifest configuration.
    /// </summary>
    public sealed class ManifestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the generated file. The default is `manifest` with the extension suffix corresponding to the `Manifest.type`.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        [Input("muxStreams", required: true)]
        private InputList<string>? _muxStreams;

        /// <summary>
        /// List of user given `MuxStream.key`s that should appear in this manifest. When `Manifest.type` is `HLS`, a media manifest with name `MuxStream.key` and `.m3u8` extension is generated for each element of the `Manifest.mux_streams`.
        /// </summary>
        public InputList<string> MuxStreams
        {
            get => _muxStreams ?? (_muxStreams = new InputList<string>());
            set => _muxStreams = value;
        }

        /// <summary>
        /// Type of the manifest, can be `HLS` or `DASH`.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Transcoder.V1.ManifestType> Type { get; set; } = null!;

        public ManifestArgs()
        {
        }
    }
}