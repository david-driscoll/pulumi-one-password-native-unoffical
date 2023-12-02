// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unoffical

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MembershipItem struct {
	pulumi.CustomResourceState

	Attachments OutAttachmentMapOutput `pulumi:"attachments"`
	Category    pulumi.StringOutput    `pulumi:"category"`
	ExpiryDate  pulumi.StringPtrOutput `pulumi:"expiryDate"`
	Fields      OutFieldMapOutput      `pulumi:"fields"`
	Group       pulumi.StringPtrOutput `pulumi:"group"`
	MemberId    pulumi.StringPtrOutput `pulumi:"memberId"`
	MemberName  pulumi.StringPtrOutput `pulumi:"memberName"`
	MemberSince pulumi.StringPtrOutput `pulumi:"memberSince"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	Pin         pulumi.StringPtrOutput `pulumi:"pin"`
	References  OutFieldMapOutput      `pulumi:"references"`
	Sections    OutSectionMapOutput    `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags      pulumi.StringArrayOutput `pulumi:"tags"`
	Telephone pulumi.StringPtrOutput   `pulumi:"telephone"`
	// The title of the item.
	Title pulumi.StringOutput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault   pulumi.StringOutput    `pulumi:"vault"`
	Website pulumi.StringPtrOutput `pulumi:"website"`
}

// NewMembershipItem registers a new resource with the given unique name, arguments, and options.
func NewMembershipItem(ctx *pulumi.Context,
	name string, args *MembershipItemArgs, opts ...pulumi.ResourceOption) (*MembershipItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("Membership")
	if args.Pin != nil {
		args.Pin = pulumi.ToSecret(args.Pin).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"pin",
		"references",
		"sections",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource MembershipItem
	err := ctx.RegisterResource("one-password-native-unoffical:index:MembershipItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembershipItem gets an existing MembershipItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembershipItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembershipItemState, opts ...pulumi.ResourceOption) (*MembershipItem, error) {
	var resource MembershipItem
	err := ctx.ReadResource("one-password-native-unoffical:index:MembershipItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MembershipItem resources.
type membershipItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type MembershipItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (MembershipItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipItemState)(nil)).Elem()
}

type membershipItemArgs struct {
	// The category of the vault the item is in.
	Category         *string                          `pulumi:"category"`
	ExpiryDate       *string                          `pulumi:"expiryDate"`
	Fields           map[string]Field                 `pulumi:"fields"`
	Group            *string                          `pulumi:"group"`
	InputAttachments map[string]pulumi.AssetOrArchive `pulumi:"inputAttachments"`
	MemberId         *string                          `pulumi:"memberId"`
	MemberName       *string                          `pulumi:"memberName"`
	MemberSince      *string                          `pulumi:"memberSince"`
	Notes            *string                          `pulumi:"notes"`
	Pin              *string                          `pulumi:"pin"`
	Sections         map[string]Section               `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags      []string `pulumi:"tags"`
	Telephone *string  `pulumi:"telephone"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault   string  `pulumi:"vault"`
	Website *string `pulumi:"website"`
}

// The set of arguments for constructing a MembershipItem resource.
type MembershipItemArgs struct {
	// The category of the vault the item is in.
	Category         pulumi.StringPtrInput
	ExpiryDate       pulumi.StringPtrInput
	Fields           FieldMapInput
	Group            pulumi.StringPtrInput
	InputAttachments pulumi.AssetOrArchiveMapInput
	MemberId         pulumi.StringPtrInput
	MemberName       pulumi.StringPtrInput
	MemberSince      pulumi.StringPtrInput
	Notes            pulumi.StringPtrInput
	Pin              pulumi.StringPtrInput
	Sections         SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags      pulumi.StringArrayInput
	Telephone pulumi.StringPtrInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault   pulumi.StringInput
	Website pulumi.StringPtrInput
}

func (MembershipItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipItemArgs)(nil)).Elem()
}

func (r *MembershipItem) GetAttachment(ctx *pulumi.Context, args *MembershipItemGetAttachmentArgs) (MembershipItemGetAttachmentResultOutput, error) {
	out, err := ctx.Call("one-password-native-unoffical:index:MembershipItem/attachment", args, MembershipItemGetAttachmentResultOutput{}, r)
	if err != nil {
		return MembershipItemGetAttachmentResultOutput{}, err
	}
	return out.(MembershipItemGetAttachmentResultOutput), nil
}

type membershipItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name string `pulumi:"name"`
}

// The set of arguments for the GetAttachment method of the MembershipItem resource.
type MembershipItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name pulumi.StringInput
}

func (MembershipItemGetAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipItemGetAttachmentArgs)(nil)).Elem()
}

// The resolved reference value
type MembershipItemGetAttachmentResult struct {
	// the value of the attachment
	Value string `pulumi:"value"`
}

