// This file was auto-generated by Fern from our API Definition.

package client

import (
	agents "github.com/FlatFilers/flatfile-go/agents"
	cells "github.com/FlatFilers/flatfile-go/cells"
	commits "github.com/FlatFilers/flatfile-go/commits"
	core "github.com/FlatFilers/flatfile-go/core"
	documents "github.com/FlatFilers/flatfile-go/documents"
	environments "github.com/FlatFilers/flatfile-go/environments"
	events "github.com/FlatFilers/flatfile-go/events"
	files "github.com/FlatFilers/flatfile-go/files"
	guests "github.com/FlatFilers/flatfile-go/guests"
	jobs "github.com/FlatFilers/flatfile-go/jobs"
	records "github.com/FlatFilers/flatfile-go/records"
	roles "github.com/FlatFilers/flatfile-go/roles"
	secrets "github.com/FlatFilers/flatfile-go/secrets"
	sheets "github.com/FlatFilers/flatfile-go/sheets"
	snapshots "github.com/FlatFilers/flatfile-go/snapshots"
	spaces "github.com/FlatFilers/flatfile-go/spaces"
	users "github.com/FlatFilers/flatfile-go/users"
	versions "github.com/FlatFilers/flatfile-go/versions"
	workbooks "github.com/FlatFilers/flatfile-go/workbooks"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Agents       *agents.Client
	Cells        *cells.Client
	Commits      *commits.Client
	Documents    *documents.Client
	Environments *environments.Client
	Events       *events.Client
	Files        *files.Client
	Guests       *guests.Client
	Jobs         *jobs.Client
	Records      *records.Client
	Roles        *roles.Client
	Secrets      *secrets.Client
	Sheets       *sheets.Client
	Snapshots    *snapshots.Client
	Spaces       *spaces.Client
	Users        *users.Client
	Versions     *versions.Client
	Workbooks    *workbooks.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:      options.BaseURL,
		caller:       core.NewCaller(options.HTTPClient),
		header:       options.ToHeader(),
		Agents:       agents.NewClient(opts...),
		Cells:        cells.NewClient(opts...),
		Commits:      commits.NewClient(opts...),
		Documents:    documents.NewClient(opts...),
		Environments: environments.NewClient(opts...),
		Events:       events.NewClient(opts...),
		Files:        files.NewClient(opts...),
		Guests:       guests.NewClient(opts...),
		Jobs:         jobs.NewClient(opts...),
		Records:      records.NewClient(opts...),
		Roles:        roles.NewClient(opts...),
		Secrets:      secrets.NewClient(opts...),
		Sheets:       sheets.NewClient(opts...),
		Snapshots:    snapshots.NewClient(opts...),
		Spaces:       spaces.NewClient(opts...),
		Users:        users.NewClient(opts...),
		Versions:     versions.NewClient(opts...),
		Workbooks:    workbooks.NewClient(opts...),
	}
}
