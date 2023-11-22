// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
)

type ListFilesRequest struct {
	SpaceId *string `json:"-"`
	// Number of jobs to return in a page (default 20)
	PageSize *int `json:"-"`
	// Based on pageSize, which page of jobs to return
	PageNumber *int `json:"-"`
	// The storage mode of file to fetch, defaults to "import"
	Mode *Mode `json:"-"`
}

type Action struct {
	Slug             *string             `json:"slug,omitempty"`
	Operation        *string             `json:"operation,omitempty"`
	Mode             *ActionMode         `json:"mode,omitempty"`
	Label            string              `json:"label"`
	Tooltip          *string             `json:"tooltip,omitempty"`
	Type             *string             `json:"type,omitempty"`
	Description      *string             `json:"description,omitempty"`
	Schedule         *ActionSchedule     `json:"schedule,omitempty"`
	Primary          *bool               `json:"primary,omitempty"`
	Confirm          *bool               `json:"confirm,omitempty"`
	Icon             *string             `json:"icon,omitempty"`
	RequireAllValid  *bool               `json:"requireAllValid,omitempty"`
	RequireSelection *bool               `json:"requireSelection,omitempty"`
	InputForm        *InputForm          `json:"inputForm,omitempty"`
	Constraints      []*ActionConstraint `json:"constraints,omitempty"`

	_rawJSON json.RawMessage
}

func (a *Action) UnmarshalJSON(data []byte) error {
	type unmarshaler Action
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = Action(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *Action) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type FileResponse struct {
	Data *File `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (f *FileResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler FileResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FileResponse(value)
	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *FileResponse) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type ListFilesResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       []*File     `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListFilesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFilesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFilesResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFilesResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type Mode string

const (
	ModeImport Mode = "import"
	ModeExport Mode = "export"
)

func NewModeFromString(s string) (Mode, error) {
	switch s {
	case "import":
		return ModeImport, nil
	case "export":
		return ModeExport, nil
	}
	var t Mode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (m Mode) Ptr() *Mode {
	return &m
}

type ModelFileStatusEnum string

const (
	ModelFileStatusEnumPartial  ModelFileStatusEnum = "partial"
	ModelFileStatusEnumComplete ModelFileStatusEnum = "complete"
	ModelFileStatusEnumArchived ModelFileStatusEnum = "archived"
	ModelFileStatusEnumPurged   ModelFileStatusEnum = "purged"
	ModelFileStatusEnumFailed   ModelFileStatusEnum = "failed"
)

func NewModelFileStatusEnumFromString(s string) (ModelFileStatusEnum, error) {
	switch s {
	case "partial":
		return ModelFileStatusEnumPartial, nil
	case "complete":
		return ModelFileStatusEnumComplete, nil
	case "archived":
		return ModelFileStatusEnumArchived, nil
	case "purged":
		return ModelFileStatusEnumPurged, nil
	case "failed":
		return ModelFileStatusEnumFailed, nil
	}
	var t ModelFileStatusEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (m ModelFileStatusEnum) Ptr() *ModelFileStatusEnum {
	return &m
}

type UpdateFileRequest struct {
	WorkbookId *WorkbookId `json:"workbookId,omitempty"`
	// The name of the file
	Name *string `json:"name,omitempty"`
	// The storage mode of file to update
	Mode *Mode `json:"mode,omitempty"`
	// Status of the file
	Status *ModelFileStatusEnum `json:"status,omitempty"`
	// The actions attached to the file
	Actions []*Action `json:"actions,omitempty"`
}

type CreateFileRequest struct {
	SpaceId       SpaceId       `json:"spaceId"`
	EnvironmentId EnvironmentId `json:"environmentId"`
	// The storage mode of file to insert, defaults to "import"
	Mode *Mode `json:"mode,omitempty"`
	// The actions attached to the file
	Actions []*Action `json:"actions,omitempty"`
}
