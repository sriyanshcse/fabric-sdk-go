// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/api (interfaces: CoreProviders,SvcProviders,Providers,CoreProviderFactory,ServiceProviderFactory,SessionClientFactory)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	channel "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	context "github.com/hyperledger/fabric-sdk-go/pkg/context"
	api "github.com/hyperledger/fabric-sdk-go/pkg/context/api"
	core "github.com/hyperledger/fabric-sdk-go/pkg/context/api/core"
	fab "github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab"
	api0 "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/api"
)

// MockCoreProviders is a mock of CoreProviders interface
type MockCoreProviders struct {
	ctrl     *gomock.Controller
	recorder *MockCoreProvidersMockRecorder
}

// MockCoreProvidersMockRecorder is the mock recorder for MockCoreProviders
type MockCoreProvidersMockRecorder struct {
	mock *MockCoreProviders
}

// NewMockCoreProviders creates a new mock instance
func NewMockCoreProviders(ctrl *gomock.Controller) *MockCoreProviders {
	mock := &MockCoreProviders{ctrl: ctrl}
	mock.recorder = &MockCoreProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoreProviders) EXPECT() *MockCoreProvidersMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockCoreProviders) Config() core.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(core.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockCoreProvidersMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCoreProviders)(nil).Config))
}

// CryptoSuite mocks base method
func (m *MockCoreProviders) CryptoSuite() core.CryptoSuite {
	ret := m.ctrl.Call(m, "CryptoSuite")
	ret0, _ := ret[0].(core.CryptoSuite)
	return ret0
}

// CryptoSuite indicates an expected call of CryptoSuite
func (mr *MockCoreProvidersMockRecorder) CryptoSuite() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoSuite", reflect.TypeOf((*MockCoreProviders)(nil).CryptoSuite))
}

// FabricProvider mocks base method
func (m *MockCoreProviders) FabricProvider() api0.FabricProvider {
	ret := m.ctrl.Call(m, "FabricProvider")
	ret0, _ := ret[0].(api0.FabricProvider)
	return ret0
}

// FabricProvider indicates an expected call of FabricProvider
func (mr *MockCoreProvidersMockRecorder) FabricProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FabricProvider", reflect.TypeOf((*MockCoreProviders)(nil).FabricProvider))
}

// SigningManager mocks base method
func (m *MockCoreProviders) SigningManager() api.SigningManager {
	ret := m.ctrl.Call(m, "SigningManager")
	ret0, _ := ret[0].(api.SigningManager)
	return ret0
}

// SigningManager indicates an expected call of SigningManager
func (mr *MockCoreProvidersMockRecorder) SigningManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningManager", reflect.TypeOf((*MockCoreProviders)(nil).SigningManager))
}

// StateStore mocks base method
func (m *MockCoreProviders) StateStore() api.KVStore {
	ret := m.ctrl.Call(m, "StateStore")
	ret0, _ := ret[0].(api.KVStore)
	return ret0
}

// StateStore indicates an expected call of StateStore
func (mr *MockCoreProvidersMockRecorder) StateStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateStore", reflect.TypeOf((*MockCoreProviders)(nil).StateStore))
}

// MockSvcProviders is a mock of SvcProviders interface
type MockSvcProviders struct {
	ctrl     *gomock.Controller
	recorder *MockSvcProvidersMockRecorder
}

// MockSvcProvidersMockRecorder is the mock recorder for MockSvcProviders
type MockSvcProvidersMockRecorder struct {
	mock *MockSvcProviders
}

// NewMockSvcProviders creates a new mock instance
func NewMockSvcProviders(ctrl *gomock.Controller) *MockSvcProviders {
	mock := &MockSvcProviders{ctrl: ctrl}
	mock.recorder = &MockSvcProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSvcProviders) EXPECT() *MockSvcProvidersMockRecorder {
	return m.recorder
}

// ChannelProvider mocks base method
func (m *MockSvcProviders) ChannelProvider() fab.ChannelProvider {
	ret := m.ctrl.Call(m, "ChannelProvider")
	ret0, _ := ret[0].(fab.ChannelProvider)
	return ret0
}

