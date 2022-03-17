// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1.Inputs
{

    /// <summary>
    /// Oracle table.
    /// </summary>
    public sealed class OracleTableArgs : Pulumi.ResourceArgs
    {
        [Input("oracleColumns")]
        private InputList<Inputs.OracleColumnArgs>? _oracleColumns;

        /// <summary>
        /// Oracle columns in the schema. When unspecified as part of inclue/exclude lists, includes/excludes everything.
        /// </summary>
        public InputList<Inputs.OracleColumnArgs> OracleColumns
        {
            get => _oracleColumns ?? (_oracleColumns = new InputList<Inputs.OracleColumnArgs>());
            set => _oracleColumns = value;
        }

        /// <summary>
        /// Table name.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public OracleTableArgs()
        {
        }
    }
}
