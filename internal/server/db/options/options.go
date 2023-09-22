package options

type Option interface {
	Set(opt *Options)
}

type Options struct {
	Where map[string]interface{}
}

func ParseOptions(option ...Option) *Options {
	opts := &Options{}
	for _, opt := range option {
		opt.Set(opts)
	}
	return opts
}
