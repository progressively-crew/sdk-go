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
	boolFlag map[string]bool
	flags    map[string]interface{}
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

	flags := GetFlags(url)
	sdk := &Sdk{boolFlag: make(map[string]bool), flags: flags}

	return sdk
}

func (sdk *Sdk) Evaluate(flagKey string) interface{} {
	return sdk.flags[flagKey]
}
