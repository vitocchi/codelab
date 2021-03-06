#!/bin/bash

declare -a codelabs;
codelabs=(
    "gocon-2018-autumn-setup"   "1nBB9Sl-mnKE5htc_Fbx_-Q3VUY1YIfF_C3Wi_HFgzoo"
    "gocon-2019-spring-setup"   "1SBE8eGYPdnZVuMffOA0t4OT_6sxndTjNqmEOv-d61mA"
    "google-cloud-shell-go"     "1tmQBU8WtUIQ_Yb2STllFCa3lE8DQsMFnkikvcjMjF1c"
    "google-app-engine-go"      "1AjOewYLEtWy-x8RGQtxQe1vMCTKia15CK3iKczmfoQM"
    "google-cloud-functions-go" "1uwGLNnjLfFky0mvQuAKu7CIVTt6jO2rIMD4AkxrRUj4"
)

for ((i = 0; i < ${#codelabs[@]}; i+=2)) {
    echo "generate ${codelabs[i]}..."
    echo claat export -prefix="https://storage.googleapis.com" ${codelabs[i+1]}
    claat export -prefix="https://storage.googleapis.com" ${codelabs[i+1]}
}
