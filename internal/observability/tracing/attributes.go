// Copyright 2021-2022 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package tracing

import "go.opentelemetry.io/otel/attribute"

const (
	requestIDKey     = attribute.Key("cerbos.request.id")
	reqResourceIDKey = attribute.Key("cerbos.request.resource_id")
	policyFQNKey     = attribute.Key("cerbos.policy.fqn")
	policyNameKey    = attribute.Key("cerbos.policy.name")
	policyScopeKey   = attribute.Key("cerbos.policy.scope")
	policyVersionKey = attribute.Key("cerbos.policy.version")
)

var (
	RequestID     = requestIDKey.String
	ReqResourceID = reqResourceIDKey.String
	PolicyFQN     = policyFQNKey.String
	PolicyName    = policyNameKey.String
	PolicyScope   = policyScopeKey.String
	PolicyVersion = policyVersionKey.String
)
