// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Inputs
{

    /// <summary>
    /// Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files. Static file handlers describe which files in the application directory are static files, and which URLs serve them.
    /// </summary>
    public sealed class StaticFilesHandlerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether files should also be uploaded as code data. By default, files declared in static file handlers are uploaded as static data and are only served to end users; they cannot be read by the application. If enabled, uploads are charged against both your code and static data storage resource quotas.
        /// </summary>
        [Input("applicationReadable")]
        public Input<bool>? ApplicationReadable { get; set; }

        /// <summary>
        /// Time a static file served by this handler should be cached by web proxies and browsers.
        /// </summary>
        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        [Input("httpHeaders")]
        private InputMap<string>? _httpHeaders;

        /// <summary>
        /// HTTP headers to use for all responses from these URLs.
        /// </summary>
        public InputMap<string> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputMap<string>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// MIME type used to serve all files served by this handler.Defaults to file-specific MIME types, which are derived from each file's filename extension.
        /// </summary>
        [Input("mimeType")]
        public Input<string>? MimeType { get; set; }

        /// <summary>
        /// Path to the static files matched by the URL pattern, from the application root directory. The path can refer to text matched in groupings in the URL pattern.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Whether this handler should match the request if the file referenced by the handler does not exist.
        /// </summary>
        [Input("requireMatchingFile")]
        public Input<bool>? RequireMatchingFile { get; set; }

        /// <summary>
        /// Regular expression that matches the file paths for all files that should be referenced by this handler.
        /// </summary>
        [Input("uploadPathRegex")]
        public Input<string>? UploadPathRegex { get; set; }

        public StaticFilesHandlerArgs()
        {
        }
    }
}
