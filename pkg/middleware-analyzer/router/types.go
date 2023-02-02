// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package router

import (
	rabbitinterfaces "github.com/dvonthenen/rabbitmq-manager/pkg/interfaces"

	interfaces "github.com/dvonthenen/enterprise-reference-implementation/pkg/middleware-analyzer/interfaces"
)

/*
	Subscriber handlers
*/
type HandlerOptions struct {
	Manager  *rabbitinterfaces.Manager
	Callback *interfaces.InsightCallback
}

type ConversationInitHandler struct {
	manager  *rabbitinterfaces.Manager
	callback *interfaces.InsightCallback

	appMessage *rabbitinterfaces.Publisher
}

type EntityHandler struct {
	manager  *rabbitinterfaces.Manager
	callback *interfaces.InsightCallback
}

type InsightHandler struct {
	manager  *rabbitinterfaces.Manager
	callback *interfaces.InsightCallback
}

type MessageHandler struct {
	manager  *rabbitinterfaces.Manager
	callback *interfaces.InsightCallback
}

type TopicHandler struct {
	manager  *rabbitinterfaces.Manager
	callback *interfaces.InsightCallback
}

type TrackerHandler struct {
	manager  *rabbitinterfaces.Manager
	callback *interfaces.InsightCallback
}

type ConversationTeardownHandler struct {
	manager  *rabbitinterfaces.Manager
	callback *interfaces.InsightCallback
}