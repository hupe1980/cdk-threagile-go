package plusaws

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdktg.plus_aws.ApplicationLoadBalancer",
		reflect.TypeOf((*ApplicationLoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ciaTriad", GoGetter: "CiaTriad"},
			_jsii_.MemberMethod{JsiiMethod: "communicatesWith", GoMethod: "CommunicatesWith"},
			_jsii_.MemberProperty{JsiiProperty: "customDevelopedParts", GoGetter: "CustomDevelopedParts"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormatsAccepted", GoGetter: "DataFormatsAccepted"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "highestAvailability", GoGetter: "HighestAvailability"},
			_jsii_.MemberProperty{JsiiProperty: "humanUse", GoGetter: "HumanUse"},
			_jsii_.MemberProperty{JsiiProperty: "internet", GoGetter: "Internet"},
			_jsii_.MemberMethod{JsiiMethod: "isTrafficForwarding", GoMethod: "IsTrafficForwarding"},
			_jsii_.MemberMethod{JsiiMethod: "isWebApplication", GoMethod: "IsWebApplication"},
			_jsii_.MemberMethod{JsiiMethod: "isWebService", GoMethod: "IsWebService"},
			_jsii_.MemberProperty{JsiiProperty: "machine", GoGetter: "Machine"},
			_jsii_.MemberProperty{JsiiProperty: "multiTenant", GoGetter: "MultiTenant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberMethod{JsiiMethod: "processes", GoMethod: "Processes"},
			_jsii_.MemberProperty{JsiiProperty: "redundant", GoGetter: "Redundant"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberMethod{JsiiMethod: "stores", GoMethod: "Stores"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "technology", GoGetter: "Technology"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustBoundary", GoGetter: "TrustBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationLoadBalancer{}
			_jsii_.InitJsiiProxy(&j.Type__cdktgTechnicalAsset)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.plus_aws.ApplicationLoadBalancerProps",
		reflect.TypeOf((*ApplicationLoadBalancerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.plus_aws.Cloud",
		reflect.TypeOf((*Cloud)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTechnicalAssets", GoMethod: "AddTechnicalAssets"},
			_jsii_.MemberMethod{JsiiMethod: "addTrustBoundary", GoMethod: "AddTrustBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "isNetworkBoundary", GoMethod: "IsNetworkBoundary"},
			_jsii_.MemberMethod{JsiiMethod: "isWithinCloud", GoMethod: "IsWithinCloud"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_Cloud{}
			_jsii_.InitJsiiProxy(&j.Type__cdktgTrustBoundary)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.plus_aws.CloudProps",
		reflect.TypeOf((*CloudProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.plus_aws.SecurityGroup",
		reflect.TypeOf((*SecurityGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTechnicalAssets", GoMethod: "AddTechnicalAssets"},
			_jsii_.MemberMethod{JsiiMethod: "addTrustBoundary", GoMethod: "AddTrustBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "isNetworkBoundary", GoMethod: "IsNetworkBoundary"},
			_jsii_.MemberMethod{JsiiMethod: "isWithinCloud", GoMethod: "IsWithinCloud"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktgTrustBoundary)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.plus_aws.SecurityGroupProps",
		reflect.TypeOf((*SecurityGroupProps)(nil)).Elem(),
	)
}
