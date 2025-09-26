package client

import (
	"context"
	"fmt"
	"io"
)

// ErrServiceNotImplemented is returned when a service is not yet implemented
var ErrServiceNotImplemented = fmt.Errorf("this service is not yet implemented in the current SDK version")

// notImplementedAdService implements AdService with not-implemented errors
type notImplementedAdService struct{}

func (s *notImplementedAdService) Create(ctx context.Context, req *AdCreateRequest) (*AdCreateResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAdService) Get(ctx context.Context, req *AdGetRequest) (*AdGetResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAdService) Update(ctx context.Context, req *AdUpdateRequest) (*AdUpdateResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAdService) Delete(ctx context.Context, req *AdDeleteRequest) (*AdDeleteResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAdService) UpdateStatus(ctx context.Context, req *AdStatusUpdateRequest) (*AdStatusUpdateResponse, error) {
	return nil, ErrServiceNotImplemented
}

// notImplementedAdGroupService implements AdGroupService with not-implemented errors
type notImplementedAdGroupService struct{}

func (s *notImplementedAdGroupService) Create(ctx context.Context, req *AdGroupCreateRequest) (*AdGroupCreateResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAdGroupService) Get(ctx context.Context, req *AdGroupGetRequest) (*AdGroupGetResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAdGroupService) Update(ctx context.Context, req *AdGroupUpdateRequest) (*AdGroupUpdateResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAdGroupService) Delete(ctx context.Context, req *AdGroupDeleteRequest) (*AdGroupDeleteResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAdGroupService) UpdateStatus(ctx context.Context, req *AdGroupStatusUpdateRequest) (*AdGroupStatusUpdateResponse, error) {
	return nil, ErrServiceNotImplemented
}

// notImplementedAudienceService implements AudienceService with not-implemented errors
type notImplementedAudienceService struct{}

func (s *notImplementedAudienceService) CreateCustomAudience(ctx context.Context, req *CustomAudienceCreateRequest) (*CustomAudienceCreateResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAudienceService) GetCustomAudiences(ctx context.Context, req *CustomAudienceGetRequest) (*CustomAudienceGetResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAudienceService) UpdateCustomAudience(ctx context.Context, req *CustomAudienceUpdateRequest) (*CustomAudienceUpdateResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAudienceService) DeleteCustomAudience(ctx context.Context, req *CustomAudienceDeleteRequest) (*CustomAudienceDeleteResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedAudienceService) CreateLookalikeAudience(ctx context.Context, req *LookalikeAudienceCreateRequest) (*LookalikeAudienceCreateResponse, error) {
	return nil, ErrServiceNotImplemented
}

// notImplementedCreativeService implements CreativeService with not-implemented errors
type notImplementedCreativeService struct{}

func (s *notImplementedCreativeService) UploadImage(ctx context.Context, req *ImageUploadRequest) (*ImageUploadResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedCreativeService) UploadVideo(ctx context.Context, req *VideoUploadRequest) (*VideoUploadResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedCreativeService) GetCreatives(ctx context.Context, req *CreativeGetRequest) (*CreativeGetResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedCreativeService) UpdateCreative(ctx context.Context, req *CreativeUpdateRequest) (*CreativeUpdateResponse, error) {
	return nil, ErrServiceNotImplemented
}

// notImplementedReportingService implements ReportingService with not-implemented errors
type notImplementedReportingService struct{}

func (s *notImplementedReportingService) GetBasicReports(ctx context.Context, req *ReportingRequest) (*ReportingResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedReportingService) GetAudienceReports(ctx context.Context, req *AudienceReportingRequest) (*AudienceReportingResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedReportingService) CreateAsyncReport(ctx context.Context, req *AsyncReportRequest) (*AsyncReportResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedReportingService) GetAsyncReportStatus(ctx context.Context, taskID string) (*AsyncReportStatusResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedReportingService) DownloadAsyncReport(ctx context.Context, taskID string) (io.ReadCloser, error) {
	return nil, ErrServiceNotImplemented
}

// notImplementedBCService implements BCService with not-implemented errors
type notImplementedBCService struct{}

func (s *notImplementedBCService) GetBusinessCenters(ctx context.Context) (*BusinessCentersResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedBCService) GetBusinessCenterInfo(ctx context.Context, bcID string) (*BusinessCenterInfo, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedBCService) GetAdvertisersInBC(ctx context.Context, bcID string) (*BCAdvertisersResponse, error) {
	return nil, ErrServiceNotImplemented
}

func (s *notImplementedBCService) TransferAdvertiser(ctx context.Context, req *TransferAdvertiserRequest) (*TransferAdvertiserResponse, error) {
	return nil, ErrServiceNotImplemented
}
