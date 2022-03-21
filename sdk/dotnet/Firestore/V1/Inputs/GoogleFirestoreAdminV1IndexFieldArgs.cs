// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Firestore.V1.Inputs
{

    /// <summary>
    /// A field in an index. The field_path describes which field is indexed, the value_mode describes how the field value is indexed.
    /// </summary>
    public sealed class GoogleFirestoreAdminV1IndexFieldArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates that this field supports operations on `array_value`s.
        /// </summary>
        [Input("arrayConfig")]
        public Input<Pulumi.GoogleNative.Firestore.V1.GoogleFirestoreAdminV1IndexFieldArrayConfig>? ArrayConfig { get; set; }

        /// <summary>
        /// Can be __name__. For single field indexes, this must match the name of the field or may be omitted.
        /// </summary>
        [Input("fieldPath")]
        public Input<string>? FieldPath { get; set; }

        /// <summary>
        /// Indicates that this field supports ordering by the specified order or comparing using =, !=, &lt;, &lt;=, &gt;, &gt;=.
        /// </summary>
        [Input("order")]
        public Input<Pulumi.GoogleNative.Firestore.V1.GoogleFirestoreAdminV1IndexFieldOrder>? Order { get; set; }

        public GoogleFirestoreAdminV1IndexFieldArgs()
        {
        }
    }
}
