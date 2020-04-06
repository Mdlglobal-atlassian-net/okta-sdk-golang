/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"fmt"
	"github.com/okta/okta-sdk-golang/okta/query"
	"time"
)

type UserFactor interface {
	IsUserFactorInstance() bool
}

type FactorResource resource

type Factor struct {
	Embedded                  interface{}          `json:"_embedded,omitempty"`
	Links                     interface{}          `json:"_links,omitempty"`
	Created                   *time.Time           `json:"created,omitempty"`
	Device                    string               `json:"device,omitempty"`
	DeviceType                string               `json:"deviceType,omitempty"`
	FactorProfileId           string               `json:"factorProfileId,omitempty"`
	FactorType                string               `json:"factorType,omitempty"`
	Id                        string               `json:"id,omitempty"`
	LastUpdated               *time.Time           `json:"lastUpdated,omitempty"`
	MfaStateTokenId           string               `json:"mfaStateTokenId,omitempty"`
	Profile                   *UserFactorProfile   `json:"profile,omitempty"`
	Provider                  string               `json:"provider,omitempty"`
	RechallengeExistingFactor *bool                `json:"rechallengeExistingFactor,omitempty"`
	SessionId                 string               `json:"sessionId,omitempty"`
	Status                    string               `json:"status,omitempty"`
	TokenLifetimeSeconds      int64                `json:"tokenLifetimeSeconds,omitempty"`
	UserId                    string               `json:"userId,omitempty"`
	Verify                    *VerifyFactorRequest `json:"verify,omitempty"`
}

func NewFactor() *Factor {
	return &Factor{}
}

func (a *Factor) IsUserFactorInstance() bool {
	return true
}

// Unenrolls an existing factor for the specified user, allowing the user to enroll a new factor.
func (m *FactorResource) DeleteFactor(userId string, factorId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates all the enrolled factors for the specified user
func (m *FactorResource) ListFactors(userId string) ([]UserFactor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor []UserFactor

	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}

	return factor, resp, nil
}

// Enrolls a user with a supported factor.
func (m *FactorResource) EnrollFactor(userId string, body UserFactor, qp *query.Params) (UserFactor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	factor := body

	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}

	return factor, resp, nil
}

// Enumerates all the supported factors that can be enrolled for the specified user
func (m *FactorResource) ListSupportedFactors(userId string) ([]UserFactor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/catalog", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor []UserFactor

	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}

	return factor, resp, nil
}

// Enumerates all available security questions for a user&#x27;s &#x60;question&#x60; factor
func (m *FactorResource) ListSupportedSecurityQuestions(userId string) ([]*SecurityQuestion, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/questions", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var securityQuestion []*SecurityQuestion

	resp, err := m.client.requestExecutor.Do(req, &securityQuestion)
	if err != nil {
		return nil, resp, err
	}

	return securityQuestion, resp, nil
}

// Fetches a factor for the specified user
func (m *FactorResource) GetFactor(userId string, factorId string, factorInstance UserFactor) (UserFactor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	factor := factorInstance

	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}

	return factor, resp, nil
}

// The &#x60;sms&#x60; and &#x60;token:software:totp&#x60; factor types require activation to complete the enrollment process.
func (m *FactorResource) ActivateFactor(userId string, factorId string, body VerifyFactorRequest, factorInstance UserFactor) (UserFactor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/lifecycle/activate", userId, factorId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	factor := factorInstance

	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}

	return factor, resp, nil
}

// Polls factors verification transaction for status.
func (m *FactorResource) GetFactorTransactionStatus(userId string, factorId string, transactionId string) (*VerifyUserFactorResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/transactions/%v", userId, factorId, transactionId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var verifyUserFactorResponse *VerifyUserFactorResponse

	resp, err := m.client.requestExecutor.Do(req, &verifyUserFactorResponse)
	if err != nil {
		return nil, resp, err
	}

	return verifyUserFactorResponse, resp, nil
}

// Verifies an OTP for a &#x60;token&#x60; or &#x60;token:hardware&#x60; factor
func (m *FactorResource) VerifyFactor(userId string, factorId string, body VerifyFactorRequest, factorInstance UserFactor, qp *query.Params) (*VerifyUserFactorResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/verify", userId, factorId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var verifyUserFactorResponse *VerifyUserFactorResponse

	resp, err := m.client.requestExecutor.Do(req, &verifyUserFactorResponse)
	if err != nil {
		return nil, resp, err
	}

	return verifyUserFactorResponse, resp, nil
}
