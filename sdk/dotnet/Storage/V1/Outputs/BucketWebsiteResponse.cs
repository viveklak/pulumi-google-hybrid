// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Outputs
{

    /// <summary>
    /// The bucket's website configuration, controlling how the service behaves when accessing bucket contents as a web site. See the Static Website Examples for more information.
    /// </summary>
    [OutputType]
    public sealed class BucketWebsiteResponse
    {
        /// <summary>
        /// If the requested object path is missing, the service will ensure the path has a trailing '/', append this suffix, and attempt to retrieve the resulting object. This allows the creation of index.html objects to represent directory pages.
        /// </summary>
        public readonly string MainPageSuffix;
        /// <summary>
        /// If the requested object path is missing, and any mainPageSuffix object is missing, if applicable, the service will return the named object from this bucket as the content for a 404 Not Found result.
        /// </summary>
        public readonly string NotFoundPage;

        [OutputConstructor]
        private BucketWebsiteResponse(
            string mainPageSuffix,

            string notFoundPage)
        {
            MainPageSuffix = mainPageSuffix;
            NotFoundPage = notFoundPage;
        }
    }
}
