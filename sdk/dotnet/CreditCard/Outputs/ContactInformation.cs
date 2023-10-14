// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword.CreditCard.Outputs
{

    [OutputType]
    public sealed class ContactInformation
    {
        public readonly string? IssuingBank;
        public readonly string? PhoneIntl;
        public readonly string? PhoneLocal;
        public readonly string? PhoneTollFree;
        public readonly string? Website;

        [OutputConstructor]
        private ContactInformation(
            string? issuingBank,

            string? phoneIntl,

            string? phoneLocal,

            string? phoneTollFree,

            string? website)
        {
            IssuingBank = issuingBank;
            PhoneIntl = phoneIntl;
            PhoneLocal = phoneLocal;
            PhoneTollFree = phoneTollFree;
            Website = website;
        }
    }
}
