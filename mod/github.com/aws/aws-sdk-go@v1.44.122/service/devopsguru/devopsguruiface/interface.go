// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package devopsguruiface provides an interface to enable mocking the Amazon DevOps Guru service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package devopsguruiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/devopsguru"
)

// DevOpsGuruAPI provides an interface to enable mocking the
// devopsguru.DevOpsGuru service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon DevOps Guru.
//	func myFunc(svc devopsguruiface.DevOpsGuruAPI) bool {
//	    // Make svc.AddNotificationChannel request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := devopsguru.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockDevOpsGuruClient struct {
//	    devopsguruiface.DevOpsGuruAPI
//	}
//	func (m *mockDevOpsGuruClient) AddNotificationChannel(input *devopsguru.AddNotificationChannelInput) (*devopsguru.AddNotificationChannelOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockDevOpsGuruClient{}
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
type DevOpsGuruAPI interface {
	AddNotificationChannel(*devopsguru.AddNotificationChannelInput) (*devopsguru.AddNotificationChannelOutput, error)
	AddNotificationChannelWithContext(aws.Context, *devopsguru.AddNotificationChannelInput, ...request.Option) (*devopsguru.AddNotificationChannelOutput, error)
	AddNotificationChannelRequest(*devopsguru.AddNotificationChannelInput) (*request.Request, *devopsguru.AddNotificationChannelOutput)

	DeleteInsight(*devopsguru.DeleteInsightInput) (*devopsguru.DeleteInsightOutput, error)
	DeleteInsightWithContext(aws.Context, *devopsguru.DeleteInsightInput, ...request.Option) (*devopsguru.DeleteInsightOutput, error)
	DeleteInsightRequest(*devopsguru.DeleteInsightInput) (*request.Request, *devopsguru.DeleteInsightOutput)

	DescribeAccountHealth(*devopsguru.DescribeAccountHealthInput) (*devopsguru.DescribeAccountHealthOutput, error)
	DescribeAccountHealthWithContext(aws.Context, *devopsguru.DescribeAccountHealthInput, ...request.Option) (*devopsguru.DescribeAccountHealthOutput, error)
	DescribeAccountHealthRequest(*devopsguru.DescribeAccountHealthInput) (*request.Request, *devopsguru.DescribeAccountHealthOutput)

	DescribeAccountOverview(*devopsguru.DescribeAccountOverviewInput) (*devopsguru.DescribeAccountOverviewOutput, error)
	DescribeAccountOverviewWithContext(aws.Context, *devopsguru.DescribeAccountOverviewInput, ...request.Option) (*devopsguru.DescribeAccountOverviewOutput, error)
	DescribeAccountOverviewRequest(*devopsguru.DescribeAccountOverviewInput) (*request.Request, *devopsguru.DescribeAccountOverviewOutput)

	DescribeAnomaly(*devopsguru.DescribeAnomalyInput) (*devopsguru.DescribeAnomalyOutput, error)
	DescribeAnomalyWithContext(aws.Context, *devopsguru.DescribeAnomalyInput, ...request.Option) (*devopsguru.DescribeAnomalyOutput, error)
	DescribeAnomalyRequest(*devopsguru.DescribeAnomalyInput) (*request.Request, *devopsguru.DescribeAnomalyOutput)

	DescribeEventSourcesConfig(*devopsguru.DescribeEventSourcesConfigInput) (*devopsguru.DescribeEventSourcesConfigOutput, error)
	DescribeEventSourcesConfigWithContext(aws.Context, *devopsguru.DescribeEventSourcesConfigInput, ...request.Option) (*devopsguru.DescribeEventSourcesConfigOutput, error)
	DescribeEventSourcesConfigRequest(*devopsguru.DescribeEventSourcesConfigInput) (*request.Request, *devopsguru.DescribeEventSourcesConfigOutput)

	DescribeFeedback(*devopsguru.DescribeFeedbackInput) (*devopsguru.DescribeFeedbackOutput, error)
	DescribeFeedbackWithContext(aws.Context, *devopsguru.DescribeFeedbackInput, ...request.Option) (*devopsguru.DescribeFeedbackOutput, error)
	DescribeFeedbackRequest(*devopsguru.DescribeFeedbackInput) (*request.Request, *devopsguru.DescribeFeedbackOutput)

	DescribeInsight(*devopsguru.DescribeInsightInput) (*devopsguru.DescribeInsightOutput, error)
	DescribeInsightWithContext(aws.Context, *devopsguru.DescribeInsightInput, ...request.Option) (*devopsguru.DescribeInsightOutput, error)
	DescribeInsightRequest(*devopsguru.DescribeInsightInput) (*request.Request, *devopsguru.DescribeInsightOutput)

	DescribeOrganizationHealth(*devopsguru.DescribeOrganizationHealthInput) (*devopsguru.DescribeOrganizationHealthOutput, error)
	DescribeOrganizationHealthWithContext(aws.Context, *devopsguru.DescribeOrganizationHealthInput, ...request.Option) (*devopsguru.DescribeOrganizationHealthOutput, error)
	DescribeOrganizationHealthRequest(*devopsguru.DescribeOrganizationHealthInput) (*request.Request, *devopsguru.DescribeOrganizationHealthOutput)

	DescribeOrganizationOverview(*devopsguru.DescribeOrganizationOverviewInput) (*devopsguru.DescribeOrganizationOverviewOutput, error)
	DescribeOrganizationOverviewWithContext(aws.Context, *devopsguru.DescribeOrganizationOverviewInput, ...request.Option) (*devopsguru.DescribeOrganizationOverviewOutput, error)
	DescribeOrganizationOverviewRequest(*devopsguru.DescribeOrganizationOverviewInput) (*request.Request, *devopsguru.DescribeOrganizationOverviewOutput)

	DescribeOrganizationResourceCollectionHealth(*devopsguru.DescribeOrganizationResourceCollectionHealthInput) (*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, error)
	DescribeOrganizationResourceCollectionHealthWithContext(aws.Context, *devopsguru.DescribeOrganizationResourceCollectionHealthInput, ...request.Option) (*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, error)
	DescribeOrganizationResourceCollectionHealthRequest(*devopsguru.DescribeOrganizationResourceCollectionHealthInput) (*request.Request, *devopsguru.DescribeOrganizationResourceCollectionHealthOutput)

	DescribeOrganizationResourceCollectionHealthPages(*devopsguru.DescribeOrganizationResourceCollectionHealthInput, func(*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, bool) bool) error
	DescribeOrganizationResourceCollectionHealthPagesWithContext(aws.Context, *devopsguru.DescribeOrganizationResourceCollectionHealthInput, func(*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, bool) bool, ...request.Option) error

	DescribeResourceCollectionHealth(*devopsguru.DescribeResourceCollectionHealthInput) (*devopsguru.DescribeResourceCollectionHealthOutput, error)
	DescribeResourceCollectionHealthWithContext(aws.Context, *devopsguru.DescribeResourceCollectionHealthInput, ...request.Option) (*devopsguru.DescribeResourceCollectionHealthOutput, error)
	DescribeResourceCollectionHealthRequest(*devopsguru.DescribeResourceCollectionHealthInput) (*request.Request, *devopsguru.DescribeResourceCollectionHealthOutput)

	DescribeResourceCollectionHealthPages(*devopsguru.DescribeResourceCollectionHealthInput, func(*devopsguru.DescribeResourceCollectionHealthOutput, bool) bool) error
	DescribeResourceCollectionHealthPagesWithContext(aws.Context, *devopsguru.DescribeResourceCollectionHealthInput, func(*devopsguru.DescribeResourceCollectionHealthOutput, bool) bool, ...request.Option) error

	DescribeServiceIntegration(*devopsguru.DescribeServiceIntegrationInput) (*devopsguru.DescribeServiceIntegrationOutput, error)
	DescribeServiceIntegrationWithContext(aws.Context, *devopsguru.DescribeServiceIntegrationInput, ...request.Option) (*devopsguru.DescribeServiceIntegrationOutput, error)
	DescribeServiceIntegrationRequest(*devopsguru.DescribeServiceIntegrationInput) (*request.Request, *devopsguru.DescribeServiceIntegrationOutput)

	GetCostEstimation(*devopsguru.GetCostEstimationInput) (*devopsguru.GetCostEstimationOutput, error)
	GetCostEstimationWithContext(aws.Context, *devopsguru.GetCostEstimationInput, ...request.Option) (*devopsguru.GetCostEstimationOutput, error)
	GetCostEstimationRequest(*devopsguru.GetCostEstimationInput) (*request.Request, *devopsguru.GetCostEstimationOutput)

	GetCostEstimationPages(*devopsguru.GetCostEstimationInput, func(*devopsguru.GetCostEstimationOutput, bool) bool) error
	GetCostEstimationPagesWithContext(aws.Context, *devopsguru.GetCostEstimationInput, func(*devopsguru.GetCostEstimationOutput, bool) bool, ...request.Option) error

	GetResourceCollection(*devopsguru.GetResourceCollectionInput) (*devopsguru.GetResourceCollectionOutput, error)
	GetResourceCollectionWithContext(aws.Context, *devopsguru.GetResourceCollectionInput, ...request.Option) (*devopsguru.GetResourceCollectionOutput, error)
	GetResourceCollectionRequest(*devopsguru.GetResourceCollectionInput) (*request.Request, *devopsguru.GetResourceCollectionOutput)

	GetResourceCollectionPages(*devopsguru.GetResourceCollectionInput, func(*devopsguru.GetResourceCollectionOutput, bool) bool) error
	GetResourceCollectionPagesWithContext(aws.Context, *devopsguru.GetResourceCollectionInput, func(*devopsguru.GetResourceCollectionOutput, bool) bool, ...request.Option) error

	ListAnomaliesForInsight(*devopsguru.ListAnomaliesForInsightInput) (*devopsguru.ListAnomaliesForInsightOutput, error)
	ListAnomaliesForInsightWithContext(aws.Context, *devopsguru.ListAnomaliesForInsightInput, ...request.Option) (*devopsguru.ListAnomaliesForInsightOutput, error)
	ListAnomaliesForInsightRequest(*devopsguru.ListAnomaliesForInsightInput) (*request.Request, *devopsguru.ListAnomaliesForInsightOutput)

	ListAnomaliesForInsightPages(*devopsguru.ListAnomaliesForInsightInput, func(*devopsguru.ListAnomaliesForInsightOutput, bool) bool) error
	ListAnomaliesForInsightPagesWithContext(aws.Context, *devopsguru.ListAnomaliesForInsightInput, func(*devopsguru.ListAnomaliesForInsightOutput, bool) bool, ...request.Option) error

	ListAnomalousLogGroups(*devopsguru.ListAnomalousLogGroupsInput) (*devopsguru.ListAnomalousLogGroupsOutput, error)
	ListAnomalousLogGroupsWithContext(aws.Context, *devopsguru.ListAnomalousLogGroupsInput, ...request.Option) (*devopsguru.ListAnomalousLogGroupsOutput, error)
	ListAnomalousLogGroupsRequest(*devopsguru.ListAnomalousLogGroupsInput) (*request.Request, *devopsguru.ListAnomalousLogGroupsOutput)

	ListAnomalousLogGroupsPages(*devopsguru.ListAnomalousLogGroupsInput, func(*devopsguru.ListAnomalousLogGroupsOutput, bool) bool) error
	ListAnomalousLogGroupsPagesWithContext(aws.Context, *devopsguru.ListAnomalousLogGroupsInput, func(*devopsguru.ListAnomalousLogGroupsOutput, bool) bool, ...request.Option) error

	ListEvents(*devopsguru.ListEventsInput) (*devopsguru.ListEventsOutput, error)
	ListEventsWithContext(aws.Context, *devopsguru.ListEventsInput, ...request.Option) (*devopsguru.ListEventsOutput, error)
	ListEventsRequest(*devopsguru.ListEventsInput) (*request.Request, *devopsguru.ListEventsOutput)

	ListEventsPages(*devopsguru.ListEventsInput, func(*devopsguru.ListEventsOutput, bool) bool) error
	ListEventsPagesWithContext(aws.Context, *devopsguru.ListEventsInput, func(*devopsguru.ListEventsOutput, bool) bool, ...request.Option) error

	ListInsights(*devopsguru.ListInsightsInput) (*devopsguru.ListInsightsOutput, error)
	ListInsightsWithContext(aws.Context, *devopsguru.ListInsightsInput, ...request.Option) (*devopsguru.ListInsightsOutput, error)
	ListInsightsRequest(*devopsguru.ListInsightsInput) (*request.Request, *devopsguru.ListInsightsOutput)

	ListInsightsPages(*devopsguru.ListInsightsInput, func(*devopsguru.ListInsightsOutput, bool) bool) error
	ListInsightsPagesWithContext(aws.Context, *devopsguru.ListInsightsInput, func(*devopsguru.ListInsightsOutput, bool) bool, ...request.Option) error

	ListMonitoredResources(*devopsguru.ListMonitoredResourcesInput) (*devopsguru.ListMonitoredResourcesOutput, error)
	ListMonitoredResourcesWithContext(aws.Context, *devopsguru.ListMonitoredResourcesInput, ...request.Option) (*devopsguru.ListMonitoredResourcesOutput, error)
	ListMonitoredResourcesRequest(*devopsguru.ListMonitoredResourcesInput) (*request.Request, *devopsguru.ListMonitoredResourcesOutput)

	ListMonitoredResourcesPages(*devopsguru.ListMonitoredResourcesInput, func(*devopsguru.ListMonitoredResourcesOutput, bool) bool) error
	ListMonitoredResourcesPagesWithContext(aws.Context, *devopsguru.ListMonitoredResourcesInput, func(*devopsguru.ListMonitoredResourcesOutput, bool) bool, ...request.Option) error

	ListNotificationChannels(*devopsguru.ListNotificationChannelsInput) (*devopsguru.ListNotificationChannelsOutput, error)
	ListNotificationChannelsWithContext(aws.Context, *devopsguru.ListNotificationChannelsInput, ...request.Option) (*devopsguru.ListNotificationChannelsOutput, error)
	ListNotificationChannelsRequest(*devopsguru.ListNotificationChannelsInput) (*request.Request, *devopsguru.ListNotificationChannelsOutput)

	ListNotificationChannelsPages(*devopsguru.ListNotificationChannelsInput, func(*devopsguru.ListNotificationChannelsOutput, bool) bool) error
	ListNotificationChannelsPagesWithContext(aws.Context, *devopsguru.ListNotificationChannelsInput, func(*devopsguru.ListNotificationChannelsOutput, bool) bool, ...request.Option) error

	ListOrganizationInsights(*devopsguru.ListOrganizationInsightsInput) (*devopsguru.ListOrganizationInsightsOutput, error)
	ListOrganizationInsightsWithContext(aws.Context, *devopsguru.ListOrganizationInsightsInput, ...request.Option) (*devopsguru.ListOrganizationInsightsOutput, error)
	ListOrganizationInsightsRequest(*devopsguru.ListOrganizationInsightsInput) (*request.Request, *devopsguru.ListOrganizationInsightsOutput)

	ListOrganizationInsightsPages(*devopsguru.ListOrganizationInsightsInput, func(*devopsguru.ListOrganizationInsightsOutput, bool) bool) error
	ListOrganizationInsightsPagesWithContext(aws.Context, *devopsguru.ListOrganizationInsightsInput, func(*devopsguru.ListOrganizationInsightsOutput, bool) bool, ...request.Option) error

	ListRecommendations(*devopsguru.ListRecommendationsInput) (*devopsguru.ListRecommendationsOutput, error)
	ListRecommendationsWithContext(aws.Context, *devopsguru.ListRecommendationsInput, ...request.Option) (*devopsguru.ListRecommendationsOutput, error)
	ListRecommendationsRequest(*devopsguru.ListRecommendationsInput) (*request.Request, *devopsguru.ListRecommendationsOutput)

	ListRecommendationsPages(*devopsguru.ListRecommendationsInput, func(*devopsguru.ListRecommendationsOutput, bool) bool) error
	ListRecommendationsPagesWithContext(aws.Context, *devopsguru.ListRecommendationsInput, func(*devopsguru.ListRecommendationsOutput, bool) bool, ...request.Option) error

	PutFeedback(*devopsguru.PutFeedbackInput) (*devopsguru.PutFeedbackOutput, error)
	PutFeedbackWithContext(aws.Context, *devopsguru.PutFeedbackInput, ...request.Option) (*devopsguru.PutFeedbackOutput, error)
	PutFeedbackRequest(*devopsguru.PutFeedbackInput) (*request.Request, *devopsguru.PutFeedbackOutput)

	RemoveNotificationChannel(*devopsguru.RemoveNotificationChannelInput) (*devopsguru.RemoveNotificationChannelOutput, error)
	RemoveNotificationChannelWithContext(aws.Context, *devopsguru.RemoveNotificationChannelInput, ...request.Option) (*devopsguru.RemoveNotificationChannelOutput, error)
	RemoveNotificationChannelRequest(*devopsguru.RemoveNotificationChannelInput) (*request.Request, *devopsguru.RemoveNotificationChannelOutput)

	SearchInsights(*devopsguru.SearchInsightsInput) (*devopsguru.SearchInsightsOutput, error)
	SearchInsightsWithContext(aws.Context, *devopsguru.SearchInsightsInput, ...request.Option) (*devopsguru.SearchInsightsOutput, error)
	SearchInsightsRequest(*devopsguru.SearchInsightsInput) (*request.Request, *devopsguru.SearchInsightsOutput)

	SearchInsightsPages(*devopsguru.SearchInsightsInput, func(*devopsguru.SearchInsightsOutput, bool) bool) error
	SearchInsightsPagesWithContext(aws.Context, *devopsguru.SearchInsightsInput, func(*devopsguru.SearchInsightsOutput, bool) bool, ...request.Option) error

	SearchOrganizationInsights(*devopsguru.SearchOrganizationInsightsInput) (*devopsguru.SearchOrganizationInsightsOutput, error)
	SearchOrganizationInsightsWithContext(aws.Context, *devopsguru.SearchOrganizationInsightsInput, ...request.Option) (*devopsguru.SearchOrganizationInsightsOutput, error)
	SearchOrganizationInsightsRequest(*devopsguru.SearchOrganizationInsightsInput) (*request.Request, *devopsguru.SearchOrganizationInsightsOutput)

	SearchOrganizationInsightsPages(*devopsguru.SearchOrganizationInsightsInput, func(*devopsguru.SearchOrganizationInsightsOutput, bool) bool) error
	SearchOrganizationInsightsPagesWithContext(aws.Context, *devopsguru.SearchOrganizationInsightsInput, func(*devopsguru.SearchOrganizationInsightsOutput, bool) bool, ...request.Option) error

	StartCostEstimation(*devopsguru.StartCostEstimationInput) (*devopsguru.StartCostEstimationOutput, error)
	StartCostEstimationWithContext(aws.Context, *devopsguru.StartCostEstimationInput, ...request.Option) (*devopsguru.StartCostEstimationOutput, error)
	StartCostEstimationRequest(*devopsguru.StartCostEstimationInput) (*request.Request, *devopsguru.StartCostEstimationOutput)

	UpdateEventSourcesConfig(*devopsguru.UpdateEventSourcesConfigInput) (*devopsguru.UpdateEventSourcesConfigOutput, error)
	UpdateEventSourcesConfigWithContext(aws.Context, *devopsguru.UpdateEventSourcesConfigInput, ...request.Option) (*devopsguru.UpdateEventSourcesConfigOutput, error)
	UpdateEventSourcesConfigRequest(*devopsguru.UpdateEventSourcesConfigInput) (*request.Request, *devopsguru.UpdateEventSourcesConfigOutput)

	UpdateResourceCollection(*devopsguru.UpdateResourceCollectionInput) (*devopsguru.UpdateResourceCollectionOutput, error)
	UpdateResourceCollectionWithContext(aws.Context, *devopsguru.UpdateResourceCollectionInput, ...request.Option) (*devopsguru.UpdateResourceCollectionOutput, error)
	UpdateResourceCollectionRequest(*devopsguru.UpdateResourceCollectionInput) (*request.Request, *devopsguru.UpdateResourceCollectionOutput)

	UpdateServiceIntegration(*devopsguru.UpdateServiceIntegrationInput) (*devopsguru.UpdateServiceIntegrationOutput, error)
	UpdateServiceIntegrationWithContext(aws.Context, *devopsguru.UpdateServiceIntegrationInput, ...request.Option) (*devopsguru.UpdateServiceIntegrationOutput, error)
	UpdateServiceIntegrationRequest(*devopsguru.UpdateServiceIntegrationInput) (*request.Request, *devopsguru.UpdateServiceIntegrationOutput)
}

var _ DevOpsGuruAPI = (*devopsguru.DevOpsGuru)(nil)
