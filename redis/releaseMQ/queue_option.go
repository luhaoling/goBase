package queue

// Option 定义了一个函数闭包实现的选项模式，用于生成具有不同选项的 option 结构，其中包含了可配置的字段 topic 和 handler

type Option func(*Options)

type Options struct {
	topic   string
	handler handlerFunc
}

func WithTopic(topic string) Option {
	return func(opts *Options) {
		opts.topic = topic
	}
}

func WithHandler(handler handlerFunc) Option {
	return func(opts *Options) {
		opts.handler = handler
	}
}
