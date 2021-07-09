// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1Beta1.Inputs
{

    /// <summary>
    /// The mapping for the `Job.edit_list` atoms with audio `EditAtom.inputs`.
    /// </summary>
    public sealed class AudioAtomArgs : Pulumi.ResourceArgs
    {
        [Input("channels")]
        private InputList<Inputs.AudioChannelArgs>? _channels;

        /// <summary>
        /// List of `Channel`s for this audio stream. for in-depth explanation.
        /// </summary>
        public InputList<Inputs.AudioChannelArgs> Channels
        {
            get => _channels ?? (_channels = new InputList<Inputs.AudioChannelArgs>());
            set => _channels = value;
        }

        /// <summary>
        /// The `EditAtom.key` that references the atom with audio inputs in the `Job.edit_list`.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public AudioAtomArgs()
        {
        }
    }
}
