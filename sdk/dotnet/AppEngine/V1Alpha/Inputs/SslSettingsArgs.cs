// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Alpha.Inputs
{

    /// <summary>
    /// SSL configuration for a DomainMapping resource.
    /// </summary>
    public sealed class SslSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will remove SSL support.By default, a managed certificate is automatically created for every domain mapping. To omit SSL support or to configure SSL manually, specify no_managed_certificate on a CREATE or UPDATE request. You must be authorized to administer the AuthorizedCertificate resource to manually map it to a DomainMapping resource. Example: 12345.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        public SslSettingsArgs()
        {
        }
    }
}
