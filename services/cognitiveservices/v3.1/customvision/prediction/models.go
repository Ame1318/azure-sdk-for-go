package prediction

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/gofrs/uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.1/customvision/prediction"

// BoundingBox bounding box that defines a region of an image.
type BoundingBox struct {
	// Left - Coordinate of the left boundary.
	Left *float64 `json:"left,omitempty"`
	// Top - Coordinate of the top boundary.
	Top *float64 `json:"top,omitempty"`
	// Width - Width.
	Width *float64 `json:"width,omitempty"`
	// Height - Height.
	Height *float64 `json:"height,omitempty"`
}

// CustomVisionError ...
type CustomVisionError struct {
	// Code - The error code. Possible values include: 'NoError', 'BadRequest', 'BadRequestExceededBatchSize', 'BadRequestNotSupported', 'BadRequestInvalidIds', 'BadRequestProjectName', 'BadRequestProjectNameNotUnique', 'BadRequestProjectDescription', 'BadRequestProjectUnknownDomain', 'BadRequestProjectUnknownClassification', 'BadRequestProjectUnsupportedDomainTypeChange', 'BadRequestProjectUnsupportedExportPlatform', 'BadRequestProjectImagePreprocessingSettings', 'BadRequestProjectDuplicated', 'BadRequestIterationName', 'BadRequestIterationNameNotUnique', 'BadRequestIterationDescription', 'BadRequestIterationIsNotTrained', 'BadRequestIterationValidationFailed', 'BadRequestWorkspaceCannotBeModified', 'BadRequestWorkspaceNotDeletable', 'BadRequestTagName', 'BadRequestTagNameNotUnique', 'BadRequestTagDescription', 'BadRequestTagType', 'BadRequestMultipleNegativeTag', 'BadRequestMultipleGeneralProductTag', 'BadRequestImageTags', 'BadRequestImageRegions', 'BadRequestNegativeAndRegularTagOnSameImage', 'BadRequestUnsupportedDomain', 'BadRequestRequiredParamIsNull', 'BadRequestIterationIsPublished', 'BadRequestInvalidPublishName', 'BadRequestInvalidPublishTarget', 'BadRequestUnpublishFailed', 'BadRequestIterationNotPublished', 'BadRequestSubscriptionAPI', 'BadRequestExceedProjectLimit', 'BadRequestExceedIterationPerProjectLimit', 'BadRequestExceedTagPerProjectLimit', 'BadRequestExceedTagPerImageLimit', 'BadRequestExceededQuota', 'BadRequestCannotMigrateProjectWithName', 'BadRequestNotLimitedTrial', 'BadRequestImageBatch', 'BadRequestImageStream', 'BadRequestImageURL', 'BadRequestImageFormat', 'BadRequestImageSizeBytes', 'BadRequestImageDimensions', 'BadRequestImageExceededCount', 'BadRequestTrainingNotNeeded', 'BadRequestTrainingNotNeededButTrainingPipelineUpdated', 'BadRequestTrainingValidationFailed', 'BadRequestClassificationTrainingValidationFailed', 'BadRequestMultiClassClassificationTrainingValidationFailed', 'BadRequestMultiLabelClassificationTrainingValidationFailed', 'BadRequestDetectionTrainingValidationFailed', 'BadRequestTrainingAlreadyInProgress', 'BadRequestDetectionTrainingNotAllowNegativeTag', 'BadRequestInvalidEmailAddress', 'BadRequestDomainNotSupportedForAdvancedTraining', 'BadRequestExportPlatformNotSupportedForAdvancedTraining', 'BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining', 'BadRequestExportValidationFailed', 'BadRequestExportAlreadyInProgress', 'BadRequestPredictionIdsMissing', 'BadRequestPredictionIdsExceededCount', 'BadRequestPredictionTagsExceededCount', 'BadRequestPredictionResultsExceededCount', 'BadRequestPredictionInvalidApplicationName', 'BadRequestPredictionInvalidQueryParameters', 'BadRequestInvalidImportToken', 'BadRequestExportWhileTraining', 'BadRequestImageMetadataKey', 'BadRequestImageMetadataValue', 'BadRequestOperationNotSupported', 'BadRequestInvalidArtifactURI', 'BadRequestCustomerManagedKeyRevoked', 'BadRequestInvalid', 'UnsupportedMediaType', 'Forbidden', 'ForbiddenUser', 'ForbiddenUserResource', 'ForbiddenUserSignupDisabled', 'ForbiddenUserSignupAllowanceExceeded', 'ForbiddenUserDoesNotExist', 'ForbiddenUserDisabled', 'ForbiddenUserInsufficientCapability', 'ForbiddenDRModeEnabled', 'ForbiddenInvalid', 'NotFound', 'NotFoundProject', 'NotFoundProjectDefaultIteration', 'NotFoundIteration', 'NotFoundIterationPerformance', 'NotFoundTag', 'NotFoundImage', 'NotFoundDomain', 'NotFoundApimSubscription', 'NotFoundInvalid', 'Conflict', 'ConflictInvalid', 'ErrorUnknown', 'ErrorIterationCopyFailed', 'ErrorPreparePerformanceMigrationFailed', 'ErrorProjectInvalidWorkspace', 'ErrorProjectInvalidPipelineConfiguration', 'ErrorProjectInvalidDomain', 'ErrorProjectTrainingRequestFailed', 'ErrorProjectImportRequestFailed', 'ErrorProjectExportRequestFailed', 'ErrorFeaturizationServiceUnavailable', 'ErrorFeaturizationQueueTimeout', 'ErrorFeaturizationInvalidFeaturizer', 'ErrorFeaturizationAugmentationUnavailable', 'ErrorFeaturizationUnrecognizedJob', 'ErrorFeaturizationAugmentationError', 'ErrorExporterInvalidPlatform', 'ErrorExporterInvalidFeaturizer', 'ErrorExporterInvalidClassifier', 'ErrorPredictionServiceUnavailable', 'ErrorPredictionModelNotFound', 'ErrorPredictionModelNotCached', 'ErrorPrediction', 'ErrorPredictionStorage', 'ErrorRegionProposal', 'ErrorUnknownBaseModel', 'ErrorInvalid'
	Code CustomVisionErrorCodes `json:"code,omitempty"`
	// Message - A message explaining the error reported by the service.
	Message *string `json:"message,omitempty"`
}

