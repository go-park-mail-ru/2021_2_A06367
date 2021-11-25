// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB8a1138eDecodeGithubComGoParkMailRu20212A06367InternalModels(in *jlexer.Lexer, out *Actors) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "height":
			out.Height = float32(in.Float32())
		case "date_of_birth":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateOfBirth).UnmarshalJSON(data))
			}
		case "description":
			out.Description = string(in.String())
		case "genres":
			if in.IsNull() {
				in.Skip()
				out.Genres = nil
			} else {
				in.Delim('[')
				if out.Genres == nil {
					if !in.IsDelim(']') {
						out.Genres = make([]string, 0, 4)
					} else {
						out.Genres = []string{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Genres = append(out.Genres, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB8a1138eEncodeGithubComGoParkMailRu20212A06367InternalModels(out *jwriter.Writer, in Actors) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.Id).MarshalText())
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.Surname != "" {
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	if in.Avatar != "" {
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	if in.Height != 0 {
		const prefix string = ",\"height\":"
		out.RawString(prefix)
		out.Float32(float32(in.Height))
	}
	if true {
		const prefix string = ",\"date_of_birth\":"
		out.RawString(prefix)
		out.Raw((in.DateOfBirth).MarshalJSON())
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	if len(in.Genres) != 0 {
		const prefix string = ",\"genres\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Genres {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Actors) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB8a1138eEncodeGithubComGoParkMailRu20212A06367InternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Actors) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB8a1138eEncodeGithubComGoParkMailRu20212A06367InternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Actors) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB8a1138eDecodeGithubComGoParkMailRu20212A06367InternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Actors) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB8a1138eDecodeGithubComGoParkMailRu20212A06367InternalModels(l, v)
}
