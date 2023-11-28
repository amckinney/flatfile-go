// This file was auto-generated by Fern from our API Definition.

package guests

import (
	context "context"
	fmt "fmt"
	flatfilego "github.com/FlatFilers/flatfile-go"
	core "github.com/FlatFilers/flatfile-go/core"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Returns all guests
func (c *Client) List(ctx context.Context, request *flatfilego.ListGuestsRequest) (*flatfilego.ListGuestsResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "guests"

	queryParams := make(url.Values)
	queryParams.Add("spaceId", fmt.Sprintf("%v", request.SpaceId))
	if request.Email != nil {
		queryParams.Add("email", fmt.Sprintf("%v", *request.Email))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.ListGuestsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Guests are only there to upload, edit, and download files and perform their tasks in a specific Space.
func (c *Client) Create(ctx context.Context, request []*flatfilego.GuestConfig) (*flatfilego.CreateGuestResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "guests"

	var response *flatfilego.CreateGuestResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns a single guest
//
// ID of guest to return
func (c *Client) Get(ctx context.Context, guestId flatfilego.GuestId) (*flatfilego.GuestResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"guests/%v", guestId)

	var response *flatfilego.GuestResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Deletes a single guest
//
// ID of guest to return
func (c *Client) Delete(ctx context.Context, guestId flatfilego.GuestId) (*flatfilego.Success, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"guests/%v", guestId)

	var response *flatfilego.Success
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodDelete,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Updates a single guest, for example to change name or email
//
// ID of guest to return
func (c *Client) Update(ctx context.Context, guestId flatfilego.GuestId, request *flatfilego.GuestConfigUpdate) (*flatfilego.GuestResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"guests/%v", guestId)

	var response *flatfilego.GuestResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPatch,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns a single guest token
//
// ID of guest to return
func (c *Client) GetGuestToken(ctx context.Context, guestId flatfilego.GuestId, request *flatfilego.GetGuestTokenRequest) (*flatfilego.GuestTokenResponse, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"guests/%v/token", guestId)

	queryParams := make(url.Values)
	if request.SpaceId != nil {
		queryParams.Add("spaceId", fmt.Sprintf("%v", request.SpaceId))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *flatfilego.GuestTokenResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Guests can be created as a named guest on the Space or there’s a global link that will let anonymous guests into the space.
func (c *Client) Invite(ctx context.Context, request []*flatfilego.Invite) (*flatfilego.Success, error) {
	baseURL := "https://api.x.flatfile.com/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "invitations"

	var response *flatfilego.Success
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
