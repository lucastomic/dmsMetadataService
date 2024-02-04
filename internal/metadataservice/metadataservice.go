package metadataservice

import (
	"context"
	"mime/multipart"

	"github.com/lucastomic/dmsMetadataService/internal/domain"
	"github.com/lucastomic/dmsMetadataService/internal/idgenerator"
	"github.com/lucastomic/dmsMetadataService/internal/logging"
	"github.com/lucastomic/dmsMetadataService/internal/storageserviceurl"
)

// MetadataService defines an interface for services capable of retrieving metadata for files.
// It abstracts the logic necessary to generate metadata, including generating unique identifiers
// for files and determining their appropriate storage URLs.
type MetadataService interface {
	// Get generates and returns metadata for a given file. It uses an ID generator to assign a unique
	// identifier to the file and a storage URL service to determine where the file should be uploaded.
	Get(context.Context, *multipart.File) (domain.Metadata, error)
}

// metadataService implements the MetadataService interface, providing functionality
// to generate and retrieve metadata for files. It utilizes a logger for error logging,
// a storage URL service to determine storage locations, and an ID generator for assigning
// unique identifiers to files.
type metadataService struct {
	logger            logging.Logger                      // logger facilitates logging throughout the metadata service.
	storageURLService storageserviceurl.StorageURLService // storageURLService is used to find storage URLs for files.
	idgenerator       idgenerator.IDGenerator[int64]      // idgenerator is responsible for generating unique IDs for files.
}

// New initializes a new instance of metadataService with the specified logger, storage URL service,
// and ID generator. This function returns a MetadataService, ready to generate and retrieve file metadata.
func New(
	logger logging.Logger,
	storageURLService storageserviceurl.StorageURLService,
	idgenerator idgenerator.IDGenerator[int64],
) MetadataService {
	return &metadataService{
		logger, storageURLService, idgenerator,
	}
}

// Get generates metadata for a given file, including a unique identifier and a storage URL.
// It leverages the injected ID generator to create a unique ID for the file and queries
// the storage URL service to determine the appropriate upload URL.
//
// If the process encounters any errors, such as failing to retrieve a storage URL, it logs
// the error and returns it to the caller. Otherwise, it returns the generated metadata
// containing the file ID and storage URL.
func (m metadataService) Get(ctx context.Context, file *multipart.File) (domain.Metadata, error) {
	id := m.idgenerator.GenerateID()
	uploadUrl, err := m.storageURLService.FindStorageURL(ctx, *file)
	if err != nil {
		m.logger.Error(ctx, "Error retrieving storage URL: %s", err)
		return domain.Metadata{}, err
	}
	return domain.Metadata{
		Id:        id,
		UploadURL: uploadUrl,
	}, nil
}
