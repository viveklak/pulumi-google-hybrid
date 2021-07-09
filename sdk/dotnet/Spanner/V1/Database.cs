// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Spanner.V1
{
    /// <summary>
    /// Creates a new Cloud Spanner database and starts to prepare it for serving. The returned long-running operation will have a name of the format `/operations/` and can be used to track preparation of the database. The metadata field type is CreateDatabaseMetadata. The response field type is Database, if successful.
    /// </summary>
    [GoogleNativeResourceType("google-native:spanner/v1:Database")]
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// If exists, the time at which the database creation started.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Earliest timestamp at which older versions of the data can be read. This value is continuously updated by Cloud Spanner and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
        /// </summary>
        [Output("earliestVersionTime")]
        public Output<string> EarliestVersionTime { get; private set; } = null!;

        /// <summary>
        /// For databases that are using customer managed encryption, this field contains the encryption configuration for the database. For databases that are using Google default or other types of encryption, this field is empty.
        /// </summary>
        [Output("encryptionConfig")]
        public Output<Outputs.EncryptionConfigResponse> EncryptionConfig { get; private set; } = null!;

        /// <summary>
        /// For databases that are using customer managed encryption, this field contains the encryption information for the database, such as encryption state and the Cloud KMS key versions that are in use. For databases that are using Google default or other types of encryption, this field is empty. This field is propagated lazily from the backend. There might be a delay from when a key version is being used and when it appears in this field.
        /// </summary>
        [Output("encryptionInfo")]
        public Output<ImmutableArray<Outputs.EncryptionInfoResponse>> EncryptionInfo { get; private set; } = null!;

        /// <summary>
        /// The name of the database. Values are of the form `projects//instances//databases/`, where `` is as specified in the `CREATE DATABASE` statement. This name can be passed to other API methods to identify the database.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Applicable only for restored databases. Contains information about the restore source.
        /// </summary>
        [Output("restoreInfo")]
        public Output<Outputs.RestoreInfoResponse> RestoreInfo { get; private set; } = null!;

        /// <summary>
        /// The current database state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The period in which Cloud Spanner retains all versions of data for the database. This is the same as the value of version_retention_period database option set using UpdateDatabaseDdl. Defaults to 1 hour, if not set.
        /// </summary>
        [Output("versionRetentionPeriod")]
        public Output<string> VersionRetentionPeriod { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("google-native:spanner/v1:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:spanner/v1:Database", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Database(name, id, options);
        }
    }

    public sealed class DatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `CREATE DATABASE` statement, which specifies the ID of the new database. The database ID must conform to the regular expression `a-z*[a-z0-9]` and be between 2 and 30 characters in length. If the database ID is a reserved word or if it contains a hyphen, the database ID must be enclosed in backticks (`` ` ``).
        /// </summary>
        [Input("createStatement", required: true)]
        public Input<string> CreateStatement { get; set; } = null!;

        /// <summary>
        /// Optional. The encryption configuration for the database. If this field is not specified, Cloud Spanner will encrypt/decrypt all data at rest using Google default encryption.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.EncryptionConfigArgs>? EncryptionConfig { get; set; }

        [Input("extraStatements")]
        private InputList<string>? _extraStatements;

        /// <summary>
        /// Optional. A list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc. These statements execute atomically with the creation of the database: if there is an error in any statement, the database is not created.
        /// </summary>
        public InputList<string> ExtraStatements
        {
            get => _extraStatements ?? (_extraStatements = new InputList<string>());
            set => _extraStatements = value;
        }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public DatabaseArgs()
        {
        }
    }
}
