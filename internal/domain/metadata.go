package domain

// Metadata represents the metadata associated with a file in the distributed file storage system.
// It includes information necessary for storing and retrieving the file, such as its unique identifier
// and the URL to which the file should be uploaded.
type Metadata struct {
	// UploadURL specifies the URL where the file should be uploaded.
	// This URL points to a storage server where the file will be stored and can later be retrieved
	// using the unique identifier provided in the Id field.
	UploadURL string
	// Id is a unique identifier for the file within the distributed file storage system.
	// It is used to track and manage the file across the system and is typically generated
	// by the MetadataService when a new file is to be stored.
	Id int64
}
