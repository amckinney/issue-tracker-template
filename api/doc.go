// Package api defines the types used in
// the Issue Tracker's external API.
//
// The types in this package are symmetrical
// to the internal entities, but they are
// intentionally decoupled so that the two
// can evolve independently of one another.
//
// This independent evolution is particularly
// important to support a wide range of API
// encodings; adding support for Protocol
// Buffers, Thrift, Avro, and others will
// only require changes to the http.Handler
// tier rather than requiring a total rewrite
// of your application.
package api
