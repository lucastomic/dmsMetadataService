package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/lucastomic/dmsMetadataService/internal/controller/apitypes"
	"github.com/lucastomic/dmsMetadataService/internal/errs"
	"github.com/lucastomic/dmsMetadataService/internal/logging"
	"github.com/lucastomic/dmsMetadataService/internal/metadataservice"
)

// MetadataController handles operations related to metadata management in the distributed file storage system.
// It utilizes a metadata service for processing metadata requests and a common controller for handling shared functionality.
type MetadataController struct {
	logger          logging.Logger                  // logger for logging messages throughout the metadata controller.
	metadataservice metadataservice.MetadataService // metadataservice provides operations for managing file metadata.
	common          CommonController                // common provides shared handler functionalities such as error parsing.
}

// NewMetadataController creates a new instance of MetadataController with the provided logger and metadata service.
// This function returns a Controller interface that abstracts the implementation details of metadata operations.
func NewMetadataController(
	logger logging.Logger,
	metadataservice metadataservice.MetadataService,
) Controller {
	return &MetadataController{
		logger,
		metadataservice,
		CommonController{},
	}
}

// Router defines routes for metadata operations and associates them with handler functions.
// Currently, it supports a single route for retrieving metadata about a file.
func (c *MetadataController) Router() apitypes.Router {
	return []apitypes.Route{
		{
			Path:    "/file",
			Method:  "GET",
			Handler: c.GetMetadata,
		},
	}
}

// GetMetadata handles the request for retrieving metadata of a specific file.
// It parses and validates the upload request, and if successful, retrieves the file metadata using the metadata service.
// Returns a response with the storage URL and file ID or an error response if any step fails.
func (c *MetadataController) GetMetadata(
	w http.ResponseWriter,
	req *http.Request,
) apitypes.Response {
	file, err := c.parseAndValidateUploadReq(req, w)
	if err != nil {
		return c.common.ParseError(req.Context(), req, w, err)
	}
	metadata, err := c.metadataservice.Get(req.Context(), &file)
	if err != nil {
		return c.common.ParseError(req.Context(), req, w, err)
	}
	return apitypes.Response{
		Status: http.StatusOK,
		Content: map[string]any{
			"storageURL": metadata.UploadURL,
			"id":         metadata.Id,
		},
	}
}

// parseAndValidateUploadReq validates the file upload request ensuring the file size does not exceed the maximum allowed size.
// It returns a pointer to the uploaded file and an error if the validation fails or no file is provided.
func (c *MetadataController) parseAndValidateUploadReq(
	req *http.Request,
	w http.ResponseWriter,
) (multipart.File, error) {
	const maxUploadSize = 10 << 20 // 10MB
	req.Body = http.MaxBytesReader(w, req.Body, maxUploadSize)

	if err := req.ParseMultipartForm(maxUploadSize); err != nil {
		return nil, fmt.Errorf(
			"%w:%s",
			errs.ErrInvalidInput,
			"The uploaded file is too big. Maximum file size is 10MB.",
		)
	}

	file, _, err := req.FormFile("uploadFile")
	if err != nil {
		return nil, fmt.Errorf(
			"%w:%s",
			errs.ErrInvalidInput,
			"No file provided in field 'uploadFile'",
		)
	}
	defer file.Close()

	return file, nil
}
