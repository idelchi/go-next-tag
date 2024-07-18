package main

import (
	"encoding/json"

	"github.com/showa-93/go-mask"
)

// PrintJSON returns a pretty-printed JSON representation of the provided object.
func PrintJSON(obj any) string {
	bytes, err := json.MarshalIndent(obj, "  ", "    ")
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// PrintJSONMasked returns a pretty-printed JSON string representation of the provided object with masked sensitive
// fields.
func PrintJSONMasked(obj any) string {
	return PrintJSON(JSONMasked(obj))
}

// JSONMasked returns a pretty-printed JSON representation of the provided object with masked sensitive fields.
func JSONMasked(obj any) any {
	masker := mask.NewMasker()

	masker.SetMaskChar("-")

	masker.RegisterMaskStringFunc(mask.MaskTypeFilled, masker.MaskFilledString)

	t, err := mask.Mask(obj)
	if err != nil {
		return err.Error()
	}

	return t
}
