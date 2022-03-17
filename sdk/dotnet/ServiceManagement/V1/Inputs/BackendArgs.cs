// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Inputs
{

    /// <summary>
    /// `Backend` defines the backend configuration for a service.
    /// </summary>
    public sealed class BackendArgs : Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.BackendRuleArgs>? _rules;

        /// <summary>
        /// A list of API backend rules that apply to individual API methods. **NOTE:** All service configuration rules follow "last one wins" order.
        /// </summary>
        public InputList<Inputs.BackendRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BackendRuleArgs>());
            set => _rules = value;
        }

        public BackendArgs()
        {
        }
    }
}
