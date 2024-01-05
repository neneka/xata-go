// SPDX-License-Identifier: Apache-2.0

// This file was auto-generated by Fern from our API Definition.

package api

// Object representing a file in an array
type InputFileEntry struct {
	// Base64 encoded content
	Base64Content *string `json:"base64Content,omitempty"`
	// Enable public access to the file
	EnablePublicUrl *bool       `json:"enablePublicUrl,omitempty"`
	Id              *FileItemId `json:"id,omitempty"`
	MediaType       *MediaType  `json:"mediaType,omitempty"`
	Name            *FileName   `json:"name,omitempty"`
	// Time to live for signed URLs
	SignedUrlTimeout *int `json:"signedUrlTimeout,omitempty"`
	// Time to live for upload URLs
	UploadUrlTimeout *int `json:"uploadUrlTimeout,omitempty"`
}
