// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Outputs
{

    /// <summary>
    /// Custom static error page to be served when an error occurs.
    /// </summary>
    [OutputType]
    public sealed class ErrorHandlerResponse
    {
        /// <summary>
        /// Error condition this handler applies to.
        /// </summary>
        public readonly string ErrorCode;
        /// <summary>
        /// MIME type of file. Defaults to text/html.
        /// </summary>
        public readonly string MimeType;
        /// <summary>
        /// Static file content to be served for this error.
        /// </summary>
        public readonly string StaticFile;

        [OutputConstructor]
        private ErrorHandlerResponse(
            string errorCode,

            string mimeType,

            string staticFile)
        {
            ErrorCode = errorCode;
            MimeType = mimeType;
            StaticFile = staticFile;
        }
    }
}
