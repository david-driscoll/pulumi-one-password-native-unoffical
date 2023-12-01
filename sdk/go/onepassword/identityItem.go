// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/Sacro/pulumi-onepassword/sdk/go/onepassword/identity"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IdentityItem struct {
	pulumi.CustomResourceState

	Address         identity.AddressSectionPtrOutput         `pulumi:"address"`
	Attachments     OutFieldMapOutput                        `pulumi:"attachments"`
	Category        pulumi.StringOutput                      `pulumi:"category"`
	Fields          OutFieldMapOutput                        `pulumi:"fields"`
	Identification  identity.IdentificationSectionPtrOutput  `pulumi:"identification"`
	InternetDetails identity.InternetDetailsSectionPtrOutput `pulumi:"internetDetails"`
	Notes           pulumi.StringPtrOutput                   `pulumi:"notes"`
	References      OutFieldMapOutput                        `pulumi:"references"`
	Sections        OutSectionMapOutput                      `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewIdentityItem registers a new resource with the given unique name, arguments, and options.
func NewIdentityItem(ctx *pulumi.Context,
	name string, args *IdentityItemArgs, opts ...pulumi.ResourceOption) (*IdentityItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("Identity")
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"references",
		"sections",
	})
	opts = append(opts, secrets)
	var resource IdentityItem
	err := ctx.RegisterResource("one-password-native:index:IdentityItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityItem gets an existing IdentityItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityItemState, opts ...pulumi.ResourceOption) (*IdentityItem, error) {
	var resource IdentityItem
	err := ctx.ReadResource("one-password-native:index:IdentityItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityItem resources.
type identityItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type IdentityItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (IdentityItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityItemState)(nil)).Elem()
}

type identityItemArgs struct {
	Address     *identity.AddressSection         `pulumi:"address"`
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category        *string                          `pulumi:"category"`
	Fields          map[string]Field                 `pulumi:"fields"`
	Identification  *identity.IdentificationSection  `pulumi:"identification"`
	InternetDetails *identity.InternetDetailsSection `pulumi:"internetDetails"`
	Notes           *string                          `pulumi:"notes"`
	Sections        map[string]Section               `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a IdentityItem resource.
type IdentityItemArgs struct {
	Address     identity.AddressSectionPtrInput
	Attachments pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category        pulumi.StringPtrInput
	Fields          FieldMapInput
	Identification  identity.IdentificationSectionPtrInput
	InternetDetails identity.InternetDetailsSectionPtrInput
	Notes           pulumi.StringPtrInput
	Sections        SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (IdentityItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityItemArgs)(nil)).Elem()
}

func (r *IdentityItem) GetAttachment(ctx *pulumi.Context, args *IdentityItemGetAttachmentArgs) (IdentityItemGetAttachmentResultOutput, error) {
	out, err := ctx.Call("one-password-native:index:IdentityItem/attachment", args, IdentityItemGetAttachmentResultOutput{}, r)
	if err != nil {
		return IdentityItemGetAttachmentResultOutput{}, err
	}
	return out.(IdentityItemGetAttachmentResultOutput), nil
}

type identityItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name string `pulumi:"name"`
}

// The set of arguments for the GetAttachment method of the IdentityItem resource.
type IdentityItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name pulumi.StringInput
}

func (IdentityItemGetAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityItemGetAttachmentArgs)(nil)).Elem()
}

// The resolved reference value
type IdentityItemGetAttachmentResult struct {
	// the value of the attachment
	Value string `pulumi:"value"`
}

type IdentityItemGetAttachmentResultOutput struct{ *pulumi.OutputState }

func (IdentityItemGetAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityItemGetAttachmentResult)(nil)).Elem()
}

// the value of the attachment
func (o IdentityItemGetAttachmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityItemGetAttachmentResult) string { return v.Value }).(pulumi.StringOutput)
}

type IdentityItemInput interface {
	pulumi.Input

	ToIdentityItemOutput() IdentityItemOutput
	ToIdentityItemOutputWithContext(ctx context.Context) IdentityItemOutput
}

func (*IdentityItem) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityItem)(nil)).Elem()
}

