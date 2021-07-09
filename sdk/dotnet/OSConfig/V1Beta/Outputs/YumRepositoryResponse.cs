// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    [OutputType]
    public sealed class YumRepositoryResponse
    {
        /// <summary>
        /// The location of the repository directory.
        /// </summary>
        public readonly string BaseUrl;
        /// <summary>
        /// The display name of the repository.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// URIs of GPG keys.
        /// </summary>
        public readonly ImmutableArray<string> GpgKeys;

        [OutputConstructor]
        private YumRepositoryResponse(
            string baseUrl,

            string displayName,

            ImmutableArray<string> gpgKeys)
        {
            BaseUrl = baseUrl;
            DisplayName = displayName;
            GpgKeys = gpgKeys;
        }
    }
}
