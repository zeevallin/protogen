package protoc

import "fmt"

// ErrConfigType happens when the wrong configuration for a language is provided
type ErrConfigType struct {
	t interface{}
}

func (e ErrConfigType) Error() string {
	return fmt.Sprintf("protoc language config type invalid: %T", e.t)
}

// ErrUnknownLanguage happens when protoc does not recognise the provided language
type ErrUnknownLanguage struct {
	msg interface{}
}

func (e ErrUnknownLanguage) Error() string {
	return fmt.Sprintf("protoc does not support language: %v", e.msg)
}

// ErrFailedProtocExec happens when protoc could not execute
type ErrFailedProtocExec struct {
	msg interface{}
}

func (e ErrFailedProtocExec) Error() string {
	return fmt.Sprintf("protoc could not execute: %v", e.msg)
}

// ErrProtocMissing happens when the protoc binary does not exist
type ErrProtocMissing struct {
	msg interface{}
}

func (e ErrProtocMissing) Error() string {
	return fmt.Sprintf("protoc does not exist: %v", e.msg)
}

// ErrExtensionMissing happens when the protoc extension binary does not exist
type ErrExtensionMissing struct {
	lang string
	msg  interface{}
}

func (e ErrExtensionMissing) Error() string {
	return fmt.Sprintf("protoc-gen-%s does not exist: %v", e.lang, e.msg)
}
