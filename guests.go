// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/FlatFilers/flatfile-go/core"
)

type ListGuestsRequest struct {
	// ID of space to return
	SpaceId SpaceId `json:"-"`
	// Email of guest to return
	Email *string `json:"-"`
}

// Guest ID
type GuestId = string

type CreateGuestResponse struct {
	Data []*Guest `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CreateGuestResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateGuestResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateGuestResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateGuestResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Configurations for the guests
type GuestConfig struct {
	EnvironmentId EnvironmentId `json:"environmentId"`
	Email         string        `json:"email"`
	Name          string        `json:"name"`
	Spaces        []*GuestSpace `json:"spaces,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GuestConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestConfig(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestConfig) String() string {
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

// Properties used to update an existing guest
type GuestConfigUpdate struct {
	EnvironmentId *EnvironmentId `json:"environmentId,omitempty"`
	Email         *string        `json:"email,omitempty"`
	Name          *string        `json:"name,omitempty"`
	Spaces        []*GuestSpace  `json:"spaces,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GuestConfigUpdate) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestConfigUpdate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestConfigUpdate(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestConfigUpdate) String() string {
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

type GuestResponse struct {
	Data *Guest `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GuestResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GuestResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GuestResponse(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GuestResponse) String() string {
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

type Invite struct {
	GuestId GuestId `json:"guestId"`
	SpaceId SpaceId `json:"spaceId"`
	// The name of the person or company sending the invitation
	FromName *string `json:"fromName,omitempty"`
	// Message to send with the invite
	Message *string `json:"message,omitempty"`

	_rawJSON json.RawMessage
}

func (i *Invite) UnmarshalJSON(data []byte) error {
	type unmarshaler Invite
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = Invite(value)
	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *Invite) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type ListGuestsResponse struct {
	Data []*Guest `json:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListGuestsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListGuestsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListGuestsResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListGuestsResponse) String() string {
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
