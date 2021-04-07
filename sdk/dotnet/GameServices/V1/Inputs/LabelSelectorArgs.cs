// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.GameServices.V1.Inputs
{

    /// <summary>
    /// The label selector, used to group labels on the resources.
    /// </summary>
    public sealed class LabelSelectorArgs : Pulumi.ResourceArgs
    {
        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels for this selector.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        public LabelSelectorArgs()
        {
        }
    }
}