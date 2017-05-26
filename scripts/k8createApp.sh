#!/usr/bin/env bash




#gcloud auth application-default login

#gcloud container clusters get-credentials awesome-cluster --zone europe-west1-c --project advse-167708

#kubectl proxy


############################## api

kubectl create --save-config -f ./ase_api/api.yaml; kubectl create --save-config -f ./ase_api/api_service.yaml;



##################### compute

kubectl create --save-config -f ./ase_compute/compute.yaml;



##################### queue

kubectl create --save-config -f ./ase_queue/queue.yaml; kubectl create --save-config -f ./ase_queue/queue_service.yaml;



############################## timeseries

kubectl create --save-config -f ./ase_timeseries/timeseries.yaml; kubectl create --save-config -f ./ase_timeseries/timeseries_service_db.yaml; kubectl create --save-config -f ./ase_timeseries/timeseries_service_admin.yaml;



############################## autoscaler

kubectl create --save-config -f ./ase_autoscaler/autoscaler.yaml;



##################### twitter

kubectl create --save-config -f ./ase_twitter/twitter.yaml;



##################### tweets

kubectl create --save-config -f ./ase_load_tweets/tweets.yaml; kubectl create --save-config -f ./ase_load_tweets/tweets_service.yaml;


