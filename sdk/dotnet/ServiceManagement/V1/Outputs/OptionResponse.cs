// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.ServiceManagement.V1.Outputs
{

    [OutputType]
    public sealed class OptionResponse
    {
        /// <summary>
        /// The option's name. For protobuf built-in options (options defined in descriptor.proto), this is the short name. For example, `"map_entry"`. For custom options, it should be the fully-qualified name. For example, `"google.api.http"`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The option's value packed in an Any message. If the value is a primitive, the corresponding wrapper type defined in google/protobuf/wrappers.proto should be used. If the value is an enum, it should be stored as an int32 value using the google.protobuf.Int32Value type.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Value;

        [OutputConstructor]
        private OptionResponse(
            string name,

            ImmutableDictionary<string, string> value)
        {
            Name = name;
            Value = value;
        }
    }
}