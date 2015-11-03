package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestErrorInvalidBody(c *gin.Context) error {
	return RestError(c, "invalid request body")
}

func RestErrorNotFound(c *gin.Context) error {
	return RestError(c, "not found")
}

func RestErrorAppNotFound(c *gin.Context) error {
	return RestError(c, "app not found")
}

func RestErrorUnauthorized(c *gin.Context) error {
	return RestError(c, "not authorized")
}

func RestError(c *gin.Context, err interface{}) error {
	status := http.StatusInternalServerError

	var msg string
	switch t := err.(type) {
	case error:
		msg = t.Error()
	case string:
		msg = t
	}

	getError := func() error {
		return errors.New(msg)
	}

	if msg == "not found" || msg == "app not found" {
		status = http.StatusNotFound
	} else if msg == "invalid request body" {
		status = http.StatusBadRequest
	} else if msg == "not authorized" {
		status = http.StatusUnauthorized
	} else if msg == "ns not found" {
		//mongo throws this error when a collection does not exist and we call drop
		return getError()
	}

	requestError := getError()
	c.AbortWithError(status, requestError)

	return requestError
}
