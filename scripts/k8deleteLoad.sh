#!/usr/bin/env bash




#gcloud auth application-default login

#gcloud container clusters get-credentials awesome-cluster --zone europe-west1-c --project advse-167708

#kubectl proxy


##################### replay

kubectl delete -f ./ase_load_replay/replay.yaml;



##################### tweets

kubectl delete -f ./ase_load_tweets/tweets.yaml;
kubectl delete -f ./ase_load_tweets/tweets_service.yaml;


