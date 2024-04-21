package ring

import (
	"io"
	"net/http"

	"github.com/pasztorpisti/qs"
    "github.com/mitchellh/mapstructure"
)

type (
    // MultipartFormBinder parses the request body as "multipart/form-data",
    // and binds it to the provided model using the \`multipart:""\` struct
    // tag.
    // 
    // See [mapstructure.Decoder] for more decoding details.
    MultipartFormBinder[T any] struct{
        MaxMemory int64
    }

    // URLEncodedFormBinder parses the request body as "application/x-www-formurlencoded",
    // and binds it to the provided model using the \`qs:""\` struct
    // tag.
    //
    // See [qs.QSUnmarshaler] for more decoding details
    URLEncodedFormBinder[T any] struct {
        Options *qs.UnmarshalOptions
    }
)

// Bind implements Binder.
func (binder *MultipartFormBinder[T]) Bind(r *http.Request, model T) error {
    if err := r.ParseMultipartForm(binder.MaxMemory); err != nil {
        return err
    }
    
    reader, err := r.MultipartReader()
    if err != nil {
        return err
    }

    form, err := reader.ReadForm(binder.MaxMemory)
    if err != nil {
        return err
    }

    dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
        TagName: "multipart",
        Result: model,
    })

    if err != nil {
        return err
    }

    if err := dec.Decode(form.Value); err != nil {
        return err
    }

    if err := dec.Decode(form.File); err != nil {
        return err
    }

    return nil
}

// Bind implements Binder.
func (binder *URLEncodedFormBinder[T]) Bind(r *http.Request, model T) error {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        return err
    }

    dec := qs.NewUnmarshaler(binder.Options)
    err = dec.Unmarshal(model, string(body))
    if err != nil {
        return err
    }

    return nil
}

var (
    _ Binder[any] = (*MultipartFormBinder[any])(nil)
    _ Binder[any] = (*URLEncodedFormBinder[any])(nil)
)
