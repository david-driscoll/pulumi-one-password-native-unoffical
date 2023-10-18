// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OnePasswordNative.RewardProgram.Outputs
{

    [OutputType]
    public sealed class MoreInformationSection
    {
        public readonly string? CustomerServicePhone;
        public readonly string? MemberIdAdditional;
        public readonly string? MemberSince;
        public readonly string? PhoneForReservations;
        public readonly string? Website;

        [OutputConstructor]
        private MoreInformationSection(
            string? customerServicePhone,

            string? memberIdAdditional,

            string? memberSince,

            string? phoneForReservations,

            string? website)
        {
            CustomerServicePhone = customerServicePhone;
            MemberIdAdditional = memberIdAdditional;
            MemberSince = memberSince;
            PhoneForReservations = phoneForReservations;
            Website = website;
        }
    }
}
