// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    /// <summary>
    /// Command describes a step performed as part of the build pipeline.
    /// </summary>
    [OutputType]
    public sealed class CommandResponse
    {
        /// <summary>
        /// Command-line arguments used when executing this command.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// Working directory (relative to project source root) used when running this command.
        /// </summary>
        public readonly string Dir;
        /// <summary>
        /// Environment variables set before running this command.
        /// </summary>
        public readonly ImmutableArray<string> Env;
        /// <summary>
        /// Name of the command, as presented on the command line, or if the command is packaged as a Docker container, as presented to `docker pull`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID(s) of the command(s) that this command depends on.
        /// </summary>
        public readonly ImmutableArray<string> WaitFor;

        [OutputConstructor]
        private CommandResponse(
            ImmutableArray<string> args,

            string dir,

            ImmutableArray<string> env,

            string name,

            ImmutableArray<string> waitFor)
        {
            Args = args;
            Dir = dir;
            Env = env;
            Name = name;
            WaitFor = waitFor;
        }
    }
}
