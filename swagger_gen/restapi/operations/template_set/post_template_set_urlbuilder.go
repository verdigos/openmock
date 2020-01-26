// Code generated by go-swagger; DO NOT EDIT.

package template_set

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PostTemplateSetURL generates an URL for the post template set operation
type PostTemplateSetURL struct {
	SetKey string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostTemplateSetURL) WithBasePath(bp string) *PostTemplateSetURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostTemplateSetURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PostTemplateSetURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/template_sets/{setKey}"

	setKey := o.SetKey
	if setKey != "" {
		_path = strings.Replace(_path, "{setKey}", setKey, -1)
	} else {
		return nil, errors.New("setKey is required on PostTemplateSetURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PostTemplateSetURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PostTemplateSetURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PostTemplateSetURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostTemplateSetURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostTemplateSetURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PostTemplateSetURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
