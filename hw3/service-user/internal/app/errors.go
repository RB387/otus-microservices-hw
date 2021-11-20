package app

import "errors"

var ValidationError = errors.New("invalid request, validation failed")
var NotFound = errors.New("page not found")
