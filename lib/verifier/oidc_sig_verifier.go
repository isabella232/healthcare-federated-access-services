// Copyright 2020 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package verifier

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes" /* copybara-comment */
	"google.golang.org/grpc/status" /* copybara-comment */
	"gopkg.in/square/go-jose.v2/jwt" /* copybara-comment */
	"github.com/coreos/go-oidc" /* copybara-comment */
	"github.com/GoogleCloudPlatform/healthcare-federated-access-services/lib/errutil" /* copybara-comment: errutil */
	"github.com/GoogleCloudPlatform/healthcare-federated-access-services/lib/ga4gh" /* copybara-comment: ga4gh */
)

type oidcSigVerifier struct {
	issuer   string
	verifier *oidc.IDTokenVerifier
}

// newOIDCSigVerifier creates a new oidc token extractClaimsAndVerifySignature.
func newOIDCSigVerifier(ctx context.Context, issuer string) (*oidcSigVerifier, error) {
	p, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, errutil.WithErrorReason(errCreateVerifierFailed, status.Errorf(codes.Unavailable, "create oidc failed: %v", err))
	}

	v := p.Verifier(&oidc.Config{
		// Skip client claims check if no client claims passed in.
		SkipClientIDCheck: true,
		// Expire check and issuer check will do explicitly.
		SkipExpiryCheck: true,
		SkipIssuerCheck: true,
	})

	return &oidcSigVerifier{
		issuer:   issuer,
		verifier: v,
	}, nil
}

func (s *oidcSigVerifier) ExtractClaims(ctx context.Context, token string, claims interface{}) (*ga4gh.StdClaims, error) {
	// extracts the unsafe claims here to allow following step to validate issuer, timestamp.
	d, err := ga4gh.NewStdClaimsFromJWT(token)
	if err != nil {
		return nil, errutil.WithErrorReason(errParseFailed, status.Errorf(codes.Unauthenticated, "NewStdClaimsFromJWT() failed: %v", err))
	}

	if claims != nil {
		if err := unsafeClaimsFromJWTToken(token, claims); err != nil {
			return nil, err
		}
	}

	return d, nil
}

func (s *oidcSigVerifier) VerifySig(ctx context.Context, token string) error {
	// ensure token does not include a jku header
	tok, err := jwt.ParseSigned(token)
	if err != nil {
		return fmt.Errorf("parse jwt failed: %v", err)
	}

	if len(tok.Headers) != 1 {
		return fmt.Errorf("not single header jwt")
	}

	if _, ok := tok.Headers[0].ExtraHeaders["jku"]; ok {
		return fmt.Errorf("token should not have jku header")
	}

	_, err = s.verifier.Verify(ctx, token)
	return err
}

func (s *oidcSigVerifier) Issuer() string {
	return s.issuer
}
