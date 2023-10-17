// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/softwarelicense"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SoftwareLicenseItem struct {
	pulumi.CustomResourceState

	Attachments OutFieldMapOutput                         `pulumi:"attachments"`
	Category    pulumi.StringOutput                       `pulumi:"category"`
	Customer    softwarelicense.CustomerSectionPtrOutput  `pulumi:"customer"`
	Fields      OutFieldMapOutput                         `pulumi:"fields"`
	LicenseKey  pulumi.StringPtrOutput                    `pulumi:"licenseKey"`
	Notes       pulumi.StringPtrOutput                    `pulumi:"notes"`
	Order       softwarelicense.OrderSectionPtrOutput     `pulumi:"order"`
	Publisher   softwarelicense.PublisherSectionPtrOutput `pulumi:"publisher"`
	References  OutFieldMapOutput                         `pulumi:"references"`
	Sections    OutSectionMapOutput                       `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault   pulumi.StringOutput    `pulumi:"vault"`
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewSoftwareLicenseItem registers a new resource with the given unique name, arguments, and options.
func NewSoftwareLicenseItem(ctx *pulumi.Context,
	name string, args *SoftwareLicenseItemArgs, opts ...pulumi.ResourceOption) (*SoftwareLicenseItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("Software License")
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"references",
		"sections",
	})
	opts = append(opts, secrets)
	var resource SoftwareLicenseItem
	err := ctx.RegisterResource("onepassword:index:SoftwareLicenseItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSoftwareLicenseItem gets an existing SoftwareLicenseItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSoftwareLicenseItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SoftwareLicenseItemState, opts ...pulumi.ResourceOption) (*SoftwareLicenseItem, error) {
	var resource SoftwareLicenseItem
	err := ctx.ReadResource("onepassword:index:SoftwareLicenseItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SoftwareLicenseItem resources.
type softwareLicenseItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type SoftwareLicenseItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (SoftwareLicenseItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareLicenseItemState)(nil)).Elem()
}

type softwareLicenseItemArgs struct {
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category   *string                           `pulumi:"category"`
	Customer   *softwarelicense.CustomerSection  `pulumi:"customer"`
	Fields     map[string]Field                  `pulumi:"fields"`
	LicenseKey *string                           `pulumi:"licenseKey"`
	Notes      *string                           `pulumi:"notes"`
	Order      *softwarelicense.OrderSection     `pulumi:"order"`
	Publisher  *softwarelicense.PublisherSection `pulumi:"publisher"`
	Sections   map[string]Section                `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault   string  `pulumi:"vault"`
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a SoftwareLicenseItem resource.
type SoftwareLicenseItemArgs struct {
	Attachments pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category   pulumi.StringPtrInput
	Customer   softwarelicense.CustomerSectionPtrInput
	Fields     FieldMapInput
	LicenseKey pulumi.StringPtrInput
	Notes      pulumi.StringPtrInput
	Order      softwarelicense.OrderSectionPtrInput
	Publisher  softwarelicense.PublisherSectionPtrInput
	Sections   SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault   pulumi.StringInput
	Version pulumi.StringPtrInput
}

func (SoftwareLicenseItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareLicenseItemArgs)(nil)).Elem()
}

func (r *SoftwareLicenseItem) Attachment(ctx *pulumi.Context, args *SoftwareLicenseItemAttachmentArgs) (SoftwareLicenseItemAttachmentResultOutput, error) {
	out, err := ctx.Call("onepassword:index:SoftwareLicenseItem/attachment", args, SoftwareLicenseItemAttachmentResultOutput{}, r)
	if err != nil {
		return SoftwareLicenseItemAttachmentResultOutput{}, err
	}
	return out.(SoftwareLicenseItemAttachmentResultOutput), nil
}

type softwareLicenseItemAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name string `pulumi:"name"`
}

// The set of arguments for the Attachment method of the SoftwareLicenseItem resource.
type SoftwareLicenseItemAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name pulumi.StringInput
}

func (SoftwareLicenseItemAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareLicenseItemAttachmentArgs)(nil)).Elem()
}

// The resolved reference value
type SoftwareLicenseItemAttachmentResult struct {
	// the value of the attachment
	Value string `pulumi:"value"`
}

type SoftwareLicenseItemAttachmentResultOutput struct{ *pulumi.OutputState }

func (SoftwareLicenseItemAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareLicenseItemAttachmentResult)(nil)).Elem()
}

// the value of the attachment
func (o SoftwareLicenseItemAttachmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SoftwareLicenseItemAttachmentResult) string { return v.Value }).(pulumi.StringOutput)
}

type SoftwareLicenseItemInput interface {
	pulumi.Input

	ToSoftwareLicenseItemOutput() SoftwareLicenseItemOutput
	ToSoftwareLicenseItemOutputWithContext(ctx context.Context) SoftwareLicenseItemOutput
}

func (*SoftwareLicenseItem) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareLicenseItem)(nil)).Elem()
}

func (i *SoftwareLicenseItem) ToSoftwareLicenseItemOutput() SoftwareLicenseItemOutput {
	return i.ToSoftwareLicenseItemOutputWithContext(context.Background())
}

func (i *SoftwareLicenseItem) ToSoftwareLicenseItemOutputWithContext(ctx context.Context) SoftwareLicenseItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareLicenseItemOutput)
}

