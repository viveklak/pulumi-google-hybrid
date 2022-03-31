// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Privateca.V1
{
    /// <summary>
    /// The log type that this config enables.
    /// </summary>
    [EnumType]
    public readonly struct AuditLogConfigLogType : IEquatable<AuditLogConfigLogType>
    {
        private readonly string _value;

        private AuditLogConfigLogType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default case. Should never be this.
        /// </summary>
        public static AuditLogConfigLogType LogTypeUnspecified { get; } = new AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED");
        /// <summary>
        /// Admin reads. Example: CloudIAM getIamPolicy
        /// </summary>
        public static AuditLogConfigLogType AdminRead { get; } = new AuditLogConfigLogType("ADMIN_READ");
        /// <summary>
        /// Data writes. Example: CloudSQL Users create
        /// </summary>
        public static AuditLogConfigLogType DataWrite { get; } = new AuditLogConfigLogType("DATA_WRITE");
        /// <summary>
        /// Data reads. Example: CloudSQL Users list
        /// </summary>
        public static AuditLogConfigLogType DataRead { get; } = new AuditLogConfigLogType("DATA_READ");

        public static bool operator ==(AuditLogConfigLogType left, AuditLogConfigLogType right) => left.Equals(right);
        public static bool operator !=(AuditLogConfigLogType left, AuditLogConfigLogType right) => !left.Equals(right);

        public static explicit operator string(AuditLogConfigLogType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AuditLogConfigLogType other && Equals(other);
        public bool Equals(AuditLogConfigLogType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. Immutable. The Tier of this CaPool.
    /// </summary>
    [EnumType]
    public readonly struct CaPoolTier : IEquatable<CaPoolTier>
    {
        private readonly string _value;

        private CaPoolTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static CaPoolTier TierUnspecified { get; } = new CaPoolTier("TIER_UNSPECIFIED");
        /// <summary>
        /// Enterprise tier.
        /// </summary>
        public static CaPoolTier Enterprise { get; } = new CaPoolTier("ENTERPRISE");
        /// <summary>
        /// DevOps tier.
        /// </summary>
        public static CaPoolTier Devops { get; } = new CaPoolTier("DEVOPS");

        public static bool operator ==(CaPoolTier left, CaPoolTier right) => left.Equals(right);
        public static bool operator !=(CaPoolTier left, CaPoolTier right) => !left.Equals(right);

        public static explicit operator string(CaPoolTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CaPoolTier other && Equals(other);
        public bool Equals(CaPoolTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. Immutable. The Type of this CertificateAuthority.
    /// </summary>
    [EnumType]
    public readonly struct CertificateAuthorityType : IEquatable<CertificateAuthorityType>
    {
        private readonly string _value;

        private CertificateAuthorityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static CertificateAuthorityType TypeUnspecified { get; } = new CertificateAuthorityType("TYPE_UNSPECIFIED");
        /// <summary>
        /// Self-signed CA.
        /// </summary>
        public static CertificateAuthorityType SelfSigned { get; } = new CertificateAuthorityType("SELF_SIGNED");
        /// <summary>
        /// Subordinate CA. Could be issued by a Private CA CertificateAuthority or an unmanaged CA.
        /// </summary>
        public static CertificateAuthorityType Subordinate { get; } = new CertificateAuthorityType("SUBORDINATE");

        public static bool operator ==(CertificateAuthorityType left, CertificateAuthorityType right) => left.Equals(right);
        public static bool operator !=(CertificateAuthorityType left, CertificateAuthorityType right) => !left.Equals(right);

        public static explicit operator string(CertificateAuthorityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateAuthorityType other && Equals(other);
        public bool Equals(CertificateAuthorityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CertificateExtensionConstraintsKnownExtensionsItem : IEquatable<CertificateExtensionConstraintsKnownExtensionsItem>
    {
        private readonly string _value;

        private CertificateExtensionConstraintsKnownExtensionsItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static CertificateExtensionConstraintsKnownExtensionsItem KnownCertificateExtensionUnspecified { get; } = new CertificateExtensionConstraintsKnownExtensionsItem("KNOWN_CERTIFICATE_EXTENSION_UNSPECIFIED");
        /// <summary>
        /// Refers to a certificate's Key Usage extension, as described in [RFC 5280 section 4.2.1.3](https://tools.ietf.org/html/rfc5280#section-4.2.1.3). This corresponds to the KeyUsage.base_key_usage field.
        /// </summary>
        public static CertificateExtensionConstraintsKnownExtensionsItem BaseKeyUsage { get; } = new CertificateExtensionConstraintsKnownExtensionsItem("BASE_KEY_USAGE");
        /// <summary>
        /// Refers to a certificate's Extended Key Usage extension, as described in [RFC 5280 section 4.2.1.12](https://tools.ietf.org/html/rfc5280#section-4.2.1.12). This corresponds to the KeyUsage.extended_key_usage message.
        /// </summary>
        public static CertificateExtensionConstraintsKnownExtensionsItem ExtendedKeyUsage { get; } = new CertificateExtensionConstraintsKnownExtensionsItem("EXTENDED_KEY_USAGE");
        /// <summary>
        /// Refers to a certificate's Basic Constraints extension, as described in [RFC 5280 section 4.2.1.9](https://tools.ietf.org/html/rfc5280#section-4.2.1.9). This corresponds to the X509Parameters.ca_options field.
        /// </summary>
        public static CertificateExtensionConstraintsKnownExtensionsItem CaOptions { get; } = new CertificateExtensionConstraintsKnownExtensionsItem("CA_OPTIONS");
        /// <summary>
        /// Refers to a certificate's Policy object identifiers, as described in [RFC 5280 section 4.2.1.4](https://tools.ietf.org/html/rfc5280#section-4.2.1.4). This corresponds to the X509Parameters.policy_ids field.
        /// </summary>
        public static CertificateExtensionConstraintsKnownExtensionsItem PolicyIds { get; } = new CertificateExtensionConstraintsKnownExtensionsItem("POLICY_IDS");
        /// <summary>
        /// Refers to OCSP servers in a certificate's Authority Information Access extension, as described in [RFC 5280 section 4.2.2.1](https://tools.ietf.org/html/rfc5280#section-4.2.2.1), This corresponds to the X509Parameters.aia_ocsp_servers field.
        /// </summary>
        public static CertificateExtensionConstraintsKnownExtensionsItem AiaOcspServers { get; } = new CertificateExtensionConstraintsKnownExtensionsItem("AIA_OCSP_SERVERS");

        public static bool operator ==(CertificateExtensionConstraintsKnownExtensionsItem left, CertificateExtensionConstraintsKnownExtensionsItem right) => left.Equals(right);
        public static bool operator !=(CertificateExtensionConstraintsKnownExtensionsItem left, CertificateExtensionConstraintsKnownExtensionsItem right) => !left.Equals(right);

        public static explicit operator string(CertificateExtensionConstraintsKnownExtensionsItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateExtensionConstraintsKnownExtensionsItem other && Equals(other);
        public bool Equals(CertificateExtensionConstraintsKnownExtensionsItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Immutable. Specifies how the Certificate's identity fields are to be decided. If this is omitted, the `DEFAULT` subject mode will be used.
    /// </summary>
    [EnumType]
    public readonly struct CertificateSubjectMode : IEquatable<CertificateSubjectMode>
    {
        private readonly string _value;

        private CertificateSubjectMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static CertificateSubjectMode SubjectRequestModeUnspecified { get; } = new CertificateSubjectMode("SUBJECT_REQUEST_MODE_UNSPECIFIED");
        /// <summary>
        /// The default mode used in most cases. Indicates that the certificate's Subject and/or SubjectAltNames are specified in the certificate request. This mode requires the caller to have the `privateca.certificates.create` permission.
        /// </summary>
        public static CertificateSubjectMode Default { get; } = new CertificateSubjectMode("DEFAULT");
        /// <summary>
        /// A mode reserved for special cases. Indicates that the certificate should have one or more SPIFFE SubjectAltNames set by the service based on the caller's identity. This mode will ignore any explicitly specified Subject and/or SubjectAltNames in the certificate request. This mode requires the caller to have the `privateca.certificates.createForSelf` permission.
        /// </summary>
        public static CertificateSubjectMode ReflectedSpiffe { get; } = new CertificateSubjectMode("REFLECTED_SPIFFE");

        public static bool operator ==(CertificateSubjectMode left, CertificateSubjectMode right) => left.Equals(right);
        public static bool operator !=(CertificateSubjectMode left, CertificateSubjectMode right) => !left.Equals(right);

        public static explicit operator string(CertificateSubjectMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateSubjectMode other && Equals(other);
        public bool Equals(CertificateSubjectMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional. A signature algorithm that must be used. If this is omitted, any EC-based signature algorithm will be allowed.
    /// </summary>
    [EnumType]
    public readonly struct EcKeyTypeSignatureAlgorithm : IEquatable<EcKeyTypeSignatureAlgorithm>
    {
        private readonly string _value;

        private EcKeyTypeSignatureAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified. Signifies that any signature algorithm may be used.
        /// </summary>
        public static EcKeyTypeSignatureAlgorithm EcSignatureAlgorithmUnspecified { get; } = new EcKeyTypeSignatureAlgorithm("EC_SIGNATURE_ALGORITHM_UNSPECIFIED");
        /// <summary>
        /// Refers to the Elliptic Curve Digital Signature Algorithm over the NIST P-256 curve.
        /// </summary>
        public static EcKeyTypeSignatureAlgorithm EcdsaP256 { get; } = new EcKeyTypeSignatureAlgorithm("ECDSA_P256");
        /// <summary>
        /// Refers to the Elliptic Curve Digital Signature Algorithm over the NIST P-384 curve.
        /// </summary>
        public static EcKeyTypeSignatureAlgorithm EcdsaP384 { get; } = new EcKeyTypeSignatureAlgorithm("ECDSA_P384");
        /// <summary>
        /// Refers to the Edwards-curve Digital Signature Algorithm over curve 25519, as described in RFC 8410.
        /// </summary>
        public static EcKeyTypeSignatureAlgorithm Eddsa25519 { get; } = new EcKeyTypeSignatureAlgorithm("EDDSA_25519");

        public static bool operator ==(EcKeyTypeSignatureAlgorithm left, EcKeyTypeSignatureAlgorithm right) => left.Equals(right);
        public static bool operator !=(EcKeyTypeSignatureAlgorithm left, EcKeyTypeSignatureAlgorithm right) => !left.Equals(right);

        public static explicit operator string(EcKeyTypeSignatureAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EcKeyTypeSignatureAlgorithm other && Equals(other);
        public bool Equals(EcKeyTypeSignatureAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The algorithm to use for creating a managed Cloud KMS key for a for a simplified experience. All managed keys will be have their ProtectionLevel as `HSM`.
    /// </summary>
    [EnumType]
    public readonly struct KeyVersionSpecAlgorithm : IEquatable<KeyVersionSpecAlgorithm>
    {
        private readonly string _value;

        private KeyVersionSpecAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static KeyVersionSpecAlgorithm SignHashAlgorithmUnspecified { get; } = new KeyVersionSpecAlgorithm("SIGN_HASH_ALGORITHM_UNSPECIFIED");
        /// <summary>
        /// maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PSS_2048_SHA256
        /// </summary>
        public static KeyVersionSpecAlgorithm RsaPss2048Sha256 { get; } = new KeyVersionSpecAlgorithm("RSA_PSS_2048_SHA256");
        /// <summary>
        /// maps to CryptoKeyVersionAlgorithm. RSA_SIGN_PSS_3072_SHA256
        /// </summary>
        public static KeyVersionSpecAlgorithm RsaPss3072Sha256 { get; } = new KeyVersionSpecAlgorithm("RSA_PSS_3072_SHA256");
        /// <summary>
        /// maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PSS_4096_SHA256
        /// </summary>
        public static KeyVersionSpecAlgorithm RsaPss4096Sha256 { get; } = new KeyVersionSpecAlgorithm("RSA_PSS_4096_SHA256");
        /// <summary>
        /// maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PKCS1_2048_SHA256
        /// </summary>
        public static KeyVersionSpecAlgorithm RsaPkcs12048Sha256 { get; } = new KeyVersionSpecAlgorithm("RSA_PKCS1_2048_SHA256");
        /// <summary>
        /// maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PKCS1_3072_SHA256
        /// </summary>
        public static KeyVersionSpecAlgorithm RsaPkcs13072Sha256 { get; } = new KeyVersionSpecAlgorithm("RSA_PKCS1_3072_SHA256");
        /// <summary>
        /// maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PKCS1_4096_SHA256
        /// </summary>
        public static KeyVersionSpecAlgorithm RsaPkcs14096Sha256 { get; } = new KeyVersionSpecAlgorithm("RSA_PKCS1_4096_SHA256");
        /// <summary>
        /// maps to CryptoKeyVersionAlgorithm.EC_SIGN_P256_SHA256
        /// </summary>
        public static KeyVersionSpecAlgorithm EcP256Sha256 { get; } = new KeyVersionSpecAlgorithm("EC_P256_SHA256");
        /// <summary>
        /// maps to CryptoKeyVersionAlgorithm.EC_SIGN_P384_SHA384
        /// </summary>
        public static KeyVersionSpecAlgorithm EcP384Sha384 { get; } = new KeyVersionSpecAlgorithm("EC_P384_SHA384");

        public static bool operator ==(KeyVersionSpecAlgorithm left, KeyVersionSpecAlgorithm right) => left.Equals(right);
        public static bool operator !=(KeyVersionSpecAlgorithm left, KeyVersionSpecAlgorithm right) => !left.Equals(right);

        public static explicit operator string(KeyVersionSpecAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyVersionSpecAlgorithm other && Equals(other);
        public bool Equals(KeyVersionSpecAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. The format of the public key.
    /// </summary>
    [EnumType]
    public readonly struct PublicKeyFormat : IEquatable<PublicKeyFormat>
    {
        private readonly string _value;

        private PublicKeyFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default unspecified value.
        /// </summary>
        public static PublicKeyFormat KeyFormatUnspecified { get; } = new PublicKeyFormat("KEY_FORMAT_UNSPECIFIED");
        /// <summary>
        /// The key is PEM-encoded as defined in [RFC 7468](https://tools.ietf.org/html/rfc7468). It can be any of the following: a PEM-encoded PKCS#1/RFC 3447 RSAPublicKey structure, an RFC 5280 [SubjectPublicKeyInfo](https://tools.ietf.org/html/rfc5280#section-4.1) or a PEM-encoded X.509 certificate signing request (CSR). If a [SubjectPublicKeyInfo](https://tools.ietf.org/html/rfc5280#section-4.1) is specified, it can contain a A PEM-encoded PKCS#1/RFC 3447 RSAPublicKey or a NIST P-256/secp256r1/prime256v1 or P-384 key. If a CSR is specified, it will used solely for the purpose of extracting the public key. When generated by the service, it will always be an RFC 5280 [SubjectPublicKeyInfo](https://tools.ietf.org/html/rfc5280#section-4.1) structure containing an algorithm identifier and a key.
        /// </summary>
        public static PublicKeyFormat Pem { get; } = new PublicKeyFormat("PEM");

        public static bool operator ==(PublicKeyFormat left, PublicKeyFormat right) => left.Equals(right);
        public static bool operator !=(PublicKeyFormat left, PublicKeyFormat right) => !left.Equals(right);

        public static explicit operator string(PublicKeyFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PublicKeyFormat other && Equals(other);
        public bool Equals(PublicKeyFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
