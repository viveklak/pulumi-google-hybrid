// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    public sealed class GoogleCloudApigeeV1ReportPropertyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of the property
        /// </summary>
        [Input("property")]
        public Input<string>? Property { get; set; }

        [Input("value")]
        private InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>? _value;

        /// <summary>
        /// property values
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1AttributeArgs> Value
        {
            get => _value ?? (_value = new InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>());
            set => _value = value;
        }

        public GoogleCloudApigeeV1ReportPropertyArgs()
        {
        }
    }
}
