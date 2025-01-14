// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/description/v4 (interfaces: Machine,MachinePortRanges,UnitPortRanges)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	description "github.com/juju/description/v4"
	names "github.com/juju/names/v4"
)

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// AddBlockDevice mocks base method.
func (m *MockMachine) AddBlockDevice(arg0 description.BlockDeviceArgs) description.BlockDevice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBlockDevice", arg0)
	ret0, _ := ret[0].(description.BlockDevice)
	return ret0
}

// AddBlockDevice indicates an expected call of AddBlockDevice.
func (mr *MockMachineMockRecorder) AddBlockDevice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlockDevice", reflect.TypeOf((*MockMachine)(nil).AddBlockDevice), arg0)
}

// AddContainer mocks base method.
func (m *MockMachine) AddContainer(arg0 description.MachineArgs) description.Machine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddContainer", arg0)
	ret0, _ := ret[0].(description.Machine)
	return ret0
}

// AddContainer indicates an expected call of AddContainer.
func (mr *MockMachineMockRecorder) AddContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddContainer", reflect.TypeOf((*MockMachine)(nil).AddContainer), arg0)
}

// AddOpenedPortRange mocks base method.
func (m *MockMachine) AddOpenedPortRange(arg0 description.OpenedPortRangeArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddOpenedPortRange", arg0)
}

// AddOpenedPortRange indicates an expected call of AddOpenedPortRange.
func (mr *MockMachineMockRecorder) AddOpenedPortRange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOpenedPortRange", reflect.TypeOf((*MockMachine)(nil).AddOpenedPortRange), arg0)
}

// Annotations mocks base method.
func (m *MockMachine) Annotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Annotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Annotations indicates an expected call of Annotations.
func (mr *MockMachineMockRecorder) Annotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Annotations", reflect.TypeOf((*MockMachine)(nil).Annotations))
}

// Base mocks base method.
func (m *MockMachine) Base() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Base")
	ret0, _ := ret[0].(string)
	return ret0
}

// Base indicates an expected call of Base.
func (mr *MockMachineMockRecorder) Base() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Base", reflect.TypeOf((*MockMachine)(nil).Base))
}

// BlockDevices mocks base method.
func (m *MockMachine) BlockDevices() []description.BlockDevice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockDevices")
	ret0, _ := ret[0].([]description.BlockDevice)
	return ret0
}

// BlockDevices indicates an expected call of BlockDevices.
func (mr *MockMachineMockRecorder) BlockDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockDevices", reflect.TypeOf((*MockMachine)(nil).BlockDevices))
}

// Constraints mocks base method.
func (m *MockMachine) Constraints() description.Constraints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Constraints")
	ret0, _ := ret[0].(description.Constraints)
	return ret0
}

// Constraints indicates an expected call of Constraints.
func (mr *MockMachineMockRecorder) Constraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constraints", reflect.TypeOf((*MockMachine)(nil).Constraints))
}

// ContainerType mocks base method.
func (m *MockMachine) ContainerType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContainerType indicates an expected call of ContainerType.
func (mr *MockMachineMockRecorder) ContainerType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockMachine)(nil).ContainerType))
}

// Containers mocks base method.
func (m *MockMachine) Containers() []description.Machine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Containers")
	ret0, _ := ret[0].([]description.Machine)
	return ret0
}

// Containers indicates an expected call of Containers.
func (mr *MockMachineMockRecorder) Containers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Containers", reflect.TypeOf((*MockMachine)(nil).Containers))
}

// Id mocks base method.
func (m *MockMachine) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockMachineMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockMachine)(nil).Id))
}

// Instance mocks base method.
func (m *MockMachine) Instance() description.CloudInstance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instance")
	ret0, _ := ret[0].(description.CloudInstance)
	return ret0
}

// Instance indicates an expected call of Instance.
func (mr *MockMachineMockRecorder) Instance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instance", reflect.TypeOf((*MockMachine)(nil).Instance))
}

// Jobs mocks base method.
func (m *MockMachine) Jobs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Jobs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Jobs indicates an expected call of Jobs.
func (mr *MockMachineMockRecorder) Jobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jobs", reflect.TypeOf((*MockMachine)(nil).Jobs))
}

// MachineAddresses mocks base method.
func (m *MockMachine) MachineAddresses() []description.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineAddresses")
	ret0, _ := ret[0].([]description.Address)
	return ret0
}

// MachineAddresses indicates an expected call of MachineAddresses.
func (mr *MockMachineMockRecorder) MachineAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineAddresses", reflect.TypeOf((*MockMachine)(nil).MachineAddresses))
}

// Nonce mocks base method.
func (m *MockMachine) Nonce() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nonce")
	ret0, _ := ret[0].(string)
	return ret0
}

