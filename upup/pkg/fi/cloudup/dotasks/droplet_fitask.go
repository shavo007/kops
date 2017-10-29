/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ""fitask" -type=Droplet"; DO NOT EDIT

package dotasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// Droplet

// JSON marshalling boilerplate
type realDroplet Droplet

// UnmarshalJSON implements conversion to JSON, supporitng an alternate specification of the object as a string
func (o *Droplet) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.Name = &jsonName
		return nil
	}

	var r realDroplet
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = Droplet(r)
	return nil
}

var _ fi.HasLifecycle = &Droplet{}

// GetLifecycle returns the Lifecycle of the object, implementing fi.HasLifecycle
func (o *Droplet) GetLifecycle() *fi.Lifecycle {
	return o.Lifecycle
}

var _ fi.HasName = &Droplet{}

// GetName returns the Name of the object, implementing fi.HasName
func (o *Droplet) GetName() *string {
	return o.Name
}

// SetName sets the Name of the object, implementing fi.SetName
func (o *Droplet) SetName(name string) {
	o.Name = &name
}

// String is the stringer function for the task, producing readable output using fi.TaskAsString
func (o *Droplet) String() string {
	return fi.TaskAsString(o)
}