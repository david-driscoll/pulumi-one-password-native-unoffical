// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/creditcard"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCreditCard(ctx *pulumi.Context, args *GetCreditCardArgs, opts ...pulumi.InvokeOption) (*GetCreditCardResult, error) {
	var rv GetCreditCardResult
	err := ctx.Invoke("one-password-native:index:GetCreditCard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetCreditCardArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetCreditCardResult struct {
	AdditionalDetails  *creditcard.AdditionalDetailsSection  `pulumi:"additionalDetails"`
	Attachments        map[string]OutField                   `pulumi:"attachments"`
	CardholderName     *string                               `pulumi:"cardholderName"`
	Category           string                                `pulumi:"category"`
	ContactInformation *creditcard.ContactInformationSection `pulumi:"contactInformation"`
	ExpiryDate         *string                               `pulumi:"expiryDate"`
	Fields             map[string]OutField                   `pulumi:"fields"`
	Notes              *string                               `pulumi:"notes"`
	Number             *string                               `pulumi:"number"`
	References         map[string]OutField                   `pulumi:"references"`
	Sections           map[string]OutSection                 `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string  `pulumi:"title"`
	Type  *string `pulumi:"type"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid      string  `pulumi:"uuid"`
	ValidFrom *string `pulumi:"validFrom"`
	// The UUID of the vault the item is in.
	Vault              string  `pulumi:"vault"`
	VerificationNumber *string `pulumi:"verificationNumber"`
}

func GetCreditCardOutput(ctx *pulumi.Context, args GetCreditCardOutputArgs, opts ...pulumi.InvokeOption) GetCreditCardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCreditCardResult, error) {
			args := v.(GetCreditCardArgs)
			r, err := GetCreditCard(ctx, &args, opts...)
			var s GetCreditCardResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCreditCardResultOutput)
}

type GetCreditCardOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetCreditCardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCreditCardArgs)(nil)).Elem()
}

type GetCreditCardResultOutput struct{ *pulumi.OutputState }

func (GetCreditCardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCreditCardResult)(nil)).Elem()
}

func (o GetCreditCardResultOutput) ToGetCreditCardResultOutput() GetCreditCardResultOutput {
	return o
}

func (o GetCreditCardResultOutput) ToGetCreditCardResultOutputWithContext(ctx context.Context) GetCreditCardResultOutput {
	return o
}

func (o GetCreditCardResultOutput) AdditionalDetails() creditcard.AdditionalDetailsSectionPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *creditcard.AdditionalDetailsSection { return v.AdditionalDetails }).(creditcard.AdditionalDetailsSectionPtrOutput)
}

func (o GetCreditCardResultOutput) Attachments() OutFieldMapOutput {
	return o.ApplyT(func(v GetCreditCardResult) map[string]OutField { return v.Attachments }).(OutFieldMapOutput)
}

func (o GetCreditCardResultOutput) CardholderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *string { return v.CardholderName }).(pulumi.StringPtrOutput)
}

func (o GetCreditCardResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetCreditCardResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetCreditCardResultOutput) ContactInformation() creditcard.ContactInformationSectionPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *creditcard.ContactInformationSection { return v.ContactInformation }).(creditcard.ContactInformationSectionPtrOutput)
}

func (o GetCreditCardResultOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *string { return v.ExpiryDate }).(pulumi.StringPtrOutput)
}

func (o GetCreditCardResultOutput) Fields() OutFieldMapOutput {
	return o.ApplyT(func(v GetCreditCardResult) map[string]OutField { return v.Fields }).(OutFieldMapOutput)
}

func (o GetCreditCardResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetCreditCardResultOutput) Number() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *string { return v.Number }).(pulumi.StringPtrOutput)
}

func (o GetCreditCardResultOutput) References() OutFieldMapOutput {
	return o.ApplyT(func(v GetCreditCardResult) map[string]OutField { return v.References }).(OutFieldMapOutput)
}

func (o GetCreditCardResultOutput) Sections() OutSectionMapOutput {
	return o.ApplyT(func(v GetCreditCardResult) map[string]OutSection { return v.Sections }).(OutSectionMapOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetCreditCardResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCreditCardResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetCreditCardResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetCreditCardResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetCreditCardResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetCreditCardResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetCreditCardResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o GetCreditCardResultOutput) ValidFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *string { return v.ValidFrom }).(pulumi.StringPtrOutput)
}

// The UUID of the vault the item is in.
func (o GetCreditCardResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v GetCreditCardResult) string { return v.Vault }).(pulumi.StringOutput)
}

func (o GetCreditCardResultOutput) VerificationNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCreditCardResult) *string { return v.VerificationNumber }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCreditCardResultOutput{})
}
