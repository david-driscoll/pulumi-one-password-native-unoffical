// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unoffical

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDatabase(ctx *pulumi.Context, args *GetDatabaseArgs, opts ...pulumi.InvokeOption) (*GetDatabaseResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatabaseResult
	err := ctx.Invoke("one-password-native-unoffical:index:GetDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDatabaseArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetDatabaseResult struct {
	Alias             *string                  `pulumi:"alias"`
	Attachments       map[string]OutAttachment `pulumi:"attachments"`
	Category          string                   `pulumi:"category"`
	ConnectionOptions *string                  `pulumi:"connectionOptions"`
	Database          *string                  `pulumi:"database"`
	Fields            map[string]OutField      `pulumi:"fields"`
	Notes             *string                  `pulumi:"notes"`
	Password          *string                  `pulumi:"password"`
	Port              *string                  `pulumi:"port"`
	References        map[string]OutField      `pulumi:"references"`
	Sections          map[string]OutSection    `pulumi:"sections"`
	Server            *string                  `pulumi:"server"`
	Sid               *string                  `pulumi:"sid"`
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

func GetDatabaseOutput(ctx *pulumi.Context, args GetDatabaseOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseResult, error) {
			args := v.(GetDatabaseArgs)
			r, err := GetDatabase(ctx, &args, opts...)
			var s GetDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabaseResultOutput)
}

type GetDatabaseOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseArgs)(nil)).Elem()
}

type GetDatabaseResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseResult)(nil)).Elem()
}

func (o GetDatabaseResultOutput) ToGetDatabaseResultOutput() GetDatabaseResultOutput {
	return o
}

func (o GetDatabaseResultOutput) ToGetDatabaseResultOutputWithContext(ctx context.Context) GetDatabaseResultOutput {
	return o
}

func (o GetDatabaseResultOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Attachments() OutAttachmentMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]OutAttachment { return v.Attachments }).(OutAttachmentMapOutput)
}

func (o GetDatabaseResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetDatabaseResultOutput) ConnectionOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.ConnectionOptions }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Fields() OutFieldMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]OutField { return v.Fields }).(OutFieldMapOutput)
}

func (o GetDatabaseResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Port }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) References() OutFieldMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]OutField { return v.References }).(OutFieldMapOutput)
}

func (o GetDatabaseResultOutput) Sections() OutSectionMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]OutSection { return v.Sections }).(OutSectionMapOutput)
}

func (o GetDatabaseResultOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetDatabaseResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDatabaseResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetDatabaseResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetDatabaseResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetDatabaseResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The UUID of the vault the item is in.
func (o GetDatabaseResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Vault }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseResultOutput{})
}
