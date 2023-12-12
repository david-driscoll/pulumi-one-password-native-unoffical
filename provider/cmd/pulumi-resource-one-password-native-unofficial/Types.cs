
using System.Collections.Immutable;
using pulumi_resource_one_password_native_unofficial.Domain;

namespace pulumi_resource_one_password_native_unofficial;
public static class ItemType
{
    public static string APICredentialItem { get; } = "one-password-native-unofficial:index:APICredentialItem";
public static string BankAccountItem { get; } = "one-password-native-unofficial:index:BankAccountItem";
public static string CreditCardItem { get; } = "one-password-native-unofficial:index:CreditCardItem";
public static string CryptoWalletItem { get; } = "one-password-native-unofficial:index:CryptoWalletItem";
public static string DatabaseItem { get; } = "one-password-native-unofficial:index:DatabaseItem";
public static string DocumentItem { get; } = "one-password-native-unofficial:index:DocumentItem";
public static string DriverLicenseItem { get; } = "one-password-native-unofficial:index:DriverLicenseItem";
public static string EmailAccountItem { get; } = "one-password-native-unofficial:index:EmailAccountItem";
public static string IdentityItem { get; } = "one-password-native-unofficial:index:IdentityItem";
public static string Item { get; } = "one-password-native-unofficial:index:Item";
public static string LoginItem { get; } = "one-password-native-unofficial:index:LoginItem";
public static string MedicalRecordItem { get; } = "one-password-native-unofficial:index:MedicalRecordItem";
public static string MembershipItem { get; } = "one-password-native-unofficial:index:MembershipItem";
public static string OutdoorLicenseItem { get; } = "one-password-native-unofficial:index:OutdoorLicenseItem";
public static string PassportItem { get; } = "one-password-native-unofficial:index:PassportItem";
public static string PasswordItem { get; } = "one-password-native-unofficial:index:PasswordItem";
public static string RewardProgramItem { get; } = "one-password-native-unofficial:index:RewardProgramItem";
public static string SSHKeyItem { get; } = "one-password-native-unofficial:index:SSHKeyItem";
public static string SecureNoteItem { get; } = "one-password-native-unofficial:index:SecureNoteItem";
public static string ServerItem { get; } = "one-password-native-unofficial:index:ServerItem";
public static string SocialSecurityNumberItem { get; } = "one-password-native-unofficial:index:SocialSecurityNumberItem";
public static string SoftwareLicenseItem { get; } = "one-password-native-unofficial:index:SoftwareLicenseItem";
public static string WirelessRouterItem { get; } = "one-password-native-unofficial:index:WirelessRouterItem";
public static string GetItem { get; } = "one-password-native-unofficial:index:GetItem";
public static string GetVault { get; } = "one-password-native-unofficial:index:GetVault";
public static string GetSecretReference { get; } = "one-password-native-unofficial:index:GetSecretReference";
public static string Read { get; } = "one-password-native-unofficial:index:Read";
public static string Inject { get; } = "one-password-native-unofficial:index:Inject";
public static string GetAttachment { get; } = "one-password-native-unofficial:index:GetAttachment";
public static string GetAPICredential { get; } = "one-password-native-unofficial:index:GetAPICredential";
public static string GetBankAccount { get; } = "one-password-native-unofficial:index:GetBankAccount";
public static string GetCreditCard { get; } = "one-password-native-unofficial:index:GetCreditCard";
public static string GetCryptoWallet { get; } = "one-password-native-unofficial:index:GetCryptoWallet";
public static string GetDatabase { get; } = "one-password-native-unofficial:index:GetDatabase";
public static string GetDocument { get; } = "one-password-native-unofficial:index:GetDocument";
public static string GetDriverLicense { get; } = "one-password-native-unofficial:index:GetDriverLicense";
public static string GetEmailAccount { get; } = "one-password-native-unofficial:index:GetEmailAccount";
public static string GetIdentity { get; } = "one-password-native-unofficial:index:GetIdentity";
public static string GetLogin { get; } = "one-password-native-unofficial:index:GetLogin";
public static string GetMedicalRecord { get; } = "one-password-native-unofficial:index:GetMedicalRecord";
public static string GetMembership { get; } = "one-password-native-unofficial:index:GetMembership";
public static string GetOutdoorLicense { get; } = "one-password-native-unofficial:index:GetOutdoorLicense";
public static string GetPassport { get; } = "one-password-native-unofficial:index:GetPassport";
public static string GetPassword { get; } = "one-password-native-unofficial:index:GetPassword";
public static string GetRewardProgram { get; } = "one-password-native-unofficial:index:GetRewardProgram";
public static string GetSSHKey { get; } = "one-password-native-unofficial:index:GetSSHKey";
public static string GetSecureNote { get; } = "one-password-native-unofficial:index:GetSecureNote";
public static string GetServer { get; } = "one-password-native-unofficial:index:GetServer";
public static string GetSocialSecurityNumber { get; } = "one-password-native-unofficial:index:GetSocialSecurityNumber";
public static string GetSoftwareLicense { get; } = "one-password-native-unofficial:index:GetSoftwareLicense";
public static string GetWirelessRouter { get; } = "one-password-native-unofficial:index:GetWirelessRouter";
}
public static partial class TemplateMetadata
{
    private static ImmutableArray<ResourceType> ResourceTypes = [
        new("one-password-native-unofficial:index:APICredentialItem", "API Credential", "API_CREDENTIAL", TransformInputsToAPICredential, TransformOutputsToAPICredential, [("notes", null), ("username", null), ("credential", null), ("type", null), ("filename", null), ("validFrom", null), ("expires", null), ("hostname", null)]),
new("one-password-native-unofficial:index:BankAccountItem", "Bank Account", "BANK_ACCOUNT", TransformInputsToBankAccount, TransformOutputsToBankAccount, [("notes", null), ("bankName", null), ("nameOnAccount", null), ("type", null), ("routingNumber", null), ("accountNumber", null), ("swift", null), ("iban", null), ("pin", null), ("phone", "branchInformation"), ("address", "branchInformation")]),
new("one-password-native-unofficial:index:CreditCardItem", "Credit Card", "CREDIT_CARD", TransformInputsToCreditCard, TransformOutputsToCreditCard, [("notes", null), ("cardholderName", null), ("type", null), ("number", null), ("verificationNumber", null), ("expiryDate", null), ("validFrom", null), ("issuingBank", "contactInformation"), ("phoneLocal", "contactInformation"), ("phoneTollFree", "contactInformation"), ("phoneIntl", "contactInformation"), ("website", "contactInformation"), ("pin", "additionalDetails"), ("creditLimit", "additionalDetails"), ("cashWithdrawalLimit", "additionalDetails"), ("interestRate", "additionalDetails"), ("issueNumber", "additionalDetails")]),
new("one-password-native-unofficial:index:CryptoWalletItem", "Crypto Wallet", "CUSTOM", TransformInputsToCryptoWallet, TransformOutputsToCryptoWallet, [("notes", null), ("recoveryPhrase", null), ("password", null), ("walletAddress", "wallet")]),
new("one-password-native-unofficial:index:DatabaseItem", "Database", "DATABASE", TransformInputsToDatabase, TransformOutputsToDatabase, [("notes", null), ("type", null), ("server", null), ("port", null), ("database", null), ("username", null), ("password", null), ("sid", null), ("alias", null), ("connectionOptions", null)]),
new("one-password-native-unofficial:index:DocumentItem", "Document", "DOCUMENT", TransformInputsToDocument, TransformOutputsToDocument, [("notes", null)]),
new("one-password-native-unofficial:index:DriverLicenseItem", "Driver License", "DRIVER_LICENSE", TransformInputsToDriverLicense, TransformOutputsToDriverLicense, [("notes", null), ("fullName", null), ("address", null), ("dateOfBirth", null), ("gender", null), ("height", null), ("number", null), ("licenseClass", null), ("conditionsRestrictions", null), ("state", null), ("country", null), ("expiryDate", null)]),
new("one-password-native-unofficial:index:EmailAccountItem", "Email Account", "EMAIL_ACCOUNT", TransformInputsToEmailAccount, TransformOutputsToEmailAccount, [("notes", null), ("type", null), ("username", null), ("server", null), ("portNumber", null), ("password", null), ("security", null), ("authMethod", null), ("smtpServer", "smtp"), ("portNumber", "smtp"), ("username", "smtp"), ("password", "smtp"), ("security", "smtp"), ("authMethod", "smtp"), ("provider", "contactInformation"), ("providersWebsite", "contactInformation"), ("phoneLocal", "contactInformation"), ("phoneTollFree", "contactInformation")]),
new("one-password-native-unofficial:index:IdentityItem", "Identity", "IDENTITY", TransformInputsToIdentity, TransformOutputsToIdentity, [("notes", null), ("firstName", "identification"), ("initial", "identification"), ("lastName", "identification"), ("gender", "identification"), ("birthDate", "identification"), ("occupation", "identification"), ("company", "identification"), ("department", "identification"), ("jobTitle", "identification"), ("address", "address"), ("defaultPhone", "address"), ("home", "address"), ("cell", "address"), ("business", "address"), ("username", "internetDetails"), ("reminderQuestion", "internetDetails"), ("reminderAnswer", "internetDetails"), ("email", "internetDetails"), ("website", "internetDetails"), ("icq", "internetDetails"), ("skype", "internetDetails"), ("aolAim", "internetDetails"), ("yahoo", "internetDetails"), ("msn", "internetDetails"), ("forumSignature", "internetDetails")]),
new("one-password-native-unofficial:index:Item", "Item", "SECURE_NOTE", TransformInputsToItem, TransformOutputsToItem, [("notes", null)]),
new("one-password-native-unofficial:index:LoginItem", "Login", "LOGIN", TransformInputsToLogin, TransformOutputsToLogin, [("username", null), ("password", null), ("notes", null)]),
new("one-password-native-unofficial:index:MedicalRecordItem", "Medical Record", "MEDICAL_RECORD", TransformInputsToMedicalRecord, TransformOutputsToMedicalRecord, [("notes", null), ("date", null), ("location", null), ("healthcareProfessional", null), ("patient", null), ("reasonForVisit", null), ("medication", "medication"), ("dosage", "medication"), ("medicationNotes", "medication")]),
new("one-password-native-unofficial:index:MembershipItem", "Membership", "MEMBERSHIP", TransformInputsToMembership, TransformOutputsToMembership, [("notes", null), ("group", null), ("website", null), ("telephone", null), ("memberName", null), ("memberSince", null), ("expiryDate", null), ("memberId", null), ("pin", null)]),
new("one-password-native-unofficial:index:OutdoorLicenseItem", "Outdoor License", "OUTDOOR_LICENSE", TransformInputsToOutdoorLicense, TransformOutputsToOutdoorLicense, [("notes", null), ("fullName", null), ("validFrom", null), ("expires", null), ("approvedWildlife", null), ("maximumQuota", null), ("state", null), ("country", null)]),
new("one-password-native-unofficial:index:PassportItem", "Passport", "PASSPORT", TransformInputsToPassport, TransformOutputsToPassport, [("notes", null), ("type", null), ("issuingCountry", null), ("number", null), ("fullName", null), ("gender", null), ("nationality", null), ("issuingAuthority", null), ("dateOfBirth", null), ("placeOfBirth", null), ("issuedOn", null), ("expiryDate", null)]),
new("one-password-native-unofficial:index:PasswordItem", "Password", "PASSWORD", TransformInputsToPassword, TransformOutputsToPassword, [("password", null), ("notes", null)]),
new("one-password-native-unofficial:index:RewardProgramItem", "Reward Program", "REWARD_PROGRAM", TransformInputsToRewardProgram, TransformOutputsToRewardProgram, [("notes", null), ("companyName", null), ("memberName", null), ("memberId", null), ("pin", null), ("memberIdAdditional", "moreInformation"), ("memberSince", "moreInformation"), ("customerServicePhone", "moreInformation"), ("phoneForReservations", "moreInformation"), ("website", "moreInformation")]),
new("one-password-native-unofficial:index:SSHKeyItem", "SSH Key", "SSH_KEY", TransformInputsToSSHKey, TransformOutputsToSSHKey, [("notes", null), ("privateKey", null)]),
new("one-password-native-unofficial:index:SecureNoteItem", "Secure Note", "SECURE_NOTE", TransformInputsToSecureNote, TransformOutputsToSecureNote, [("notes", null)]),
new("one-password-native-unofficial:index:ServerItem", "Server", "SERVER", TransformInputsToServer, TransformOutputsToServer, [("notes", null), ("url", null), ("username", null), ("password", null), ("adminConsoleUrl", "adminConsole"), ("adminConsoleUsername", "adminConsole"), ("consolePassword", "adminConsole"), ("name", "hostingProvider"), ("website", "hostingProvider"), ("supportUrl", "hostingProvider"), ("supportPhone", "hostingProvider")]),
new("one-password-native-unofficial:index:SocialSecurityNumberItem", "Social Security Number", "SOCIAL_SECURITY_NUMBER", TransformInputsToSocialSecurityNumber, TransformOutputsToSocialSecurityNumber, [("notes", null), ("name", null), ("number", null)]),
new("one-password-native-unofficial:index:SoftwareLicenseItem", "Software License", "SOFTWARE_LICENSE", TransformInputsToSoftwareLicense, TransformOutputsToSoftwareLicense, [("notes", null), ("version", null), ("licenseKey", null), ("licensedTo", "customer"), ("registeredEmail", "customer"), ("company", "customer"), ("downloadPage", "publisher"), ("publisher", "publisher"), ("website", "publisher"), ("retailPrice", "publisher"), ("supportEmail", "publisher"), ("purchaseDate", "order"), ("orderNumber", "order"), ("orderTotal", "order")]),
new("one-password-native-unofficial:index:WirelessRouterItem", "Wireless Router", "WIRELESS_ROUTER", TransformInputsToWirelessRouter, TransformOutputsToWirelessRouter, [("notes", null), ("baseStationName", null), ("baseStationPassword", null), ("serverIpAddress", null), ("airPortId", null), ("networkName", null), ("wirelessSecurity", null), ("wirelessNetworkPassword", null), ("attachedStoragePassword", null)])];
    private static ImmutableArray<FunctionType> FunctionTypes = [
        new("one-password-native-unofficial:index:GetAPICredential", "API Credential", "API_CREDENTIAL", TransformOutputsToAPICredential, [("notes", null), ("username", null), ("credential", null), ("type", null), ("filename", null), ("validFrom", null), ("expires", null), ("hostname", null)]),
new("one-password-native-unofficial:index:GetBankAccount", "Bank Account", "BANK_ACCOUNT", TransformOutputsToBankAccount, [("notes", null), ("bankName", null), ("nameOnAccount", null), ("type", null), ("routingNumber", null), ("accountNumber", null), ("swift", null), ("iban", null), ("pin", null), ("phone", "branchInformation"), ("address", "branchInformation")]),
new("one-password-native-unofficial:index:GetCreditCard", "Credit Card", "CREDIT_CARD", TransformOutputsToCreditCard, [("notes", null), ("cardholderName", null), ("type", null), ("number", null), ("verificationNumber", null), ("expiryDate", null), ("validFrom", null), ("issuingBank", "contactInformation"), ("phoneLocal", "contactInformation"), ("phoneTollFree", "contactInformation"), ("phoneIntl", "contactInformation"), ("website", "contactInformation"), ("pin", "additionalDetails"), ("creditLimit", "additionalDetails"), ("cashWithdrawalLimit", "additionalDetails"), ("interestRate", "additionalDetails"), ("issueNumber", "additionalDetails")]),
new("one-password-native-unofficial:index:GetCryptoWallet", "Crypto Wallet", "CUSTOM", TransformOutputsToCryptoWallet, [("notes", null), ("recoveryPhrase", null), ("password", null), ("walletAddress", "wallet")]),
new("one-password-native-unofficial:index:GetDatabase", "Database", "DATABASE", TransformOutputsToDatabase, [("notes", null), ("type", null), ("server", null), ("port", null), ("database", null), ("username", null), ("password", null), ("sid", null), ("alias", null), ("connectionOptions", null)]),
new("one-password-native-unofficial:index:GetDocument", "Document", "DOCUMENT", TransformOutputsToDocument, [("notes", null)]),
new("one-password-native-unofficial:index:GetDriverLicense", "Driver License", "DRIVER_LICENSE", TransformOutputsToDriverLicense, [("notes", null), ("fullName", null), ("address", null), ("dateOfBirth", null), ("gender", null), ("height", null), ("number", null), ("licenseClass", null), ("conditionsRestrictions", null), ("state", null), ("country", null), ("expiryDate", null)]),
new("one-password-native-unofficial:index:GetEmailAccount", "Email Account", "EMAIL_ACCOUNT", TransformOutputsToEmailAccount, [("notes", null), ("type", null), ("username", null), ("server", null), ("portNumber", null), ("password", null), ("security", null), ("authMethod", null), ("smtpServer", "smtp"), ("portNumber", "smtp"), ("username", "smtp"), ("password", "smtp"), ("security", "smtp"), ("authMethod", "smtp"), ("provider", "contactInformation"), ("providersWebsite", "contactInformation"), ("phoneLocal", "contactInformation"), ("phoneTollFree", "contactInformation")]),
new("one-password-native-unofficial:index:GetIdentity", "Identity", "IDENTITY", TransformOutputsToIdentity, [("notes", null), ("firstName", "identification"), ("initial", "identification"), ("lastName", "identification"), ("gender", "identification"), ("birthDate", "identification"), ("occupation", "identification"), ("company", "identification"), ("department", "identification"), ("jobTitle", "identification"), ("address", "address"), ("defaultPhone", "address"), ("home", "address"), ("cell", "address"), ("business", "address"), ("username", "internetDetails"), ("reminderQuestion", "internetDetails"), ("reminderAnswer", "internetDetails"), ("email", "internetDetails"), ("website", "internetDetails"), ("icq", "internetDetails"), ("skype", "internetDetails"), ("aolAim", "internetDetails"), ("yahoo", "internetDetails"), ("msn", "internetDetails"), ("forumSignature", "internetDetails")]),
new("one-password-native-unofficial:index:GetItem", "Item", "SECURE_NOTE", TransformOutputsToItem, [("notes", null)]),
new("one-password-native-unofficial:index:GetLogin", "Login", "LOGIN", TransformOutputsToLogin, [("username", null), ("password", null), ("notes", null)]),
new("one-password-native-unofficial:index:GetMedicalRecord", "Medical Record", "MEDICAL_RECORD", TransformOutputsToMedicalRecord, [("notes", null), ("date", null), ("location", null), ("healthcareProfessional", null), ("patient", null), ("reasonForVisit", null), ("medication", "medication"), ("dosage", "medication"), ("medicationNotes", "medication")]),
new("one-password-native-unofficial:index:GetMembership", "Membership", "MEMBERSHIP", TransformOutputsToMembership, [("notes", null), ("group", null), ("website", null), ("telephone", null), ("memberName", null), ("memberSince", null), ("expiryDate", null), ("memberId", null), ("pin", null)]),
new("one-password-native-unofficial:index:GetOutdoorLicense", "Outdoor License", "OUTDOOR_LICENSE", TransformOutputsToOutdoorLicense, [("notes", null), ("fullName", null), ("validFrom", null), ("expires", null), ("approvedWildlife", null), ("maximumQuota", null), ("state", null), ("country", null)]),
new("one-password-native-unofficial:index:GetPassport", "Passport", "PASSPORT", TransformOutputsToPassport, [("notes", null), ("type", null), ("issuingCountry", null), ("number", null), ("fullName", null), ("gender", null), ("nationality", null), ("issuingAuthority", null), ("dateOfBirth", null), ("placeOfBirth", null), ("issuedOn", null), ("expiryDate", null)]),
new("one-password-native-unofficial:index:GetPassword", "Password", "PASSWORD", TransformOutputsToPassword, [("password", null), ("notes", null)]),
new("one-password-native-unofficial:index:GetRewardProgram", "Reward Program", "REWARD_PROGRAM", TransformOutputsToRewardProgram, [("notes", null), ("companyName", null), ("memberName", null), ("memberId", null), ("pin", null), ("memberIdAdditional", "moreInformation"), ("memberSince", "moreInformation"), ("customerServicePhone", "moreInformation"), ("phoneForReservations", "moreInformation"), ("website", "moreInformation")]),
new("one-password-native-unofficial:index:GetSSHKey", "SSH Key", "SSH_KEY", TransformOutputsToSSHKey, [("notes", null), ("privateKey", null)]),
new("one-password-native-unofficial:index:GetSecureNote", "Secure Note", "SECURE_NOTE", TransformOutputsToSecureNote, [("notes", null)]),
new("one-password-native-unofficial:index:GetServer", "Server", "SERVER", TransformOutputsToServer, [("notes", null), ("url", null), ("username", null), ("password", null), ("adminConsoleUrl", "adminConsole"), ("adminConsoleUsername", "adminConsole"), ("consolePassword", "adminConsole"), ("name", "hostingProvider"), ("website", "hostingProvider"), ("supportUrl", "hostingProvider"), ("supportPhone", "hostingProvider")]),
new("one-password-native-unofficial:index:GetSocialSecurityNumber", "Social Security Number", "SOCIAL_SECURITY_NUMBER", TransformOutputsToSocialSecurityNumber, [("notes", null), ("name", null), ("number", null)]),
new("one-password-native-unofficial:index:GetSoftwareLicense", "Software License", "SOFTWARE_LICENSE", TransformOutputsToSoftwareLicense, [("notes", null), ("version", null), ("licenseKey", null), ("licensedTo", "customer"), ("registeredEmail", "customer"), ("company", "customer"), ("downloadPage", "publisher"), ("publisher", "publisher"), ("website", "publisher"), ("retailPrice", "publisher"), ("supportEmail", "publisher"), ("purchaseDate", "order"), ("orderNumber", "order"), ("orderTotal", "order")]),
new("one-password-native-unofficial:index:GetWirelessRouter", "Wireless Router", "WIRELESS_ROUTER", TransformOutputsToWirelessRouter, [("notes", null), ("baseStationName", null), ("baseStationPassword", null), ("serverIpAddress", null), ("airPortId", null), ("networkName", null), ("wirelessSecurity", null), ("wirelessNetworkPassword", null), ("attachedStoragePassword", null)])];
}
