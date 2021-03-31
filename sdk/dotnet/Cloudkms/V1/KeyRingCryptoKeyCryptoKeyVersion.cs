// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Cloudkms.V1
{
    /// <summary>
    /// Create a new CryptoKeyVersion in a CryptoKey. The server will assign the next sequential id. If unset, state will be set to ENABLED.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:cloudkms/v1:KeyRingCryptoKeyCryptoKeyVersion")]
    public partial class KeyRingCryptoKeyCryptoKeyVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a KeyRingCryptoKeyCryptoKeyVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyRingCryptoKeyCryptoKeyVersion(string name, KeyRingCryptoKeyCryptoKeyVersionArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:cloudkms/v1:KeyRingCryptoKeyCryptoKeyVersion", name, args ?? new KeyRingCryptoKeyCryptoKeyVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyRingCryptoKeyCryptoKeyVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:cloudkms/v1:KeyRingCryptoKeyCryptoKeyVersion", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing KeyRingCryptoKeyCryptoKeyVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyRingCryptoKeyCryptoKeyVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KeyRingCryptoKeyCryptoKeyVersion(name, id, options);
        }
    }

    public sealed class KeyRingCryptoKeyCryptoKeyVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Output only. Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only provided for key versions with protection_level HSM.
        /// </summary>
        [Input("attestation")]
        public Input<Inputs.KeyOperationAttestationArgs>? Attestation { get; set; }

        /// <summary>
        /// Output only. The time at which this CryptoKeyVersion was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("cryptoKeysId", required: true)]
        public Input<string> CryptoKeysId { get; set; } = null!;

        /// <summary>
        /// Output only. The time this CryptoKeyVersion's key material was destroyed. Only present if state is DESTROYED.
        /// </summary>
        [Input("destroyEventTime")]
        public Input<string>? DestroyEventTime { get; set; }

        /// <summary>
        /// Output only. The time this CryptoKeyVersion's key material is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.
        /// </summary>
        [Input("destroyTime")]
        public Input<string>? DestroyTime { get; set; }

        /// <summary>
        /// ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
        /// </summary>
        [Input("externalProtectionLevelOptions")]
        public Input<Inputs.ExternalProtectionLevelOptionsArgs>? ExternalProtectionLevelOptions { get; set; }

        /// <summary>
        /// Output only. The time this CryptoKeyVersion's key material was generated.
        /// </summary>
        [Input("generateTime")]
        public Input<string>? GenerateTime { get; set; }

        /// <summary>
        /// Output only. The root cause of an import failure. Only present if state is IMPORT_FAILED.
        /// </summary>
        [Input("importFailureReason")]
        public Input<string>? ImportFailureReason { get; set; }

        /// <summary>
        /// Output only. The name of the ImportJob used to import this CryptoKeyVersion. Only present if the underlying key material was imported.
        /// </summary>
        [Input("importJob")]
        public Input<string>? ImportJob { get; set; }

        /// <summary>
        /// Output only. The time at which this CryptoKeyVersion's key material was imported.
        /// </summary>
        [Input("importTime")]
        public Input<string>? ImportTime { get; set; }

        [Input("keyRingsId", required: true)]
        public Input<string> KeyRingsId { get; set; } = null!;

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        /// <summary>
        /// Output only. The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// Output only. The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        /// </summary>
        [Input("protectionLevel")]
        public Input<string>? ProtectionLevel { get; set; }

        /// <summary>
        /// The current state of the CryptoKeyVersion.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public KeyRingCryptoKeyCryptoKeyVersionArgs()
        {
        }
    }
}
