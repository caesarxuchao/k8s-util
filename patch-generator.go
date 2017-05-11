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

	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/util/strategicpatch"
)

//var podA = v1.ReplicationController{
//	ObjectMeta: v1.ObjectMeta{
//		//		Finalizers: []string{
//		//			"x/y",
//		//			"w/z",
//		//		},
//		OwnerReferences: []v1.OwnerReference{
//			{UID: "1", BlockOwnerDeletion: &trueVar},
//			{UID: "2"},
//			{UID: "3"},
//		},
//	},
//}

var nodeA = v1.Node{
	Status: v1.NodeStatus{
		VolumesInUse: []v1.UniqueVolumeName{
			"1",
			"2",
		},
	},
}

var nodeB = v1.Node{
	Status: v1.NodeStatus{
		VolumesInUse: []v1.UniqueVolumeName{
			"1",
			"3",
		},
	},
}
var trueVar = true

//var podB = v1.ReplicationController{
//	ObjectMeta: v1.ObjectMeta{
//		OwnerReferences: []v1.OwnerReference{
//			{UID: "1", BlockOwnerDeletion: &trueVar},
//			{UID: "2"},
//			{UID: "3"},
//		},
//	},
//}

func main() {
	original, err := json.Marshal(nodeA)
	if err != nil {
		panic(err)
		return
	}
	modified, _ := json.Marshal(nodeB)
	if err != nil {
		panic(err)
		return
	}
	patch, err := strategicpatch.CreateStrategicMergePatch(original, modified, v1.Node{})
	if err != nil {
		panic(err)
	}
	fmt.Println("patch=", string(patch))
	patched, err := strategicpatch.StrategicMergePatch(original, patch, v1.Node{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(patched))
}
