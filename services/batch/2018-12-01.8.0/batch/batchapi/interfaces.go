package batchapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/batch/2018-12-01.8.0/batch"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/gofrs/uuid"
)

// ApplicationClientAPI contains the set of methods on the ApplicationClient type.
type ApplicationClientAPI interface {
	Get(ctx context.Context, applicationID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.ApplicationSummary, err error)
	List(ctx context.Context, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.ApplicationListResultPage, err error)
	ListComplete(ctx context.Context, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.ApplicationListResultIterator, err error)
}

var _ ApplicationClientAPI = (*batch.ApplicationClient)(nil)

// PoolClientAPI contains the set of methods on the PoolClient type.
type PoolClientAPI interface {
	Add(ctx context.Context, pool batch.PoolAddParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	Delete(ctx context.Context, poolID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	DisableAutoScale(ctx context.Context, poolID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	EnableAutoScale(ctx context.Context, poolID string, poolEnableAutoScaleParameter batch.PoolEnableAutoScaleParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	EvaluateAutoScale(ctx context.Context, poolID string, poolEvaluateAutoScaleParameter batch.PoolEvaluateAutoScaleParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.AutoScaleRun, err error)
	Exists(ctx context.Context, poolID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Get(ctx context.Context, poolID string, selectParameter string, expand string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result batch.CloudPool, err error)
	GetAllLifetimeStatistics(ctx context.Context, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.PoolStatistics, err error)
	List(ctx context.Context, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudPoolListResultPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudPoolListResultIterator, err error)
	ListUsageMetrics(ctx context.Context, startTime *date.Time, endTime *date.Time, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.PoolListUsageMetricsResultPage, err error)
	ListUsageMetricsComplete(ctx context.Context, startTime *date.Time, endTime *date.Time, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.PoolListUsageMetricsResultIterator, err error)
	Patch(ctx context.Context, poolID string, poolPatchParameter batch.PoolPatchParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	RemoveNodes(ctx context.Context, poolID string, nodeRemoveParameter batch.NodeRemoveParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Resize(ctx context.Context, poolID string, poolResizeParameter batch.PoolResizeParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	StopResize(ctx context.Context, poolID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	UpdateProperties(ctx context.Context, poolID string, poolUpdatePropertiesParameter batch.PoolUpdatePropertiesParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
}

var _ PoolClientAPI = (*batch.PoolClient)(nil)

// AccountClientAPI contains the set of methods on the AccountClient type.
type AccountClientAPI interface {
	ListNodeAgentSkus(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.AccountListNodeAgentSkusResultPage, err error)
	ListNodeAgentSkusComplete(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.AccountListNodeAgentSkusResultIterator, err error)
	ListPoolNodeCounts(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.PoolNodeCountsListResultPage, err error)
	ListPoolNodeCountsComplete(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.PoolNodeCountsListResultIterator, err error)
}

var _ AccountClientAPI = (*batch.AccountClient)(nil)

// JobClientAPI contains the set of methods on the JobClient type.
type JobClientAPI interface {
	Add(ctx context.Context, job batch.JobAddParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	Delete(ctx context.Context, jobID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Disable(ctx context.Context, jobID string, jobDisableParameter batch.JobDisableParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Enable(ctx context.Context, jobID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Get(ctx context.Context, jobID string, selectParameter string, expand string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result batch.CloudJob, err error)
	GetAllLifetimeStatistics(ctx context.Context, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.JobStatistics, err error)
	GetTaskCounts(ctx context.Context, jobID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.TaskCounts, err error)
	List(ctx context.Context, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudJobListResultPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudJobListResultIterator, err error)
	ListFromJobSchedule(ctx context.Context, jobScheduleID string, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudJobListResultPage, err error)
	ListFromJobScheduleComplete(ctx context.Context, jobScheduleID string, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudJobListResultIterator, err error)
	ListPreparationAndReleaseTaskStatus(ctx context.Context, jobID string, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudJobListPreparationAndReleaseTaskStatusResultPage, err error)
	ListPreparationAndReleaseTaskStatusComplete(ctx context.Context, jobID string, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudJobListPreparationAndReleaseTaskStatusResultIterator, err error)
	Patch(ctx context.Context, jobID string, jobPatchParameter batch.JobPatchParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Terminate(ctx context.Context, jobID string, jobTerminateParameter *batch.JobTerminateParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Update(ctx context.Context, jobID string, jobUpdateParameter batch.JobUpdateParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
}

var _ JobClientAPI = (*batch.JobClient)(nil)

// CertificateClientAPI contains the set of methods on the CertificateClient type.
type CertificateClientAPI interface {
	Add(ctx context.Context, certificate batch.CertificateAddParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	CancelDeletion(ctx context.Context, thumbprintAlgorithm string, thumbprint string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	Delete(ctx context.Context, thumbprintAlgorithm string, thumbprint string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	Get(ctx context.Context, thumbprintAlgorithm string, thumbprint string, selectParameter string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.Certificate, err error)
	List(ctx context.Context, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CertificateListResultPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CertificateListResultIterator, err error)
}

var _ CertificateClientAPI = (*batch.CertificateClient)(nil)

// FileClientAPI contains the set of methods on the FileClient type.
type FileClientAPI interface {
	DeleteFromComputeNode(ctx context.Context, poolID string, nodeID string, filePath string, recursive *bool, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	DeleteFromTask(ctx context.Context, jobID string, taskID string, filePath string, recursive *bool, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	GetFromComputeNode(ctx context.Context, poolID string, nodeID string, filePath string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ocpRange string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result batch.ReadCloser, err error)
	GetFromTask(ctx context.Context, jobID string, taskID string, filePath string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ocpRange string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result batch.ReadCloser, err error)
	GetPropertiesFromComputeNode(ctx context.Context, poolID string, nodeID string, filePath string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	GetPropertiesFromTask(ctx context.Context, jobID string, taskID string, filePath string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	ListFromComputeNode(ctx context.Context, poolID string, nodeID string, filter string, recursive *bool, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.NodeFileListResultPage, err error)
	ListFromComputeNodeComplete(ctx context.Context, poolID string, nodeID string, filter string, recursive *bool, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.NodeFileListResultIterator, err error)
	ListFromTask(ctx context.Context, jobID string, taskID string, filter string, recursive *bool, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.NodeFileListResultPage, err error)
	ListFromTaskComplete(ctx context.Context, jobID string, taskID string, filter string, recursive *bool, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.NodeFileListResultIterator, err error)
}

var _ FileClientAPI = (*batch.FileClient)(nil)

// JobScheduleClientAPI contains the set of methods on the JobScheduleClient type.
type JobScheduleClientAPI interface {
	Add(ctx context.Context, cloudJobSchedule batch.JobScheduleAddParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	Delete(ctx context.Context, jobScheduleID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Disable(ctx context.Context, jobScheduleID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Enable(ctx context.Context, jobScheduleID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Exists(ctx context.Context, jobScheduleID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Get(ctx context.Context, jobScheduleID string, selectParameter string, expand string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result batch.CloudJobSchedule, err error)
	List(ctx context.Context, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudJobScheduleListResultPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudJobScheduleListResultIterator, err error)
	Patch(ctx context.Context, jobScheduleID string, jobSchedulePatchParameter batch.JobSchedulePatchParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Terminate(ctx context.Context, jobScheduleID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Update(ctx context.Context, jobScheduleID string, jobScheduleUpdateParameter batch.JobScheduleUpdateParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
}

var _ JobScheduleClientAPI = (*batch.JobScheduleClient)(nil)

// TaskClientAPI contains the set of methods on the TaskClient type.
type TaskClientAPI interface {
	Add(ctx context.Context, jobID string, task batch.TaskAddParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	AddCollection(ctx context.Context, jobID string, taskCollection batch.TaskAddCollectionParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.TaskAddCollectionResult, err error)
	Delete(ctx context.Context, jobID string, taskID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Get(ctx context.Context, jobID string, taskID string, selectParameter string, expand string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result batch.CloudTask, err error)
	List(ctx context.Context, jobID string, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudTaskListResultPage, err error)
	ListComplete(ctx context.Context, jobID string, filter string, selectParameter string, expand string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudTaskListResultIterator, err error)
	ListSubtasks(ctx context.Context, jobID string, taskID string, selectParameter string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.CloudTaskListSubtasksResult, err error)
	Reactivate(ctx context.Context, jobID string, taskID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Terminate(ctx context.Context, jobID string, taskID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
	Update(ctx context.Context, jobID string, taskID string, taskUpdateParameter batch.TaskUpdateParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, ifMatch string, ifNoneMatch string, ifModifiedSince *date.TimeRFC1123, ifUnmodifiedSince *date.TimeRFC1123) (result autorest.Response, err error)
}

var _ TaskClientAPI = (*batch.TaskClient)(nil)

// ComputeNodeClientAPI contains the set of methods on the ComputeNodeClient type.
type ComputeNodeClientAPI interface {
	AddUser(ctx context.Context, poolID string, nodeID string, userParameter batch.ComputeNodeUser, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	DeleteUser(ctx context.Context, poolID string, nodeID string, userName string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	DisableScheduling(ctx context.Context, poolID string, nodeID string, nodeDisableSchedulingParameter *batch.NodeDisableSchedulingParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	EnableScheduling(ctx context.Context, poolID string, nodeID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	Get(ctx context.Context, poolID string, nodeID string, selectParameter string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.ComputeNode, err error)
	GetRemoteDesktop(ctx context.Context, poolID string, nodeID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.ReadCloser, err error)
	GetRemoteLoginSettings(ctx context.Context, poolID string, nodeID string, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.ComputeNodeGetRemoteLoginSettingsResult, err error)
	List(ctx context.Context, poolID string, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.ComputeNodeListResultPage, err error)
	ListComplete(ctx context.Context, poolID string, filter string, selectParameter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.ComputeNodeListResultIterator, err error)
	Reboot(ctx context.Context, poolID string, nodeID string, nodeRebootParameter *batch.NodeRebootParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	Reimage(ctx context.Context, poolID string, nodeID string, nodeReimageParameter *batch.NodeReimageParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	UpdateUser(ctx context.Context, poolID string, nodeID string, userName string, nodeUpdateUserParameter batch.NodeUpdateUserParameter, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result autorest.Response, err error)
	UploadBatchServiceLogs(ctx context.Context, poolID string, nodeID string, uploadBatchServiceLogsConfiguration batch.UploadBatchServiceLogsConfiguration, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result batch.UploadBatchServiceLogsResult, err error)
}

var _ ComputeNodeClientAPI = (*batch.ComputeNodeClient)(nil)
