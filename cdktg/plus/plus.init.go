package plus

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdktg.plus.Browser",
		reflect.TypeOf((*Browser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ciaTriad", GoGetter: "CiaTriad"},
			_jsii_.MemberMethod{JsiiMethod: "communicatesWith", GoMethod: "CommunicatesWith"},
			_jsii_.MemberProperty{JsiiProperty: "customDevelopedParts", GoGetter: "CustomDevelopedParts"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormatsAccepted", GoGetter: "DataFormatsAccepted"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "highestAvailability", GoGetter: "HighestAvailability"},
			_jsii_.MemberProperty{JsiiProperty: "highestIntegrity", GoGetter: "HighestIntegrity"},
			_jsii_.MemberProperty{JsiiProperty: "humanUse", GoGetter: "HumanUse"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
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
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberMethod{JsiiMethod: "stores", GoMethod: "Stores"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "technology", GoGetter: "Technology"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustBoundary", GoGetter: "TrustBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
		},
		func() interface{} {
			j := jsiiProxy_Browser{}
			_jsii_.InitJsiiProxy(&j.Type__cdktgTechnicalAsset)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.plus.BrowserProps",
		reflect.TypeOf((*BrowserProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.plus.StorageType",
		reflect.TypeOf((*StorageType)(nil)).Elem(),
		map[string]interface{}{
			"CLOUD_PROVIDER": StorageType_CLOUD_PROVIDER,
			"CONTAINER_PLATFORM": StorageType_CONTAINER_PLATFORM,
			"DATABASE": StorageType_DATABASE,
			"FILESYSTEM": StorageType_FILESYSTEM,
			"IN_MEMORY": StorageType_IN_MEMORY,
			"SERVICE_REGISTRY": StorageType_SERVICE_REGISTRY,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.plus.Vault",
		reflect.TypeOf((*Vault)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ciaTriad", GoGetter: "CiaTriad"},
			_jsii_.MemberMethod{JsiiMethod: "communicatesWith", GoMethod: "CommunicatesWith"},
			_jsii_.MemberProperty{JsiiProperty: "configurationSecrets", GoGetter: "ConfigurationSecrets"},
			_jsii_.MemberProperty{JsiiProperty: "customDevelopedParts", GoGetter: "CustomDevelopedParts"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormatsAccepted", GoGetter: "DataFormatsAccepted"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "highestAvailability", GoGetter: "HighestAvailability"},
			_jsii_.MemberProperty{JsiiProperty: "highestIntegrity", GoGetter: "HighestIntegrity"},
			_jsii_.MemberProperty{JsiiProperty: "humanUse", GoGetter: "HumanUse"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "internet", GoGetter: "Internet"},
			_jsii_.MemberMethod{JsiiMethod: "isTrafficForwarding", GoMethod: "IsTrafficForwarding"},
			_jsii_.MemberMethod{JsiiMethod: "isUsedBy", GoMethod: "IsUsedBy"},
			_jsii_.MemberMethod{JsiiMethod: "isWebApplication", GoMethod: "IsWebApplication"},
			_jsii_.MemberMethod{JsiiMethod: "isWebService", GoMethod: "IsWebService"},
			_jsii_.MemberProperty{JsiiProperty: "machine", GoGetter: "Machine"},
			_jsii_.MemberProperty{JsiiProperty: "multiTenant", GoGetter: "MultiTenant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberMethod{JsiiMethod: "processes", GoMethod: "Processes"},
			_jsii_.MemberProperty{JsiiProperty: "redundant", GoGetter: "Redundant"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberMethod{JsiiMethod: "stores", GoMethod: "Stores"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "technology", GoGetter: "Technology"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustBoundary", GoGetter: "TrustBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "vaultStorage", GoGetter: "VaultStorage"},
		},
		func() interface{} {
			j := jsiiProxy_Vault{}
			_jsii_.InitJsiiProxy(&j.Type__cdktgTechnicalAsset)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.plus.VaultProps",
		reflect.TypeOf((*VaultProps)(nil)).Elem(),
	)
}
