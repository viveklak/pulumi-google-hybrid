// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Outputs
{

    [OutputType]
    public sealed class PySparkJobResponse
    {
        /// <summary>
        /// Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
        /// </summary>
        public readonly ImmutableArray<string> ArchiveUris;
        /// <summary>
        /// Optional. The arguments to pass to the driver. Do not include arguments, such as --conf, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.
        /// </summary>
        public readonly ImmutableArray<string> FileUris;
        /// <summary>
        /// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.
        /// </summary>
        public readonly ImmutableArray<string> JarFileUris;
        /// <summary>
        /// Optional. The runtime log config for job execution.
        /// </summary>
        public readonly Outputs.LoggingConfigResponse LoggingConfig;
        /// <summary>
        /// The HCFS URI of the main Python file to use as the driver. Must be a .py file.
        /// </summary>
        public readonly string MainPythonFileUri;
        /// <summary>
        /// Optional. A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        /// <summary>
        /// Optional. HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.
        /// </summary>
        public readonly ImmutableArray<string> PythonFileUris;

        [OutputConstructor]
        private PySparkJobResponse(
            ImmutableArray<string> archiveUris,

            ImmutableArray<string> args,

            ImmutableArray<string> fileUris,

            ImmutableArray<string> jarFileUris,

            Outputs.LoggingConfigResponse loggingConfig,

            string mainPythonFileUri,

            ImmutableDictionary<string, string> properties,

            ImmutableArray<string> pythonFileUris)
        {
            ArchiveUris = archiveUris;
            Args = args;
            FileUris = fileUris;
            JarFileUris = jarFileUris;
            LoggingConfig = loggingConfig;
            MainPythonFileUri = mainPythonFileUri;
            Properties = properties;
            PythonFileUris = pythonFileUris;
        }
    }
}