// ChannelProvider indicates an expected call of ChannelProvider
func (mr *MockSvcProvidersMockRecorder) ChannelProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelProvider", reflect.TypeOf((*MockSvcProviders)(nil).ChannelProvider))
}

// DiscoveryProvider mocks base method
func (m *MockSvcProviders) DiscoveryProvider() fab.DiscoveryProvider {
	ret := m.ctrl.Call(m, "DiscoveryProvider")
	ret0, _ := ret[0].(fab.DiscoveryProvider)
	return ret0
}

// DiscoveryProvider indicates an expected call of DiscoveryProvider
func (mr *MockSvcProvidersMockRecorder) DiscoveryProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryProvider", reflect.TypeOf((*MockSvcProviders)(nil).DiscoveryProvider))
}

// SelectionProvider mocks base method
func (m *MockSvcProviders) SelectionProvider() fab.SelectionProvider {
	ret := m.ctrl.Call(m, "SelectionProvider")
	ret0, _ := ret[0].(fab.SelectionProvider)
	return ret0
}

// SelectionProvider indicates an expected call of SelectionProvider
func (mr *MockSvcProvidersMockRecorder) SelectionProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectionProvider", reflect.TypeOf((*MockSvcProviders)(nil).SelectionProvider))
}

// MockProviders is a mock of Providers interface
type MockProviders struct {
	ctrl     *gomock.Controller
	recorder *MockProvidersMockRecorder
}

// MockProvidersMockRecorder is the mock recorder for MockProviders
type MockProvidersMockRecorder struct {
	mock *MockProviders
}

// NewMockProviders creates a new mock instance
func NewMockProviders(ctrl *gomock.Controller) *MockProviders {
	mock := &MockProviders{ctrl: ctrl}
	mock.recorder = &MockProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProviders) EXPECT() *MockProvidersMockRecorder {
	return m.recorder
}

// ChannelProvider mocks base method
func (m *MockProviders) ChannelProvider() fab.ChannelProvider {
	ret := m.ctrl.Call(m, "ChannelProvider")
	ret0, _ := ret[0].(fab.ChannelProvider)
	return ret0
}

// ChannelProvider indicates an expected call of ChannelProvider
func (mr *MockProvidersMockRecorder) ChannelProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelProvider", reflect.TypeOf((*MockProviders)(nil).ChannelProvider))
}

// Config mocks base method
func (m *MockProviders) Config() core.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(core.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockProvidersMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockProviders)(nil).Config))
}

// CryptoSuite mocks base method
func (m *MockProviders) CryptoSuite() core.CryptoSuite {
	ret := m.ctrl.Call(m, "CryptoSuite")
	ret0, _ := ret[0].(core.CryptoSuite)
	return ret0
}

// CryptoSuite indicates an expected call of CryptoSuite
func (mr *MockProvidersMockRecorder) CryptoSuite() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoSuite", reflect.TypeOf((*MockProviders)(nil).CryptoSuite))
}

// DiscoveryProvider mocks base method
func (m *MockProviders) DiscoveryProvider() fab.DiscoveryProvider {
	ret := m.ctrl.Call(m, "DiscoveryProvider")
	ret0, _ := ret[0].(fab.DiscoveryProvider)
	return ret0
}

// DiscoveryProvider indicates an expected call of DiscoveryProvider
func (mr *MockProvidersMockRecorder) DiscoveryProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryProvider", reflect.TypeOf((*MockProviders)(nil).DiscoveryProvider))
}

// FabricProvider mocks base method
func (m *MockProviders) FabricProvider() api0.FabricProvider {
	ret := m.ctrl.Call(m, "FabricProvider")
	ret0, _ := ret[0].(api0.FabricProvider)
	return ret0
}

// FabricProvider indicates an expected call of FabricProvider
func (mr *MockProvidersMockRecorder) FabricProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FabricProvider", reflect.TypeOf((*MockProviders)(nil).FabricProvider))
}

// SelectionProvider mocks base method
func (m *MockProviders) SelectionProvider() fab.SelectionProvider {
	ret := m.ctrl.Call(m, "SelectionProvider")
	ret0, _ := ret[0].(fab.SelectionProvider)
	return ret0
}

