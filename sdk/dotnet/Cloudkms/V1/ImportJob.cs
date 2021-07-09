// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1
{
    /// <summary>
    /// Create a new ImportJob within a KeyRing. ImportJob.import_method is required.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudkms/v1:ImportJob")]
    public partial class ImportJob : Pulumi.CustomResource
    {
        /// <summary>
        /// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen ImportMethod is one with a protection level of HSM.
        /// </summary>
        [Output("attestation")]
        public Output<Outputs.KeyOperationAttestationResponse> Attestation { get; private set; } = null!;

        /// <summary>
        /// The time at which this ImportJob was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time this ImportJob expired. Only present if state is EXPIRED.
        /// </summary>
        [Output("expireEventTime")]
        public Output<string> ExpireEventTime { get; private set; } = null!;

        /// <summary>
        /// The time at which this ImportJob is scheduled for expiration and can no longer be used to import key material.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// The time this ImportJob's key material was generated.
        /// </summary>
        [Output("generateTime")]
        public Output<string> GenerateTime { get; private set; } = null!;

        /// <summary>
        /// Immutable. The wrapping method to be used for incoming key material.
        /// </summary>
        [Output("importMethod")]
        public Output<string> ImportMethod { get; private set; } = null!;

        /// <summary>
        /// The resource name for this ImportJob in the format `projects/*/locations/*/keyRings/*/importJobs/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
        /// </summary>
        [Output("protectionLevel")]
        public Output<string> ProtectionLevel { get; private set; } = null!;

        /// <summary>
        /// The public key with which to wrap key material prior to import. Only returned if state is ACTIVE.
        /// </summary>
        [Output("publicKey")]
        public Output<Outputs.WrappingPublicKeyResponse> PublicKey { get; private set; } = null!;

        /// <summary>
        /// The current state of the ImportJob, indicating if it can be used.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a ImportJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImportJob(string name, ImportJobArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudkms/v1:ImportJob", name, args ?? new ImportJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImportJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudkms/v1:ImportJob", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ImportJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImportJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ImportJob(name, id, options);
        }
    }

    public sealed class ImportJobArgs : Pulumi.ResourceArgs
    {
        [Input("importJobId", required: true)]
        public Input<string> ImportJobId { get; set; } = null!;

        /// <summary>
        /// Immutable. The wrapping method to be used for incoming key material.
        /// </summary>
        [Input("importMethod", required: true)]
        public Input<Pulumi.GoogleNative.Cloudkms.V1.ImportJobImportMethod> ImportMethod { get; set; } = null!;

        [Input("keyRingId", required: true)]
        public Input<string> KeyRingId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
        /// </summary>
        [Input("protectionLevel", required: true)]
        public Input<Pulumi.GoogleNative.Cloudkms.V1.ImportJobProtectionLevel> ProtectionLevel { get; set; } = null!;

        public ImportJobArgs()
        {
        }
    }
}
