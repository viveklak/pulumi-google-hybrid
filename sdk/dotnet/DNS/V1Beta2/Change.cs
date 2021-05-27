// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2
{
    /// <summary>
    /// Atomically updates the ResourceRecordSet collection.
    /// </summary>
    [GoogleNativeResourceType("google-native:dns/v1beta2:Change")]
    public partial class Change : Pulumi.CustomResource
    {
        /// <summary>
        /// Which ResourceRecordSets to add?
        /// </summary>
        [Output("additions")]
        public Output<ImmutableArray<Outputs.ResourceRecordSetResponse>> Additions { get; private set; } = null!;

        /// <summary>
        /// Which ResourceRecordSets to remove? Must match existing data exactly.
        /// </summary>
        [Output("deletions")]
        public Output<ImmutableArray<Outputs.ResourceRecordSetResponse>> Deletions { get; private set; } = null!;

        /// <summary>
        /// If the DNS queries for the zone will be served.
        /// </summary>
        [Output("isServing")]
        public Output<bool> IsServing { get; private set; } = null!;

        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The time that this operation was started by the server (output only). This is in RFC3339 text format.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// Status of the operation (output only). A status of "done" means that the request to update the authoritative servers has been sent, but the servers might not be updated yet.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Change resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Change(string name, ChangeArgs args, CustomResourceOptions? options = null)
            : base("google-native:dns/v1beta2:Change", name, args ?? new ChangeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Change(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dns/v1beta2:Change", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Change resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Change Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Change(name, id, options);
        }
    }

    public sealed class ChangeArgs : Pulumi.ResourceArgs
    {
        [Input("additions")]
        private InputList<Inputs.ResourceRecordSetArgs>? _additions;

        /// <summary>
        /// Which ResourceRecordSets to add?
        /// </summary>
        public InputList<Inputs.ResourceRecordSetArgs> Additions
        {
            get => _additions ?? (_additions = new InputList<Inputs.ResourceRecordSetArgs>());
            set => _additions = value;
        }

        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        [Input("deletions")]
        private InputList<Inputs.ResourceRecordSetArgs>? _deletions;

        /// <summary>
        /// Which ResourceRecordSets to remove? Must match existing data exactly.
        /// </summary>
        public InputList<Inputs.ResourceRecordSetArgs> Deletions
        {
            get => _deletions ?? (_deletions = new InputList<Inputs.ResourceRecordSetArgs>());
            set => _deletions = value;
        }

        /// <summary>
        /// Unique identifier for the resource; defined by the server (output only).
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// If the DNS queries for the zone will be served.
        /// </summary>
        [Input("isServing")]
        public Input<bool>? IsServing { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("managedZone", required: true)]
        public Input<string> ManagedZone { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The time that this operation was started by the server (output only). This is in RFC3339 text format.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Status of the operation (output only). A status of "done" means that the request to update the authoritative servers has been sent, but the servers might not be updated yet.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ChangeArgs()
        {
        }
    }
}
