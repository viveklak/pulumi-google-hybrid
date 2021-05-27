// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.Alpha
{
    /// <summary>
    /// Creates a type provider.
    /// </summary>
    [GoogleNativeResourceType("google-native:deploymentmanager/alpha:TypeProvider")]
    public partial class TypeProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// Allows resource handling overrides for specific collections
        /// </summary>
        [Output("collectionOverrides")]
        public Output<ImmutableArray<Outputs.CollectionOverrideResponse>> CollectionOverrides { get; private set; } = null!;

        /// <summary>
        /// Credential used when interacting with this type.
        /// </summary>
        [Output("credential")]
        public Output<Outputs.CredentialResponse> Credential { get; private set; } = null!;

        /// <summary>
        /// List of up to 2 custom certificate authority roots to use for TLS authentication when making calls on behalf of this type provider. If set, TLS authentication will exclusively use these roots instead of relying on publicly trusted certificate authorities when validating TLS certificate authenticity. The certificates must be in base64-encoded PEM format. The maximum size of each certificate must not exceed 10KB.
        /// </summary>
        [Output("customCertificateAuthorityRoots")]
        public Output<ImmutableArray<string>> CustomCertificateAuthorityRoots { get; private set; } = null!;

        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Descriptor Url for the this type provider.
        /// </summary>
        [Output("descriptorUrl")]
        public Output<string> DescriptorUrl { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("insertTime")]
        public Output<string> InsertTime { get; private set; } = null!;

        /// <summary>
        /// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.TypeProviderLabelEntryResponse>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Operation that most recently ran, or is currently running, on this type provider.
        /// </summary>
        [Output("operation")]
        public Output<Outputs.OperationResponse> Operation { get; private set; } = null!;

        /// <summary>
        /// Options to apply when handling any resources in this service.
        /// </summary>
        [Output("options")]
        public Output<Outputs.OptionsResponse> Options { get; private set; } = null!;

        /// <summary>
        /// Self link for the type provider.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a TypeProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TypeProvider(string name, TypeProviderArgs args, CustomResourceOptions? options = null)
            : base("google-native:deploymentmanager/alpha:TypeProvider", name, args ?? new TypeProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TypeProvider(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:deploymentmanager/alpha:TypeProvider", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing TypeProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TypeProvider Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TypeProvider(name, id, options);
        }
    }

    public sealed class TypeProviderArgs : Pulumi.ResourceArgs
    {
        [Input("collectionOverrides")]
        private InputList<Inputs.CollectionOverrideArgs>? _collectionOverrides;

        /// <summary>
        /// Allows resource handling overrides for specific collections
        /// </summary>
        public InputList<Inputs.CollectionOverrideArgs> CollectionOverrides
        {
            get => _collectionOverrides ?? (_collectionOverrides = new InputList<Inputs.CollectionOverrideArgs>());
            set => _collectionOverrides = value;
        }

        /// <summary>
        /// Credential used when interacting with this type.
        /// </summary>
        [Input("credential")]
        public Input<Inputs.CredentialArgs>? Credential { get; set; }

        [Input("customCertificateAuthorityRoots")]
        private InputList<string>? _customCertificateAuthorityRoots;

        /// <summary>
        /// List of up to 2 custom certificate authority roots to use for TLS authentication when making calls on behalf of this type provider. If set, TLS authentication will exclusively use these roots instead of relying on publicly trusted certificate authorities when validating TLS certificate authenticity. The certificates must be in base64-encoded PEM format. The maximum size of each certificate must not exceed 10KB.
        /// </summary>
        public InputList<string> CustomCertificateAuthorityRoots
        {
            get => _customCertificateAuthorityRoots ?? (_customCertificateAuthorityRoots = new InputList<string>());
            set => _customCertificateAuthorityRoots = value;
        }

        /// <summary>
        /// An optional textual description of the resource; provided by the client when the resource is created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Descriptor Url for the this type provider.
        /// </summary>
        [Input("descriptorUrl")]
        public Input<string>? DescriptorUrl { get; set; }

        /// <summary>
        /// Unique identifier for the resource defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("insertTime")]
        public Input<string>? InsertTime { get; set; }

        [Input("labels")]
        private InputList<Inputs.TypeProviderLabelEntryArgs>? _labels;

        /// <summary>
        /// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`
        /// </summary>
        public InputList<Inputs.TypeProviderLabelEntryArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.TypeProviderLabelEntryArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Operation that most recently ran, or is currently running, on this type provider.
        /// </summary>
        [Input("operation")]
        public Input<Inputs.OperationArgs>? Operation { get; set; }

        /// <summary>
        /// Options to apply when handling any resources in this service.
        /// </summary>
        [Input("options")]
        public Input<Inputs.OptionsArgs>? Options { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Self link for the type provider.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public TypeProviderArgs()
        {
        }
    }
}
