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
    /// A filter to target VM instances for patching. The targeted VMs must meet all criteria specified. So if both labels and zones are specified, the patch job targets only VMs with those labels and in those zones.
    /// </summary>
    public sealed class PatchInstanceFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target all VM instances in the project. If true, no other criteria is permitted.
        /// </summary>
        [Input("all")]
        public Input<bool>? All { get; set; }

        [Input("groupLabels")]
        private InputList<Inputs.PatchInstanceFilterGroupLabelArgs>? _groupLabels;

        /// <summary>
        /// Targets VM instances matching at least one of these label sets. This allows targeting of disparate groups, for example "env=prod or env=staging".
        /// </summary>
        public InputList<Inputs.PatchInstanceFilterGroupLabelArgs> GroupLabels
        {
            get => _groupLabels ?? (_groupLabels = new InputList<Inputs.PatchInstanceFilterGroupLabelArgs>());
            set => _groupLabels = value;
        }

        [Input("instanceNamePrefixes")]
        private InputList<string>? _instanceNamePrefixes;

        /// <summary>
        /// Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group VMs when targeting configs, for example prefix="prod-".
        /// </summary>
        public InputList<string> InstanceNamePrefixes
        {
            get => _instanceNamePrefixes ?? (_instanceNamePrefixes = new InputList<string>());
            set => _instanceNamePrefixes = value;
        }

        [Input("instances")]
        private InputList<string>? _instances;

        /// <summary>
        /// Targets any of the VM instances specified. Instances are specified by their URI in the form `zones/[ZONE]/instances/[INSTANCE_NAME]`, `projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`, or `https://www.googleapis.com/compute/v1/projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`
        /// </summary>
        public InputList<string> Instances
        {
            get => _instances ?? (_instances = new InputList<string>());
            set => _instances = value;
        }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public PatchInstanceFilterArgs()
        {
        }
    }
}
