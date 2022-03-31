// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Inputs
{

    /// <summary>
    /// A type definition for some HL7v2 type (incl. Segments and Datatypes).
    /// </summary>
    public sealed class TypeArgs : Pulumi.ResourceArgs
    {
        [Input("fields")]
        private InputList<Inputs.FieldArgs>? _fields;

        /// <summary>
        /// The (sub) fields this type has (if not primitive).
        /// </summary>
        public InputList<Inputs.FieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.FieldArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// The name of this type. This would be the segment or datatype name. For example, "PID" or "XPN".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If this is a primitive type then this field is the type of the primitive For example, STRING. Leave unspecified for composite types.
        /// </summary>
        [Input("primitive")]
        public Input<Pulumi.GoogleNative.Healthcare.V1.TypePrimitive>? Primitive { get; set; }

        public TypeArgs()
        {
        }
    }
}
