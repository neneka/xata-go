// This file was auto-generated by Fern from our API Definition.

package api

type AskTableResponse struct {
	// The answer to the input question
	Answer string `json:"answer"`
	// The session ID for the chat session.
	SessionId string `json:"sessionId"`
}