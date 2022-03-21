// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    public sealed class KnowledgeBaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The KB name (generally of the form KB[0-9]+ i.e. KB123456).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A link to the KB in the Windows update catalog - https://www.catalog.update.microsoft.com/
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public KnowledgeBaseArgs()
        {
        }
    }
}
