// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson46cd36d5DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsLastWeekLastWeek1List) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsLastWeekLastWeek1List, 0, 8)
			} else {
				*out = GetFwLeaderboardsLastWeekLastWeek1List{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsLastWeekLastWeek1
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson46cd36d5EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsLastWeekLastWeek1List) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsLastWeekLastWeek1List) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson46cd36d5EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsLastWeekLastWeek1List) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson46cd36d5EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsLastWeekLastWeek1List) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson46cd36d5DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsLastWeekLastWeek1List) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson46cd36d5DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson46cd36d5DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsLastWeekLastWeek1) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "amount":
			out.Amount = int32(in.Int32())
		case "faction_id":
			out.FactionId = int32(in.Int32())
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
func easyjson46cd36d5EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsLastWeekLastWeek1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Amount))
	}
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsLastWeekLastWeek1) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson46cd36d5EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsLastWeekLastWeek1) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson46cd36d5EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsLastWeekLastWeek1) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson46cd36d5DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsLastWeekLastWeek1) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson46cd36d5DecodeGithubComAntihaxGoesiEsi1(l, v)
}