package options

type WhereOption struct {
	key   string
	value interface{}
}

func (o *WhereOption) Set(opts *Options) {
	if opts.Where == nil {
		opts.Where = make(map[string]interface{})
	}
	opts.Where[o.key] = o.value
}

func Where(key string, value interface{}) *WhereOption {
	return &WhereOption{
		key:   key,
		value: value,
	}
}

func WhereId(value interface{}) *WhereOption {
	return &WhereOption{
		key:   "id",
		value: value,
	}
}

func WhereName(value interface{}) *WhereOption {
	return &WhereOption{
		key:   "name",
		value: value,
	}
}
