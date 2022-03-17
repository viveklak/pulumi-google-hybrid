// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Inputs
{

    /// <summary>
    /// The matcher that is based on node metadata presented by xDS clients.
    /// </summary>
    public sealed class MetadataLabelMatcherArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how matching should be done. Supported values are: MATCH_ANY: At least one of the Labels specified in the matcher should match the metadata presented by xDS client. MATCH_ALL: The metadata presented by the xDS client should contain all of the labels specified here. The selection is determined based on the best match. For example, suppose there are three EndpointPolicy resources P1, P2 and P3 and if P1 has a the matcher as MATCH_ANY , P2 has MATCH_ALL , and P3 has MATCH_ALL . If a client with label connects, the config from P1 will be selected. If a client with label connects, the config from P2 will be selected. If a client with label connects, the config from P3 will be selected. If there is more than one best match, (for example, if a config P4 with selector exists and if a client with label connects), an error will be thrown.
        /// </summary>
        [Input("metadataLabelMatchCriteria")]
        public Input<Pulumi.GoogleNative.NetworkServices.V1Beta1.MetadataLabelMatcherMetadataLabelMatchCriteria>? MetadataLabelMatchCriteria { get; set; }

        [Input("metadataLabels")]
        private InputList<Inputs.MetadataLabelsArgs>? _metadataLabels;

        /// <summary>
        /// The list of label value pairs that must match labels in the provided metadata based on filterMatchCriteria This list can have at most 64 entries. The list can be empty if the match criteria is MATCH_ANY, to specify a wildcard match (i.e this matches any client).
        /// </summary>
        public InputList<Inputs.MetadataLabelsArgs> MetadataLabels
        {
            get => _metadataLabels ?? (_metadataLabels = new InputList<Inputs.MetadataLabelsArgs>());
            set => _metadataLabels = value;
        }

        public MetadataLabelMatcherArgs()
        {
        }
    }
}
