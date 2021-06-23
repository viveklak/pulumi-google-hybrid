// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// A set of instance tags.
    /// </summary>
    public sealed class TagsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a fingerprint for this request, which is essentially a hash of the tags' contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update tags. You must always provide an up-to-date fingerprint hash in order to update or change tags. To see the latest fingerprint, make get() request to the instance.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        [Input("items")]
        private InputList<string>? _items;

        /// <summary>
        /// An array of tags. Each tag must be 1-63 characters long, and comply with RFC1035.
        /// </summary>
        public InputList<string> Items
        {
            get => _items ?? (_items = new InputList<string>());
            set => _items = value;
        }

        public TagsArgs()
        {
        }
    }
}
