package common

import (
	"github.com/juju/go-oracle-cloud/api"
	"github.com/juju/go-oracle-cloud/common"
	"github.com/juju/go-oracle-cloud/response"
)

type Instancer interface {
	InstanceDetails(string) (response.Instance, error)
}

type Composer interface {
	ComposeName(string) string
}

type RulesAPI interface {
	AllSecRules([]api.Filter) (response.AllSecRules, error)
	DeleteSecRule(string) error
	CreateSecRule(api.SecRuleParams) (response.SecRule, error)
}

type AclAPI interface {
	AclDetails(string) (response.Acl, error)
	CreateAcl(string, string, bool, []string) (response.Acl, error)
	DeleteAcl(string) error
}

type SecIpAPI interface {
	AllSecIpLists([]api.Filter) (response.AllSecIpLists, error)
	CreateSecIpList(string, string, []string) (response.SecIpList, error)
	AllDefaultSecIpLists([]api.Filter) (response.AllSecIpLists, error)
}

type IpAddressPrefixSetAPI interface {
	CreateIpAddressPrefixSet(
		api.IpAddressPrefixSetParams) (response.IpAddressPrefixSet, error)

	AllIpAddressPrefixSets([]api.Filter) (response.AllIpAddressPrefixSets, error)
}

type SecListAPI interface {
	SecListDetails(string) (response.SecList, error)
	DeleteSecList(string) error
	CreateSecList(string, string,
		common.SecRuleAction, common.SecRuleAction) (response.SecList, error)
}

type SecRulesAPI interface {
	AllSecurityRules([]api.Filter) (response.AllSecurityRules, error)
	DeleteSecurityRule(string) error
	CreateSecurityRule(api.SecurityRuleParams) (response.SecurityRule, error)
}

type ApplicationsAPI interface {
	AllSecApplications([]api.Filter) (response.AllSecApplications, error)
	DefaultSecApplications([]api.Filter) (response.AllSecApplications, error)
	CreateSecApplication(api.SecApplicationParams) (response.SecApplication, error)
}

type AssociationAPI interface {
	AllSecAssociations([]api.Filter) (response.AllSecAssociations, error)
}

type StorageVolumeAPI interface {
	CreateStorageVolume(p api.StorageVolumeParams) (resp response.StorageVolume, err error)
	DeleteStorageVolume(name string) (err error)
	StorageVolumeDetails(name string) (resp response.StorageVolume, err error)
	AllStorageVolumes(filter []api.Filter) (resp response.AllStorageVolumes, err error)
	UpdateStorageVolume(p api.StorageVolumeParams, currentName string) (resp response.StorageVolume, err error)
}

type StorageAttachmentAPI interface {
	CreateStorageAttachment(p api.StorageAttachmentParams) (response.StorageAttachment, error)
	DeleteStorageAttachment(name string) error
	StorageAttachmentDetails(name string) (response.StorageAttachment, error)
	AllStorageAttachments(filter []api.Filter) (response.AllStorageAttachments, error)
}