// ImagePrediction result of an image prediction request.
type ImagePrediction struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Prediction Id.
	ID *uuid.UUID `json:"id,omitempty"`
	// Project - READ-ONLY; Project Id.
	Project *uuid.UUID `json:"project,omitempty"`
	// Iteration - READ-ONLY; Iteration Id.
	Iteration *uuid.UUID `json:"iteration,omitempty"`
	// Created - READ-ONLY; Date this prediction was created.
	Created *date.Time `json:"created,omitempty"`
	// Predictions - READ-ONLY; List of predictions.
	Predictions *[]Model `json:"predictions,omitempty"`
}

// ImageURL image url.
type ImageURL struct {
	// URL - Url of the image.
	URL *string `json:"url,omitempty"`
}

// Model prediction result.
type Model struct {
	// Probability - READ-ONLY; Probability of the tag.
	Probability *float64 `json:"probability,omitempty"`
	// TagID - READ-ONLY; Id of the predicted tag.
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// TagName - READ-ONLY; Name of the predicted tag.
	TagName *string `json:"tagName,omitempty"`
	// BoundingBox - READ-ONLY; Bounding box of the prediction.
	BoundingBox *BoundingBox `json:"boundingBox,omitempty"`
	// TagType - READ-ONLY; Type of the predicted tag. Possible values include: 'Regular', 'Negative', 'GeneralProduct'
	TagType TagType `json:"tagType,omitempty"`
}
