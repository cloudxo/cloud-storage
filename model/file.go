package model

import (
	"io"
)

// ID to which one needs it
type ID struct {
	ID string `json:"id,omitempty"`
}

// FileModel for filestore service
type FileModel struct {
	ParentID      string    `json:"parentID,omitempty"`
	SourcesID     string    `json:"sourcesID,omitempty"`
	DestinationID string    `json:"destinationID,omitempty"`
	Source        string    `json:"source,omitempty"`
	Sources       []string  `json:"sources,omitempty"`
	Destination   string    `json:"destination,omitempty"`
	Destinations  []string  `json:"destinations,omitempty"`
	Name          string    `json:"name,omitempty"`
	MimeType      string    `json:"mimeType,omitempty"`
	Path          string    `json:"path,omitempty"`
	Content       io.Reader `json:"content,omitempty"`
	Query         string    `json:"query,omitempty"`
}

// OneDriveItem for OneDrive item object
type OneDriveItem struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// ListOneDriveItem for OneDrive list item object
type ListOneDriveItem struct {
	Value []OneDriveItem `json:"value,omitempty"`
}

// CreateOneDriveFolder for create folder
type CreateOneDriveFolder struct {
	Name                           string `json:"name,omitempty"`
	MicrosoftGraphConflictBehavior string `json:"@microsoft.graph.conflictBehavior,omitempty"`
}

type MoveOneDriveItem struct {
	Name            string `json:"name,omitempty"`
	ParentReference ID     `json:"parentReference,omitempty"`
}
