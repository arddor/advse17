#!/usr/bin/env bash




#gcloud auth application-default login

#gcloud container clusters get-credentials awesome-cluster --zone europe-west1-c --project advse-167708

#kubectl proxy


############################## api

kubectl delete -f ./ase_api/api.yaml;
kubectl delete -f ./ase_api/api_service.yaml;



##################### compute

kubectl delete -f ./ase_compute/compute.yaml;



##################### replay

kubectl delete -f ./ase_load_replay/replay.yaml;



##################### tweets

kubectl delete -f ./ase_load_tweets/tweets.yaml;
kubectl delete -f ./ase_load_tweets/tweets_service.yaml;



##################### queue

kubectl delete -f ./ase_queue/queue.yaml;
kubectl delete -f ./ase_queue/queue_service.yaml;



############################## timeseries

kubectl delete -f ./ase_timeseries/timeseries.yaml
kubectl delete -f ./ase_timeseries/timeseries_service_admin.yaml
kubectl delete -f ./ase_timeseries/timeseries_service_db.yaml



##################### twitter

kubectl delete -f ./ase_twitter/twitter.yaml;



##################### web

kubectl delete -f ./ase_web/web.yaml;
kubectl delete -f ./ase_web/web_service.yaml;



