// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Securitycenter.V1
{
    /// <summary>
    /// Creates a source.
    /// </summary>
    [GoogleNativeResourceType("google-native:securitycenter/v1:OrganizationSource")]
    public partial class OrganizationSource : Pulumi.CustomResource
    {
        /// <summary>
        /// The canonical name of the finding. It's either "organizations/{organization_id}/sources/{source_id}", "folders/{folder_id}/sources/{source_id}" or "projects/{project_number}/sources/{source_id}", depending on the closest CRM ancestor of the resource associated with the finding.
        /// </summary>
        [Output("canonicalName")]
        public Output<string> CanonicalName { get; private set; } = null!;

        /// <summary>
        /// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated or insecure libraries."
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationSource(string name, OrganizationSourceArgs args, CustomResourceOptions? options = null)
            : base("google-native:securitycenter/v1:OrganizationSource", name, args ?? new OrganizationSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationSource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:securitycenter/v1:OrganizationSource", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationSource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationSource(name, id, options);
        }
    }

    public sealed class OrganizationSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The canonical name of the finding. It's either "organizations/{organization_id}/sources/{source_id}", "folders/{folder_id}/sources/{source_id}" or "projects/{project_number}/sources/{source_id}", depending on the closest CRM ancestor of the resource associated with the finding.
        /// </summary>
        [Input("canonicalName")]
        public Input<string>? CanonicalName { get; set; }

        /// <summary>
        /// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated or insecure libraries."
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public OrganizationSourceArgs()
        {
        }
    }
}
