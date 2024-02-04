package storageserviceurl

import (
	"context"
	"mime/multipart"
)

// StorageURLService defines the interface for a service that finds and returns storage URLs for file uploads.
type StorageURLService interface {
	// FindStorageURL determines the best storage server URL for uploading a given file.
	// It takes a context for handling request cancellation and a multipart.File representing the file to be uploaded.
	// Returns the URL of the storage server where the file should be uploaded, or an error if no suitable URL can be found
	// or if there is a problem accessing the necessary information to make that determination.
	FindStorageURL(context.Context, multipart.File) (string, error)
}
