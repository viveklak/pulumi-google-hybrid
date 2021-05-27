// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SourceRepo.V1
{
    /// <summary>
    /// Creates a repo in the given project with the given name. If the named repository already exists, `CreateRepo` returns `ALREADY_EXISTS`.
    /// </summary>
    [GoogleNativeResourceType("google-native:sourcerepo/v1:Repo")]
    public partial class Repo : Pulumi.CustomResource
    {
        /// <summary>
        /// How this repository mirrors a repository managed by another service. Read-only field.
        /// </summary>
        [Output("mirrorConfig")]
        public Output<Outputs.MirrorConfigResponse> MirrorConfig { get; private set; } = null!;

        /// <summary>
        /// Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
        /// </summary>
        [Output("pubsubConfigs")]
        public Output<ImmutableDictionary<string, string>> PubsubConfigs { get; private set; } = null!;

        /// <summary>
        /// The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
        /// </summary>
        [Output("size")]
        public Output<string> Size { get; private set; } = null!;

        /// <summary>
        /// URL to clone the repository from Google Cloud Source Repositories. Read-only field.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Repo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repo(string name, RepoArgs args, CustomResourceOptions? options = null)
            : base("google-native:sourcerepo/v1:Repo", name, args ?? new RepoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repo(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:sourcerepo/v1:Repo", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Repo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repo Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Repo(name, id, options);
        }
    }

    public sealed class RepoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How this repository mirrors a repository managed by another service. Read-only field.
        /// </summary>
        [Input("mirrorConfig")]
        public Input<Inputs.MirrorConfigArgs>? MirrorConfig { get; set; }

        /// <summary>
        /// Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("pubsubConfigs")]
        private InputMap<string>? _pubsubConfigs;

        /// <summary>
        /// How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
        /// </summary>
        public InputMap<string> PubsubConfigs
        {
            get => _pubsubConfigs ?? (_pubsubConfigs = new InputMap<string>());
            set => _pubsubConfigs = value;
        }

        /// <summary>
        /// The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// URL to clone the repository from Google Cloud Source Repositories. Read-only field.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public RepoArgs()
        {
        }
    }
}
