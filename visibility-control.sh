set -o errexit
set -o nounset
set -o pipefail

base_path="$K1/staging/src/k8s.io/client-go/pkg/apis"
Groups="apps
extensions
authentication
autoscaling
certificates
authorization
settings
batch
storage
policy
rbac"

for group in $Groups; do
    path=$base_path/$group
    build_file=$path/BUILD
    if [ ! -e $build_file ]; then
        echo "$build_file not found"
        continue
    fi
    if grep -q " visibility =" $build_file; then
        echo "visibility rules already exist in $build_file"
        continue
    fi
    
    sed -i "s|deps = \[|\
visibility = \[\n\
        \"//vendor/k8s.io/client-go/pkg/api/v1:__pkg__\",\n\
        \"//vendor/k8s.io/client-go/pkg/apis:__subpackages__\",\n\
    ],\n\
    deps = \[|g" $build_file
done


# install package
for group in $Groups; do
    path=$base_path/$group
    build_file=$path/install/BUILD
    if [ ! -e $build_file ]; then
        echo "$build_file not found"
        continue
    fi
    if grep -q " visibility =" $build_file; then
        echo "visibility rules already exist in $build_file"
        continue
    fi
    
    sed -i "s|deps = \[|\
visibility = \[\"//visibility:private\"\],\n\
    deps = \[|g" $build_file
done
