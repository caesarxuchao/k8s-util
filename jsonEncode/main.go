package main

import (
	"encoding/json"
	"fmt"

	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/api/v1"
)

func main() {
	listJson, err := json.Marshal(v1.PodList{TypeMeta: unversioned.TypeMeta{Kind: "podlist"}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", listJson)
}
