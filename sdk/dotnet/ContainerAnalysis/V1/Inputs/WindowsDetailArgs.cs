// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    public sealed class WindowsDetailArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [CPE URI](https://cpe.mitre.org/specification/) this vulnerability affects.
        /// </summary>
        [Input("cpeUri", required: true)]
        public Input<string> CpeUri { get; set; } = null!;

        /// <summary>
        /// The description of this vulnerability.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fixingKbs", required: true)]
        private InputList<Inputs.KnowledgeBaseArgs>? _fixingKbs;

        /// <summary>
        /// The names of the KBs which have hotfixes to mitigate this vulnerability. Note that there may be multiple hotfixes (and thus multiple KBs) that mitigate a given vulnerability. Currently any listed KBs presence is considered a fix.
        /// </summary>
        public InputList<Inputs.KnowledgeBaseArgs> FixingKbs
        {
            get => _fixingKbs ?? (_fixingKbs = new InputList<Inputs.KnowledgeBaseArgs>());
            set => _fixingKbs = value;
        }

        /// <summary>
        /// The name of this vulnerability.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public WindowsDetailArgs()
        {
        }
    }
}