package gosubmit

import (
	"regexp"
)

type Input interface {
	Name() string
	Type() string
	Value() string
	Options() []string
	Fill(val string) (value string, ok bool)
	Required() bool
	Multiple() bool
}

type anyInput struct {
	name      string
	inputType string
	value     string
	required  bool
	multiple  bool
}

func (i anyInput) Name() string {
	return i.name
}

func (i anyInput) Type() string {
	return i.inputType
}

func (i anyInput) Value() string {
	return i.value
}

func (i anyInput) Required() bool {
	return i.required
}

func (i anyInput) Multiple() bool {
	return i.multiple
}

func (i anyInput) Options() (values []string) {
	return
}

type FileInput struct {
	anyInput
}

func (f FileInput) Fill(val string) (value string, ok bool) {
	return "", false
}

type TextInput struct {
	anyInput
	validator *regexp.Regexp
}

func (i TextInput) Fill(val string) (value string, ok bool) {
	ok = true
	value = val
	if i.validator == nil {
		return
	}
	ok = i.validator.MatchString(value)
	return
}

type inputWithOptions struct {
	anyInput
	options []string
}

func (i inputWithOptions) Options() []string {
	return i.options
}

func (i inputWithOptions) Fill(val string) (value string, ok bool) {
	ok = false
	value = val
	for _, opt := range i.options {
		if opt == value {
			ok = true
			break
		}
	}
	return
}

type Checkbox struct {
	inputWithOptions
}

type Radio struct {
	inputWithOptions
}

type Select struct {
	inputWithOptions
}

type Button struct {
	Name  string
	Value string
}
