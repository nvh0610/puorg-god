package kafka_manager

import (
	"crawler_go_facebook/app/entities/kafka_msg"
	"encoding/json"
)

type Serializer interface {
	Encode(interface{}) []byte
	Decode([]byte) interface{}
}

type DefaultSerializer struct {
}

func (s *DefaultSerializer) Encode(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func (s *DefaultSerializer) Decode(v []byte) interface{} {
	var x interface{}
	_ = json.Unmarshal(v, &x)
	return x
}

type ProfileSerializer struct {
}

func (s *ProfileSerializer) Encode(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func (s *ProfileSerializer) Decode(v []byte) interface{} {
	var x kafka_msg.ProfileMsg
	_ = json.Unmarshal(v, &x)
	return x
}

type CommentSerializer struct {
}

func (s *CommentSerializer) Encode(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func (s *CommentSerializer) Decode(v []byte) interface{} {
	var x kafka_msg.CommentMsg
	_ = json.Unmarshal(v, &x)
	return x
}

type UserLikedSerializer struct {
}

func (s *UserLikedSerializer) Encode(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func (s *UserLikedSerializer) Decode(v []byte) interface{} {
	var x kafka_msg.UserLikedMsg
	_ = json.Unmarshal(v, &x)
	return x
}

type UrlToIdSerializer struct {
}

func (s *UrlToIdSerializer) Encode(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func (s *UrlToIdSerializer) Decode(v []byte) interface{} {
	var x kafka_msg.UrlToIdMsg
	_ = json.Unmarshal(v, &x)
	return x
}

type ReactorSharerSerializer struct {
}

func (s *ReactorSharerSerializer) Encode(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func (s *ReactorSharerSerializer) Decode(v []byte) interface{} {
	var x kafka_msg.ReactorSharerMsg
	_ = json.Unmarshal(v, &x)
	return x
}

type PostSerializer struct {
}

func (s *PostSerializer) Encode(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func (s *PostSerializer) Decode(v []byte) interface{} {
	var x kafka_msg.PostMsg
	_ = json.Unmarshal(v, &x)
	return x
}
