// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// Represents a Goo package repository. These is added to a repo file that is stored at C:/ProgramData/GooGet/repos/google_osconfig.repo.
    /// </summary>
    [OutputType]
    public sealed class GooRepositoryResponse
    {
        /// <summary>
        /// The name of the repository.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The url of the repository.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GooRepositoryResponse(
            string name,

            string url)
        {
            Name = name;
            Url = url;
        }
    }
}
