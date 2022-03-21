// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Inputs
{

    /// <summary>
    /// These values are used to create the distinguished name and subject alternative name fields in an X.509 certificate.
    /// </summary>
    public sealed class SubjectConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains distinguished name fields such as the common name, location and organization.
        /// </summary>
        [Input("subject", required: true)]
        public Input<Inputs.SubjectArgs> Subject { get; set; } = null!;

        /// <summary>
        /// Optional. The subject alternative name fields.
        /// </summary>
        [Input("subjectAltName")]
        public Input<Inputs.SubjectAltNamesArgs>? SubjectAltName { get; set; }

        public SubjectConfigArgs()
        {
        }
    }
}
