//
// Copyright (c) 2020-2020 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//
package devfile_registry

import "github.com/eclipse-che/che-operator/pkg/deploy"

func init() {
	err := deploy.InitTestDefaultsFromDeployment("../../../deploy/operator.yaml")
	if err != nil {
		panic(err)
	}
}
