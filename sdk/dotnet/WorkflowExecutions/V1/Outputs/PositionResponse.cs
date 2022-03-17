// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WorkflowExecutions.V1.Outputs
{

    /// <summary>
    /// Position contains source position information about the stack trace element such as line number, column number and length of the code block in bytes.
    /// </summary>
    [OutputType]
    public sealed class PositionResponse
    {
        /// <summary>
        /// The source code column position (of the line) the current instruction was generated from.
        /// </summary>
        public readonly string Column;
        /// <summary>
        /// The number of bytes of source code making up this stack trace element.
        /// </summary>
        public readonly string Length;
        /// <summary>
        /// The source code line number the current instruction was generated from.
        /// </summary>
        public readonly string Line;

        [OutputConstructor]
        private PositionResponse(
            string column,

            string length,

            string line)
        {
            Column = column;
            Length = length;
            Line = line;
        }
    }
}
