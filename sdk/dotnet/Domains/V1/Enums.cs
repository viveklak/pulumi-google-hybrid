// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Domains.V1
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
    /// Required. Privacy setting for the contacts associated with the `Registration`.
    /// </summary>
    [EnumType]
    public readonly struct ContactSettingsPrivacy : IEquatable<ContactSettingsPrivacy>
    {
        private readonly string _value;

        private ContactSettingsPrivacy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The contact privacy settings are undefined.
        /// </summary>
        public static ContactSettingsPrivacy ContactPrivacyUnspecified { get; } = new ContactSettingsPrivacy("CONTACT_PRIVACY_UNSPECIFIED");
        /// <summary>
        /// All the data from `ContactSettings` is publicly available. When setting this option, you must also provide a `PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT` in the `contact_notices` field of the request.
        /// </summary>
        public static ContactSettingsPrivacy PublicContactData { get; } = new ContactSettingsPrivacy("PUBLIC_CONTACT_DATA");
        /// <summary>
        /// None of the data from `ContactSettings` is publicly available. Instead, proxy contact data is published for your domain. Email sent to the proxy email address is forwarded to the registrant's email address. Cloud Domains provides this privacy proxy service at no additional cost.
        /// </summary>
        public static ContactSettingsPrivacy PrivateContactData { get; } = new ContactSettingsPrivacy("PRIVATE_CONTACT_DATA");
        /// <summary>
        /// Some data from `ContactSettings` is publicly available. The actual information redacted depends on the domain. For details, see [the registration privacy article](https://support.google.com/domains/answer/3251242).
        /// </summary>
        public static ContactSettingsPrivacy RedactedContactData { get; } = new ContactSettingsPrivacy("REDACTED_CONTACT_DATA");

        public static bool operator ==(ContactSettingsPrivacy left, ContactSettingsPrivacy right) => left.Equals(right);
        public static bool operator !=(ContactSettingsPrivacy left, ContactSettingsPrivacy right) => !left.Equals(right);

        public static explicit operator string(ContactSettingsPrivacy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContactSettingsPrivacy other && Equals(other);
        public bool Equals(ContactSettingsPrivacy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The algorithm used to generate the referenced DNSKEY.
    /// </summary>
    [EnumType]
    public readonly struct DsRecordAlgorithm : IEquatable<DsRecordAlgorithm>
    {
        private readonly string _value;

        private DsRecordAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The algorithm is unspecified.
        /// </summary>
        public static DsRecordAlgorithm AlgorithmUnspecified { get; } = new DsRecordAlgorithm("ALGORITHM_UNSPECIFIED");
        /// <summary>
        /// RSA/MD5. Cannot be used for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Rsamd5 { get; } = new DsRecordAlgorithm("RSAMD5");
        /// <summary>
        /// Diffie-Hellman. Cannot be used for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Dh { get; } = new DsRecordAlgorithm("DH");
        /// <summary>
        /// DSA/SHA1. Not recommended for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Dsa { get; } = new DsRecordAlgorithm("DSA");
        /// <summary>
        /// ECC. Not recommended for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Ecc { get; } = new DsRecordAlgorithm("ECC");
        /// <summary>
        /// RSA/SHA-1. Not recommended for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Rsasha1 { get; } = new DsRecordAlgorithm("RSASHA1");
        /// <summary>
        /// DSA-NSEC3-SHA1. Not recommended for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Dsansec3sha1 { get; } = new DsRecordAlgorithm("DSANSEC3SHA1");
        /// <summary>
        /// RSA/SHA1-NSEC3-SHA1. Not recommended for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Rsasha1nsec3sha1 { get; } = new DsRecordAlgorithm("RSASHA1NSEC3SHA1");
        /// <summary>
        /// RSA/SHA-256.
        /// </summary>
        public static DsRecordAlgorithm Rsasha256 { get; } = new DsRecordAlgorithm("RSASHA256");
        /// <summary>
        /// RSA/SHA-512.
        /// </summary>
        public static DsRecordAlgorithm Rsasha512 { get; } = new DsRecordAlgorithm("RSASHA512");
        /// <summary>
        /// GOST R 34.10-2001.
        /// </summary>
        public static DsRecordAlgorithm Eccgost { get; } = new DsRecordAlgorithm("ECCGOST");
        /// <summary>
        /// ECDSA Curve P-256 with SHA-256.
        /// </summary>
        public static DsRecordAlgorithm Ecdsap256sha256 { get; } = new DsRecordAlgorithm("ECDSAP256SHA256");
        /// <summary>
        /// ECDSA Curve P-384 with SHA-384.
        /// </summary>
        public static DsRecordAlgorithm Ecdsap384sha384 { get; } = new DsRecordAlgorithm("ECDSAP384SHA384");
        /// <summary>
        /// Ed25519.
        /// </summary>
        public static DsRecordAlgorithm Ed25519 { get; } = new DsRecordAlgorithm("ED25519");
        /// <summary>
        /// Ed448.
        /// </summary>
        public static DsRecordAlgorithm Ed448 { get; } = new DsRecordAlgorithm("ED448");
        /// <summary>
        /// Reserved for Indirect Keys. Cannot be used for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Indirect { get; } = new DsRecordAlgorithm("INDIRECT");
        /// <summary>
        /// Private algorithm. Cannot be used for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Privatedns { get; } = new DsRecordAlgorithm("PRIVATEDNS");
        /// <summary>
        /// Private algorithm OID. Cannot be used for new deployments.
        /// </summary>
        public static DsRecordAlgorithm Privateoid { get; } = new DsRecordAlgorithm("PRIVATEOID");

        public static bool operator ==(DsRecordAlgorithm left, DsRecordAlgorithm right) => left.Equals(right);
        public static bool operator !=(DsRecordAlgorithm left, DsRecordAlgorithm right) => !left.Equals(right);

        public static explicit operator string(DsRecordAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DsRecordAlgorithm other && Equals(other);
        public bool Equals(DsRecordAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The hash function used to generate the digest of the referenced DNSKEY.
    /// </summary>
    [EnumType]
    public readonly struct DsRecordDigestType : IEquatable<DsRecordDigestType>
    {
        private readonly string _value;

        private DsRecordDigestType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The DigestType is unspecified.
        /// </summary>
        public static DsRecordDigestType DigestTypeUnspecified { get; } = new DsRecordDigestType("DIGEST_TYPE_UNSPECIFIED");
        /// <summary>
        /// SHA-1. Not recommended for new deployments.
        /// </summary>
        public static DsRecordDigestType Sha1 { get; } = new DsRecordDigestType("SHA1");
        /// <summary>
        /// SHA-256.
        /// </summary>
        public static DsRecordDigestType Sha256 { get; } = new DsRecordDigestType("SHA256");
        /// <summary>
        /// GOST R 34.11-94.
        /// </summary>
        public static DsRecordDigestType Gost3411 { get; } = new DsRecordDigestType("GOST3411");
        /// <summary>
        /// SHA-384.
        /// </summary>
        public static DsRecordDigestType Sha384 { get; } = new DsRecordDigestType("SHA384");

        public static bool operator ==(DsRecordDigestType left, DsRecordDigestType right) => left.Equals(right);
        public static bool operator !=(DsRecordDigestType left, DsRecordDigestType right) => !left.Equals(right);

        public static explicit operator string(DsRecordDigestType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DsRecordDigestType other && Equals(other);
        public bool Equals(DsRecordDigestType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. The state of DS records for this domain. Used to enable or disable automatic DNSSEC.
    /// </summary>
    [EnumType]
    public readonly struct GoogleDomainsDnsDsState : IEquatable<GoogleDomainsDnsDsState>
    {
        private readonly string _value;

        private GoogleDomainsDnsDsState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// DS state is unspecified.
        /// </summary>
        public static GoogleDomainsDnsDsState DsStateUnspecified { get; } = new GoogleDomainsDnsDsState("DS_STATE_UNSPECIFIED");
        /// <summary>
        /// DNSSEC is disabled for this domain. No DS records for this domain are published in the parent DNS zone.
        /// </summary>
        public static GoogleDomainsDnsDsState DsRecordsUnpublished { get; } = new GoogleDomainsDnsDsState("DS_RECORDS_UNPUBLISHED");
        /// <summary>
        /// DNSSEC is enabled for this domain. Appropriate DS records for this domain are published in the parent DNS zone. This option is valid only if the DNS zone referenced in the `Registration`'s `dns_provider` field is already DNSSEC-signed.
        /// </summary>
        public static GoogleDomainsDnsDsState DsRecordsPublished { get; } = new GoogleDomainsDnsDsState("DS_RECORDS_PUBLISHED");

        public static bool operator ==(GoogleDomainsDnsDsState left, GoogleDomainsDnsDsState right) => left.Equals(right);
        public static bool operator !=(GoogleDomainsDnsDsState left, GoogleDomainsDnsDsState right) => !left.Equals(right);

        public static explicit operator string(GoogleDomainsDnsDsState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleDomainsDnsDsState other && Equals(other);
        public bool Equals(GoogleDomainsDnsDsState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Controls whether the domain can be transferred to another registrar.
    /// </summary>
    [EnumType]
    public readonly struct ManagementSettingsTransferLockState : IEquatable<ManagementSettingsTransferLockState>
    {
        private readonly string _value;

        private ManagementSettingsTransferLockState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The state is unspecified.
        /// </summary>
        public static ManagementSettingsTransferLockState TransferLockStateUnspecified { get; } = new ManagementSettingsTransferLockState("TRANSFER_LOCK_STATE_UNSPECIFIED");
        /// <summary>
        /// The domain is unlocked and can be transferred to another registrar.
        /// </summary>
        public static ManagementSettingsTransferLockState Unlocked { get; } = new ManagementSettingsTransferLockState("UNLOCKED");
        /// <summary>
        /// The domain is locked and cannot be transferred to another registrar.
        /// </summary>
        public static ManagementSettingsTransferLockState Locked { get; } = new ManagementSettingsTransferLockState("LOCKED");

        public static bool operator ==(ManagementSettingsTransferLockState left, ManagementSettingsTransferLockState right) => left.Equals(right);
        public static bool operator !=(ManagementSettingsTransferLockState left, ManagementSettingsTransferLockState right) => !left.Equals(right);

        public static explicit operator string(ManagementSettingsTransferLockState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManagementSettingsTransferLockState other && Equals(other);
        public bool Equals(ManagementSettingsTransferLockState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct RegistrationContactNoticesItem : IEquatable<RegistrationContactNoticesItem>
    {
        private readonly string _value;

        private RegistrationContactNoticesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The notice is undefined.
        /// </summary>
        public static RegistrationContactNoticesItem ContactNoticeUnspecified { get; } = new RegistrationContactNoticesItem("CONTACT_NOTICE_UNSPECIFIED");
        /// <summary>
        /// Required when setting the `privacy` field of `ContactSettings` to `PUBLIC_CONTACT_DATA`, which exposes contact data publicly.
        /// </summary>
        public static RegistrationContactNoticesItem PublicContactDataAcknowledgement { get; } = new RegistrationContactNoticesItem("PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT");

        public static bool operator ==(RegistrationContactNoticesItem left, RegistrationContactNoticesItem right) => left.Equals(right);
        public static bool operator !=(RegistrationContactNoticesItem left, RegistrationContactNoticesItem right) => !left.Equals(right);

        public static explicit operator string(RegistrationContactNoticesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RegistrationContactNoticesItem other && Equals(other);
        public bool Equals(RegistrationContactNoticesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct RegistrationDomainNoticesItem : IEquatable<RegistrationDomainNoticesItem>
    {
        private readonly string _value;

        private RegistrationDomainNoticesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The notice is undefined.
        /// </summary>
        public static RegistrationDomainNoticesItem DomainNoticeUnspecified { get; } = new RegistrationDomainNoticesItem("DOMAIN_NOTICE_UNSPECIFIED");
        /// <summary>
        /// Indicates that the domain is preloaded on the HTTP Strict Transport Security list in browsers. Serving a website on such domain requires an SSL certificate. For details, see [how to get an SSL certificate](https://support.google.com/domains/answer/7638036).
        /// </summary>
        public static RegistrationDomainNoticesItem HstsPreloaded { get; } = new RegistrationDomainNoticesItem("HSTS_PRELOADED");

        public static bool operator ==(RegistrationDomainNoticesItem left, RegistrationDomainNoticesItem right) => left.Equals(right);
        public static bool operator !=(RegistrationDomainNoticesItem left, RegistrationDomainNoticesItem right) => !left.Equals(right);

        public static explicit operator string(RegistrationDomainNoticesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RegistrationDomainNoticesItem other && Equals(other);
        public bool Equals(RegistrationDomainNoticesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