type MembershipItemGetAttachmentResultOutput struct{ *pulumi.OutputState }

func (MembershipItemGetAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipItemGetAttachmentResult)(nil)).Elem()
}

// the value of the attachment
func (o MembershipItemGetAttachmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v MembershipItemGetAttachmentResult) string { return v.Value }).(pulumi.StringOutput)
}

type MembershipItemInput interface {
	pulumi.Input

	ToMembershipItemOutput() MembershipItemOutput
	ToMembershipItemOutputWithContext(ctx context.Context) MembershipItemOutput
}

func (*MembershipItem) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipItem)(nil)).Elem()
}

func (i *MembershipItem) ToMembershipItemOutput() MembershipItemOutput {
	return i.ToMembershipItemOutputWithContext(context.Background())
}

func (i *MembershipItem) ToMembershipItemOutputWithContext(ctx context.Context) MembershipItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipItemOutput)
}

// MembershipItemArrayInput is an input type that accepts MembershipItemArray and MembershipItemArrayOutput values.
// You can construct a concrete instance of `MembershipItemArrayInput` via:
//
//	MembershipItemArray{ MembershipItemArgs{...} }
type MembershipItemArrayInput interface {
	pulumi.Input

	ToMembershipItemArrayOutput() MembershipItemArrayOutput
	ToMembershipItemArrayOutputWithContext(context.Context) MembershipItemArrayOutput
}

type MembershipItemArray []MembershipItemInput

func (MembershipItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MembershipItem)(nil)).Elem()
}

func (i MembershipItemArray) ToMembershipItemArrayOutput() MembershipItemArrayOutput {
	return i.ToMembershipItemArrayOutputWithContext(context.Background())
}

func (i MembershipItemArray) ToMembershipItemArrayOutputWithContext(ctx context.Context) MembershipItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipItemArrayOutput)
}

// MembershipItemMapInput is an input type that accepts MembershipItemMap and MembershipItemMapOutput values.
// You can construct a concrete instance of `MembershipItemMapInput` via:
//
//	MembershipItemMap{ "key": MembershipItemArgs{...} }
type MembershipItemMapInput interface {
	pulumi.Input

	ToMembershipItemMapOutput() MembershipItemMapOutput
	ToMembershipItemMapOutputWithContext(context.Context) MembershipItemMapOutput
}

type MembershipItemMap map[string]MembershipItemInput

func (MembershipItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MembershipItem)(nil)).Elem()
}

func (i MembershipItemMap) ToMembershipItemMapOutput() MembershipItemMapOutput {
	return i.ToMembershipItemMapOutputWithContext(context.Background())
}

func (i MembershipItemMap) ToMembershipItemMapOutputWithContext(ctx context.Context) MembershipItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipItemMapOutput)
}

type MembershipItemOutput struct{ *pulumi.OutputState }

func (MembershipItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipItem)(nil)).Elem()
}

func (o MembershipItemOutput) ToMembershipItemOutput() MembershipItemOutput {
	return o
}

func (o MembershipItemOutput) ToMembershipItemOutputWithContext(ctx context.Context) MembershipItemOutput {
	return o
}

type MembershipItemArrayOutput struct{ *pulumi.OutputState }

func (MembershipItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MembershipItem)(nil)).Elem()
}

func (o MembershipItemArrayOutput) ToMembershipItemArrayOutput() MembershipItemArrayOutput {
	return o
}

func (o MembershipItemArrayOutput) ToMembershipItemArrayOutputWithContext(ctx context.Context) MembershipItemArrayOutput {
	return o
}

func (o MembershipItemArrayOutput) Index(i pulumi.IntInput) MembershipItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MembershipItem {
		return vs[0].([]*MembershipItem)[vs[1].(int)]
	}).(MembershipItemOutput)
}

type MembershipItemMapOutput struct{ *pulumi.OutputState }

func (MembershipItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MembershipItem)(nil)).Elem()
}

func (o MembershipItemMapOutput) ToMembershipItemMapOutput() MembershipItemMapOutput {
	return o
}

func (o MembershipItemMapOutput) ToMembershipItemMapOutputWithContext(ctx context.Context) MembershipItemMapOutput {
	return o
}

func (o MembershipItemMapOutput) MapIndex(k pulumi.StringInput) MembershipItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MembershipItem {
		return vs[0].(map[string]*MembershipItem)[vs[1].(string)]
	}).(MembershipItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipItemInput)(nil)).Elem(), &MembershipItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipItemArrayInput)(nil)).Elem(), MembershipItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipItemMapInput)(nil)).Elem(), MembershipItemMap{})
	pulumi.RegisterOutputType(MembershipItemOutput{})
	pulumi.RegisterOutputType(MembershipItemGetAttachmentResultOutput{})
	pulumi.RegisterOutputType(MembershipItemArrayOutput{})
	pulumi.RegisterOutputType(MembershipItemMapOutput{})
}
