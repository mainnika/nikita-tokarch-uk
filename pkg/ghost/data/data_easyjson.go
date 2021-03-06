// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package data

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	template "html/template"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData(in *jlexer.Lexer, out *Posts) {
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
		case "posts":
			if in.IsNull() {
				in.Skip()
				out.Posts = nil
			} else {
				in.Delim('[')
				if out.Posts == nil {
					if !in.IsDelim(']') {
						out.Posts = make([]Post, 0, 0)
					} else {
						out.Posts = []Post{}
					}
				} else {
					out.Posts = (out.Posts)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Post
					(v1).UnmarshalEasyJSON(in)
					out.Posts = append(out.Posts, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "meta":
			(out.Meta).UnmarshalEasyJSON(in)
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
func easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData(out *jwriter.Writer, in Posts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"posts\":"
		out.RawString(prefix[1:])
		if in.Posts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Posts {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"meta\":"
		out.RawString(prefix)
		(in.Meta).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Posts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Posts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData(l, v)
}
func easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData1(in *jlexer.Lexer, out *Post) {
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
			out.ID = string(in.String())
		case "uuid":
			out.UUID = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "html":
			out.HTML = template.HTML(in.String())
		case "feature_image":
			out.FImage = template.URL(in.String())
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
func easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData1(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix)
		out.String(string(in.UUID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"html\":"
		out.RawString(prefix)
		out.String(string(in.HTML))
	}
	{
		const prefix string = ",\"feature_image\":"
		out.RawString(prefix)
		out.String(string(in.FImage))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Post) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Post) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData1(l, v)
}
func easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData2(in *jlexer.Lexer, out *Pagination) {
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
		case "page":
			out.Page = int(in.Int())
		case "limit":
			out.Limit = int(in.Int())
		case "pages":
			out.Pages = int(in.Int())
		case "total":
			out.Total = int(in.Int())
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
func easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData2(out *jwriter.Writer, in Pagination) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"page\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Page))
	}
	{
		const prefix string = ",\"limit\":"
		out.RawString(prefix)
		out.Int(int(in.Limit))
	}
	{
		const prefix string = ",\"pages\":"
		out.RawString(prefix)
		out.Int(int(in.Pages))
	}
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix)
		out.Int(int(in.Total))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Pagination) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Pagination) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData2(l, v)
}
func easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData3(in *jlexer.Lexer, out *Pages) {
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
		case "pages":
			if in.IsNull() {
				in.Skip()
				out.Pages = nil
			} else {
				in.Delim('[')
				if out.Pages == nil {
					if !in.IsDelim(']') {
						out.Pages = make([]Post, 0, 0)
					} else {
						out.Pages = []Post{}
					}
				} else {
					out.Pages = (out.Pages)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Post
					(v4).UnmarshalEasyJSON(in)
					out.Pages = append(out.Pages, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "meta":
			(out.Meta).UnmarshalEasyJSON(in)
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
func easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData3(out *jwriter.Writer, in Pages) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pages\":"
		out.RawString(prefix[1:])
		if in.Pages == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Pages {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"meta\":"
		out.RawString(prefix)
		(in.Meta).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Pages) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Pages) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData3(l, v)
}
func easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData4(in *jlexer.Lexer, out *Meta) {
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
		case "pagination":
			(out.Pagination).UnmarshalEasyJSON(in)
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
func easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData4(out *jwriter.Writer, in Meta) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pagination\":"
		out.RawString(prefix[1:])
		(in.Pagination).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Meta) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData4(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Meta) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeCodeTokarchUkMainnikaNikitaTokarchUkFrontendGhostData4(l, v)
}