// Nonce indicates an expected call of Nonce.
func (mr *MockMachineMockRecorder) Nonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockMachine)(nil).Nonce))
}

// OpenedPortRanges mocks base method.
func (m *MockMachine) OpenedPortRanges() description.MachinePortRanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenedPortRanges")
	ret0, _ := ret[0].(description.MachinePortRanges)
	return ret0
}

// OpenedPortRanges indicates an expected call of OpenedPortRanges.
func (mr *MockMachineMockRecorder) OpenedPortRanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenedPortRanges", reflect.TypeOf((*MockMachine)(nil).OpenedPortRanges))
}

// PasswordHash mocks base method.
func (m *MockMachine) PasswordHash() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordHash")
	ret0, _ := ret[0].(string)
	return ret0
}

// PasswordHash indicates an expected call of PasswordHash.
func (mr *MockMachineMockRecorder) PasswordHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordHash", reflect.TypeOf((*MockMachine)(nil).PasswordHash))
}

// Placement mocks base method.
func (m *MockMachine) Placement() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Placement")
	ret0, _ := ret[0].(string)
	return ret0
}

// Placement indicates an expected call of Placement.
func (mr *MockMachineMockRecorder) Placement() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Placement", reflect.TypeOf((*MockMachine)(nil).Placement))
}

// PreferredPrivateAddress mocks base method.
func (m *MockMachine) PreferredPrivateAddress() description.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreferredPrivateAddress")
	ret0, _ := ret[0].(description.Address)
	return ret0
}

// PreferredPrivateAddress indicates an expected call of PreferredPrivateAddress.
func (mr *MockMachineMockRecorder) PreferredPrivateAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreferredPrivateAddress", reflect.TypeOf((*MockMachine)(nil).PreferredPrivateAddress))
}

// PreferredPublicAddress mocks base method.
func (m *MockMachine) PreferredPublicAddress() description.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreferredPublicAddress")
	ret0, _ := ret[0].(description.Address)
	return ret0
}

// PreferredPublicAddress indicates an expected call of PreferredPublicAddress.
func (mr *MockMachineMockRecorder) PreferredPublicAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreferredPublicAddress", reflect.TypeOf((*MockMachine)(nil).PreferredPublicAddress))
}

// ProviderAddresses mocks base method.
func (m *MockMachine) ProviderAddresses() []description.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProviderAddresses")
	ret0, _ := ret[0].([]description.Address)
	return ret0
}

// ProviderAddresses indicates an expected call of ProviderAddresses.
func (mr *MockMachineMockRecorder) ProviderAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProviderAddresses", reflect.TypeOf((*MockMachine)(nil).ProviderAddresses))
}

// SetAddresses mocks base method.
func (m *MockMachine) SetAddresses(arg0, arg1 []description.AddressArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAddresses", arg0, arg1)
}

// SetAddresses indicates an expected call of SetAddresses.
func (mr *MockMachineMockRecorder) SetAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAddresses", reflect.TypeOf((*MockMachine)(nil).SetAddresses), arg0, arg1)
}

// SetAnnotations mocks base method.
func (m *MockMachine) SetAnnotations(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotations", arg0)
}

// SetAnnotations indicates an expected call of SetAnnotations.
func (mr *MockMachineMockRecorder) SetAnnotations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockMachine)(nil).SetAnnotations), arg0)
}

// SetConstraints mocks base method.
func (m *MockMachine) SetConstraints(arg0 description.ConstraintsArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConstraints", arg0)
}

// SetConstraints indicates an expected call of SetConstraints.
func (mr *MockMachineMockRecorder) SetConstraints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConstraints", reflect.TypeOf((*MockMachine)(nil).SetConstraints), arg0)
}

// SetInstance mocks base method.
func (m *MockMachine) SetInstance(arg0 description.CloudInstanceArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInstance", arg0)
}

// SetInstance indicates an expected call of SetInstance.
func (mr *MockMachineMockRecorder) SetInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInstance", reflect.TypeOf((*MockMachine)(nil).SetInstance), arg0)
}

// SetPreferredAddresses mocks base method.
func (m *MockMachine) SetPreferredAddresses(arg0, arg1 description.AddressArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPreferredAddresses", arg0, arg1)
}

// SetPreferredAddresses indicates an expected call of SetPreferredAddresses.
func (mr *MockMachineMockRecorder) SetPreferredAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPreferredAddresses", reflect.TypeOf((*MockMachine)(nil).SetPreferredAddresses), arg0, arg1)
}

// SetStatus mocks base method.
func (m *MockMachine) SetStatus(arg0 description.StatusArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStatus", arg0)
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockMachineMockRecorder) SetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockMachine)(nil).SetStatus), arg0)
}

