// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Inputs
{

    /// <summary>
    /// Setting a parameter value.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3FulfillmentSetParameterActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display name of the parameter.
        /// </summary>
        [Input("parameter")]
        public Input<string>? Parameter { get; set; }

        /// <summary>
        /// The new value of the parameter. A null value clears the parameter.
        /// </summary>
        [Input("value")]
        public Input<object>? Value { get; set; }

        public GoogleCloudDialogflowCxV3FulfillmentSetParameterActionArgs()
        {
        }
    }
}
