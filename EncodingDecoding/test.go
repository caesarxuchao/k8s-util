package main

import (
	"k8s.io/kubernetes/pkg/api"
	_ "k8s.io/kubernetes/pkg/api/install"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/runtime"
)

func main() {
	pod := &api.Pod{}
	pod.Name = "hello"
	codec := api.Codecs.LegacyCodec(v1.SchemeGroupVersion)
	data, err := runtime.Encode(codec, pod)
	if err != nil {
		panic(err)
	}
	newPod := &api.Pod{}
	_, _, err = codec.Decode(data, nil, newPod)
	if err != nil {
		panic(err)
	}
}
