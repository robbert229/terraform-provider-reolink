package reolinkapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Camera
type Camera struct {
	// BaseURL is the url to connect with.
	BaseURL string
	// Username is the username to connect with.
	Username string
	// Password is the password to connect with.
	Password string
	// Transport is the http transport that is used when talking to the camera.
	Transport http.RoundTripper
	// DeferLogin indicates that the login process should be deferred.
	DeferLogin bool
	// Profile indicates if we should connect to the main, or sub camera profiles.
	Profile string
}

type cameraCfg struct {
	BaseURL    string
	Username   string
	Password   string
	DeferLogin bool
	Profile    string
	Transport  http.RoundTripper
}

func (v cameraCfg) validate() error {
	if v.Username == "" {
		return fmt.Errorf("no username given")
	}

	_, err := url.Parse(v.BaseURL)
	if err != nil {
		return fmt.Errorf("invalid base url given: %w", err)
	}

	return nil
}

// CameraOption is the type for options to the Camera constructor.
type CameraOption func(cfg *cameraCfg) error

// WithUsername configures the username to use to connect to the camera.
func WithUsername(username string) CameraOption {
	return func(cfg *cameraCfg) error {
		cfg.Username = username
		return nil
	}
}

// WithPassword configures the password to use to connect to the camera.
func WithPassword(password string) CameraOption {
	return func(cfg *cameraCfg) error {
		cfg.Password = password
		return nil
	}
}

// WithDeferLogin configures the camera connection to defer login.
func WithDeferLogin(deferLogin bool) CameraOption {
	return func(cfg *cameraCfg) error {
		cfg.DeferLogin = deferLogin
		return nil
	}
}

// WithURL configures the url that the camera can be found at.
func WithURL(url string) CameraOption {
	return func(cfg *cameraCfg) error {
		cfg.URL = url
		return nil
	}
}

// WithTransport configures the http.Transport that should be used to talk to
// the camera.
func WithTransport(transport http.RoundTripper) CameraOption {
	return func(cfg *cameraCfg) error {
		cfg.Transport = transport
		return nil
	}
}

// NewCamera creates a new camera object.
func NewCamera(opts ...CameraOption) (*Camera, error) {
	cfg := &cameraCfg{
		Transport: http.DefaultTransport,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	err := cfg.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid configuration: %+w", err)
	}

	return &Camera{
		BaseURL:    cfg.BaseURL,
		Username:   cfg.Username,
		Password:   cfg.Password,
		DeferLogin: cfg.DeferLogin,
		Profile:    cfg.Profile,
		Transport:  cfg.Transport,
	}, nil
}

type CommandBody[T any] struct {
	Cmd    string `json:"cmd"`
	Action int64  `json:"action"`
	Param  T      `json:"param"`
}

type CommandParam struct {
	Cmd   string `json:"cmd"`
	Token string `json:"token"`
}

type Command[T any] struct {
	Body  CommandBody[T]
	Param CommandParam
}

// GetURL returns the url of the api.
func (c *Camera) GetURL() (string, error) {
	return url.JoinPath(c.BaseURL, "/cgi-bin/api.cgi")
}

func GetURLWithParams[T any](baseURL string, params T) (string, error) {
	var values url.Values
	values.Add()
}

func (c *Camera) login(ctx context.Context) (string, error) {
	url, err := c.GetURL()
	if err != nil {
		return "", fmt.Errorf("")
	}

	urlWithParams, err := GetURLWithParams(url, CommandParam[nil]{})

	http.NewRequestWithContext(ctx, http.MethodPost, url, body)
}

func (c *Camera) Logout(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

// func (c *Camera) ExecuteCommand[T any](ctx context.Context, data CommandBody[T], multi bool) (interface{}, error) {

// }