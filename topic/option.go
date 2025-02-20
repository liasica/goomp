// Copyright (C) goomp. 2025-present.
//
// Created at 2025-02-20, by liasica

package topic

type Options struct {
	page int
}

type Option interface {
	apply(*Options)
}

type optionFunc func(*Options)

func (f optionFunc) apply(o *Options) {
	f(o)
}

func WithPage(page int) Option {
	return optionFunc(func(o *Options) {
		o.page = page
	})
}
