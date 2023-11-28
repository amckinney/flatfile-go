// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
)

type ListJobsRequest struct {
	EnvironmentId *EnvironmentId `json:"-"`
	SpaceId       *SpaceId       `json:"-"`
	WorkbookId    *WorkbookId    `json:"-"`
	FileId        *FileId        `json:"-"`
	ParentId      *JobId         `json:"-"`
	// Number of jobs to return in a page (default 20)
	PageSize *int `json:"-"`
	// Based on pageSize, which page of jobs to return
	PageNumber    *int           `json:"-"`
	SortDirection *SortDirection `json:"-"`
}

// Pipeline Job ID
type JobId = string

// Details about the user who acknowledged the job
type JobAckDetails struct {
	Info                  *string    `json:"info,omitempty"`
	Progress              *int       `json:"progress,omitempty"`
	EstimatedCompletionAt *time.Time `json:"estimatedCompletionAt,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobAckDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler JobAckDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobAckDetails(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobAckDetails) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// Info about the reason the job was canceled
type JobCancelDetails struct {
	Info *string `json:"info,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobCancelDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler JobCancelDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobCancelDetails(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobCancelDetails) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// Outcome summary of a job
type JobCompleteDetails struct {
	Outcome *JobOutcome `json:"outcome,omitempty"`
	Info    *string     `json:"info,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobCompleteDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler JobCompleteDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobCompleteDetails(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobCompleteDetails) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// A single unit of work that a pipeline will execute
type JobConfig struct {
	// The type of job
	Type JobType `json:"type,omitempty"`
	// the type of operation to perform on the data. For example, "export".
	Operation   string           `json:"operation"`
	Source      JobSource        `json:"source"`
	Destination *JobDestination  `json:"destination,omitempty"`
	Config      *JobUpdateConfig `json:"config,omitempty"`
	// the type of trigger to use for this job
	Trigger *Trigger `json:"trigger,omitempty"`
	// the status of the job
	Status *JobStatus `json:"status,omitempty"`
	// the progress of the job
	Progress *float64 `json:"progress,omitempty"`
	FileId   *FileId  `json:"fileId,omitempty"`
	// the mode of the job
	Mode *JobMode `json:"mode,omitempty"`
	// Input parameters for this job type.
	Input map[string]interface{} `json:"input,omitempty"`
	// Subject parameters for this job type.
	Subject *JobSubject `json:"subject,omitempty"`
	// Outcome summary of job.
	Outcome map[string]interface{} `json:"outcome,omitempty"`
	// Current status of job in text
	Info *string `json:"info,omitempty"`
	// Indicates if Flatfile is managing the control flow of this job or if it is being manually tracked.
	Managed *bool `json:"managed,omitempty"`
	// The id of the environment this job belongs to
	EnvironmentId *EnvironmentId `json:"environmentId,omitempty"`
	// The part number of this job
	Part *int `json:"part,omitempty"`
	// The data for this part of the job
	PartData map[string]interface{} `json:"partData,omitempty"`
	// The execution mode for this part of the job
	PartExecution *JobPartExecution `json:"partExecution,omitempty"`
	// The id of the parent job
	ParentId *JobId `json:"parentId,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler JobConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobConfig(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobConfig) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobExecutionPlanConfigRequest struct {
	FieldMapping              []*Edge             `json:"fieldMapping,omitempty"`
	UnmappedSourceFields      []*SourceField      `json:"unmappedSourceFields,omitempty"`
	UnmappedDestinationFields []*DestinationField `json:"unmappedDestinationFields,omitempty"`
	FileId                    FileId              `json:"fileId"`
	JobId                     JobId               `json:"jobId"`

	_rawJSON json.RawMessage
}

func (j *JobExecutionPlanConfigRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler JobExecutionPlanConfigRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobExecutionPlanConfigRequest(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobExecutionPlanConfigRequest) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobExecutionPlanRequest struct {
	FieldMapping              []*Edge             `json:"fieldMapping,omitempty"`
	UnmappedSourceFields      []*SourceField      `json:"unmappedSourceFields,omitempty"`
	UnmappedDestinationFields []*DestinationField `json:"unmappedDestinationFields,omitempty"`
	FileId                    FileId              `json:"fileId"`
	JobId                     JobId               `json:"jobId"`

	_rawJSON json.RawMessage
}

func (j *JobExecutionPlanRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler JobExecutionPlanRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobExecutionPlanRequest(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobExecutionPlanRequest) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobPlanResponse struct {
	Data *JobPlan `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobPlanResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler JobPlanResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobPlanResponse(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobPlanResponse) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobResponse struct {
	Data *Job `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler JobResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobResponse(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobResponse) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// Info about the reason the job was split
type JobSplitDetails struct {
	Parts         *JobParts `json:"parts,omitempty"`
	RunInParallel *bool     `json:"runInParallel,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobSplitDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler JobSplitDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobSplitDetails(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobSplitDetails) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

// A single unit of work that will be executed
type JobUpdate struct {
	Config *JobUpdateConfig `json:"config,omitempty"`
	// the status of the job
	Status *JobStatus `json:"status,omitempty"`
	// the progress of the job
	Progress *float64 `json:"progress,omitempty"`
	// the time that the job's outcome has been acknowledged by a user
	OutcomeAcknowledgedAt *time.Time `json:"outcomeAcknowledgedAt,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobUpdate) UnmarshalJSON(data []byte) error {
	type unmarshaler JobUpdate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobUpdate(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobUpdate) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type ListJobsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       []*Job      `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListJobsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListJobsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListJobsResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListJobsResponse) String() string {
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
