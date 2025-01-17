// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetes

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccKubernetesDataSourceNamespace_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKubernetesDataSourceNamespaceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.kubernetes_namespace.test", "metadata.0.name", "kube-system"),
					resource.TestCheckResourceAttrSet("data.kubernetes_namespace.test", "metadata.0.generation"),
					resource.TestCheckResourceAttrSet("data.kubernetes_namespace.test", "metadata.0.resource_version"),
					resource.TestCheckResourceAttrSet("data.kubernetes_namespace.test", "metadata.0.uid"),
					resource.TestCheckResourceAttr("data.kubernetes_namespace.test", "spec.#", "1"),
					resource.TestCheckResourceAttr("data.kubernetes_namespace.test", "spec.0.finalizers.#", "1"),
					resource.TestCheckResourceAttr("data.kubernetes_namespace.test", "spec.0.finalizers.0", "kubernetes"),
				),
			},
		},
	})
}

func testAccKubernetesDataSourceNamespaceConfig_basic() string {
	return `
data "kubernetes_namespace" "test" {
  metadata {
    name = "kube-system"
  }
}
`
}
