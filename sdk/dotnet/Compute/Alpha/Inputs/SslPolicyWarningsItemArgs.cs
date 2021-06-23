// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class SslPolicyWarningsItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] A warning code, if applicable. For example, Compute Engine returns NO_RESULTS_ON_PAGE if there are no results in the response.
        /// </summary>
        [Input("code")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SslPolicyWarningsItemCode>? Code { get; set; }

        [Input("data")]
        private InputList<Inputs.SslPolicyWarningsItemDataItemArgs>? _data;

        /// <summary>
        /// [Output Only] Metadata about this warning in key: value format. For example: "data": [ { "key": "scope", "value": "zones/us-east1-d" } 
        /// </summary>
        public InputList<Inputs.SslPolicyWarningsItemDataItemArgs> Data
        {
            get => _data ?? (_data = new InputList<Inputs.SslPolicyWarningsItemDataItemArgs>());
            set => _data = value;
        }

        /// <summary>
        /// [Output Only] A human-readable description of the warning code.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        public SslPolicyWarningsItemArgs()
        {
        }
    }
}
