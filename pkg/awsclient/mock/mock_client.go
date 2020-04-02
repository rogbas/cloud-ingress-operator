// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/awsclient/client.go

// Package mock_awsclient is a generated GoMock package.
package mock_awsclient

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	elb "github.com/aws/aws-sdk-go/service/elb"
	elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
	route53 "github.com/aws/aws-sdk-go/service/route53"
	gomock "github.com/golang/mock/gomock"
	awsclient "github.com/openshift/cloud-ingress-operator/pkg/awsclient"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ApplySecurityGroupsToLoadBalancer mocks base method
func (m *MockClient) ApplySecurityGroupsToLoadBalancer(arg0 *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplySecurityGroupsToLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.ApplySecurityGroupsToLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplySecurityGroupsToLoadBalancer indicates an expected call of ApplySecurityGroupsToLoadBalancer
func (mr *MockClientMockRecorder) ApplySecurityGroupsToLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplySecurityGroupsToLoadBalancer", reflect.TypeOf((*MockClient)(nil).ApplySecurityGroupsToLoadBalancer), arg0)
}

// ConfigureHealthCheck mocks base method
func (m *MockClient) ConfigureHealthCheck(arg0 *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureHealthCheck", arg0)
	ret0, _ := ret[0].(*elb.ConfigureHealthCheckOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigureHealthCheck indicates an expected call of ConfigureHealthCheck
func (mr *MockClientMockRecorder) ConfigureHealthCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureHealthCheck", reflect.TypeOf((*MockClient)(nil).ConfigureHealthCheck), arg0)
}

// CreateLoadBalancer mocks base method
func (m *MockClient) CreateLoadBalancer(arg0 *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.CreateLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancer indicates an expected call of CreateLoadBalancer
func (mr *MockClientMockRecorder) CreateLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancer", reflect.TypeOf((*MockClient)(nil).CreateLoadBalancer), arg0)
}

// CreateLoadBalancerListeners mocks base method
func (m *MockClient) CreateLoadBalancerListeners(arg0 *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancerListeners", arg0)
	ret0, _ := ret[0].(*elb.CreateLoadBalancerListenersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancerListeners indicates an expected call of CreateLoadBalancerListeners
func (mr *MockClientMockRecorder) CreateLoadBalancerListeners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancerListeners", reflect.TypeOf((*MockClient)(nil).CreateLoadBalancerListeners), arg0)
}

// DeregisterInstancesFromLoadBalancer mocks base method
func (m *MockClient) DeregisterInstancesFromLoadBalancer(arg0 *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterInstancesFromLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.DeregisterInstancesFromLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterInstancesFromLoadBalancer indicates an expected call of DeregisterInstancesFromLoadBalancer
func (mr *MockClientMockRecorder) DeregisterInstancesFromLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterInstancesFromLoadBalancer", reflect.TypeOf((*MockClient)(nil).DeregisterInstancesFromLoadBalancer), arg0)
}

// DescribeLoadBalancers mocks base method
func (m *MockClient) DescribeLoadBalancers(arg0 *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancers", arg0)
	ret0, _ := ret[0].(*elb.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancers indicates an expected call of DescribeLoadBalancers
func (mr *MockClientMockRecorder) DescribeLoadBalancers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*MockClient)(nil).DescribeLoadBalancers), arg0)
}

// DescribeTags mocks base method
func (m *MockClient) DescribeTags(arg0 *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTags", arg0)
	ret0, _ := ret[0].(*elb.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTags indicates an expected call of DescribeTags
func (mr *MockClientMockRecorder) DescribeTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockClient)(nil).DescribeTags), arg0)
}

