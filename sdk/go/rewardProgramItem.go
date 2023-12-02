// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unoffical

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/rewardprogram"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RewardProgramItem struct {
	pulumi.CustomResourceState

	Attachments     OutAttachmentMapOutput                        `pulumi:"attachments"`
	Category        pulumi.StringOutput                           `pulumi:"category"`
	CompanyName     pulumi.StringPtrOutput                        `pulumi:"companyName"`
	Fields          OutFieldMapOutput                             `pulumi:"fields"`
	MemberId        pulumi.StringPtrOutput                        `pulumi:"memberId"`
	MemberName      pulumi.StringPtrOutput                        `pulumi:"memberName"`
	MoreInformation rewardprogram.MoreInformationSectionPtrOutput `pulumi:"moreInformation"`
	Notes           pulumi.StringPtrOutput                        `pulumi:"notes"`
	Pin             pulumi.StringPtrOutput                        `pulumi:"pin"`
	References      OutFieldMapOutput                             `pulumi:"references"`
	Sections        OutSectionMapOutput                           `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewRewardProgramItem registers a new resource with the given unique name, arguments, and options.
func NewRewardProgramItem(ctx *pulumi.Context,
	name string, args *RewardProgramItemArgs, opts ...pulumi.ResourceOption) (*RewardProgramItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("Reward Program")
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
	var resource RewardProgramItem
	err := ctx.RegisterResource("one-password-native-unoffical:index:RewardProgramItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRewardProgramItem gets an existing RewardProgramItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRewardProgramItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RewardProgramItemState, opts ...pulumi.ResourceOption) (*RewardProgramItem, error) {
	var resource RewardProgramItem
	err := ctx.ReadResource("one-password-native-unoffical:index:RewardProgramItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RewardProgramItem resources.
type rewardProgramItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type RewardProgramItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (RewardProgramItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*rewardProgramItemState)(nil)).Elem()
}

type rewardProgramItemArgs struct {
	// The category of the vault the item is in.
	Category         *string                               `pulumi:"category"`
	CompanyName      *string                               `pulumi:"companyName"`
	Fields           map[string]Field                      `pulumi:"fields"`
	InputAttachments map[string]pulumi.AssetOrArchive      `pulumi:"inputAttachments"`
	MemberId         *string                               `pulumi:"memberId"`
	MemberName       *string                               `pulumi:"memberName"`
	MoreInformation  *rewardprogram.MoreInformationSection `pulumi:"moreInformation"`
	Notes            *string                               `pulumi:"notes"`
	Pin              *string                               `pulumi:"pin"`
	Sections         map[string]Section                    `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a RewardProgramItem resource.
type RewardProgramItemArgs struct {
	// The category of the vault the item is in.
	Category         pulumi.StringPtrInput
	CompanyName      pulumi.StringPtrInput
	Fields           FieldMapInput
	InputAttachments pulumi.AssetOrArchiveMapInput
	MemberId         pulumi.StringPtrInput
	MemberName       pulumi.StringPtrInput
	MoreInformation  rewardprogram.MoreInformationSectionPtrInput
	Notes            pulumi.StringPtrInput
	Pin              pulumi.StringPtrInput
	Sections         SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (RewardProgramItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rewardProgramItemArgs)(nil)).Elem()
}

func (r *RewardProgramItem) GetAttachment(ctx *pulumi.Context, args *RewardProgramItemGetAttachmentArgs) (RewardProgramItemGetAttachmentResultOutput, error) {
	out, err := ctx.Call("one-password-native-unoffical:index:RewardProgramItem/attachment", args, RewardProgramItemGetAttachmentResultOutput{}, r)
	if err != nil {
		return RewardProgramItemGetAttachmentResultOutput{}, err
	}
	return out.(RewardProgramItemGetAttachmentResultOutput), nil
}

type rewardProgramItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name string `pulumi:"name"`
}

// The set of arguments for the GetAttachment method of the RewardProgramItem resource.
type RewardProgramItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name pulumi.StringInput
}

func (RewardProgramItemGetAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rewardProgramItemGetAttachmentArgs)(nil)).Elem()
}

// The resolved reference value
type RewardProgramItemGetAttachmentResult struct {
	// the value of the attachment
	Value string `pulumi:"value"`
}

type RewardProgramItemGetAttachmentResultOutput struct{ *pulumi.OutputState }

func (RewardProgramItemGetAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RewardProgramItemGetAttachmentResult)(nil)).Elem()
}

// the value of the attachment
func (o RewardProgramItemGetAttachmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v RewardProgramItemGetAttachmentResult) string { return v.Value }).(pulumi.StringOutput)
}

type RewardProgramItemInput interface {
	pulumi.Input

	ToRewardProgramItemOutput() RewardProgramItemOutput
	ToRewardProgramItemOutputWithContext(ctx context.Context) RewardProgramItemOutput
}

