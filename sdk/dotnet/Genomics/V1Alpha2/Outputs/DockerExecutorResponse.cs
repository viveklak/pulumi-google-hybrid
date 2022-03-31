// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Genomics.V1Alpha2.Outputs
{

    /// <summary>
    /// The Docker execuctor specification.
    /// </summary>
    [OutputType]
    public sealed class DockerExecutorResponse
    {
        /// <summary>
        /// The command or newline delimited script to run. The command string will be executed within a bash shell. If the command exits with a non-zero exit code, output parameter de-localization will be skipped and the pipeline operation's `error` field will be populated. Maximum command string length is 16384.
        /// </summary>
        public readonly string Cmd;
        /// <summary>
        /// Image name from either Docker Hub or Google Container Registry. Users that run pipelines must have READ access to the image.
        /// </summary>
        public readonly string ImageName;

        [OutputConstructor]
        private DockerExecutorResponse(
            string cmd,

            string imageName)
        {
            Cmd = cmd;
            ImageName = imageName;
        }
    }
}
