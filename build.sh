#!/bin/bash

set -e

if [ $# -gt 0 ]; then
    version="$1"
else
    current=$(date "+%Y-%m-%d %H:%M:%S")
    timeStamp=$(date -d "$current" +%s)
    currentTimeStamp=$(((timeStamp * 1000 + 10#$(date "+%N") / 1000000) / 1000))
    version="$currentTimeStamp"
fi
workdir=$(dirname $(realpath $0))

arr=("v1.21" "v1.22")

for value in ${arr[@]}; do
    k8s_ver=${value}

    folder_name="polaris-controller-release_kubernetes${k8s_ver}_${version}"
    pkg_name="${folder_name}.zip"

    cd $workdir

    # 清理环境
    rm -rf ${folder_name}
    rm -f "${pkg_name}"

    # 打包
    mkdir -p ${folder_name}
    cp -r deploy/${k8s_ver} ${folder_name}
    zip -r "${pkg_name}" ${folder_name}
    #md5sum ${pkg_name} > "${pkg_name}.md5sum"

    if [[ $(uname -a | grep "Darwin" | wc -l) -eq 1 ]]; then
        md5 ${pkg_name} >"${pkg_name}.md5sum"
    else
        md5sum ${pkg_name} >"${pkg_name}.md5sum"
    fi

done