// SetStatusHistory mocks base method.
func (m *MockMachine) SetStatusHistory(arg0 []description.StatusArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStatusHistory", arg0)
}

// SetStatusHistory indicates an expected call of SetStatusHistory.
func (mr *MockMachineMockRecorder) SetStatusHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatusHistory", reflect.TypeOf((*MockMachine)(nil).SetStatusHistory), arg0)
}

// SetTools mocks base method.
func (m *MockMachine) SetTools(arg0 description.AgentToolsArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTools", arg0)
}

// SetTools indicates an expected call of SetTools.
func (mr *MockMachineMockRecorder) SetTools(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTools", reflect.TypeOf((*MockMachine)(nil).SetTools), arg0)
}

// Status mocks base method.
func (m *MockMachine) Status() description.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(description.Status)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockMachineMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockMachine)(nil).Status))
}

// StatusHistory mocks base method.
func (m *MockMachine) StatusHistory() []description.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusHistory")
	ret0, _ := ret[0].([]description.Status)
	return ret0
}

// StatusHistory indicates an expected call of StatusHistory.
func (mr *MockMachineMockRecorder) StatusHistory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusHistory", reflect.TypeOf((*MockMachine)(nil).StatusHistory))
}

// SupportedContainers mocks base method.
func (m *MockMachine) SupportedContainers() ([]string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportedContainers")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// SupportedContainers indicates an expected call of SupportedContainers.
func (mr *MockMachineMockRecorder) SupportedContainers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportedContainers", reflect.TypeOf((*MockMachine)(nil).SupportedContainers))
}

// Tag mocks base method.
func (m *MockMachine) Tag() names.MachineTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.MachineTag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockMachineMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockMachine)(nil).Tag))
}

// Tools mocks base method.
func (m *MockMachine) Tools() description.AgentTools {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tools")
	ret0, _ := ret[0].(description.AgentTools)
	return ret0
}

// Tools indicates an expected call of Tools.
func (mr *MockMachineMockRecorder) Tools() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tools", reflect.TypeOf((*MockMachine)(nil).Tools))
}

// Validate mocks base method.
func (m *MockMachine) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockMachineMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockMachine)(nil).Validate))
}

// MockMachinePortRanges is a mock of MachinePortRanges interface.
type MockMachinePortRanges struct {
	ctrl     *gomock.Controller
	recorder *MockMachinePortRangesMockRecorder
}

// MockMachinePortRangesMockRecorder is the mock recorder for MockMachinePortRanges.
type MockMachinePortRangesMockRecorder struct {
	mock *MockMachinePortRanges
}

// NewMockMachinePortRanges creates a new mock instance.
func NewMockMachinePortRanges(ctrl *gomock.Controller) *MockMachinePortRanges {
	mock := &MockMachinePortRanges{ctrl: ctrl}
	mock.recorder = &MockMachinePortRangesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachinePortRanges) EXPECT() *MockMachinePortRangesMockRecorder {
	return m.recorder
}

// ByUnit mocks base method.
func (m *MockMachinePortRanges) ByUnit() map[string]description.UnitPortRanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByUnit")
	ret0, _ := ret[0].(map[string]description.UnitPortRanges)
	return ret0
}

// ByUnit indicates an expected call of ByUnit.
func (mr *MockMachinePortRangesMockRecorder) ByUnit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByUnit", reflect.TypeOf((*MockMachinePortRanges)(nil).ByUnit))
}

// MockUnitPortRanges is a mock of UnitPortRanges interface.
type MockUnitPortRanges struct {
	ctrl     *gomock.Controller
	recorder *MockUnitPortRangesMockRecorder
}

// MockUnitPortRangesMockRecorder is the mock recorder for MockUnitPortRanges.
type MockUnitPortRangesMockRecorder struct {
	mock *MockUnitPortRanges
}

// NewMockUnitPortRanges creates a new mock instance.
func NewMockUnitPortRanges(ctrl *gomock.Controller) *MockUnitPortRanges {
	mock := &MockUnitPortRanges{ctrl: ctrl}
	mock.recorder = &MockUnitPortRangesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitPortRanges) EXPECT() *MockUnitPortRangesMockRecorder {
	return m.recorder
}

// ByEndpoint mocks base method.
func (m *MockUnitPortRanges) ByEndpoint() map[string][]description.UnitPortRange {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByEndpoint")
	ret0, _ := ret[0].(map[string][]description.UnitPortRange)
	return ret0
}

// ByEndpoint indicates an expected call of ByEndpoint.
func (mr *MockUnitPortRangesMockRecorder) ByEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByEndpoint", reflect.TypeOf((*MockUnitPortRanges)(nil).ByEndpoint))
}
