// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"sync/atomic"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/controller/network"
	gardenerkubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
)

var (
	// StatusTypeMeta is the TypeMeta of Cilium Status
	StatusTypeMeta = metav1.TypeMeta{
		APIVersion: ciliumv1alpha1.SchemeGroupVersion.String(),
		Kind:       "NetworkStatus",
	}
)

type actuator struct {
	restConfig *rest.Config
	client     client.Client

	chartRendererFactory extensionscontroller.ChartRendererFactory
	chartApplier         gardenerkubernetes.ChartApplier

	atomicShootWebhookConfig *atomic.Value
	webhookServerPort        int
}

// NewActuator creates a new Actuator that updates the status of the handled Network resources.
func NewActuator(mgr manager.Manager, chartApplier gardenerkubernetes.ChartApplier, chartRendererFactory extensionscontroller.ChartRendererFactory, shootWebhookConfig *atomic.Value, webhookServerPort int) network.Actuator {
	return &actuator{
		client:                   mgr.GetClient(),
		restConfig:               mgr.GetConfig(),
		chartApplier:             chartApplier,
		chartRendererFactory:     chartRendererFactory,
		atomicShootWebhookConfig: shootWebhookConfig,
		webhookServerPort:        webhookServerPort,
	}
}
