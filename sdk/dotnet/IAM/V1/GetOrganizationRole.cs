// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IAM.V1
{
    public static class GetOrganizationRole
    {
        /// <summary>
        /// Gets the definition of a Role.
        /// </summary>
        public static Task<GetOrganizationRoleResult> InvokeAsync(GetOrganizationRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationRoleResult>("google-native:iam/v1:getOrganizationRole", args ?? new GetOrganizationRoleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the definition of a Role.
        /// </summary>
        public static Output<GetOrganizationRoleResult> Invoke(GetOrganizationRoleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrganizationRoleResult>("google-native:iam/v1:getOrganizationRole", args ?? new GetOrganizationRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationRoleArgs : Pulumi.InvokeArgs
    {
        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("roleId", required: true)]
        public string RoleId { get; set; } = null!;

        public GetOrganizationRoleArgs()
        {
        }
    }

    public sealed class GetOrganizationRoleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public GetOrganizationRoleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrganizationRoleResult
    {
        /// <summary>
        /// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
        /// </summary>
        public readonly bool Deleted;
        /// <summary>
        /// Optional. A human-readable description for the role.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Used to perform a consistent read-modify-write.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The names of the permissions this role grants when bound in an IAM policy.
        /// </summary>
        public readonly ImmutableArray<string> IncludedPermissions;
        /// <summary>
        /// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
        /// </summary>
        public readonly string Stage;
        /// <summary>
        /// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private GetOrganizationRoleResult(
            bool deleted,

            string description,

            string etag,

            ImmutableArray<string> includedPermissions,

            string name,

            string stage,

            string title)
        {
            Deleted = deleted;
            Description = description;
            Etag = etag;
            IncludedPermissions = includedPermissions;
            Name = name;
            Stage = stage;
            Title = title;
        }
    }
}
