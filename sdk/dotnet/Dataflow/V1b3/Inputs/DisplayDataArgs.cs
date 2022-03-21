// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Inputs
{

    /// <summary>
    /// Data provided with a pipeline or transform to provide descriptive info.
    /// </summary>
    public sealed class DisplayDataArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains value if the data is of a boolean type.
        /// </summary>
        [Input("boolValue")]
        public Input<bool>? BoolValue { get; set; }

        /// <summary>
        /// Contains value if the data is of duration type.
        /// </summary>
        [Input("durationValue")]
        public Input<string>? DurationValue { get; set; }

        /// <summary>
        /// Contains value if the data is of float type.
        /// </summary>
        [Input("floatValue")]
        public Input<double>? FloatValue { get; set; }

        /// <summary>
        /// Contains value if the data is of int64 type.
        /// </summary>
        [Input("int64Value")]
        public Input<string>? Int64Value { get; set; }

        /// <summary>
        /// Contains value if the data is of java class type.
        /// </summary>
        [Input("javaClassValue")]
        public Input<string>? JavaClassValue { get; set; }

        /// <summary>
        /// The key identifying the display data. This is intended to be used as a label for the display data when viewed in a dax monitoring system.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// An optional label to display in a dax UI for the element.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The namespace for the key. This is usually a class name or programming language namespace (i.e. python module) which defines the display data. This allows a dax monitoring system to specially handle the data and perform custom rendering.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// A possible additional shorter value to display. For example a java_class_name_value of com.mypackage.MyDoFn will be stored with MyDoFn as the short_str_value and com.mypackage.MyDoFn as the java_class_name value. short_str_value can be displayed and java_class_name_value will be displayed as a tooltip.
        /// </summary>
        [Input("shortStrValue")]
        public Input<string>? ShortStrValue { get; set; }

        /// <summary>
        /// Contains value if the data is of string type.
        /// </summary>
        [Input("strValue")]
        public Input<string>? StrValue { get; set; }

        /// <summary>
        /// Contains value if the data is of timestamp type.
        /// </summary>
        [Input("timestampValue")]
        public Input<string>? TimestampValue { get; set; }

        /// <summary>
        /// An optional full URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public DisplayDataArgs()
        {
        }
    }
}