// SelectionProvider indicates an expected call of SelectionProvider
func (mr *MockProvidersMockRecorder) SelectionProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectionProvider", reflect.TypeOf((*MockProviders)(nil).SelectionProvider))
}

// SigningManager mocks base method
func (m *MockProviders) SigningManager() api.SigningManager {
	ret := m.ctrl.Call(m, "SigningManager")
	ret0, _ := ret[0].(api.SigningManager)
	return ret0
}

// SigningManager indicates an expected call of SigningManager
func (mr *MockProvidersMockRecorder) SigningManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningManager", reflect.TypeOf((*MockProviders)(nil).SigningManager))
}

// StateStore mocks base method
func (m *MockProviders) StateStore() api.KVStore {
	ret := m.ctrl.Call(m, "StateStore")
	ret0, _ := ret[0].(api.KVStore)
	return ret0
}

// StateStore indicates an expected call of StateStore
func (mr *MockProvidersMockRecorder) StateStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateStore", reflect.TypeOf((*MockProviders)(nil).StateStore))
}

// MockCoreProviderFactory is a mock of CoreProviderFactory interface
type MockCoreProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCoreProviderFactoryMockRecorder
}

// MockCoreProviderFactoryMockRecorder is the mock recorder for MockCoreProviderFactory
type MockCoreProviderFactoryMockRecorder struct {
	mock *MockCoreProviderFactory
}

// NewMockCoreProviderFactory creates a new mock instance
func NewMockCoreProviderFactory(ctrl *gomock.Controller) *MockCoreProviderFactory {
	mock := &MockCoreProviderFactory{ctrl: ctrl}
	mock.recorder = &MockCoreProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoreProviderFactory) EXPECT() *MockCoreProviderFactoryMockRecorder {
	return m.recorder
}

