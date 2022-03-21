// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Outputs
{

    /// <summary>
    /// User-specified config for running a Spark task.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1TaskSparkTaskConfigResponse
    {
        /// <summary>
        /// Optional. Cloud Storage URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
        /// </summary>
        public readonly ImmutableArray<string> ArchiveUris;
        /// <summary>
        /// Optional. Cloud Storage URIs of files to be placed in the working directory of each executor.
        /// </summary>
        public readonly ImmutableArray<string> FileUris;
        /// <summary>
        /// Optional. Infrastructure specification for the execution.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1TaskInfrastructureSpecResponse InfrastructureSpec;
        /// <summary>
        /// The name of the driver's main class. The jar file that contains the class must be in the default CLASSPATH or specified in jar_file_uris. The execution args are passed in as a sequence of named process arguments (--key=value).
        /// </summary>
        public readonly string MainClass;
        /// <summary>
        /// The Cloud Storage URI of the jar file that contains the main class. The execution args are passed in as a sequence of named process arguments (--key=value).
        /// </summary>
        public readonly string MainJarFileUri;
        /// <summary>
        /// The Gcloud Storage URI of the main Python file to use as the driver. Must be a .py file. The execution args are passed in as a sequence of named process arguments (--key=value).
        /// </summary>
        public readonly string PythonScriptFile;
        /// <summary>
        /// The query text. The execution args are used to declare a set of script variables (set key="value";).
        /// </summary>
        public readonly string SqlScript;
        /// <summary>
        /// A reference to a query file. This can be the Cloud Storage URI of the query file or it can the path to a SqlScript Content. The execution args are used to declare a set of script variables (set key="value";).
        /// </summary>
        public readonly string SqlScriptFile;

        [OutputConstructor]
        private GoogleCloudDataplexV1TaskSparkTaskConfigResponse(
            ImmutableArray<string> archiveUris,

            ImmutableArray<string> fileUris,

            Outputs.GoogleCloudDataplexV1TaskInfrastructureSpecResponse infrastructureSpec,

            string mainClass,

            string mainJarFileUri,

            string pythonScriptFile,

            string sqlScript,

            string sqlScriptFile)
        {
            ArchiveUris = archiveUris;
            FileUris = fileUris;
            InfrastructureSpec = infrastructureSpec;
            MainClass = mainClass;
            MainJarFileUri = mainJarFileUri;
            PythonScriptFile = pythonScriptFile;
            SqlScript = sqlScript;
            SqlScriptFile = sqlScriptFile;
        }
    }
}
