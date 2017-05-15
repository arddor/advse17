#!/usr/bin/env bash

# https://cloud.google.com/container-registry/docs/managing


gcloud config set project advse-167708

gcloud config list

PROJECT_ID="$(gcloud config get-value project)"
COMMIT_HASH="$(git rev-parse HEAD)"



if [[ ${PROJECT_ID} == "advse-167708" ]]; then
    docker build -t gcr.io/${PROJECT_ID}/ase_api:git_${COMMIT_HASH} ./ase_api/
    docker build -t gcr.io/${PROJECT_ID}/ase_compute:git_${COMMIT_HASH} ./ase_compute/
    docker build -t gcr.io/${PROJECT_ID}/ase_twitter:git_${COMMIT_HASH} ./ase_twitter/
    docker build -t gcr.io/${PROJECT_ID}/ase_web:git_${COMMIT_HASH} ./ase_web/
    
    echo "##################### IMAGES #####################"
    docker images | grep "git_"
    echo "##################### BUILD COMPLETE, PUSHING... #####################"
    
    
    gcloud docker -- push gcr.io/${PROJECT_ID}/ase_api:git_${COMMIT_HASH}
    gcloud docker -- push gcr.io/${PROJECT_ID}/ase_compute:git_${COMMIT_HASH}
    gcloud docker -- push gcr.io/${PROJECT_ID}/ase_twitter:git_${COMMIT_HASH}
    gcloud docker -- push gcr.io/${PROJECT_ID}/ase_web:git_${COMMIT_HASH}

    
    echo "##################### PUSH COMPLETE #####################"

    gcloud container images list

    
    
    
    
    
    
fi