func (i *IdentityItem) ToIdentityItemOutput() IdentityItemOutput {
	return i.ToIdentityItemOutputWithContext(context.Background())
}

func (i *IdentityItem) ToIdentityItemOutputWithContext(ctx context.Context) IdentityItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityItemOutput)
}

// IdentityItemArrayInput is an input type that accepts IdentityItemArray and IdentityItemArrayOutput values.
// You can construct a concrete instance of `IdentityItemArrayInput` via:
//
//	IdentityItemArray{ IdentityItemArgs{...} }
type IdentityItemArrayInput interface {
	pulumi.Input

	ToIdentityItemArrayOutput() IdentityItemArrayOutput
	ToIdentityItemArrayOutputWithContext(context.Context) IdentityItemArrayOutput
}

type IdentityItemArray []IdentityItemInput

func (IdentityItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityItem)(nil)).Elem()
}

func (i IdentityItemArray) ToIdentityItemArrayOutput() IdentityItemArrayOutput {
	return i.ToIdentityItemArrayOutputWithContext(context.Background())
}

func (i IdentityItemArray) ToIdentityItemArrayOutputWithContext(ctx context.Context) IdentityItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityItemArrayOutput)
}

// IdentityItemMapInput is an input type that accepts IdentityItemMap and IdentityItemMapOutput values.
// You can construct a concrete instance of `IdentityItemMapInput` via:
//
//	IdentityItemMap{ "key": IdentityItemArgs{...} }
type IdentityItemMapInput interface {
	pulumi.Input

	ToIdentityItemMapOutput() IdentityItemMapOutput
	ToIdentityItemMapOutputWithContext(context.Context) IdentityItemMapOutput
}

type IdentityItemMap map[string]IdentityItemInput

func (IdentityItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityItem)(nil)).Elem()
}

func (i IdentityItemMap) ToIdentityItemMapOutput() IdentityItemMapOutput {
	return i.ToIdentityItemMapOutputWithContext(context.Background())
}

func (i IdentityItemMap) ToIdentityItemMapOutputWithContext(ctx context.Context) IdentityItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityItemMapOutput)
}

type IdentityItemOutput struct{ *pulumi.OutputState }

func (IdentityItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityItem)(nil)).Elem()
}

func (o IdentityItemOutput) ToIdentityItemOutput() IdentityItemOutput {
	return o
}

func (o IdentityItemOutput) ToIdentityItemOutputWithContext(ctx context.Context) IdentityItemOutput {
	return o
}

type IdentityItemArrayOutput struct{ *pulumi.OutputState }

func (IdentityItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityItem)(nil)).Elem()
}

func (o IdentityItemArrayOutput) ToIdentityItemArrayOutput() IdentityItemArrayOutput {
	return o
}

func (o IdentityItemArrayOutput) ToIdentityItemArrayOutputWithContext(ctx context.Context) IdentityItemArrayOutput {
	return o
}

func (o IdentityItemArrayOutput) Index(i pulumi.IntInput) IdentityItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdentityItem {
		return vs[0].([]*IdentityItem)[vs[1].(int)]
	}).(IdentityItemOutput)
}

type IdentityItemMapOutput struct{ *pulumi.OutputState }

func (IdentityItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityItem)(nil)).Elem()
}

func (o IdentityItemMapOutput) ToIdentityItemMapOutput() IdentityItemMapOutput {
	return o
}

func (o IdentityItemMapOutput) ToIdentityItemMapOutputWithContext(ctx context.Context) IdentityItemMapOutput {
	return o
}

func (o IdentityItemMapOutput) MapIndex(k pulumi.StringInput) IdentityItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdentityItem {
		return vs[0].(map[string]*IdentityItem)[vs[1].(string)]
	}).(IdentityItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityItemInput)(nil)).Elem(), &IdentityItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityItemArrayInput)(nil)).Elem(), IdentityItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityItemMapInput)(nil)).Elem(), IdentityItemMap{})
	pulumi.RegisterOutputType(IdentityItemOutput{})
	pulumi.RegisterOutputType(IdentityItemGetAttachmentResultOutput{})
	pulumi.RegisterOutputType(IdentityItemArrayOutput{})
	pulumi.RegisterOutputType(IdentityItemMapOutput{})
}
