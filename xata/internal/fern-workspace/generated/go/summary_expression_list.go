// This file was auto-generated by Fern from our API Definition.

package api

// The description of the summaries you wish to receive. Set each key to be the field name
// you'd like for the summary. These names must not collide with other columns you've
// requested from `columns`; including implicit requests like `settings.*`.
//
// The value for each key needs to be an object. This object should contain one key and one
// value only. In this object, the key should be set to the summary function you wish to use
// and the value set to the column name to be summarized.
//
// The column being summarized cannot be an internal column (id, xata.*), nor the base of
// an object, i.e. if `settings` is an object with `dark_mode` as a field, you may summarize
// `settings.dark_mode` but not `settings` nor `settings.*`.
type SummaryExpressionList = map[string]SummaryExpression
