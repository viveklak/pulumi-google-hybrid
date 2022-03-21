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
    /// The action to take.
    /// </summary>
    public sealed class BucketLifecycleRuleItemActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target storage class. Required iff the type of the action is SetStorageClass.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        /// <summary>
        /// Type of the action. Currently, only Delete and SetStorageClass are supported.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public BucketLifecycleRuleItemActionArgs()
        {
        }
    }
}