// CreateCryptoSuiteProvider mocks base method
func (m *MockCoreProviderFactory) CreateCryptoSuiteProvider(arg0 core.Config) (core.CryptoSuite, error) {
	ret := m.ctrl.Call(m, "CreateCryptoSuiteProvider", arg0)
	ret0, _ := ret[0].(core.CryptoSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCryptoSuiteProvider indicates an expected call of CreateCryptoSuiteProvider
func (mr *MockCoreProviderFactoryMockRecorder) CreateCryptoSuiteProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCryptoSuiteProvider", reflect.TypeOf((*MockCoreProviderFactory)(nil).CreateCryptoSuiteProvider), arg0)
}

// CreateFabricProvider mocks base method
func (m *MockCoreProviderFactory) CreateFabricProvider(arg0 context.ProviderContext) (api0.FabricProvider, error) {
	ret := m.ctrl.Call(m, "CreateFabricProvider", arg0)
	ret0, _ := ret[0].(api0.FabricProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFabricProvider indicates an expected call of CreateFabricProvider
func (mr *MockCoreProviderFactoryMockRecorder) CreateFabricProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFabricProvider", reflect.TypeOf((*MockCoreProviderFactory)(nil).CreateFabricProvider), arg0)
}

// CreateIdentityManager mocks base method
func (m *MockCoreProviderFactory) CreateIdentityManager(arg0 string, arg1 core.CryptoSuite, arg2 core.Config) (api.IdentityManager, error) {
	ret := m.ctrl.Call(m, "CreateIdentityManager", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.IdentityManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIdentityManager indicates an expected call of CreateIdentityManager
func (mr *MockCoreProviderFactoryMockRecorder) CreateIdentityManager(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIdentityManager", reflect.TypeOf((*MockCoreProviderFactory)(nil).CreateIdentityManager), arg0, arg1, arg2)
}

// CreateSigningManager mocks base method
func (m *MockCoreProviderFactory) CreateSigningManager(arg0 core.CryptoSuite, arg1 core.Config) (api.SigningManager, error) {
	ret := m.ctrl.Call(m, "CreateSigningManager", arg0, arg1)
	ret0, _ := ret[0].(api.SigningManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSigningManager indicates an expected call of CreateSigningManager
func (mr *MockCoreProviderFactoryMockRecorder) CreateSigningManager(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSigningManager", reflect.TypeOf((*MockCoreProviderFactory)(nil).CreateSigningManager), arg0, arg1)
}

// CreateStateStoreProvider mocks base method
func (m *MockCoreProviderFactory) CreateStateStoreProvider(arg0 core.Config) (api.KVStore, error) {
	ret := m.ctrl.Call(m, "CreateStateStoreProvider", arg0)
	ret0, _ := ret[0].(api.KVStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStateStoreProvider indicates an expected call of CreateStateStoreProvider
func (mr *MockCoreProviderFactoryMockRecorder) CreateStateStoreProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStateStoreProvider", reflect.TypeOf((*MockCoreProviderFactory)(nil).CreateStateStoreProvider), arg0)
}

// MockServiceProviderFactory is a mock of ServiceProviderFactory interface
type MockServiceProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProviderFactoryMockRecorder
}

// MockServiceProviderFactoryMockRecorder is the mock recorder for MockServiceProviderFactory
type MockServiceProviderFactoryMockRecorder struct {
	mock *MockServiceProviderFactory
}

// NewMockServiceProviderFactory creates a new mock instance
func NewMockServiceProviderFactory(ctrl *gomock.Controller) *MockServiceProviderFactory {
	mock := &MockServiceProviderFactory{ctrl: ctrl}
	mock.recorder = &MockServiceProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceProviderFactory) EXPECT() *MockServiceProviderFactoryMockRecorder {
	return m.recorder
}

// CreateDiscoveryProvider mocks base method
func (m *MockServiceProviderFactory) CreateDiscoveryProvider(arg0 core.Config) (fab.DiscoveryProvider, error) {
	ret := m.ctrl.Call(m, "CreateDiscoveryProvider", arg0)
	ret0, _ := ret[0].(fab.DiscoveryProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDiscoveryProvider indicates an expected call of CreateDiscoveryProvider
func (mr *MockServiceProviderFactoryMockRecorder) CreateDiscoveryProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDiscoveryProvider", reflect.TypeOf((*MockServiceProviderFactory)(nil).CreateDiscoveryProvider), arg0)
}

// CreateSelectionProvider mocks base method
func (m *MockServiceProviderFactory) CreateSelectionProvider(arg0 core.Config) (fab.SelectionProvider, error) {
	ret := m.ctrl.Call(m, "CreateSelectionProvider", arg0)
	ret0, _ := ret[0].(fab.SelectionProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSelectionProvider indicates an expected call of CreateSelectionProvider
func (mr *MockServiceProviderFactoryMockRecorder) CreateSelectionProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSelectionProvider", reflect.TypeOf((*MockServiceProviderFactory)(nil).CreateSelectionProvider), arg0)
}

// MockSessionClientFactory is a mock of SessionClientFactory interface
type MockSessionClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSessionClientFactoryMockRecorder
}

// MockSessionClientFactoryMockRecorder is the mock recorder for MockSessionClientFactory
type MockSessionClientFactoryMockRecorder struct {
	mock *MockSessionClientFactory
}

// NewMockSessionClientFactory creates a new mock instance
func NewMockSessionClientFactory(ctrl *gomock.Controller) *MockSessionClientFactory {
	mock := &MockSessionClientFactory{ctrl: ctrl}
	mock.recorder = &MockSessionClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionClientFactory) EXPECT() *MockSessionClientFactoryMockRecorder {
	return m.recorder
}

// CreateChannelClient mocks base method
func (m *MockSessionClientFactory) CreateChannelClient(arg0 api0.Providers, arg1 context.SessionContext, arg2 string, arg3 fab.TargetFilter) (*channel.Client, error) {
	ret := m.ctrl.Call(m, "CreateChannelClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*channel.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChannelClient indicates an expected call of CreateChannelClient
func (mr *MockSessionClientFactoryMockRecorder) CreateChannelClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannelClient", reflect.TypeOf((*MockSessionClientFactory)(nil).CreateChannelClient), arg0, arg1, arg2, arg3)
}
