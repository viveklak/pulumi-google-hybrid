// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Inputs
{

    /// <summary>
    /// The owner of the object. This will always be the uploader of the object.
    /// </summary>
    public sealed class BucketObjectOwnerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The entity, in the form user-userId.
        /// </summary>
        [Input("entity")]
        public Input<string>? Entity { get; set; }

        /// <summary>
        /// The ID for the entity.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        public BucketObjectOwnerArgs()
        {
        }
    }
}
