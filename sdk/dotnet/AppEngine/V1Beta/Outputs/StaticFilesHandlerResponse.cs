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
    /// Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files. Static file handlers describe which files in the application directory are static files, and which URLs serve them.
    /// </summary>
    [OutputType]
    public sealed class StaticFilesHandlerResponse
    {
        /// <summary>
        /// Whether files should also be uploaded as code data. By default, files declared in static file handlers are uploaded as static data and are only served to end users; they cannot be read by the application. If enabled, uploads are charged against both your code and static data storage resource quotas.
        /// </summary>
        public readonly bool ApplicationReadable;
        /// <summary>
        /// Time a static file served by this handler should be cached by web proxies and browsers.
        /// </summary>
        public readonly string Expiration;
        /// <summary>
        /// HTTP headers to use for all responses from these URLs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> HttpHeaders;
        /// <summary>
        /// MIME type used to serve all files served by this handler.Defaults to file-specific MIME types, which are derived from each file's filename extension.
        /// </summary>
        public readonly string MimeType;
        /// <summary>
        /// Path to the static files matched by the URL pattern, from the application root directory. The path can refer to text matched in groupings in the URL pattern.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Whether this handler should match the request if the file referenced by the handler does not exist.
        /// </summary>
        public readonly bool RequireMatchingFile;
        /// <summary>
        /// Regular expression that matches the file paths for all files that should be referenced by this handler.
        /// </summary>
        public readonly string UploadPathRegex;

        [OutputConstructor]
        private StaticFilesHandlerResponse(
            bool applicationReadable,

            string expiration,

            ImmutableDictionary<string, string> httpHeaders,

            string mimeType,

            string path,

            bool requireMatchingFile,

            string uploadPathRegex)
        {
            ApplicationReadable = applicationReadable;
            Expiration = expiration;
            HttpHeaders = httpHeaders;
            MimeType = mimeType;
            Path = path;
            RequireMatchingFile = requireMatchingFile;
            UploadPathRegex = uploadPathRegex;
        }
    }
}
