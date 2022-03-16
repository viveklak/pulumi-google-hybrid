// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Outputs
{

    /// <summary>
    /// Docker image that is used to create a container and start a VM instance for the version that you deploy. Only applicable for instances running in the App Engine flexible environment.
    /// </summary>
    [OutputType]
    public sealed class ContainerInfoResponse
    {
        /// <summary>
        /// URI to the hosted container image in Google Container Registry. The URI must be fully qualified and include a tag or digest. Examples: "gcr.io/my-project/image:tag" or "gcr.io/my-project/image@digest"
        /// </summary>
        public readonly string Image;

        [OutputConstructor]
        private ContainerInfoResponse(string image)
        {
            Image = image;
        }
    }
}