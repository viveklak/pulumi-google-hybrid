// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1Beta
{
    public static class GetServicePerimeter
    {
        /// <summary>
        /// Get a Service Perimeter by resource name.
        /// </summary>
        public static Task<GetServicePerimeterResult> InvokeAsync(GetServicePerimeterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServicePerimeterResult>("google-native:accesscontextmanager/v1beta:getServicePerimeter", args ?? new GetServicePerimeterArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Service Perimeter by resource name.
        /// </summary>
        public static Output<GetServicePerimeterResult> Invoke(GetServicePerimeterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServicePerimeterResult>("google-native:accesscontextmanager/v1beta:getServicePerimeter", args ?? new GetServicePerimeterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServicePerimeterArgs : Pulumi.InvokeArgs
    {
        [Input("accessPolicyId", required: true)]
        public string AccessPolicyId { get; set; } = null!;

        [Input("servicePerimeterId", required: true)]
        public string ServicePerimeterId { get; set; } = null!;

        public GetServicePerimeterArgs()
        {
        }
    }

    public sealed class GetServicePerimeterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("accessPolicyId", required: true)]
        public Input<string> AccessPolicyId { get; set; } = null!;

        [Input("servicePerimeterId", required: true)]
        public Input<string> ServicePerimeterId { get; set; } = null!;

        public GetServicePerimeterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServicePerimeterResult
    {
        /// <summary>
        /// Description of the `ServicePerimeter` and its use. Does not affect behavior.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
        /// </summary>
        public readonly string PerimeterType;
        /// <summary>
        /// Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
        /// </summary>
        public readonly Outputs.ServicePerimeterConfigResponse Status;
        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private GetServicePerimeterResult(
            string description,

            string name,

            string perimeterType,

            Outputs.ServicePerimeterConfigResponse status,

            string title)
        {
            Description = description;
            Name = name;
            PerimeterType = perimeterType;
            Status = status;
            Title = title;
        }
    }
}
