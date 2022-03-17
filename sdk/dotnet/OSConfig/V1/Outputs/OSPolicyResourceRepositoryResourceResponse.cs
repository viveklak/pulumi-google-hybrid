// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Outputs
{

    /// <summary>
    /// A resource that manages a package repository.
    /// </summary>
    [OutputType]
    public sealed class OSPolicyResourceRepositoryResourceResponse
    {
        /// <summary>
        /// An Apt Repository.
        /// </summary>
        public readonly Outputs.OSPolicyResourceRepositoryResourceAptRepositoryResponse Apt;
        /// <summary>
        /// A Goo Repository.
        /// </summary>
        public readonly Outputs.OSPolicyResourceRepositoryResourceGooRepositoryResponse Goo;
        /// <summary>
        /// A Yum Repository.
        /// </summary>
        public readonly Outputs.OSPolicyResourceRepositoryResourceYumRepositoryResponse Yum;
        /// <summary>
        /// A Zypper Repository.
        /// </summary>
        public readonly Outputs.OSPolicyResourceRepositoryResourceZypperRepositoryResponse Zypper;

        [OutputConstructor]
        private OSPolicyResourceRepositoryResourceResponse(
            Outputs.OSPolicyResourceRepositoryResourceAptRepositoryResponse apt,

            Outputs.OSPolicyResourceRepositoryResourceGooRepositoryResponse goo,

            Outputs.OSPolicyResourceRepositoryResourceYumRepositoryResponse yum,

            Outputs.OSPolicyResourceRepositoryResourceZypperRepositoryResponse zypper)
        {
            Apt = apt;
            Goo = goo;
            Yum = yum;
            Zypper = zypper;
        }
    }
}
