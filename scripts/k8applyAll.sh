#!/usr/bin/env bash




#gcloud auth application-default login

#gcloud container clusters get-credentials awesome-cluster --zone europe-west1-c --project advse-167708

#kubectl proxy


############################## api

kubectl apply -f ./ase_api/api.yaml;
kubectl apply -f ./ase_api/api_service.yaml;



##################### compute

kubectl apply -f ./ase_compute/compute.yaml;



##################### queue

kubectl apply -f ./ase_queue/queue.yaml;
kubectl apply -f ./ase_queue/queue_service.yaml;



############################## timeseries

kubectl apply -f ./ase_timeseries/timeseries.yaml
#kubectl apply -f ./ase_timeseries/timeseries_service_admin.yaml
kubectl apply -f ./ase_timeseries/timeseries_service_db.yaml



############################## autoscaler

kubectl apply -f ./ase_autoscaler/autoscaler.yaml



##################### twitter

kubectl apply -f ./ase_twitter/twitter.yaml;



##################### web

#kubectl apply -f ./ase_web/web.yaml;
#kubectl apply -f ./ase_web/web_service.yaml;



##################### replay

#kubectl apply -f ./ase_load_replay/replay.yaml;



##################### tweets

kubectl apply -f ./ase_load_tweets/tweets.yaml;
kubectl apply -f ./ase_load_tweets/tweets_service.yaml;



