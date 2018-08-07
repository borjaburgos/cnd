package model

import (
	"fmt"
)

const (

	// InvalidJSON is returned when a json string can't be read
	InvalidJSON AppErrorCode = "InvalidJson"

	// InvalidYAML is returned when a yaml string can't be read
	InvalidYAML AppErrorCode = "InvalidYAML"

	// InvalidBase64 is returned when the a base64 string can't be decoded
	InvalidBase64 AppErrorCode = "InvalidBase64"

	// InvalidForm is returned when a form string can't be read
	InvalidForm AppErrorCode = "InvalidForm"

	// MissingID is returned when an entity doesn't have an ID
	MissingID AppErrorCode = "MissingID"

	// MissingName is returned when an entity doesn't have a name
	MissingName AppErrorCode = "MissingName"

	// InvalidName is returned when an entity doesn't have a valid name
	InvalidName AppErrorCode = "InvalidName"

	// UniqueName is returned when another entity in the same scope has the name name
	UniqueName AppErrorCode = "UniqueName"

	// MissingUsers is returned when a project settings don't have any users
	MissingUsers AppErrorCode = "MissingUsers"

	// MissingAdministrators is returned when a project settings don't have any admins
	MissingAdministrators AppErrorCode = "MissingAdministrators"

	// MissingProviderType is returned when the project settings doesn't include a provider
	MissingProviderType AppErrorCode = "MissingProviderType"

	// InvalidProviderType is returned when the project settings doesn't include a valid provider
	InvalidProviderType AppErrorCode = "InvalidProviderType"

	// MissingProviderAccessKey is returned when a project settings doesn't include an access key
	MissingProviderAccessKey AppErrorCode = "MissingProviderAccessKey"

	// MissingProviderSecretKey is returned when a project settings doesn't include a secret key
	MissingProviderSecretKey AppErrorCode = "MissingProviderSecretKey"

	// InvalidKubernetesConfiguration is returned when a project kubernetes settings are not valid
	InvalidKubernetesConfiguration AppErrorCode = "InvalidKubernetesConfiguration"

	// InvalidProviderConfiguration is returned when a project provider settings are not valid
	InvalidProviderConfiguration AppErrorCode = "InvalidProviderConfiguration"

	// InvalidReplicaCount is returned when the service manifest doesn't include a valid number of replicas
	InvalidReplicaCount AppErrorCode = "InvalidReplicaCount"

	// ProjectNotEmpty is returned when a project still has active services
	ProjectNotEmpty AppErrorCode = "ProjectNotEmpty"

	// HasActiveProjects is returned when a user still has active projects
	HasActiveProjects AppErrorCode = "HasActiveProjects"

	// InsertFailed happens when an entity couldn't be saved
	InsertFailed AppErrorCode = "InsertFailed"

	// UpdateFailed happens when an entity couldn't be updated
	UpdateFailed AppErrorCode = "UpdateFailed"

	// DeleteFailed happens when an entity couldn't be deleted
	DeleteFailed AppErrorCode = "DeleteFailed"

	// EntityNotFound happens when an entity couldn't be found.
	EntityNotFound AppErrorCode = "EntityNotFound"

	// EntityForbidden happens when an entity is found, but the caller doesn't have permission to access it
	EntityForbidden AppErrorCode = "EntityForbidden"

	// InvalidServiceStatus happens when the service is not in a valid state for the action requested
	InvalidServiceStatus AppErrorCode = "InvalidServiceStatus"

	// MissingManifest happens when the service doesn't have a manifest
	MissingManifest AppErrorCode = "MissingManifest"

	// MissingProject is returned when the request is missing the project ID
	MissingProject AppErrorCode = "MissingProject"

	// MissingProjectSettings is returned when the request is missing the project ID
	MissingProjectSettings AppErrorCode = "MissingProjectSettings"

	// InvalidURL is returned when the request contains an invalid URL
	InvalidURL AppErrorCode = "InvalidURL"

	// InvalidEmail happens when the project settings have a malformed email
	InvalidEmail AppErrorCode = "InvalidEmail"

	// InvalidRole happens when the user invitation has an invalid role
	InvalidRole AppErrorCode = "InvalidRole"

	// MissingToken is returned when an api request doesn't include a token
	MissingToken AppErrorCode = "MissingToken"

	// InvalidToken is returned when the api token is not valid
	InvalidToken AppErrorCode = "InvalidToken"

	//FailToSendEmail is returned when the email provider was not able to send an email
	FailToSendEmail AppErrorCode = "FailToSendEmail"

	//InternalServerError is returned when we don't know what happened
	InternalServerError AppErrorCode = "InternalServerError"
)

//AppErrorCode is the type of error generated by an api call
type AppErrorCode string

//AppError is an error send back from an API call.
type AppError struct {
	Code    AppErrorCode `json:"code"`
	Status  int          `json:"status"`
	Message string       `json:"-"`
}

// ToError converts a to an error
func (a *AppError) ToError() error {
	return fmt.Errorf("%s - %s", a.Code, a.Message)
}