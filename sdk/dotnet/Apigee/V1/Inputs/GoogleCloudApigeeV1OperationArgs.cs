// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// Operation represents the pairing of REST resource path and the actions (verbs) allowed on the resource path.
    /// </summary>
    public sealed class GoogleCloudApigeeV1OperationArgs : Pulumi.ResourceArgs
    {
        [Input("methods")]
        private InputList<string>? _methods;

        /// <summary>
        /// methods refers to the REST verbs as in https://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html. When none specified, all verb types are allowed.
        /// </summary>
        public InputList<string> Methods
        {
            get => _methods ?? (_methods = new InputList<string>());
            set => _methods = value;
        }

        /// <summary>
        /// resource represents REST resource path associated with the proxy/remote service.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        public GoogleCloudApigeeV1OperationArgs()
        {
        }
    }
}
