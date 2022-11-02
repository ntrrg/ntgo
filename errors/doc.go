// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

// Package errors provides utilities for error handling. This package aims to
// improve error context information and readability.
//
// Some good rules when creating error messages are:
//
// * Use a unique code. Helps for scoping and improves error comparison.
//
// * Let clients (GUIs, services, etc...) handle user-friendly messages and
// translations.
//
// * Be concise. Error descriptions are for developers, use plain English;
// codes are for machines, use simple symbols for maximizing compatibility.
//
// * Wrap errors. The most context the better, returning single errors makes
// visual debugging opaque.
//
// Since Go errors are flexible and doesn't enforce any structure, following
// this rules is a developer choice. Enforcing an error message syntax reduces
// ambiguity, which helps both humans and machines.
//
// Regularly errors should be defined at package initialization, as exported
// identifiers. Exported errors allow users to know exactly what went wrong by
// using utilities like errors.Is, errors.As, Of, All, Any and many others.
//
// Parse, MustParse and Error.New ensure that errors aligns to the syntax
// enforced by this package. New allows creation of errors without syntax
// enforcement, therefore it should be used only for very specific cases.
//
// # Error syntax
//
//	error     = "[" code "] " reason [ ": " wrapped ] .
//	code      = code_text { [ "/" ] code_text } .
//	reason    = unicode_value | byte_value .
//	wrapped   = error | ( unicode_value | byte_value ) .
//	code_text = code_char { code_char } .
//	code_char = "a" … "z" | "0" … "9" | "_" | "-" | "." .
//
// ## Examples
//
// Simple error:
//
//	[500] internal server error
//
// Scoped error:
//
//	[storage/tx/done] transaction has already been committed or rolled back
//
// Wrapped error:
//
//	[net/http] can't start server: listen tcp :80: bind: address already in use
package errors

// API Status: stable
