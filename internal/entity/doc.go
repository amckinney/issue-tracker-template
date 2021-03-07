// Package entity defines the internal
// entities used in the Issue Tracker
// application's business logic.
//
// The types in this package are symmetrical
// to the internal models, but they are
// intentionally decoupled so that the
// business logic objects do not take
// on persistence-level concerns, such
// as a database record's primary key.
//
// Similar to the justification for
// decoupling the API types from the
// entity types, the decoupling between
// the entity package and the model package
// makes it easier to swap out persistence
// technologies, such as PostgresSQL, MongoDB,
// and others.
package entity
