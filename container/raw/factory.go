// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package raw

import (
	"log"

	"github.com/google/cadvisor/container"
)

type rawFactory struct {
}

func (self *rawFactory) String() string {
	return "raw"
}

func (self *rawFactory) NewContainerHandler(name string) (container.ContainerHandler, error) {
	return newRawContainerHandler(name)
}

// The raw factory can handle any container.
func (self *rawFactory) CanHandle(name string) bool {
	return true
}

func Register() error {
	log.Printf("Registering Raw factory")
	container.RegisterContainerHandlerFactory(new(rawFactory))
	return nil
}