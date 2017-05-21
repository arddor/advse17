#!/usr/bin/env bash

# https://cloud.google.com/container-registry/docs/managing


gcloud config set project advse-167708

gcloud config list

PROJECT_ID="$(gcloud config get-value project)"
COMMIT_HASH="$(git rev-parse HEAD)"



if [[ ${PROJECT_ID} == "advse-167708" ]]; then
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_api:latest -f ase_api_Dockerfile .
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_compute:latest -f ase_compute_Dockerfile .
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_load_replay:latest ./ase_load_replay/
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_load_tweets:latest ./ase_load_tweets/
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_twitter:latest -f ase_twitter_Dockerfile .
    docker build -t eu.gcr.io/${PROJECT_ID}/ase_web:latest ./ase_web/
    
    echo "##################### IMAGES #####################"
    docker images | grep "ase_"
    echo "##################### BUILD COMPLETE, PUSHING... #####################"
    
    
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_api:latest
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_compute:latest
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_load_replay:latest
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_load_tweets:latest
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_twitter:latest
    gcloud docker -- push eu.gcr.io/${PROJECT_ID}/ase_web:latest

    
    echo "##################### PUSH COMPLETE #####################"

    gcloud container images list
    
fi

