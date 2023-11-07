// SPDX-License-Identifier: Apache-2.0

// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

type AuthorizationCodeResponse struct {
	State       *string       `json:"state,omitempty"`
	RedirectUri *string       `json:"redirectUri,omitempty"`
	Scopes      *[]OAuthScope `json:"scopes,omitempty"`
	ClientId    *string       `json:"clientId,omitempty"`
	Expires     *time.Time    `json:"expires,omitempty"`
	Code        *string       `json:"code,omitempty"`
}
