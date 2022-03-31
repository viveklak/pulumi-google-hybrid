// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// Set of files to scan.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2FileSetResponse
    {
        /// <summary>
        /// The regex-filtered set of files to scan. Exactly one of `url` or `regex_file_set` must be set.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2CloudStorageRegexFileSetResponse RegexFileSet;
        /// <summary>
        /// The Cloud Storage url of the file(s) to scan, in the format `gs:///`. Trailing wildcard in the path is allowed. If the url ends in a trailing slash, the bucket or directory represented by the url will be scanned non-recursively (content in sub-directories will not be scanned). This means that `gs://mybucket/` is equivalent to `gs://mybucket/*`, and `gs://mybucket/directory/` is equivalent to `gs://mybucket/directory/*`. Exactly one of `url` or `regex_file_set` must be set.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GooglePrivacyDlpV2FileSetResponse(
            Outputs.GooglePrivacyDlpV2CloudStorageRegexFileSetResponse regexFileSet,

            string url)
        {
            RegexFileSet = regexFileSet;
            Url = url;
        }
    }
}
