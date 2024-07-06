// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotsitewiseiface provides an interface to enable mocking the AWS IoT SiteWise service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotsitewiseiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotsitewise"
)

// IoTSiteWiseAPI provides an interface to enable mocking the
// iotsitewise.IoTSiteWise service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS IoT SiteWise.
//	func myFunc(svc iotsitewiseiface.IoTSiteWiseAPI) bool {
//	    // Make svc.AssociateAssets request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := iotsitewise.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockIoTSiteWiseClient struct {
//	    iotsitewiseiface.IoTSiteWiseAPI
//	}
//	func (m *mockIoTSiteWiseClient) AssociateAssets(input *iotsitewise.AssociateAssetsInput) (*iotsitewise.AssociateAssetsOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockIoTSiteWiseClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IoTSiteWiseAPI interface {
	AssociateAssets(*iotsitewise.AssociateAssetsInput) (*iotsitewise.AssociateAssetsOutput, error)
	AssociateAssetsWithContext(aws.Context, *iotsitewise.AssociateAssetsInput, ...request.Option) (*iotsitewise.AssociateAssetsOutput, error)
	AssociateAssetsRequest(*iotsitewise.AssociateAssetsInput) (*request.Request, *iotsitewise.AssociateAssetsOutput)

	AssociateTimeSeriesToAssetProperty(*iotsitewise.AssociateTimeSeriesToAssetPropertyInput) (*iotsitewise.AssociateTimeSeriesToAssetPropertyOutput, error)
	AssociateTimeSeriesToAssetPropertyWithContext(aws.Context, *iotsitewise.AssociateTimeSeriesToAssetPropertyInput, ...request.Option) (*iotsitewise.AssociateTimeSeriesToAssetPropertyOutput, error)
	AssociateTimeSeriesToAssetPropertyRequest(*iotsitewise.AssociateTimeSeriesToAssetPropertyInput) (*request.Request, *iotsitewise.AssociateTimeSeriesToAssetPropertyOutput)

	BatchAssociateProjectAssets(*iotsitewise.BatchAssociateProjectAssetsInput) (*iotsitewise.BatchAssociateProjectAssetsOutput, error)
	BatchAssociateProjectAssetsWithContext(aws.Context, *iotsitewise.BatchAssociateProjectAssetsInput, ...request.Option) (*iotsitewise.BatchAssociateProjectAssetsOutput, error)
	BatchAssociateProjectAssetsRequest(*iotsitewise.BatchAssociateProjectAssetsInput) (*request.Request, *iotsitewise.BatchAssociateProjectAssetsOutput)

	BatchDisassociateProjectAssets(*iotsitewise.BatchDisassociateProjectAssetsInput) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error)
	BatchDisassociateProjectAssetsWithContext(aws.Context, *iotsitewise.BatchDisassociateProjectAssetsInput, ...request.Option) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error)
	BatchDisassociateProjectAssetsRequest(*iotsitewise.BatchDisassociateProjectAssetsInput) (*request.Request, *iotsitewise.BatchDisassociateProjectAssetsOutput)

	BatchGetAssetPropertyAggregates(*iotsitewise.BatchGetAssetPropertyAggregatesInput) (*iotsitewise.BatchGetAssetPropertyAggregatesOutput, error)
	BatchGetAssetPropertyAggregatesWithContext(aws.Context, *iotsitewise.BatchGetAssetPropertyAggregatesInput, ...request.Option) (*iotsitewise.BatchGetAssetPropertyAggregatesOutput, error)
	BatchGetAssetPropertyAggregatesRequest(*iotsitewise.BatchGetAssetPropertyAggregatesInput) (*request.Request, *iotsitewise.BatchGetAssetPropertyAggregatesOutput)

	BatchGetAssetPropertyAggregatesPages(*iotsitewise.BatchGetAssetPropertyAggregatesInput, func(*iotsitewise.BatchGetAssetPropertyAggregatesOutput, bool) bool) error
	BatchGetAssetPropertyAggregatesPagesWithContext(aws.Context, *iotsitewise.BatchGetAssetPropertyAggregatesInput, func(*iotsitewise.BatchGetAssetPropertyAggregatesOutput, bool) bool, ...request.Option) error

	BatchGetAssetPropertyValue(*iotsitewise.BatchGetAssetPropertyValueInput) (*iotsitewise.BatchGetAssetPropertyValueOutput, error)
	BatchGetAssetPropertyValueWithContext(aws.Context, *iotsitewise.BatchGetAssetPropertyValueInput, ...request.Option) (*iotsitewise.BatchGetAssetPropertyValueOutput, error)
	BatchGetAssetPropertyValueRequest(*iotsitewise.BatchGetAssetPropertyValueInput) (*request.Request, *iotsitewise.BatchGetAssetPropertyValueOutput)

	BatchGetAssetPropertyValuePages(*iotsitewise.BatchGetAssetPropertyValueInput, func(*iotsitewise.BatchGetAssetPropertyValueOutput, bool) bool) error
	BatchGetAssetPropertyValuePagesWithContext(aws.Context, *iotsitewise.BatchGetAssetPropertyValueInput, func(*iotsitewise.BatchGetAssetPropertyValueOutput, bool) bool, ...request.Option) error

	BatchGetAssetPropertyValueHistory(*iotsitewise.BatchGetAssetPropertyValueHistoryInput) (*iotsitewise.BatchGetAssetPropertyValueHistoryOutput, error)
	BatchGetAssetPropertyValueHistoryWithContext(aws.Context, *iotsitewise.BatchGetAssetPropertyValueHistoryInput, ...request.Option) (*iotsitewise.BatchGetAssetPropertyValueHistoryOutput, error)
	BatchGetAssetPropertyValueHistoryRequest(*iotsitewise.BatchGetAssetPropertyValueHistoryInput) (*request.Request, *iotsitewise.BatchGetAssetPropertyValueHistoryOutput)

	BatchGetAssetPropertyValueHistoryPages(*iotsitewise.BatchGetAssetPropertyValueHistoryInput, func(*iotsitewise.BatchGetAssetPropertyValueHistoryOutput, bool) bool) error
	BatchGetAssetPropertyValueHistoryPagesWithContext(aws.Context, *iotsitewise.BatchGetAssetPropertyValueHistoryInput, func(*iotsitewise.BatchGetAssetPropertyValueHistoryOutput, bool) bool, ...request.Option) error

	BatchPutAssetPropertyValue(*iotsitewise.BatchPutAssetPropertyValueInput) (*iotsitewise.BatchPutAssetPropertyValueOutput, error)
	BatchPutAssetPropertyValueWithContext(aws.Context, *iotsitewise.BatchPutAssetPropertyValueInput, ...request.Option) (*iotsitewise.BatchPutAssetPropertyValueOutput, error)
	BatchPutAssetPropertyValueRequest(*iotsitewise.BatchPutAssetPropertyValueInput) (*request.Request, *iotsitewise.BatchPutAssetPropertyValueOutput)

	CreateAccessPolicy(*iotsitewise.CreateAccessPolicyInput) (*iotsitewise.CreateAccessPolicyOutput, error)
	CreateAccessPolicyWithContext(aws.Context, *iotsitewise.CreateAccessPolicyInput, ...request.Option) (*iotsitewise.CreateAccessPolicyOutput, error)
	CreateAccessPolicyRequest(*iotsitewise.CreateAccessPolicyInput) (*request.Request, *iotsitewise.CreateAccessPolicyOutput)

	CreateAsset(*iotsitewise.CreateAssetInput) (*iotsitewise.CreateAssetOutput, error)
	CreateAssetWithContext(aws.Context, *iotsitewise.CreateAssetInput, ...request.Option) (*iotsitewise.CreateAssetOutput, error)
	CreateAssetRequest(*iotsitewise.CreateAssetInput) (*request.Request, *iotsitewise.CreateAssetOutput)

	CreateAssetModel(*iotsitewise.CreateAssetModelInput) (*iotsitewise.CreateAssetModelOutput, error)
	CreateAssetModelWithContext(aws.Context, *iotsitewise.CreateAssetModelInput, ...request.Option) (*iotsitewise.CreateAssetModelOutput, error)
	CreateAssetModelRequest(*iotsitewise.CreateAssetModelInput) (*request.Request, *iotsitewise.CreateAssetModelOutput)

	CreateBulkImportJob(*iotsitewise.CreateBulkImportJobInput) (*iotsitewise.CreateBulkImportJobOutput, error)
	CreateBulkImportJobWithContext(aws.Context, *iotsitewise.CreateBulkImportJobInput, ...request.Option) (*iotsitewise.CreateBulkImportJobOutput, error)
	CreateBulkImportJobRequest(*iotsitewise.CreateBulkImportJobInput) (*request.Request, *iotsitewise.CreateBulkImportJobOutput)

	CreateDashboard(*iotsitewise.CreateDashboardInput) (*iotsitewise.CreateDashboardOutput, error)
	CreateDashboardWithContext(aws.Context, *iotsitewise.CreateDashboardInput, ...request.Option) (*iotsitewise.CreateDashboardOutput, error)
	CreateDashboardRequest(*iotsitewise.CreateDashboardInput) (*request.Request, *iotsitewise.CreateDashboardOutput)

	CreateGateway(*iotsitewise.CreateGatewayInput) (*iotsitewise.CreateGatewayOutput, error)
	CreateGatewayWithContext(aws.Context, *iotsitewise.CreateGatewayInput, ...request.Option) (*iotsitewise.CreateGatewayOutput, error)
	CreateGatewayRequest(*iotsitewise.CreateGatewayInput) (*request.Request, *iotsitewise.CreateGatewayOutput)

	CreatePortal(*iotsitewise.CreatePortalInput) (*iotsitewise.CreatePortalOutput, error)
	CreatePortalWithContext(aws.Context, *iotsitewise.CreatePortalInput, ...request.Option) (*iotsitewise.CreatePortalOutput, error)
	CreatePortalRequest(*iotsitewise.CreatePortalInput) (*request.Request, *iotsitewise.CreatePortalOutput)

	CreateProject(*iotsitewise.CreateProjectInput) (*iotsitewise.CreateProjectOutput, error)
	CreateProjectWithContext(aws.Context, *iotsitewise.CreateProjectInput, ...request.Option) (*iotsitewise.CreateProjectOutput, error)
	CreateProjectRequest(*iotsitewise.CreateProjectInput) (*request.Request, *iotsitewise.CreateProjectOutput)

	DeleteAccessPolicy(*iotsitewise.DeleteAccessPolicyInput) (*iotsitewise.DeleteAccessPolicyOutput, error)
	DeleteAccessPolicyWithContext(aws.Context, *iotsitewise.DeleteAccessPolicyInput, ...request.Option) (*iotsitewise.DeleteAccessPolicyOutput, error)
	DeleteAccessPolicyRequest(*iotsitewise.DeleteAccessPolicyInput) (*request.Request, *iotsitewise.DeleteAccessPolicyOutput)

	DeleteAsset(*iotsitewise.DeleteAssetInput) (*iotsitewise.DeleteAssetOutput, error)
	DeleteAssetWithContext(aws.Context, *iotsitewise.DeleteAssetInput, ...request.Option) (*iotsitewise.DeleteAssetOutput, error)
	DeleteAssetRequest(*iotsitewise.DeleteAssetInput) (*request.Request, *iotsitewise.DeleteAssetOutput)

	DeleteAssetModel(*iotsitewise.DeleteAssetModelInput) (*iotsitewise.DeleteAssetModelOutput, error)
	DeleteAssetModelWithContext(aws.Context, *iotsitewise.DeleteAssetModelInput, ...request.Option) (*iotsitewise.DeleteAssetModelOutput, error)
	DeleteAssetModelRequest(*iotsitewise.DeleteAssetModelInput) (*request.Request, *iotsitewise.DeleteAssetModelOutput)

	DeleteDashboard(*iotsitewise.DeleteDashboardInput) (*iotsitewise.DeleteDashboardOutput, error)
	DeleteDashboardWithContext(aws.Context, *iotsitewise.DeleteDashboardInput, ...request.Option) (*iotsitewise.DeleteDashboardOutput, error)
	DeleteDashboardRequest(*iotsitewise.DeleteDashboardInput) (*request.Request, *iotsitewise.DeleteDashboardOutput)

	DeleteGateway(*iotsitewise.DeleteGatewayInput) (*iotsitewise.DeleteGatewayOutput, error)
	DeleteGatewayWithContext(aws.Context, *iotsitewise.DeleteGatewayInput, ...request.Option) (*iotsitewise.DeleteGatewayOutput, error)
	DeleteGatewayRequest(*iotsitewise.DeleteGatewayInput) (*request.Request, *iotsitewise.DeleteGatewayOutput)

	DeletePortal(*iotsitewise.DeletePortalInput) (*iotsitewise.DeletePortalOutput, error)
	DeletePortalWithContext(aws.Context, *iotsitewise.DeletePortalInput, ...request.Option) (*iotsitewise.DeletePortalOutput, error)
	DeletePortalRequest(*iotsitewise.DeletePortalInput) (*request.Request, *iotsitewise.DeletePortalOutput)

	DeleteProject(*iotsitewise.DeleteProjectInput) (*iotsitewise.DeleteProjectOutput, error)
	DeleteProjectWithContext(aws.Context, *iotsitewise.DeleteProjectInput, ...request.Option) (*iotsitewise.DeleteProjectOutput, error)
	DeleteProjectRequest(*iotsitewise.DeleteProjectInput) (*request.Request, *iotsitewise.DeleteProjectOutput)

	DeleteTimeSeries(*iotsitewise.DeleteTimeSeriesInput) (*iotsitewise.DeleteTimeSeriesOutput, error)
	DeleteTimeSeriesWithContext(aws.Context, *iotsitewise.DeleteTimeSeriesInput, ...request.Option) (*iotsitewise.DeleteTimeSeriesOutput, error)
	DeleteTimeSeriesRequest(*iotsitewise.DeleteTimeSeriesInput) (*request.Request, *iotsitewise.DeleteTimeSeriesOutput)

	DescribeAccessPolicy(*iotsitewise.DescribeAccessPolicyInput) (*iotsitewise.DescribeAccessPolicyOutput, error)
	DescribeAccessPolicyWithContext(aws.Context, *iotsitewise.DescribeAccessPolicyInput, ...request.Option) (*iotsitewise.DescribeAccessPolicyOutput, error)
	DescribeAccessPolicyRequest(*iotsitewise.DescribeAccessPolicyInput) (*request.Request, *iotsitewise.DescribeAccessPolicyOutput)

	DescribeAsset(*iotsitewise.DescribeAssetInput) (*iotsitewise.DescribeAssetOutput, error)
	DescribeAssetWithContext(aws.Context, *iotsitewise.DescribeAssetInput, ...request.Option) (*iotsitewise.DescribeAssetOutput, error)
	DescribeAssetRequest(*iotsitewise.DescribeAssetInput) (*request.Request, *iotsitewise.DescribeAssetOutput)

	DescribeAssetModel(*iotsitewise.DescribeAssetModelInput) (*iotsitewise.DescribeAssetModelOutput, error)
	DescribeAssetModelWithContext(aws.Context, *iotsitewise.DescribeAssetModelInput, ...request.Option) (*iotsitewise.DescribeAssetModelOutput, error)
	DescribeAssetModelRequest(*iotsitewise.DescribeAssetModelInput) (*request.Request, *iotsitewise.DescribeAssetModelOutput)

	DescribeAssetProperty(*iotsitewise.DescribeAssetPropertyInput) (*iotsitewise.DescribeAssetPropertyOutput, error)
	DescribeAssetPropertyWithContext(aws.Context, *iotsitewise.DescribeAssetPropertyInput, ...request.Option) (*iotsitewise.DescribeAssetPropertyOutput, error)
	DescribeAssetPropertyRequest(*iotsitewise.DescribeAssetPropertyInput) (*request.Request, *iotsitewise.DescribeAssetPropertyOutput)

	DescribeBulkImportJob(*iotsitewise.DescribeBulkImportJobInput) (*iotsitewise.DescribeBulkImportJobOutput, error)
	DescribeBulkImportJobWithContext(aws.Context, *iotsitewise.DescribeBulkImportJobInput, ...request.Option) (*iotsitewise.DescribeBulkImportJobOutput, error)
	DescribeBulkImportJobRequest(*iotsitewise.DescribeBulkImportJobInput) (*request.Request, *iotsitewise.DescribeBulkImportJobOutput)

	DescribeDashboard(*iotsitewise.DescribeDashboardInput) (*iotsitewise.DescribeDashboardOutput, error)
	DescribeDashboardWithContext(aws.Context, *iotsitewise.DescribeDashboardInput, ...request.Option) (*iotsitewise.DescribeDashboardOutput, error)
	DescribeDashboardRequest(*iotsitewise.DescribeDashboardInput) (*request.Request, *iotsitewise.DescribeDashboardOutput)

	DescribeDefaultEncryptionConfiguration(*iotsitewise.DescribeDefaultEncryptionConfigurationInput) (*iotsitewise.DescribeDefaultEncryptionConfigurationOutput, error)
	DescribeDefaultEncryptionConfigurationWithContext(aws.Context, *iotsitewise.DescribeDefaultEncryptionConfigurationInput, ...request.Option) (*iotsitewise.DescribeDefaultEncryptionConfigurationOutput, error)
	DescribeDefaultEncryptionConfigurationRequest(*iotsitewise.DescribeDefaultEncryptionConfigurationInput) (*request.Request, *iotsitewise.DescribeDefaultEncryptionConfigurationOutput)

	DescribeGateway(*iotsitewise.DescribeGatewayInput) (*iotsitewise.DescribeGatewayOutput, error)
	DescribeGatewayWithContext(aws.Context, *iotsitewise.DescribeGatewayInput, ...request.Option) (*iotsitewise.DescribeGatewayOutput, error)
	DescribeGatewayRequest(*iotsitewise.DescribeGatewayInput) (*request.Request, *iotsitewise.DescribeGatewayOutput)

	DescribeGatewayCapabilityConfiguration(*iotsitewise.DescribeGatewayCapabilityConfigurationInput) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error)
	DescribeGatewayCapabilityConfigurationWithContext(aws.Context, *iotsitewise.DescribeGatewayCapabilityConfigurationInput, ...request.Option) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error)
	DescribeGatewayCapabilityConfigurationRequest(*iotsitewise.DescribeGatewayCapabilityConfigurationInput) (*request.Request, *iotsitewise.DescribeGatewayCapabilityConfigurationOutput)

	DescribeLoggingOptions(*iotsitewise.DescribeLoggingOptionsInput) (*iotsitewise.DescribeLoggingOptionsOutput, error)
	DescribeLoggingOptionsWithContext(aws.Context, *iotsitewise.DescribeLoggingOptionsInput, ...request.Option) (*iotsitewise.DescribeLoggingOptionsOutput, error)
	DescribeLoggingOptionsRequest(*iotsitewise.DescribeLoggingOptionsInput) (*request.Request, *iotsitewise.DescribeLoggingOptionsOutput)

	DescribePortal(*iotsitewise.DescribePortalInput) (*iotsitewise.DescribePortalOutput, error)
	DescribePortalWithContext(aws.Context, *iotsitewise.DescribePortalInput, ...request.Option) (*iotsitewise.DescribePortalOutput, error)
	DescribePortalRequest(*iotsitewise.DescribePortalInput) (*request.Request, *iotsitewise.DescribePortalOutput)

	DescribeProject(*iotsitewise.DescribeProjectInput) (*iotsitewise.DescribeProjectOutput, error)
	DescribeProjectWithContext(aws.Context, *iotsitewise.DescribeProjectInput, ...request.Option) (*iotsitewise.DescribeProjectOutput, error)
	DescribeProjectRequest(*iotsitewise.DescribeProjectInput) (*request.Request, *iotsitewise.DescribeProjectOutput)

	DescribeStorageConfiguration(*iotsitewise.DescribeStorageConfigurationInput) (*iotsitewise.DescribeStorageConfigurationOutput, error)
	DescribeStorageConfigurationWithContext(aws.Context, *iotsitewise.DescribeStorageConfigurationInput, ...request.Option) (*iotsitewise.DescribeStorageConfigurationOutput, error)
	DescribeStorageConfigurationRequest(*iotsitewise.DescribeStorageConfigurationInput) (*request.Request, *iotsitewise.DescribeStorageConfigurationOutput)

	DescribeTimeSeries(*iotsitewise.DescribeTimeSeriesInput) (*iotsitewise.DescribeTimeSeriesOutput, error)
	DescribeTimeSeriesWithContext(aws.Context, *iotsitewise.DescribeTimeSeriesInput, ...request.Option) (*iotsitewise.DescribeTimeSeriesOutput, error)
	DescribeTimeSeriesRequest(*iotsitewise.DescribeTimeSeriesInput) (*request.Request, *iotsitewise.DescribeTimeSeriesOutput)

	DisassociateAssets(*iotsitewise.DisassociateAssetsInput) (*iotsitewise.DisassociateAssetsOutput, error)
	DisassociateAssetsWithContext(aws.Context, *iotsitewise.DisassociateAssetsInput, ...request.Option) (*iotsitewise.DisassociateAssetsOutput, error)
	DisassociateAssetsRequest(*iotsitewise.DisassociateAssetsInput) (*request.Request, *iotsitewise.DisassociateAssetsOutput)

	DisassociateTimeSeriesFromAssetProperty(*iotsitewise.DisassociateTimeSeriesFromAssetPropertyInput) (*iotsitewise.DisassociateTimeSeriesFromAssetPropertyOutput, error)
	DisassociateTimeSeriesFromAssetPropertyWithContext(aws.Context, *iotsitewise.DisassociateTimeSeriesFromAssetPropertyInput, ...request.Option) (*iotsitewise.DisassociateTimeSeriesFromAssetPropertyOutput, error)
	DisassociateTimeSeriesFromAssetPropertyRequest(*iotsitewise.DisassociateTimeSeriesFromAssetPropertyInput) (*request.Request, *iotsitewise.DisassociateTimeSeriesFromAssetPropertyOutput)

	GetAssetPropertyAggregates(*iotsitewise.GetAssetPropertyAggregatesInput) (*iotsitewise.GetAssetPropertyAggregatesOutput, error)
	GetAssetPropertyAggregatesWithContext(aws.Context, *iotsitewise.GetAssetPropertyAggregatesInput, ...request.Option) (*iotsitewise.GetAssetPropertyAggregatesOutput, error)
	GetAssetPropertyAggregatesRequest(*iotsitewise.GetAssetPropertyAggregatesInput) (*request.Request, *iotsitewise.GetAssetPropertyAggregatesOutput)

	GetAssetPropertyAggregatesPages(*iotsitewise.GetAssetPropertyAggregatesInput, func(*iotsitewise.GetAssetPropertyAggregatesOutput, bool) bool) error
	GetAssetPropertyAggregatesPagesWithContext(aws.Context, *iotsitewise.GetAssetPropertyAggregatesInput, func(*iotsitewise.GetAssetPropertyAggregatesOutput, bool) bool, ...request.Option) error

	GetAssetPropertyValue(*iotsitewise.GetAssetPropertyValueInput) (*iotsitewise.GetAssetPropertyValueOutput, error)
	GetAssetPropertyValueWithContext(aws.Context, *iotsitewise.GetAssetPropertyValueInput, ...request.Option) (*iotsitewise.GetAssetPropertyValueOutput, error)
	GetAssetPropertyValueRequest(*iotsitewise.GetAssetPropertyValueInput) (*request.Request, *iotsitewise.GetAssetPropertyValueOutput)

	GetAssetPropertyValueHistory(*iotsitewise.GetAssetPropertyValueHistoryInput) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error)
	GetAssetPropertyValueHistoryWithContext(aws.Context, *iotsitewise.GetAssetPropertyValueHistoryInput, ...request.Option) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error)
	GetAssetPropertyValueHistoryRequest(*iotsitewise.GetAssetPropertyValueHistoryInput) (*request.Request, *iotsitewise.GetAssetPropertyValueHistoryOutput)

	GetAssetPropertyValueHistoryPages(*iotsitewise.GetAssetPropertyValueHistoryInput, func(*iotsitewise.GetAssetPropertyValueHistoryOutput, bool) bool) error
	GetAssetPropertyValueHistoryPagesWithContext(aws.Context, *iotsitewise.GetAssetPropertyValueHistoryInput, func(*iotsitewise.GetAssetPropertyValueHistoryOutput, bool) bool, ...request.Option) error

	GetInterpolatedAssetPropertyValues(*iotsitewise.GetInterpolatedAssetPropertyValuesInput) (*iotsitewise.GetInterpolatedAssetPropertyValuesOutput, error)
	GetInterpolatedAssetPropertyValuesWithContext(aws.Context, *iotsitewise.GetInterpolatedAssetPropertyValuesInput, ...request.Option) (*iotsitewise.GetInterpolatedAssetPropertyValuesOutput, error)
	GetInterpolatedAssetPropertyValuesRequest(*iotsitewise.GetInterpolatedAssetPropertyValuesInput) (*request.Request, *iotsitewise.GetInterpolatedAssetPropertyValuesOutput)

	GetInterpolatedAssetPropertyValuesPages(*iotsitewise.GetInterpolatedAssetPropertyValuesInput, func(*iotsitewise.GetInterpolatedAssetPropertyValuesOutput, bool) bool) error
	GetInterpolatedAssetPropertyValuesPagesWithContext(aws.Context, *iotsitewise.GetInterpolatedAssetPropertyValuesInput, func(*iotsitewise.GetInterpolatedAssetPropertyValuesOutput, bool) bool, ...request.Option) error

	ListAccessPolicies(*iotsitewise.ListAccessPoliciesInput) (*iotsitewise.ListAccessPoliciesOutput, error)
	ListAccessPoliciesWithContext(aws.Context, *iotsitewise.ListAccessPoliciesInput, ...request.Option) (*iotsitewise.ListAccessPoliciesOutput, error)
	ListAccessPoliciesRequest(*iotsitewise.ListAccessPoliciesInput) (*request.Request, *iotsitewise.ListAccessPoliciesOutput)

	ListAccessPoliciesPages(*iotsitewise.ListAccessPoliciesInput, func(*iotsitewise.ListAccessPoliciesOutput, bool) bool) error
	ListAccessPoliciesPagesWithContext(aws.Context, *iotsitewise.ListAccessPoliciesInput, func(*iotsitewise.ListAccessPoliciesOutput, bool) bool, ...request.Option) error

	ListAssetModels(*iotsitewise.ListAssetModelsInput) (*iotsitewise.ListAssetModelsOutput, error)
	ListAssetModelsWithContext(aws.Context, *iotsitewise.ListAssetModelsInput, ...request.Option) (*iotsitewise.ListAssetModelsOutput, error)
	ListAssetModelsRequest(*iotsitewise.ListAssetModelsInput) (*request.Request, *iotsitewise.ListAssetModelsOutput)

	ListAssetModelsPages(*iotsitewise.ListAssetModelsInput, func(*iotsitewise.ListAssetModelsOutput, bool) bool) error
	ListAssetModelsPagesWithContext(aws.Context, *iotsitewise.ListAssetModelsInput, func(*iotsitewise.ListAssetModelsOutput, bool) bool, ...request.Option) error

	ListAssetRelationships(*iotsitewise.ListAssetRelationshipsInput) (*iotsitewise.ListAssetRelationshipsOutput, error)
	ListAssetRelationshipsWithContext(aws.Context, *iotsitewise.ListAssetRelationshipsInput, ...request.Option) (*iotsitewise.ListAssetRelationshipsOutput, error)
	ListAssetRelationshipsRequest(*iotsitewise.ListAssetRelationshipsInput) (*request.Request, *iotsitewise.ListAssetRelationshipsOutput)

	ListAssetRelationshipsPages(*iotsitewise.ListAssetRelationshipsInput, func(*iotsitewise.ListAssetRelationshipsOutput, bool) bool) error
	ListAssetRelationshipsPagesWithContext(aws.Context, *iotsitewise.ListAssetRelationshipsInput, func(*iotsitewise.ListAssetRelationshipsOutput, bool) bool, ...request.Option) error

	ListAssets(*iotsitewise.ListAssetsInput) (*iotsitewise.ListAssetsOutput, error)
	ListAssetsWithContext(aws.Context, *iotsitewise.ListAssetsInput, ...request.Option) (*iotsitewise.ListAssetsOutput, error)
	ListAssetsRequest(*iotsitewise.ListAssetsInput) (*request.Request, *iotsitewise.ListAssetsOutput)

	ListAssetsPages(*iotsitewise.ListAssetsInput, func(*iotsitewise.ListAssetsOutput, bool) bool) error
	ListAssetsPagesWithContext(aws.Context, *iotsitewise.ListAssetsInput, func(*iotsitewise.ListAssetsOutput, bool) bool, ...request.Option) error

	ListAssociatedAssets(*iotsitewise.ListAssociatedAssetsInput) (*iotsitewise.ListAssociatedAssetsOutput, error)
	ListAssociatedAssetsWithContext(aws.Context, *iotsitewise.ListAssociatedAssetsInput, ...request.Option) (*iotsitewise.ListAssociatedAssetsOutput, error)
	ListAssociatedAssetsRequest(*iotsitewise.ListAssociatedAssetsInput) (*request.Request, *iotsitewise.ListAssociatedAssetsOutput)

	ListAssociatedAssetsPages(*iotsitewise.ListAssociatedAssetsInput, func(*iotsitewise.ListAssociatedAssetsOutput, bool) bool) error
	ListAssociatedAssetsPagesWithContext(aws.Context, *iotsitewise.ListAssociatedAssetsInput, func(*iotsitewise.ListAssociatedAssetsOutput, bool) bool, ...request.Option) error

	ListBulkImportJobs(*iotsitewise.ListBulkImportJobsInput) (*iotsitewise.ListBulkImportJobsOutput, error)
	ListBulkImportJobsWithContext(aws.Context, *iotsitewise.ListBulkImportJobsInput, ...request.Option) (*iotsitewise.ListBulkImportJobsOutput, error)
	ListBulkImportJobsRequest(*iotsitewise.ListBulkImportJobsInput) (*request.Request, *iotsitewise.ListBulkImportJobsOutput)

	ListBulkImportJobsPages(*iotsitewise.ListBulkImportJobsInput, func(*iotsitewise.ListBulkImportJobsOutput, bool) bool) error
	ListBulkImportJobsPagesWithContext(aws.Context, *iotsitewise.ListBulkImportJobsInput, func(*iotsitewise.ListBulkImportJobsOutput, bool) bool, ...request.Option) error

	ListDashboards(*iotsitewise.ListDashboardsInput) (*iotsitewise.ListDashboardsOutput, error)
	ListDashboardsWithContext(aws.Context, *iotsitewise.ListDashboardsInput, ...request.Option) (*iotsitewise.ListDashboardsOutput, error)
	ListDashboardsRequest(*iotsitewise.ListDashboardsInput) (*request.Request, *iotsitewise.ListDashboardsOutput)

	ListDashboardsPages(*iotsitewise.ListDashboardsInput, func(*iotsitewise.ListDashboardsOutput, bool) bool) error
	ListDashboardsPagesWithContext(aws.Context, *iotsitewise.ListDashboardsInput, func(*iotsitewise.ListDashboardsOutput, bool) bool, ...request.Option) error

	ListGateways(*iotsitewise.ListGatewaysInput) (*iotsitewise.ListGatewaysOutput, error)
	ListGatewaysWithContext(aws.Context, *iotsitewise.ListGatewaysInput, ...request.Option) (*iotsitewise.ListGatewaysOutput, error)
	ListGatewaysRequest(*iotsitewise.ListGatewaysInput) (*request.Request, *iotsitewise.ListGatewaysOutput)

	ListGatewaysPages(*iotsitewise.ListGatewaysInput, func(*iotsitewise.ListGatewaysOutput, bool) bool) error
	ListGatewaysPagesWithContext(aws.Context, *iotsitewise.ListGatewaysInput, func(*iotsitewise.ListGatewaysOutput, bool) bool, ...request.Option) error

	ListPortals(*iotsitewise.ListPortalsInput) (*iotsitewise.ListPortalsOutput, error)
	ListPortalsWithContext(aws.Context, *iotsitewise.ListPortalsInput, ...request.Option) (*iotsitewise.ListPortalsOutput, error)
	ListPortalsRequest(*iotsitewise.ListPortalsInput) (*request.Request, *iotsitewise.ListPortalsOutput)

	ListPortalsPages(*iotsitewise.ListPortalsInput, func(*iotsitewise.ListPortalsOutput, bool) bool) error
	ListPortalsPagesWithContext(aws.Context, *iotsitewise.ListPortalsInput, func(*iotsitewise.ListPortalsOutput, bool) bool, ...request.Option) error

	ListProjectAssets(*iotsitewise.ListProjectAssetsInput) (*iotsitewise.ListProjectAssetsOutput, error)
	ListProjectAssetsWithContext(aws.Context, *iotsitewise.ListProjectAssetsInput, ...request.Option) (*iotsitewise.ListProjectAssetsOutput, error)
	ListProjectAssetsRequest(*iotsitewise.ListProjectAssetsInput) (*request.Request, *iotsitewise.ListProjectAssetsOutput)

	ListProjectAssetsPages(*iotsitewise.ListProjectAssetsInput, func(*iotsitewise.ListProjectAssetsOutput, bool) bool) error
	ListProjectAssetsPagesWithContext(aws.Context, *iotsitewise.ListProjectAssetsInput, func(*iotsitewise.ListProjectAssetsOutput, bool) bool, ...request.Option) error

	ListProjects(*iotsitewise.ListProjectsInput) (*iotsitewise.ListProjectsOutput, error)
	ListProjectsWithContext(aws.Context, *iotsitewise.ListProjectsInput, ...request.Option) (*iotsitewise.ListProjectsOutput, error)
	ListProjectsRequest(*iotsitewise.ListProjectsInput) (*request.Request, *iotsitewise.ListProjectsOutput)

	ListProjectsPages(*iotsitewise.ListProjectsInput, func(*iotsitewise.ListProjectsOutput, bool) bool) error
	ListProjectsPagesWithContext(aws.Context, *iotsitewise.ListProjectsInput, func(*iotsitewise.ListProjectsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*iotsitewise.ListTagsForResourceInput) (*iotsitewise.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *iotsitewise.ListTagsForResourceInput, ...request.Option) (*iotsitewise.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*iotsitewise.ListTagsForResourceInput) (*request.Request, *iotsitewise.ListTagsForResourceOutput)

	ListTimeSeries(*iotsitewise.ListTimeSeriesInput) (*iotsitewise.ListTimeSeriesOutput, error)
	ListTimeSeriesWithContext(aws.Context, *iotsitewise.ListTimeSeriesInput, ...request.Option) (*iotsitewise.ListTimeSeriesOutput, error)
	ListTimeSeriesRequest(*iotsitewise.ListTimeSeriesInput) (*request.Request, *iotsitewise.ListTimeSeriesOutput)

	ListTimeSeriesPages(*iotsitewise.ListTimeSeriesInput, func(*iotsitewise.ListTimeSeriesOutput, bool) bool) error
	ListTimeSeriesPagesWithContext(aws.Context, *iotsitewise.ListTimeSeriesInput, func(*iotsitewise.ListTimeSeriesOutput, bool) bool, ...request.Option) error

	PutDefaultEncryptionConfiguration(*iotsitewise.PutDefaultEncryptionConfigurationInput) (*iotsitewise.PutDefaultEncryptionConfigurationOutput, error)
	PutDefaultEncryptionConfigurationWithContext(aws.Context, *iotsitewise.PutDefaultEncryptionConfigurationInput, ...request.Option) (*iotsitewise.PutDefaultEncryptionConfigurationOutput, error)
	PutDefaultEncryptionConfigurationRequest(*iotsitewise.PutDefaultEncryptionConfigurationInput) (*request.Request, *iotsitewise.PutDefaultEncryptionConfigurationOutput)

	PutLoggingOptions(*iotsitewise.PutLoggingOptionsInput) (*iotsitewise.PutLoggingOptionsOutput, error)
	PutLoggingOptionsWithContext(aws.Context, *iotsitewise.PutLoggingOptionsInput, ...request.Option) (*iotsitewise.PutLoggingOptionsOutput, error)
	PutLoggingOptionsRequest(*iotsitewise.PutLoggingOptionsInput) (*request.Request, *iotsitewise.PutLoggingOptionsOutput)

	PutStorageConfiguration(*iotsitewise.PutStorageConfigurationInput) (*iotsitewise.PutStorageConfigurationOutput, error)
	PutStorageConfigurationWithContext(aws.Context, *iotsitewise.PutStorageConfigurationInput, ...request.Option) (*iotsitewise.PutStorageConfigurationOutput, error)
	PutStorageConfigurationRequest(*iotsitewise.PutStorageConfigurationInput) (*request.Request, *iotsitewise.PutStorageConfigurationOutput)

	TagResource(*iotsitewise.TagResourceInput) (*iotsitewise.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *iotsitewise.TagResourceInput, ...request.Option) (*iotsitewise.TagResourceOutput, error)
	TagResourceRequest(*iotsitewise.TagResourceInput) (*request.Request, *iotsitewise.TagResourceOutput)

	UntagResource(*iotsitewise.UntagResourceInput) (*iotsitewise.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *iotsitewise.UntagResourceInput, ...request.Option) (*iotsitewise.UntagResourceOutput, error)
	UntagResourceRequest(*iotsitewise.UntagResourceInput) (*request.Request, *iotsitewise.UntagResourceOutput)

	UpdateAccessPolicy(*iotsitewise.UpdateAccessPolicyInput) (*iotsitewise.UpdateAccessPolicyOutput, error)
	UpdateAccessPolicyWithContext(aws.Context, *iotsitewise.UpdateAccessPolicyInput, ...request.Option) (*iotsitewise.UpdateAccessPolicyOutput, error)
	UpdateAccessPolicyRequest(*iotsitewise.UpdateAccessPolicyInput) (*request.Request, *iotsitewise.UpdateAccessPolicyOutput)

	UpdateAsset(*iotsitewise.UpdateAssetInput) (*iotsitewise.UpdateAssetOutput, error)
	UpdateAssetWithContext(aws.Context, *iotsitewise.UpdateAssetInput, ...request.Option) (*iotsitewise.UpdateAssetOutput, error)
	UpdateAssetRequest(*iotsitewise.UpdateAssetInput) (*request.Request, *iotsitewise.UpdateAssetOutput)

	UpdateAssetModel(*iotsitewise.UpdateAssetModelInput) (*iotsitewise.UpdateAssetModelOutput, error)
	UpdateAssetModelWithContext(aws.Context, *iotsitewise.UpdateAssetModelInput, ...request.Option) (*iotsitewise.UpdateAssetModelOutput, error)
	UpdateAssetModelRequest(*iotsitewise.UpdateAssetModelInput) (*request.Request, *iotsitewise.UpdateAssetModelOutput)

	UpdateAssetProperty(*iotsitewise.UpdateAssetPropertyInput) (*iotsitewise.UpdateAssetPropertyOutput, error)
	UpdateAssetPropertyWithContext(aws.Context, *iotsitewise.UpdateAssetPropertyInput, ...request.Option) (*iotsitewise.UpdateAssetPropertyOutput, error)
	UpdateAssetPropertyRequest(*iotsitewise.UpdateAssetPropertyInput) (*request.Request, *iotsitewise.UpdateAssetPropertyOutput)

	UpdateDashboard(*iotsitewise.UpdateDashboardInput) (*iotsitewise.UpdateDashboardOutput, error)
	UpdateDashboardWithContext(aws.Context, *iotsitewise.UpdateDashboardInput, ...request.Option) (*iotsitewise.UpdateDashboardOutput, error)
	UpdateDashboardRequest(*iotsitewise.UpdateDashboardInput) (*request.Request, *iotsitewise.UpdateDashboardOutput)

	UpdateGateway(*iotsitewise.UpdateGatewayInput) (*iotsitewise.UpdateGatewayOutput, error)
	UpdateGatewayWithContext(aws.Context, *iotsitewise.UpdateGatewayInput, ...request.Option) (*iotsitewise.UpdateGatewayOutput, error)
	UpdateGatewayRequest(*iotsitewise.UpdateGatewayInput) (*request.Request, *iotsitewise.UpdateGatewayOutput)

	UpdateGatewayCapabilityConfiguration(*iotsitewise.UpdateGatewayCapabilityConfigurationInput) (*iotsitewise.UpdateGatewayCapabilityConfigurationOutput, error)
	UpdateGatewayCapabilityConfigurationWithContext(aws.Context, *iotsitewise.UpdateGatewayCapabilityConfigurationInput, ...request.Option) (*iotsitewise.UpdateGatewayCapabilityConfigurationOutput, error)
	UpdateGatewayCapabilityConfigurationRequest(*iotsitewise.UpdateGatewayCapabilityConfigurationInput) (*request.Request, *iotsitewise.UpdateGatewayCapabilityConfigurationOutput)

	UpdatePortal(*iotsitewise.UpdatePortalInput) (*iotsitewise.UpdatePortalOutput, error)
	UpdatePortalWithContext(aws.Context, *iotsitewise.UpdatePortalInput, ...request.Option) (*iotsitewise.UpdatePortalOutput, error)
	UpdatePortalRequest(*iotsitewise.UpdatePortalInput) (*request.Request, *iotsitewise.UpdatePortalOutput)

	UpdateProject(*iotsitewise.UpdateProjectInput) (*iotsitewise.UpdateProjectOutput, error)
	UpdateProjectWithContext(aws.Context, *iotsitewise.UpdateProjectInput, ...request.Option) (*iotsitewise.UpdateProjectOutput, error)
	UpdateProjectRequest(*iotsitewise.UpdateProjectInput) (*request.Request, *iotsitewise.UpdateProjectOutput)

	WaitUntilAssetActive(*iotsitewise.DescribeAssetInput) error
	WaitUntilAssetActiveWithContext(aws.Context, *iotsitewise.DescribeAssetInput, ...request.WaiterOption) error

	WaitUntilAssetModelActive(*iotsitewise.DescribeAssetModelInput) error
	WaitUntilAssetModelActiveWithContext(aws.Context, *iotsitewise.DescribeAssetModelInput, ...request.WaiterOption) error

	WaitUntilAssetModelNotExists(*iotsitewise.DescribeAssetModelInput) error
	WaitUntilAssetModelNotExistsWithContext(aws.Context, *iotsitewise.DescribeAssetModelInput, ...request.WaiterOption) error

	WaitUntilAssetNotExists(*iotsitewise.DescribeAssetInput) error
	WaitUntilAssetNotExistsWithContext(aws.Context, *iotsitewise.DescribeAssetInput, ...request.WaiterOption) error

	WaitUntilPortalActive(*iotsitewise.DescribePortalInput) error
	WaitUntilPortalActiveWithContext(aws.Context, *iotsitewise.DescribePortalInput, ...request.WaiterOption) error

	WaitUntilPortalNotExists(*iotsitewise.DescribePortalInput) error
	WaitUntilPortalNotExistsWithContext(aws.Context, *iotsitewise.DescribePortalInput, ...request.WaiterOption) error
}

var _ IoTSiteWiseAPI = (*iotsitewise.IoTSiteWise)(nil)