func (*RewardProgramItem) ElementType() reflect.Type {
	return reflect.TypeOf((**RewardProgramItem)(nil)).Elem()
}

func (i *RewardProgramItem) ToRewardProgramItemOutput() RewardProgramItemOutput {
	return i.ToRewardProgramItemOutputWithContext(context.Background())
}

func (i *RewardProgramItem) ToRewardProgramItemOutputWithContext(ctx context.Context) RewardProgramItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RewardProgramItemOutput)
}

// RewardProgramItemArrayInput is an input type that accepts RewardProgramItemArray and RewardProgramItemArrayOutput values.
// You can construct a concrete instance of `RewardProgramItemArrayInput` via:
//
//	RewardProgramItemArray{ RewardProgramItemArgs{...} }
type RewardProgramItemArrayInput interface {
	pulumi.Input

	ToRewardProgramItemArrayOutput() RewardProgramItemArrayOutput
	ToRewardProgramItemArrayOutputWithContext(context.Context) RewardProgramItemArrayOutput
}

type RewardProgramItemArray []RewardProgramItemInput

func (RewardProgramItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RewardProgramItem)(nil)).Elem()
}

func (i RewardProgramItemArray) ToRewardProgramItemArrayOutput() RewardProgramItemArrayOutput {
	return i.ToRewardProgramItemArrayOutputWithContext(context.Background())
}

func (i RewardProgramItemArray) ToRewardProgramItemArrayOutputWithContext(ctx context.Context) RewardProgramItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RewardProgramItemArrayOutput)
}

// RewardProgramItemMapInput is an input type that accepts RewardProgramItemMap and RewardProgramItemMapOutput values.
// You can construct a concrete instance of `RewardProgramItemMapInput` via:
//
//	RewardProgramItemMap{ "key": RewardProgramItemArgs{...} }
type RewardProgramItemMapInput interface {
	pulumi.Input

	ToRewardProgramItemMapOutput() RewardProgramItemMapOutput
	ToRewardProgramItemMapOutputWithContext(context.Context) RewardProgramItemMapOutput
}

type RewardProgramItemMap map[string]RewardProgramItemInput

func (RewardProgramItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RewardProgramItem)(nil)).Elem()
}

func (i RewardProgramItemMap) ToRewardProgramItemMapOutput() RewardProgramItemMapOutput {
	return i.ToRewardProgramItemMapOutputWithContext(context.Background())
}

func (i RewardProgramItemMap) ToRewardProgramItemMapOutputWithContext(ctx context.Context) RewardProgramItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RewardProgramItemMapOutput)
}

type RewardProgramItemOutput struct{ *pulumi.OutputState }

func (RewardProgramItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RewardProgramItem)(nil)).Elem()
}

func (o RewardProgramItemOutput) ToRewardProgramItemOutput() RewardProgramItemOutput {
	return o
}

func (o RewardProgramItemOutput) ToRewardProgramItemOutputWithContext(ctx context.Context) RewardProgramItemOutput {
	return o
}

type RewardProgramItemArrayOutput struct{ *pulumi.OutputState }

func (RewardProgramItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RewardProgramItem)(nil)).Elem()
}

func (o RewardProgramItemArrayOutput) ToRewardProgramItemArrayOutput() RewardProgramItemArrayOutput {
	return o
}

func (o RewardProgramItemArrayOutput) ToRewardProgramItemArrayOutputWithContext(ctx context.Context) RewardProgramItemArrayOutput {
	return o
}

func (o RewardProgramItemArrayOutput) Index(i pulumi.IntInput) RewardProgramItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RewardProgramItem {
		return vs[0].([]*RewardProgramItem)[vs[1].(int)]
	}).(RewardProgramItemOutput)
}

type RewardProgramItemMapOutput struct{ *pulumi.OutputState }

func (RewardProgramItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RewardProgramItem)(nil)).Elem()
}

func (o RewardProgramItemMapOutput) ToRewardProgramItemMapOutput() RewardProgramItemMapOutput {
	return o
}

func (o RewardProgramItemMapOutput) ToRewardProgramItemMapOutputWithContext(ctx context.Context) RewardProgramItemMapOutput {
	return o
}

func (o RewardProgramItemMapOutput) MapIndex(k pulumi.StringInput) RewardProgramItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RewardProgramItem {
		return vs[0].(map[string]*RewardProgramItem)[vs[1].(string)]
	}).(RewardProgramItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RewardProgramItemInput)(nil)).Elem(), &RewardProgramItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*RewardProgramItemArrayInput)(nil)).Elem(), RewardProgramItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RewardProgramItemMapInput)(nil)).Elem(), RewardProgramItemMap{})
	pulumi.RegisterOutputType(RewardProgramItemOutput{})
	pulumi.RegisterOutputType(RewardProgramItemGetAttachmentResultOutput{})
	pulumi.RegisterOutputType(RewardProgramItemArrayOutput{})
	pulumi.RegisterOutputType(RewardProgramItemMapOutput{})
}
