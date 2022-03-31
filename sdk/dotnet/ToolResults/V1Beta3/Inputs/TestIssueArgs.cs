// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// An issue detected occurring during a test execution.
    /// </summary>
    public sealed class TestIssueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Category of issue. Required.
        /// </summary>
        [Input("category")]
        public Input<Pulumi.GoogleNative.ToolResults.V1Beta3.TestIssueCategory>? Category { get; set; }

        /// <summary>
        /// A brief human-readable message describing the issue. Required.
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        /// <summary>
        /// Severity of issue. Required.
        /// </summary>
        [Input("severity")]
        public Input<Pulumi.GoogleNative.ToolResults.V1Beta3.TestIssueSeverity>? Severity { get; set; }

        /// <summary>
        /// Deprecated in favor of stack trace fields inside specific warnings.
        /// </summary>
        [Input("stackTrace")]
        public Input<Inputs.StackTraceArgs>? StackTrace { get; set; }

        /// <summary>
        /// Type of issue. Required.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.ToolResults.V1Beta3.TestIssueType>? Type { get; set; }

        /// <summary>
        /// Warning message with additional details of the issue. Should always be a message from com.google.devtools.toolresults.v1.warnings
        /// </summary>
        [Input("warning")]
        public Input<Inputs.AnyArgs>? Warning { get; set; }

        public TestIssueArgs()
        {
        }
    }
}