// DeleteLoadBalancerListeners mocks base method
func (m *MockClient) DeleteLoadBalancerListeners(arg0 *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancerListeners", arg0)
	ret0, _ := ret[0].(*elb.DeleteLoadBalancerListenersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLoadBalancerListeners indicates an expected call of DeleteLoadBalancerListeners
func (mr *MockClientMockRecorder) DeleteLoadBalancerListeners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancerListeners", reflect.TypeOf((*MockClient)(nil).DeleteLoadBalancerListeners), arg0)
}

// RegisterInstancesWithLoadBalancer mocks base method
func (m *MockClient) RegisterInstancesWithLoadBalancer(arg0 *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInstancesWithLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.RegisterInstancesWithLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInstancesWithLoadBalancer indicates an expected call of RegisterInstancesWithLoadBalancer
func (mr *MockClientMockRecorder) RegisterInstancesWithLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInstancesWithLoadBalancer", reflect.TypeOf((*MockClient)(nil).RegisterInstancesWithLoadBalancer), arg0)
}

// DescribeLoadBalancersV2 mocks base method
func (m *MockClient) DescribeLoadBalancersV2(arg0 *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancersV2", arg0)
	ret0, _ := ret[0].(*elbv2.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancersV2 indicates an expected call of DescribeLoadBalancersV2
func (mr *MockClientMockRecorder) DescribeLoadBalancersV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancersV2", reflect.TypeOf((*MockClient)(nil).DescribeLoadBalancersV2), arg0)
}

// DeleteLoadBalancerV2 mocks base method
func (m *MockClient) DeleteLoadBalancerV2(arg0 *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancerV2", arg0)
	ret0, _ := ret[0].(*elbv2.DeleteLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLoadBalancerV2 indicates an expected call of DeleteLoadBalancerV2
func (mr *MockClientMockRecorder) DeleteLoadBalancerV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancerV2", reflect.TypeOf((*MockClient)(nil).DeleteLoadBalancerV2), arg0)
}

// CreateLoadBalancerV2 mocks base method
func (m *MockClient) CreateLoadBalancerV2(arg0 *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancerV2", arg0)
	ret0, _ := ret[0].(*elbv2.CreateLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancerV2 indicates an expected call of CreateLoadBalancerV2
func (mr *MockClientMockRecorder) CreateLoadBalancerV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancerV2", reflect.TypeOf((*MockClient)(nil).CreateLoadBalancerV2), arg0)
}

// CreateTargetGroupV2 mocks base method
func (m *MockClient) CreateTargetGroupV2(arg0 *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTargetGroupV2", arg0)
	ret0, _ := ret[0].(*elbv2.CreateTargetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTargetGroupV2 indicates an expected call of CreateTargetGroupV2
func (mr *MockClientMockRecorder) CreateTargetGroupV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTargetGroupV2", reflect.TypeOf((*MockClient)(nil).CreateTargetGroupV2), arg0)
}

// RegisterTargetsV2 mocks base method
func (m *MockClient) RegisterTargetsV2(arg0 *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTargetsV2", arg0)
	ret0, _ := ret[0].(*elbv2.RegisterTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTargetsV2 indicates an expected call of RegisterTargetsV2
func (mr *MockClientMockRecorder) RegisterTargetsV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTargetsV2", reflect.TypeOf((*MockClient)(nil).RegisterTargetsV2), arg0)
}

// CreateListenerV2 mocks base method
func (m *MockClient) CreateListenerV2(arg0 *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateListenerV2", arg0)
	ret0, _ := ret[0].(*elbv2.CreateListenerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateListenerV2 indicates an expected call of CreateListenerV2
func (mr *MockClientMockRecorder) CreateListenerV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateListenerV2", reflect.TypeOf((*MockClient)(nil).CreateListenerV2), arg0)
}

// DescribeTargetGroupsV2 mocks base method
func (m *MockClient) DescribeTargetGroupsV2(arg0 *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTargetGroupsV2", arg0)
	ret0, _ := ret[0].(*elbv2.DescribeTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTargetGroupsV2 indicates an expected call of DescribeTargetGroupsV2
func (mr *MockClientMockRecorder) DescribeTargetGroupsV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetGroupsV2", reflect.TypeOf((*MockClient)(nil).DescribeTargetGroupsV2), arg0)
}

// ChangeResourceRecordSets mocks base method
func (m *MockClient) ChangeResourceRecordSets(arg0 *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeResourceRecordSets", arg0)
	ret0, _ := ret[0].(*route53.ChangeResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeResourceRecordSets indicates an expected call of ChangeResourceRecordSets
func (mr *MockClientMockRecorder) ChangeResourceRecordSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeResourceRecordSets", reflect.TypeOf((*MockClient)(nil).ChangeResourceRecordSets), arg0)
}

// ListHostedZonesByName mocks base method
func (m *MockClient) ListHostedZonesByName(arg0 *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostedZonesByName", arg0)
	ret0, _ := ret[0].(*route53.ListHostedZonesByNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZonesByName indicates an expected call of ListHostedZonesByName
func (mr *MockClientMockRecorder) ListHostedZonesByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZonesByName", reflect.TypeOf((*MockClient)(nil).ListHostedZonesByName), arg0)
}

// AuthorizeSecurityGroupIngress mocks base method
func (m *MockClient) AuthorizeSecurityGroupIngress(arg0 *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeSecurityGroupIngress", arg0)
	ret0, _ := ret[0].(*ec2.AuthorizeSecurityGroupIngressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizeSecurityGroupIngress indicates an expected call of AuthorizeSecurityGroupIngress
func (mr *MockClientMockRecorder) AuthorizeSecurityGroupIngress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeSecurityGroupIngress", reflect.TypeOf((*MockClient)(nil).AuthorizeSecurityGroupIngress), arg0)
}

// CreateSecurityGroup mocks base method
func (m *MockClient) CreateSecurityGroup(arg0 *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecurityGroup", arg0)
	ret0, _ := ret[0].(*ec2.CreateSecurityGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecurityGroup indicates an expected call of CreateSecurityGroup
func (mr *MockClientMockRecorder) CreateSecurityGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityGroup", reflect.TypeOf((*MockClient)(nil).CreateSecurityGroup), arg0)
}

// DeleteSecurityGroup mocks base method
func (m *MockClient) DeleteSecurityGroup(arg0 *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityGroup", arg0)
	ret0, _ := ret[0].(*ec2.DeleteSecurityGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecurityGroup indicates an expected call of DeleteSecurityGroup
func (mr *MockClientMockRecorder) DeleteSecurityGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityGroup", reflect.TypeOf((*MockClient)(nil).DeleteSecurityGroup), arg0)
}

// DescribeSecurityGroups mocks base method
func (m *MockClient) DescribeSecurityGroups(arg0 *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSecurityGroups", arg0)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityGroups indicates an expected call of DescribeSecurityGroups
func (mr *MockClientMockRecorder) DescribeSecurityGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroups", reflect.TypeOf((*MockClient)(nil).DescribeSecurityGroups), arg0)
}

// RevokeSecurityGroupIngress mocks base method
func (m *MockClient) RevokeSecurityGroupIngress(arg0 *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecurityGroupIngress", arg0)
	ret0, _ := ret[0].(*ec2.RevokeSecurityGroupIngressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeSecurityGroupIngress indicates an expected call of RevokeSecurityGroupIngress
func (mr *MockClientMockRecorder) RevokeSecurityGroupIngress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecurityGroupIngress", reflect.TypeOf((*MockClient)(nil).RevokeSecurityGroupIngress), arg0)
}

// DescribeSubnets mocks base method
func (m *MockClient) DescribeSubnets(arg0 *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSubnets", arg0)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnets indicates an expected call of DescribeSubnets
func (mr *MockClientMockRecorder) DescribeSubnets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnets", reflect.TypeOf((*MockClient)(nil).DescribeSubnets), arg0)
}

// CreateTags mocks base method
func (m *MockClient) CreateTags(arg0 *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTags", arg0)
	ret0, _ := ret[0].(*ec2.CreateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTags indicates an expected call of CreateTags
func (mr *MockClientMockRecorder) CreateTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockClient)(nil).CreateTags), arg0)
}

// SubnetNameToSubnetIDLookup mocks base method
func (m *MockClient) SubnetNameToSubnetIDLookup(arg0 []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubnetNameToSubnetIDLookup", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubnetNameToSubnetIDLookup indicates an expected call of SubnetNameToSubnetIDLookup
func (mr *MockClientMockRecorder) SubnetNameToSubnetIDLookup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubnetNameToSubnetIDLookup", reflect.TypeOf((*MockClient)(nil).SubnetNameToSubnetIDLookup), arg0)
}

// SubnetIDToVPCLookup mocks base method
func (m *MockClient) SubnetIDToVPCLookup(arg0 []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubnetIDToVPCLookup", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubnetIDToVPCLookup indicates an expected call of SubnetIDToVPCLookup
func (mr *MockClientMockRecorder) SubnetIDToVPCLookup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubnetIDToVPCLookup", reflect.TypeOf((*MockClient)(nil).SubnetIDToVPCLookup), arg0)
}

// ApplyTagsToResources mocks base method
func (m *MockClient) ApplyTagsToResources(arg0 []string, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyTagsToResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyTagsToResources indicates an expected call of ApplyTagsToResources
func (mr *MockClientMockRecorder) ApplyTagsToResources(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyTagsToResources", reflect.TypeOf((*MockClient)(nil).ApplyTagsToResources), arg0, arg1)
}

// setLoadBalancerSecurityGroup mocks base method
func (m *MockClient) setLoadBalancerSecurityGroup(arg0 string, arg1 *ec2.SecurityGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setLoadBalancerSecurityGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// setLoadBalancerSecurityGroup indicates an expected call of setLoadBalancerSecurityGroup
func (mr *MockClientMockRecorder) setLoadBalancerSecurityGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setLoadBalancerSecurityGroup", reflect.TypeOf((*MockClient)(nil).setLoadBalancerSecurityGroup), arg0, arg1)
}

// findSecurityGroupByName mocks base method
func (m *MockClient) findSecurityGroupByName(arg0 string) (*ec2.SecurityGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findSecurityGroupByName", arg0)
	ret0, _ := ret[0].(*ec2.SecurityGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// findSecurityGroupByName indicates an expected call of findSecurityGroupByName
func (mr *MockClientMockRecorder) findSecurityGroupByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findSecurityGroupByName", reflect.TypeOf((*MockClient)(nil).findSecurityGroupByName), arg0)
}

// findSecurityGroupByID mocks base method
func (m *MockClient) findSecurityGroupByID(arg0 string) (*ec2.SecurityGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findSecurityGroupByID", arg0)
	ret0, _ := ret[0].(*ec2.SecurityGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// findSecurityGroupByID indicates an expected call of findSecurityGroupByID
func (mr *MockClientMockRecorder) findSecurityGroupByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findSecurityGroupByID", reflect.TypeOf((*MockClient)(nil).findSecurityGroupByID), arg0)
}

// createSecurityGroup mocks base method
func (m *MockClient) createSecurityGroup(arg0, arg1 string, arg2 map[string]string) (*ec2.SecurityGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "createSecurityGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ec2.SecurityGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// createSecurityGroup indicates an expected call of createSecurityGroup
func (mr *MockClientMockRecorder) createSecurityGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createSecurityGroup", reflect.TypeOf((*MockClient)(nil).createSecurityGroup), arg0, arg1, arg2)
}

// removeIngressRulesFromSecurityGroup mocks base method
func (m *MockClient) removeIngressRulesFromSecurityGroup(arg0 *ec2.SecurityGroup, arg1 []*ec2.IpPermission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "removeIngressRulesFromSecurityGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// removeIngressRulesFromSecurityGroup indicates an expected call of removeIngressRulesFromSecurityGroup
func (mr *MockClientMockRecorder) removeIngressRulesFromSecurityGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "removeIngressRulesFromSecurityGroup", reflect.TypeOf((*MockClient)(nil).removeIngressRulesFromSecurityGroup), arg0, arg1)
}

// addIngressRulesToSecurityGroup mocks base method
func (m *MockClient) addIngressRulesToSecurityGroup(arg0 *ec2.SecurityGroup, arg1 []*ec2.IpPermission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "addIngressRulesToSecurityGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// addIngressRulesToSecurityGroup indicates an expected call of addIngressRulesToSecurityGroup
func (mr *MockClientMockRecorder) addIngressRulesToSecurityGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addIngressRulesToSecurityGroup", reflect.TypeOf((*MockClient)(nil).addIngressRulesToSecurityGroup), arg0, arg1)
}

// EnsureCIDRAccess mocks base method
func (m *MockClient) EnsureCIDRAccess(arg0, arg1, arg2 string, arg3 []string, arg4 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureCIDRAccess", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureCIDRAccess indicates an expected call of EnsureCIDRAccess
func (mr *MockClientMockRecorder) EnsureCIDRAccess(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureCIDRAccess", reflect.TypeOf((*MockClient)(nil).EnsureCIDRAccess), arg0, arg1, arg2, arg3, arg4)
}

// MapToELBTags mocks base method
func (m *MockClient) MapToELBTags(arg0 map[string]string) []*elb.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapToELBTags", arg0)
	ret0, _ := ret[0].([]*elb.Tag)
	return ret0
}

// MapToELBTags indicates an expected call of MapToELBTags
func (mr *MockClientMockRecorder) MapToELBTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapToELBTags", reflect.TypeOf((*MockClient)(nil).MapToELBTags), arg0)
}

// CreateClassicELB mocks base method
func (m *MockClient) CreateClassicELB(arg0 string, arg1 []string, arg2 int64, arg3 map[string]string) (*awsclient.AWSLoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClassicELB", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*awsclient.AWSLoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClassicELB indicates an expected call of CreateClassicELB
func (mr *MockClientMockRecorder) CreateClassicELB(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClassicELB", reflect.TypeOf((*MockClient)(nil).CreateClassicELB), arg0, arg1, arg2, arg3)
}

// RemoveLoadBalancerListeners mocks base method
func (m *MockClient) RemoveLoadBalancerListeners(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLoadBalancerListeners", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLoadBalancerListeners indicates an expected call of RemoveLoadBalancerListeners
func (mr *MockClientMockRecorder) RemoveLoadBalancerListeners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLoadBalancerListeners", reflect.TypeOf((*MockClient)(nil).RemoveLoadBalancerListeners), arg0)
}

// AddLoadBalancerListeners mocks base method
func (m *MockClient) AddLoadBalancerListeners(arg0 string, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLoadBalancerListeners", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLoadBalancerListeners indicates an expected call of AddLoadBalancerListeners
func (mr *MockClientMockRecorder) AddLoadBalancerListeners(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLoadBalancerListeners", reflect.TypeOf((*MockClient)(nil).AddLoadBalancerListeners), arg0, arg1)
}

// removeListenersFromELB mocks base method
func (m *MockClient) removeListenersFromELB(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "removeListenersFromELB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// removeListenersFromELB indicates an expected call of removeListenersFromELB
func (mr *MockClientMockRecorder) removeListenersFromELB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "removeListenersFromELB", reflect.TypeOf((*MockClient)(nil).removeListenersFromELB), arg0)
}

// addListenersToELB mocks base method
func (m *MockClient) addListenersToELB(arg0 string, arg1 []*elb.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "addListenersToELB", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// addListenersToELB indicates an expected call of addListenersToELB
func (mr *MockClientMockRecorder) addListenersToELB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addListenersToELB", reflect.TypeOf((*MockClient)(nil).addListenersToELB), arg0, arg1)
}

// AddLoadBalancerInstances mocks base method
func (m *MockClient) AddLoadBalancerInstances(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLoadBalancerInstances", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLoadBalancerInstances indicates an expected call of AddLoadBalancerInstances
func (mr *MockClientMockRecorder) AddLoadBalancerInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLoadBalancerInstances", reflect.TypeOf((*MockClient)(nil).AddLoadBalancerInstances), arg0, arg1)
}

// RemoveInstancesFromLoadBalancer mocks base method
func (m *MockClient) RemoveInstancesFromLoadBalancer(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveInstancesFromLoadBalancer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveInstancesFromLoadBalancer indicates an expected call of RemoveInstancesFromLoadBalancer
func (mr *MockClientMockRecorder) RemoveInstancesFromLoadBalancer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveInstancesFromLoadBalancer", reflect.TypeOf((*MockClient)(nil).RemoveInstancesFromLoadBalancer), arg0, arg1)
}

// DoesELBExist mocks base method
func (m *MockClient) DoesELBExist(arg0 string) (bool, *awsclient.AWSLoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesELBExist", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*awsclient.AWSLoadBalancer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DoesELBExist indicates an expected call of DoesELBExist
func (mr *MockClientMockRecorder) DoesELBExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesELBExist", reflect.TypeOf((*MockClient)(nil).DoesELBExist), arg0)
}

// ListAllNLBs mocks base method
func (m *MockClient) ListAllNLBs() ([]awsclient.LoadBalancerV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllNLBs")
	ret0, _ := ret[0].([]awsclient.LoadBalancerV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllNLBs indicates an expected call of ListAllNLBs
func (mr *MockClientMockRecorder) ListAllNLBs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllNLBs", reflect.TypeOf((*MockClient)(nil).ListAllNLBs))
}

// DeleteExternalLoadBalancer mocks base method
func (m *MockClient) DeleteExternalLoadBalancer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalLoadBalancer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalLoadBalancer indicates an expected call of DeleteExternalLoadBalancer
func (mr *MockClientMockRecorder) DeleteExternalLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalLoadBalancer", reflect.TypeOf((*MockClient)(nil).DeleteExternalLoadBalancer), arg0)
}

// CreateNetworkLoadBalancer mocks base method
func (m *MockClient) CreateNetworkLoadBalancer(arg0, arg1, arg2 string) ([]awsclient.LoadBalancerV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetworkLoadBalancer", arg0, arg1, arg2)
	ret0, _ := ret[0].([]awsclient.LoadBalancerV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNetworkLoadBalancer indicates an expected call of CreateNetworkLoadBalancer
func (mr *MockClientMockRecorder) CreateNetworkLoadBalancer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetworkLoadBalancer", reflect.TypeOf((*MockClient)(nil).CreateNetworkLoadBalancer), arg0, arg1, arg2)
}

// CreateListenerForNLB mocks base method
func (m *MockClient) CreateListenerForNLB(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateListenerForNLB", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateListenerForNLB indicates an expected call of CreateListenerForNLB
func (mr *MockClientMockRecorder) CreateListenerForNLB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateListenerForNLB", reflect.TypeOf((*MockClient)(nil).CreateListenerForNLB), arg0, arg1)
}

// GetTargetGroupArn mocks base method
func (m *MockClient) GetTargetGroupArn(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetGroupArn", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetGroupArn indicates an expected call of GetTargetGroupArn
func (mr *MockClientMockRecorder) GetTargetGroupArn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetGroupArn", reflect.TypeOf((*MockClient)(nil).GetTargetGroupArn), arg0)
}

// addHealthCheck mocks base method
func (m *MockClient) addHealthCheck(arg0, arg1, arg2 string, arg3 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "addHealthCheck", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// addHealthCheck indicates an expected call of addHealthCheck
func (mr *MockClientMockRecorder) addHealthCheck(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addHealthCheck", reflect.TypeOf((*MockClient)(nil).addHealthCheck), arg0, arg1, arg2, arg3)
}

// UpsertARecord mocks base method
func (m *MockClient) UpsertARecord(arg0, arg1, arg2, arg3, arg4 string, arg5 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertARecord", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertARecord indicates an expected call of UpsertARecord
func (mr *MockClientMockRecorder) UpsertARecord(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertARecord", reflect.TypeOf((*MockClient)(nil).UpsertARecord), arg0, arg1, arg2, arg3, arg4, arg5)
}
