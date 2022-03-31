// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1Beta.Inputs
{

    /// <summary>
    /// Output only. Defines an error trace message for a ScanRun.
    /// </summary>
    public sealed class ScanRunErrorTraceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the error reason code.
        /// </summary>
        [Input("code")]
        public Input<Pulumi.GoogleNative.WebSecurityScanner.V1Beta.ScanRunErrorTraceCode>? Code { get; set; }

        /// <summary>
        /// If the scan encounters TOO_MANY_HTTP_ERRORS, this field indicates the most common HTTP error code, if such is available. For example, if this code is 404, the scan has encountered too many NOT_FOUND responses.
        /// </summary>
        [Input("mostCommonHttpErrorCode")]
        public Input<int>? MostCommonHttpErrorCode { get; set; }

        /// <summary>
        /// If the scan encounters SCAN_CONFIG_ISSUE error, this field has the error message encountered during scan configuration validation that is performed before each scan run.
        /// </summary>
        [Input("scanConfigError")]
        public Input<Inputs.ScanConfigErrorArgs>? ScanConfigError { get; set; }

        public ScanRunErrorTraceArgs()
        {
        }
    }
}
