// Copyright 2023 Enterprise Reference Implementation contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0

package router

import (
	"encoding/json"

	rabbitinterfaces "github.com/dvonthenen/rabbitmq-manager/pkg/interfaces"
	prettyjson "github.com/hokaccha/go-prettyjson"
	klog "k8s.io/klog/v2"

	shared "github.com/dvonthenen/enterprise-reference-implementation/pkg/shared"
)

func NewQuestionHandler(options HandlerOptions) *rabbitinterfaces.RabbitMessageHandler {
	var handler rabbitinterfaces.RabbitMessageHandler
	handler = QuestionHandler{
		manager:  options.Manager,
		callback: options.Callback,
	}
	return &handler
}

func (qh QuestionHandler) ProcessMessage(byData []byte) error {
	// pretty print
	prettyJson, err := prettyjson.Format(byData)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}
	klog.V(6).Infof("\n\n-------------------------------\n")
	klog.V(2).Infof("QuestionHandler:\n%v\n", string(prettyJson))
	klog.V(6).Infof("-------------------------------\n\n")

	// reform struct
	var qr shared.QuestionResult
	err = json.Unmarshal(byData, &qr)
	if err != nil {
		klog.V(1).Infof("[QuestionHandler] json.Unmarshal failed. Err: %v\n", err)
		return err
	}

	// invoke callback
	err = (*qh.callback).QuestionResult(&qr)
	if err == nil {
		klog.V(5).Infof("[QuestionHandler] Callback succeeded\n")
	} else {
		klog.V(1).Infof("[QuestionHandler] Callback failed. Err: %v\n", err)
		return err
	}

	return nil
}