// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Outputs
{

    /// <summary>
    /// The configuration for the parser. It determines how the server parses the messages.
    /// </summary>
    [OutputType]
    public sealed class ParserConfigResponse
    {
        /// <summary>
        /// Determines whether messages with no header are allowed.
        /// </summary>
        public readonly bool AllowNullHeader;
        /// <summary>
        /// Schemas used to parse messages in this store, if schematized parsing is desired.
        /// </summary>
        public readonly Outputs.SchemaPackageResponse Schema;
        /// <summary>
        /// Byte(s) to use as the segment terminator. If this is unset, '\r' is used as segment terminator, matching the HL7 version 2 specification.
        /// </summary>
        public readonly string SegmentTerminator;
        /// <summary>
        /// Immutable. Determines the version of both the default parser to be used when `schema` is not given, as well as the schematized parser used when `schema` is specified. This field is immutable after HL7v2 store creation.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private ParserConfigResponse(
            bool allowNullHeader,

            Outputs.SchemaPackageResponse schema,

            string segmentTerminator,

            string version)
        {
            AllowNullHeader = allowNullHeader;
            Schema = schema;
            SegmentTerminator = segmentTerminator;
            Version = version;
        }
    }
}
