#!/usr/bin/env bash

# https://cloud.google.com/container-registry/docs/managing


gcloud config set project advse-167708

gcloud config list

PROJECT_ID="$(gcloud config get-value project)"
COMMIT_HASH="$(git rev-parse HEAD)"



if [[ ${PROJECT_ID} == "advse-167708" ]]; then
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_api:latest ./ase_api/
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_compute:latest ./ase_compute/
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_twitter:latest ./ase_twitter/
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_web:latest ./ase_web/
    
    echo "##################### IMAGES #####################"
    docker images | grep "git_"
    echo "##################### BUILD COMPLETE, PUSHING... #####################"
    
    
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_api:latest
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_compute:latest
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_twitter:latest
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_web:latest

    
    echo "##################### PUSH COMPLETE #####################"

    gcloud container images list
    
fi

