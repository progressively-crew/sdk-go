package progressively

import (
	"encoding/base64"
	"encoding/json"
)

type Builder struct {
	fields map[string]interface{}
	apiUrl string
}

type Sdk struct {
	url   string
	flags map[string]interface{}
}

func SdkBuilder(clientKey string, apiUrl string) *Builder {
	builder := &Builder{fields: make(map[string]interface{}), apiUrl: apiUrl}

	builder.fields["clientKey"] = clientKey

	return builder
}

func (b *Builder) AddField(key string, value interface{}) *Builder {
	b.fields[key] = value

	return b
}

func (b *Builder) Build() *Sdk {
	data, err := json.Marshal(b.fields)

	if err != nil {
		panic(err)
	}

	b64 := base64.StdEncoding.EncodeToString(data)
	url := b.apiUrl + "/sdk/" + b64

	sdk := &Sdk{flags: make(map[string]interface{}), url: url}
	sdk.LoadFlags()

	return sdk
}

func (sdk *Sdk) LoadFlags() {
	flags := GetFlags(sdk.url)
	sdk.flags = flags
}

func (sdk *Sdk) Evaluate(flagKey string) interface{} {
	return sdk.flags[flagKey]
}
