// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Specifies Kerberos related configuration.
    /// </summary>
    [OutputType]
    public sealed class KerberosConfigResponse
    {
        /// <summary>
        /// Optional. The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship.
        /// </summary>
        public readonly string CrossRealmTrustAdminServer;
        /// <summary>
        /// Optional. The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship.
        /// </summary>
        public readonly string CrossRealmTrustKdc;
        /// <summary>
        /// Optional. The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust.
        /// </summary>
        public readonly string CrossRealmTrustRealm;
        /// <summary>
        /// Optional. The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster Kerberos realm and the remote trusted realm, in a cross realm trust relationship.
        /// </summary>
        public readonly string CrossRealmTrustSharedPasswordUri;
        /// <summary>
        /// Optional. Flag to indicate whether to Kerberize the cluster (default: false). Set this field to true to enable Kerberos on a cluster.
        /// </summary>
        public readonly bool EnableKerberos;
        /// <summary>
        /// Optional. The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database.
        /// </summary>
        public readonly string KdcDbKeyUri;
        /// <summary>
        /// Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc.
        /// </summary>
        public readonly string KeyPasswordUri;
        /// <summary>
        /// Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided keystore. For the self-signed certificate, this password is generated by Dataproc.
        /// </summary>
        public readonly string KeystorePasswordUri;
        /// <summary>
        /// Optional. The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.
        /// </summary>
        public readonly string KeystoreUri;
        /// <summary>
        /// Optional. The uri of the KMS key used to encrypt various sensitive files.
        /// </summary>
        public readonly string KmsKeyUri;
        /// <summary>
        /// Optional. The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm.
        /// </summary>
        public readonly string Realm;
        /// <summary>
        /// Optional. The Cloud Storage URI of a KMS encrypted file containing the root principal password.
        /// </summary>
        public readonly string RootPrincipalPasswordUri;
        /// <summary>
        /// Optional. The lifetime of the ticket granting ticket, in hours. If not specified, or user specifies 0, then default value 10 will be used.
        /// </summary>
        public readonly int TgtLifetimeHours;
        /// <summary>
        /// Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc.
        /// </summary>
        public readonly string TruststorePasswordUri;
        /// <summary>
        /// Optional. The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.
        /// </summary>
        public readonly string TruststoreUri;

        [OutputConstructor]
        private KerberosConfigResponse(
            string crossRealmTrustAdminServer,

            string crossRealmTrustKdc,

            string crossRealmTrustRealm,

            string crossRealmTrustSharedPasswordUri,

            bool enableKerberos,

            string kdcDbKeyUri,

            string keyPasswordUri,

            string keystorePasswordUri,

            string keystoreUri,

            string kmsKeyUri,

            string realm,

            string rootPrincipalPasswordUri,

            int tgtLifetimeHours,

            string truststorePasswordUri,

            string truststoreUri)
        {
            CrossRealmTrustAdminServer = crossRealmTrustAdminServer;
            CrossRealmTrustKdc = crossRealmTrustKdc;
            CrossRealmTrustRealm = crossRealmTrustRealm;
            CrossRealmTrustSharedPasswordUri = crossRealmTrustSharedPasswordUri;
            EnableKerberos = enableKerberos;
            KdcDbKeyUri = kdcDbKeyUri;
            KeyPasswordUri = keyPasswordUri;
            KeystorePasswordUri = keystorePasswordUri;
            KeystoreUri = keystoreUri;
            KmsKeyUri = kmsKeyUri;
            Realm = realm;
            RootPrincipalPasswordUri = rootPrincipalPasswordUri;
            TgtLifetimeHours = tgtLifetimeHours;
            TruststorePasswordUri = truststorePasswordUri;
            TruststoreUri = truststoreUri;
        }
    }
}
