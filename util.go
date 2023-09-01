package gobatis

import "time"

// PI64 to NullInt64
func PI64(i int64) *int64 {
	return &i
}

// PS to NullString
func PS(s string) *string {
	return &s
}

// PF64 to NullFloat64
func PF64(f float64) *float64 {
	return &f
}

// PT to NullTime
func PT(t time.Time) *time.Time {
	return &t
}

// NB to NullBool
func PB(b bool) *bool {
	return &b
}

// NI64 to NullInt64
func NI64(i int64) NullInt64 {
	return NullInt64{Int64: i, Valid: true}
}

// NS to NullString
func NS(s string) NullString {
	return NullString{String: s, Valid: true}
}

// NF64 to NullFloat64
func NF64(f float64) NullFloat64 {
	return NullFloat64{Float64: f, Valid: true}
}

// NT to NullTime
func NT(t time.Time) NullTime {
	return NullTime{Time: t, Valid: true}
}

// NB to NullBool
func NB(b bool) NullBool {
	return NullBool{Bool: b, Valid: true}
}
