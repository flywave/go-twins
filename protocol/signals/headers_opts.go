package signals

import "github.com/flywave/go-twins/protocol"

type HeaderOpt func(headers *protocol.Headers) error

func applyOptsHeader(headers *protocol.Headers, opts ...HeaderOpt) error {
	for _, o := range opts {
		if err := o(headers); err != nil {
			return err
		}
	}
	return nil
}

func NewHeaders(opts ...HeaderOpt) *protocol.Headers {
	res := &protocol.Headers{}
	res.Values = make(map[string]interface{})
	if err := applyOptsHeader(res, opts...); err != nil {
		return nil
	}
	return res
}

func NewHeadersFrom(orig *protocol.Headers, opts ...HeaderOpt) *protocol.Headers {
	if orig == nil {
		return NewHeaders(opts...)
	}
	res := &protocol.Headers{
		Values: make(map[string]interface{}),
	}
	for key, value := range orig.Values {
		res.Values[key] = value
	}
	if err := applyOptsHeader(res, opts...); err != nil {
		return nil
	}
	return res
}

func WithCorrelationId(correlationId string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderCorrelationId] = correlationId
		return nil
	}
}

func WithReplyTo(replyTo string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderReplyTo] = replyTo
		return nil
	}
}

func WithReplyTarget(replyTarget string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderReplyTarget] = replyTarget
		return nil
	}
}

func WithChannel(channel string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderChannel] = channel
		return nil
	}
}

func WithResponseRequired(isResponseRequired bool) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderResponseRequired] = isResponseRequired
		return nil
	}
}

func WithOriginator(dittoOriginator string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderOriginator] = dittoOriginator
		return nil
	}
}

func WithOrigin(origin string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderOrigin] = origin
		return nil
	}
}

func WithDryRun(isDryRun bool) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderDryRun] = isDryRun
		return nil
	}
}

func WithETag(eTag string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderETag] = eTag
		return nil
	}
}

func WithIfMatch(ifMatch string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderIfMatch] = ifMatch
		return nil
	}
}

func WithIfNoneMatch(ifNoneMatch string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderIfNoneMatch] = ifNoneMatch
		return nil
	}
}

func WithTimeout(timeout string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderTimeout] = timeout
		return nil
	}
}

func WithSchemaVersion(schemaVersion string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderSchemaVersion] = schemaVersion
		return nil
	}
}

func WithContentType(contentType string) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[protocol.HeaderContentType] = contentType
		return nil
	}
}

func WithGeneric(headerId string, value interface{}) HeaderOpt {
	return func(headers *protocol.Headers) error {
		headers.Values[headerId] = value
		return nil
	}
}
