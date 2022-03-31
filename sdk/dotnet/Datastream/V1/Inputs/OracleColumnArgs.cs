// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Inputs
{

    /// <summary>
    /// Oracle Column.
    /// </summary>
    public sealed class OracleColumnArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Column name.
        /// </summary>
        [Input("column")]
        public Input<string>? Column { get; set; }

        /// <summary>
        /// The Oracle data type.
        /// </summary>
        [Input("dataType")]
        public Input<string>? DataType { get; set; }

        /// <summary>
        /// Column encoding.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Column length.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// Whether or not the column can accept a null value.
        /// </summary>
        [Input("nullable")]
        public Input<bool>? Nullable { get; set; }

        /// <summary>
        /// The ordinal position of the column in the table.
        /// </summary>
        [Input("ordinalPosition")]
        public Input<int>? OrdinalPosition { get; set; }

        /// <summary>
        /// Column precision.
        /// </summary>
        [Input("precision")]
        public Input<int>? Precision { get; set; }

        /// <summary>
        /// Whether or not the column represents a primary key.
        /// </summary>
        [Input("primaryKey")]
        public Input<bool>? PrimaryKey { get; set; }

        /// <summary>
        /// Column scale.
        /// </summary>
        [Input("scale")]
        public Input<int>? Scale { get; set; }

        public OracleColumnArgs()
        {
        }
    }
}
