// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class RbacPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the RbacPolicy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions")]
        private InputList<Inputs.PermissionArgs>? _permissions;

        /// <summary>
        /// The list of permissions.
        /// </summary>
        public InputList<Inputs.PermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.PermissionArgs>());
            set => _permissions = value;
        }

        [Input("principals")]
        private InputList<Inputs.PrincipalArgs>? _principals;

        /// <summary>
        /// The list of principals.
        /// </summary>
        public InputList<Inputs.PrincipalArgs> Principals
        {
            get => _principals ?? (_principals = new InputList<Inputs.PrincipalArgs>());
            set => _principals = value;
        }

        public RbacPolicyArgs()
        {
        }
    }
}
