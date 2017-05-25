#!/usr/bin/env bash




#gcloud auth application-default login

#gcloud container clusters get-credentials awesome-cluster --zone europe-west1-c --project advse-167708

#kubectl proxy


############################## api

kubectl create -f ./ase_api/api.yaml;
kubectl create -f ./ase_api/api_service.yaml;



##################### compute

kubectl create -f ./ase_compute/compute.yaml;



##################### queue

kubectl create -f ./ase_queue/queue.yaml;
kubectl create -f ./ase_queue/queue_service.yaml;



############################## timeseries

kubectl create -f ./ase_timeseries/timeseries.yaml
kubectl create -f ./ase_timeseries/timeseries_service_admin.yaml
kubectl create -f ./ase_timeseries/timeseries_service_db.yaml



##################### twitter

kubectl create -f ./ase_twitter/twitter.yaml;



##################### web

kubectl create -f ./ase_web/web.yaml;
kubectl create -f ./ase_web/web_service.yaml;



