/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/util/strategicpatch"
)

var podA = v1.ReplicationController{
	//	ObjectMeta: v1.ObjectMeta{
	//		// 		OwnerReferences: []v1.OwnerReference{
	//		// 			{UID: "1"},
	//		// 		},
	//		Labels: map[string]string{
	//			"x": "a",
	//		},
	//	},
	Spec: v1.ReplicationControllerSpec{
		Selector: map[string]string{
			"x": "a",
		},
		Template: &v1.PodTemplateSpec{
			ObjectMeta: v1.ObjectMeta{
				Labels: map[string]string{
					"x": "a",
				},
			},
		},
	},
}

var trueVar = true

var podB = v1.ReplicationController{
	//	ObjectMeta: v1.ObjectMeta{
	//		// 		OwnerReferences: []v1.OwnerReference{
	//		// 			{UID: "1"},
	//		// 		},
	//		Labels: map[string]string{
	//			"x": "a",
	//		},
	//	},
	Spec: v1.ReplicationControllerSpec{
		Selector: map[string]string{
			"x":         "a",
			"uniqueKey": "1",
		},
		Template: &v1.PodTemplateSpec{
			ObjectMeta: v1.ObjectMeta{
				Labels: map[string]string{
					"x":         "a",
					"uniqueKey": "1",
				},
			},
		},
	},
}

func main() {
	original, err := json.Marshal(podA)
	if err != nil {
		glog.Error(err)
		return
	}
	modified, _ := json.Marshal(podB)
	if err != nil {
		glog.Error(err)
		return
	}
	patch, err := strategicpatch.CreateStrategicMergePatch(original, modified, v1.ReplicationController{})
	if err != nil {
		glog.Error(err)
		return
	}
	fmt.Println("patch=", string(patch))
}
