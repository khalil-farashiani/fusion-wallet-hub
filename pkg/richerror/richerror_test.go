package richerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	op := Op("initialize")
	err := New(op)

	assert.Equal(t, op, err.operation, "Operation should match")
	assert.Nil(t, err.wrappedError, "Wrapped error should be nil initially")
	assert.Empty(t, err.message, "Message should be empty initially")
	assert.Zero(t, err.kind, "Kind should be zero value initially")
	assert.Nil(t, err.meta, "Meta should be nil initially")
}

func TestWithOp(t *testing.T) {
	op := Op("start")
	newOp := Op("finish")
	err := New(op).WithOp(newOp)

	assert.Equal(t, newOp, err.operation, "Operation should be updated")
}

func TestWithError(t *testing.T) {
	baseErr := errors.New("base error")
	err := New(Op("action")).WithErr(baseErr)

	assert.Equal(t, baseErr, err.wrappedError, "Wrapped error should match the input error")
}

func TestWithMessage(t *testing.T) {
	message := "an error occurred"
	err := New(Op("log")).WithMessage(message)

	assert.Equal(t, message, err.message, "Message should match the input message")
}

func TestWithKind(t *testing.T) {
	kind := NotFound
	err := New(Op("delete")).WithKind(kind)

	assert.Equal(t, kind, err.kind, "Kind should match the input kind")
}

func TestWithMeta(t *testing.T) {
	meta := map[string]interface{}{"key": "value"}
	err := New(Op("update")).WithMeta(meta)

	assert.Equal(t, meta, err.meta, "Meta should match the input meta")
}

func TestErrorMethod(t *testing.T) {
	errWithMessage := New(Op("compute")).WithMessage("failure")
	assert.Equal(t, "failure", errWithMessage.Error(), "Error method should return the message when set")

	wrappedErr := errors.New("internal error")
	errWithWrappedErr := New(Op("fetch")).WithErr(wrappedErr)
	assert.Equal(t, "internal error", errWithWrappedErr.Error(), "Error method should return the wrapped error message when message is empty")
}

func TestKindMethod(t *testing.T) {
	errWithKind := New(Op("read")).WithKind(Forbidden)
	assert.Equal(t, Forbidden, errWithKind.Kind(), "Kind method should return the kind when set")

	innerErr := New(Op("inner")).WithKind(Invalid)
	wrappedErr := New(Op("outer")).WithErr(innerErr)
	assert.Equal(t, Invalid, wrappedErr.Kind(), "Kind method should return the kind from the wrapped error when not set on the outer error")
}

func TestMessageMethod(t *testing.T) {
	message := "operation failed"
	errWithMessage := New(Op("execute")).WithMessage(message)
	assert.Equal(t, message, errWithMessage.Message(), "Message method should return the message when set")

	innerErr := New(Op("inner")).WithMessage("inner failure")
	wrappedErr := New(Op("outer")).WithErr(innerErr)
	assert.Equal(t, "inner failure", wrappedErr.Message(), "Message method should return the message from the wrapped error when message is not set on the outer error")

	baseErr := errors.New("base error")
	errWithWrappedErr := New(Op("outer")).WithErr(baseErr)
	assert.Equal(t, "base error", errWithWrappedErr.Message(), "Message method should return the wrapped error's message when no message is set")
}
