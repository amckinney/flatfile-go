// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
	time "time"
)

type DeleteRecordsRequest struct {
	// The Record Ids param (ids) is a list of record ids that can be passed to several record endpoints allowing the user to identify specific records to INCLUDE in the query, or specific records to EXCLUDE, depending on whether or not filters are being applied. When passing a query param that filters the record dataset, such as 'searchValue', or a 'filter' of 'valid' | 'error' | 'all', the 'ids' param will EXCLUDE those records from the filtered results. For basic queries that do not filter the dataset, passing record ids in the 'ids' param will limit the dataset to INCLUDE just those specific records
	Ids []*RecordId `json:"-"`
}

type FindAndReplaceRecordRequest struct {
	Filter *Filter `json:"-"`
	// Name of field by which to filter records
	FilterField *FilterField `json:"-"`
	SearchValue *SearchValue `json:"-"`
	SearchField *SearchField `json:"-"`
	// The Record Ids param (ids) is a list of record ids that can be passed to several record endpoints allowing the user to identify specific records to INCLUDE in the query, or specific records to EXCLUDE, depending on whether or not filters are being applied. When passing a query param that filters the record dataset, such as 'searchValue', or a 'filter' of 'valid' | 'error' | 'all', the 'ids' param will EXCLUDE those records from the filtered results. For basic queries that do not filter the dataset, passing record ids in the 'ids' param will limit the dataset to INCLUDE just those specific records
	Ids []*RecordId `json:"-"`
	// A value to find for a given field in a sheet. Wrap the value in "" for exact match
	Find *CellValueUnion `json:"find,omitempty"`
	// The value to replace found values with
	Replace *CellValueUnion `json:"replace,omitempty"`
	// The value to replace found values with
	FieldKey string `json:"fieldKey"`
}

type FindAndReplaceRecordRequestDeprecated struct {
	// A unique key used to identify a field in a sheet
	FieldKey string `json:"-"`
	// A value to find for a given field in a sheet. Wrap the value in "" for exact match
	SearchValue string  `json:"-"`
	Filter      *Filter `json:"-"`
	// Number of records to return in a page (default 1000 if pageNumber included)
	PageSize *int `json:"-"`
	// Based on pageSize, which page of records to return
	PageNumber *int `json:"-"`
	// The value to replace found values with
	Replace interface{} `json:"replace,omitempty"`
}

type GetRecordsRequest struct {
	VersionId      *string        `json:"-"`
	SinceVersionId *VersionId     `json:"-"`
	SortField      *SortField     `json:"-"`
	SortDirection  *SortDirection `json:"-"`
	Filter         *Filter        `json:"-"`
	// Name of field by which to filter records
	FilterField *FilterField `json:"-"`
	SearchValue *SearchValue `json:"-"`
	SearchField *SearchField `json:"-"`
	// The Record Ids param (ids) is a list of record ids that can be passed to several record endpoints allowing the user to identify specific records to INCLUDE in the query, or specific records to EXCLUDE, depending on whether or not filters are being applied. When passing a query param that filters the record dataset, such as 'searchValue', or a 'filter' of 'valid' | 'error' | 'all', the 'ids' param will EXCLUDE those records from the filtered results. For basic queries that do not filter the dataset, passing record ids in the 'ids' param will limit the dataset to INCLUDE just those specific records. Maximum of 100 allowed.
	Ids []*RecordId `json:"-"`
	// Number of records to return in a page (default 1000 if pageNumber included)
	PageSize *int `json:"-"`
	// Based on pageSize, which page of records to return
	PageNumber *int `json:"-"`
	// **DEPRECATED** Use GET /sheets/:sheetId/counts
	IncludeCounts *bool `json:"-"`
	// The length of the record result set, returned as counts.total
	IncludeLength *bool `json:"-"`
	// If true, linked records will be included in the results. Defaults to false.
	IncludeLinks *bool `json:"-"`
	// Include error messages, defaults to false.
	IncludeMessages *bool `json:"-"`
	// if "for" is provided, the query parameters will be pulled from the event payload
	For *EventId `json:"-"`
	// An FFQL query used to filter the result set
	Q *string `json:"-"`
}

type CellValueUnion struct {
	typeName string
	String   string
	Integer  int
	Long     int64
	Double   float64
	Boolean  bool
	Date     time.Time
	DateTime time.Time
}

func NewCellValueUnionFromString(value string) *CellValueUnion {
	return &CellValueUnion{typeName: "string", String: value}
}

func NewCellValueUnionFromInteger(value int) *CellValueUnion {
	return &CellValueUnion{typeName: "integer", Integer: value}
}

func NewCellValueUnionFromLong(value int64) *CellValueUnion {
	return &CellValueUnion{typeName: "long", Long: value}
}

func NewCellValueUnionFromDouble(value float64) *CellValueUnion {
	return &CellValueUnion{typeName: "double", Double: value}
}

func NewCellValueUnionFromBoolean(value bool) *CellValueUnion {
	return &CellValueUnion{typeName: "boolean", Boolean: value}
}

func NewCellValueUnionFromDate(value time.Time) *CellValueUnion {
	return &CellValueUnion{typeName: "date", Date: value}
}

func NewCellValueUnionFromDateTime(value time.Time) *CellValueUnion {
	return &CellValueUnion{typeName: "dateTime", DateTime: value}
}

func (c *CellValueUnion) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typeName = "string"
		c.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		c.typeName = "integer"
		c.Integer = valueInteger
		return nil
	}
	var valueLong int64
	if err := json.Unmarshal(data, &valueLong); err == nil {
		c.typeName = "long"
		c.Long = valueLong
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		c.typeName = "double"
		c.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		c.typeName = "boolean"
		c.Boolean = valueBoolean
		return nil
	}
	var valueDate time.Time
	if err := json.Unmarshal(data, &valueDate); err == nil {
		c.typeName = "date"
		c.Date = valueDate
		return nil
	}
	var valueDateTime time.Time
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		c.typeName = "dateTime"
		c.DateTime = valueDateTime
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CellValueUnion) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return json.Marshal(c.String)
	case "integer":
		return json.Marshal(c.Integer)
	case "long":
		return json.Marshal(c.Long)
	case "double":
		return json.Marshal(c.Double)
	case "boolean":
		return json.Marshal(c.Boolean)
	case "date":
		return json.Marshal(c.Date)
	case "dateTime":
		return json.Marshal(c.DateTime)
	}
}

type CellValueUnionVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitLong(int64) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
	VisitDate(time.Time) error
	VisitDateTime(time.Time) error
}

func (c *CellValueUnion) Accept(visitor CellValueUnionVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return visitor.VisitString(c.String)
	case "integer":
		return visitor.VisitInteger(c.Integer)
	case "long":
		return visitor.VisitLong(c.Long)
	case "double":
		return visitor.VisitDouble(c.Double)
	case "boolean":
		return visitor.VisitBoolean(c.Boolean)
	case "date":
		return visitor.VisitDate(c.Date)
	case "dateTime":
		return visitor.VisitDateTime(c.DateTime)
	}
}

type GetRecordsResponse struct {
	Data *GetRecordsResponseData `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetRecordsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetRecordsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetRecordsResponse(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetRecordsResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}
