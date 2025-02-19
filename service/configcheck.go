// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtest"
	"go.opentelemetry.io/collector/consumer/consumererror"
)

// validateConfigFromFactories checks if all configurations for the given factories
// are satisfying the patterns used by the collector.
func validateConfigFromFactories(factories component.Factories) error {
	var errs []error

	for _, factory := range factories.Receivers {
		if err := configtest.CheckConfigStruct(factory.CreateDefaultConfig()); err != nil {
			errs = append(errs, err)
		}
	}
	for _, factory := range factories.Processors {
		if err := configtest.CheckConfigStruct(factory.CreateDefaultConfig()); err != nil {
			errs = append(errs, err)
		}
	}
	for _, factory := range factories.Exporters {
		if err := configtest.CheckConfigStruct(factory.CreateDefaultConfig()); err != nil {
			errs = append(errs, err)
		}
	}
	for _, factory := range factories.Extensions {
		if err := configtest.CheckConfigStruct(factory.CreateDefaultConfig()); err != nil {
			errs = append(errs, err)
		}
	}

	return consumererror.Combine(errs)
}
