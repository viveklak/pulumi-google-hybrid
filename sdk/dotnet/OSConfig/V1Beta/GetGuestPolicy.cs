// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta
{
    public static class GetGuestPolicy
    {
        /// <summary>
        /// Get an OS Config guest policy.
        /// </summary>
        public static Task<GetGuestPolicyResult> InvokeAsync(GetGuestPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGuestPolicyResult>("google-native:osconfig/v1beta:getGuestPolicy", args ?? new GetGuestPolicyArgs(), options.WithVersion());
    }


    public sealed class GetGuestPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("guestPolicyId", required: true)]
        public string GuestPolicyId { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetGuestPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGuestPolicyResult
    {
        /// <summary>
        /// Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        /// </summary>
        public readonly Outputs.AssignmentResponse Assignment;
        /// <summary>
        /// Time this guest policy was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Description of the guest policy. Length of the description is limited to 1024 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The etag for this guest policy. If this is provided on update, it must match the server's etag.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
        /// </summary>
        public readonly ImmutableArray<Outputs.PackageRepositoryResponse> PackageRepositories;
        /// <summary>
        /// The software packages to be managed by this policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.PackageResponse> Packages;
        /// <summary>
        /// A list of Recipes to install on the VM instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.SoftwareRecipeResponse> Recipes;
        /// <summary>
        /// Last time this guest policy was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetGuestPolicyResult(
            Outputs.AssignmentResponse assignment,

            string createTime,

            string description,

            string etag,

            string name,

            ImmutableArray<Outputs.PackageRepositoryResponse> packageRepositories,

            ImmutableArray<Outputs.PackageResponse> packages,

            ImmutableArray<Outputs.SoftwareRecipeResponse> recipes,

            string updateTime)
        {
            Assignment = assignment;
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            Name = name;
            PackageRepositories = packageRepositories;
            Packages = packages;
            Recipes = recipes;
            UpdateTime = updateTime;
        }
    }
}
