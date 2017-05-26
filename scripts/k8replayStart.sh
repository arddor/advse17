#!/usr/bin/env bash




#gcloud auth application-default login

#gcloud container clusters get-credentials awesome-cluster --zone europe-west1-c --project advse-167708

#kubectl proxy


##################### replay

kubectl create --save-config -f ./ase_load_replay/replay.yaml;



##################### tweets

kubectl create --save-config -f ./ase_load_tweets/tweets.yaml;
kubectl create --save-config -f ./ase_load_tweets/tweets_service.yaml;


