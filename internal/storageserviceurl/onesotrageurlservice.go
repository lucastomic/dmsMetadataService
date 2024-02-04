package storageserviceurl

import (
	"context"
	"mime/multipart"

	"github.com/lucastomic/dmsMetadataService/internal/environment"
	"github.com/lucastomic/dmsMetadataService/internal/logging"
)

// storageURLService implements the StorageURLService interface, providing logic to determine
// the appropriate URL for file storage. It utilizes a logger for logging operations
// and potential errors encountered during the process of finding a storage URL.
type storageURLService struct {
	logger logging.Logger // logger is used for logging information and errors.
}

// New initializes a new instance of storageURLService with the provided logger.
// This function returns a StorageURLService, ready for use in finding storage URLs.
func New(l logging.Logger) StorageURLService {
	return &storageURLService{l}
}

// FindStorageURL calculates and returns the storage URL for a given file. In its current implementation,
// it returns a predefined URL from the environment configuration without performing any dynamic calculations
// based on the file.
func (s *storageURLService) FindStorageURL(
	ctx context.Context,
	file multipart.File,
) (string, error) {
	return environment.GetStorageServiceURL(), nil
}