// SoftwareLicenseItemArrayInput is an input type that accepts SoftwareLicenseItemArray and SoftwareLicenseItemArrayOutput values.
// You can construct a concrete instance of `SoftwareLicenseItemArrayInput` via:
//
//	SoftwareLicenseItemArray{ SoftwareLicenseItemArgs{...} }
type SoftwareLicenseItemArrayInput interface {
	pulumi.Input

	ToSoftwareLicenseItemArrayOutput() SoftwareLicenseItemArrayOutput
	ToSoftwareLicenseItemArrayOutputWithContext(context.Context) SoftwareLicenseItemArrayOutput
}

type SoftwareLicenseItemArray []SoftwareLicenseItemInput

func (SoftwareLicenseItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SoftwareLicenseItem)(nil)).Elem()
}

func (i SoftwareLicenseItemArray) ToSoftwareLicenseItemArrayOutput() SoftwareLicenseItemArrayOutput {
	return i.ToSoftwareLicenseItemArrayOutputWithContext(context.Background())
}

func (i SoftwareLicenseItemArray) ToSoftwareLicenseItemArrayOutputWithContext(ctx context.Context) SoftwareLicenseItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareLicenseItemArrayOutput)
}

// SoftwareLicenseItemMapInput is an input type that accepts SoftwareLicenseItemMap and SoftwareLicenseItemMapOutput values.
// You can construct a concrete instance of `SoftwareLicenseItemMapInput` via:
//
//	SoftwareLicenseItemMap{ "key": SoftwareLicenseItemArgs{...} }
type SoftwareLicenseItemMapInput interface {
	pulumi.Input

	ToSoftwareLicenseItemMapOutput() SoftwareLicenseItemMapOutput
	ToSoftwareLicenseItemMapOutputWithContext(context.Context) SoftwareLicenseItemMapOutput
}

type SoftwareLicenseItemMap map[string]SoftwareLicenseItemInput

func (SoftwareLicenseItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SoftwareLicenseItem)(nil)).Elem()
}

func (i SoftwareLicenseItemMap) ToSoftwareLicenseItemMapOutput() SoftwareLicenseItemMapOutput {
	return i.ToSoftwareLicenseItemMapOutputWithContext(context.Background())
}

func (i SoftwareLicenseItemMap) ToSoftwareLicenseItemMapOutputWithContext(ctx context.Context) SoftwareLicenseItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareLicenseItemMapOutput)
}

type SoftwareLicenseItemOutput struct{ *pulumi.OutputState }

func (SoftwareLicenseItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareLicenseItem)(nil)).Elem()
}

func (o SoftwareLicenseItemOutput) ToSoftwareLicenseItemOutput() SoftwareLicenseItemOutput {
	return o
}

func (o SoftwareLicenseItemOutput) ToSoftwareLicenseItemOutputWithContext(ctx context.Context) SoftwareLicenseItemOutput {
	return o
}

type SoftwareLicenseItemArrayOutput struct{ *pulumi.OutputState }

func (SoftwareLicenseItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SoftwareLicenseItem)(nil)).Elem()
}

func (o SoftwareLicenseItemArrayOutput) ToSoftwareLicenseItemArrayOutput() SoftwareLicenseItemArrayOutput {
	return o
}

func (o SoftwareLicenseItemArrayOutput) ToSoftwareLicenseItemArrayOutputWithContext(ctx context.Context) SoftwareLicenseItemArrayOutput {
	return o
}

func (o SoftwareLicenseItemArrayOutput) Index(i pulumi.IntInput) SoftwareLicenseItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SoftwareLicenseItem {
		return vs[0].([]*SoftwareLicenseItem)[vs[1].(int)]
	}).(SoftwareLicenseItemOutput)
}

type SoftwareLicenseItemMapOutput struct{ *pulumi.OutputState }

func (SoftwareLicenseItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SoftwareLicenseItem)(nil)).Elem()
}

func (o SoftwareLicenseItemMapOutput) ToSoftwareLicenseItemMapOutput() SoftwareLicenseItemMapOutput {
	return o
}

func (o SoftwareLicenseItemMapOutput) ToSoftwareLicenseItemMapOutputWithContext(ctx context.Context) SoftwareLicenseItemMapOutput {
	return o
}

func (o SoftwareLicenseItemMapOutput) MapIndex(k pulumi.StringInput) SoftwareLicenseItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SoftwareLicenseItem {
		return vs[0].(map[string]*SoftwareLicenseItem)[vs[1].(string)]
	}).(SoftwareLicenseItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SoftwareLicenseItemInput)(nil)).Elem(), &SoftwareLicenseItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*SoftwareLicenseItemArrayInput)(nil)).Elem(), SoftwareLicenseItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SoftwareLicenseItemMapInput)(nil)).Elem(), SoftwareLicenseItemMap{})
	pulumi.RegisterOutputType(SoftwareLicenseItemOutput{})
	pulumi.RegisterOutputType(SoftwareLicenseItemAttachmentResultOutput{})
	pulumi.RegisterOutputType(SoftwareLicenseItemArrayOutput{})
	pulumi.RegisterOutputType(SoftwareLicenseItemMapOutput{})
}
