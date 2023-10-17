// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/emailaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEmailAccount(ctx *pulumi.Context, args *GetEmailAccountArgs, opts ...pulumi.InvokeOption) (*GetEmailAccountResult, error) {
	var rv GetEmailAccountResult
	err := ctx.Invoke("onepassword:index:GetEmailAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEmailAccountArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetEmailAccountResult struct {
	Attachments        map[string]OutField                     `pulumi:"attachments"`
	AuthMethod         *string                                 `pulumi:"authMethod"`
	Category           string                                  `pulumi:"category"`
	ContactInformation *emailaccount.ContactInformationSection `pulumi:"contactInformation"`
	Fields             map[string]OutField                     `pulumi:"fields"`
	Notes              *string                                 `pulumi:"notes"`
	Password           *string                                 `pulumi:"password"`
	PortNumber         *string                                 `pulumi:"portNumber"`
	References         map[string]OutField                     `pulumi:"references"`
	Sections           map[string]OutSection                   `pulumi:"sections"`
	Security           *string                                 `pulumi:"security"`
	Server             *string                                 `pulumi:"server"`
	Smtp               *emailaccount.SmtpSection               `pulumi:"smtp"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title    string  `pulumi:"title"`
	Type     *string `pulumi:"type"`
	Username *string `pulumi:"username"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

func GetEmailAccountOutput(ctx *pulumi.Context, args GetEmailAccountOutputArgs, opts ...pulumi.InvokeOption) GetEmailAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEmailAccountResult, error) {
			args := v.(GetEmailAccountArgs)
			r, err := GetEmailAccount(ctx, &args, opts...)
			var s GetEmailAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEmailAccountResultOutput)
}

type GetEmailAccountOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetEmailAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEmailAccountArgs)(nil)).Elem()
}

type GetEmailAccountResultOutput struct{ *pulumi.OutputState }

func (GetEmailAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEmailAccountResult)(nil)).Elem()
}

func (o GetEmailAccountResultOutput) ToGetEmailAccountResultOutput() GetEmailAccountResultOutput {
	return o
}

func (o GetEmailAccountResultOutput) ToGetEmailAccountResultOutputWithContext(ctx context.Context) GetEmailAccountResultOutput {
	return o
}

func (o GetEmailAccountResultOutput) Attachments() OutFieldMapOutput {
	return o.ApplyT(func(v GetEmailAccountResult) map[string]OutField { return v.Attachments }).(OutFieldMapOutput)
}

func (o GetEmailAccountResultOutput) AuthMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *string { return v.AuthMethod }).(pulumi.StringPtrOutput)
}

func (o GetEmailAccountResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetEmailAccountResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetEmailAccountResultOutput) ContactInformation() emailaccount.ContactInformationSectionPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *emailaccount.ContactInformationSection { return v.ContactInformation }).(emailaccount.ContactInformationSectionPtrOutput)
}

func (o GetEmailAccountResultOutput) Fields() OutFieldMapOutput {
	return o.ApplyT(func(v GetEmailAccountResult) map[string]OutField { return v.Fields }).(OutFieldMapOutput)
}

func (o GetEmailAccountResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetEmailAccountResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GetEmailAccountResultOutput) PortNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *string { return v.PortNumber }).(pulumi.StringPtrOutput)
}

func (o GetEmailAccountResultOutput) References() OutFieldMapOutput {
	return o.ApplyT(func(v GetEmailAccountResult) map[string]OutField { return v.References }).(OutFieldMapOutput)
}

func (o GetEmailAccountResultOutput) Sections() OutSectionMapOutput {
	return o.ApplyT(func(v GetEmailAccountResult) map[string]OutSection { return v.Sections }).(OutSectionMapOutput)
}

func (o GetEmailAccountResultOutput) Security() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *string { return v.Security }).(pulumi.StringPtrOutput)
}

func (o GetEmailAccountResultOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o GetEmailAccountResultOutput) Smtp() emailaccount.SmtpSectionPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *emailaccount.SmtpSection { return v.Smtp }).(emailaccount.SmtpSectionPtrOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetEmailAccountResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEmailAccountResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetEmailAccountResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetEmailAccountResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetEmailAccountResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetEmailAccountResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEmailAccountResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetEmailAccountResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetEmailAccountResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The UUID of the vault the item is in.
func (o GetEmailAccountResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v GetEmailAccountResult) string { return v.Vault }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEmailAccountResultOutput{})
}
