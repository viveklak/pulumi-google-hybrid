// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Managedidentities.V1alpha1
{
    /// <summary>
    /// Creates a Microsoft AD Domain in a given project. Operation
    /// </summary>
    [GoogleCloudResourceType("google-cloud:managedidentities/v1alpha1:Domain")]
    public partial class Domain : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:managedidentities/v1alpha1:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:managedidentities/v1alpha1:Domain", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, options);
        }
    }

    public sealed class DomainArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        /// </summary>
        [Input("auditLogsEnabled")]
        public Input<bool>? AuditLogsEnabled { get; set; }

        [Input("authorizedNetworks")]
        private InputList<string>? _authorizedNetworks;

        /// <summary>
        /// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
        /// </summary>
        public InputList<string> AuthorizedNetworks
        {
            get => _authorizedNetworks ?? (_authorizedNetworks = new InputList<string>());
            set => _authorizedNetworks = value;
        }

        /// <summary>
        /// Output only. The time the instance was created. Synthetic field is populated automatically by CCFE. go/ccfe-synthetic-field-user-guide
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("domainsId", required: true)]
        public Input<string> DomainsId { get; set; } = null!;

        /// <summary>
        /// Output only. Fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory that is set up on an internal network.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Resource labels to represent user provided metadata
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
        /// </summary>
        [Input("managedIdentitiesAdminName")]
        public Input<string>? ManagedIdentitiesAdminName { get; set; }

        /// <summary>
        /// Output only. Unique name of the domain in this scope including projects and location using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        /// </summary>
        [Input("reservedIpRange")]
        public Input<string>? ReservedIpRange { get; set; }

        /// <summary>
        /// Output only. The current state of this domain.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Output only. Additional information about the current status of this domain, if available.
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        [Input("trusts")]
        private InputList<Inputs.TrustArgs>? _trusts;

        /// <summary>
        /// Output only. The current trusts associated with the domain.
        /// </summary>
        public InputList<Inputs.TrustArgs> Trusts
        {
            get => _trusts ?? (_trusts = new InputList<Inputs.TrustArgs>());
            set => _trusts = value;
        }

        /// <summary>
        /// Output only. Last update time. Synthetic field is populated automatically by CCFE.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public DomainArgs()
        {
        }
    }
}
