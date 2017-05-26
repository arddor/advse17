# Deployment Instructions (Google Cloud Platform)

The following commands will assume that the application is run on google container engine (GCE), although it is possible to deploy the application on every kubernetes environment.


## Requirements

* Clone the source code from Github https://github.com/arddor/advse17
* Go to [Container Engine](https://console.cloud.google.com/project/_/kubernetes/list?_ga=1.236200065.1403830672.1494830303) and create or select a project
* Make sure you have the google cloud console on your [workstation](https://cloud.google.com/sdk/docs/quickstarts)
* Install kubectl and set the default settings on your gcould utility
    * Best to follow this [tutorial](https://cloud.google.com/container-engine/docs/tutorials/hello-node)

```gcloud components install kubectl```
```gcloud config set project PROJECT_ID``` (the one chosen above)
```gcloud config set compute/zone europe-west1-c```

* Create a [PersistentVolume](https://cloud.google.com/container-engine/docs/tutorials/persistent-disk/) for our DB.

```gcloud compute disks create --size 200GB rethink-disk```
    

## Deployment

Given that the previous step was executed successfully the deployment scripts in the script folder can be used. Conceptually they first build for each micro service the docker container and upload them to the google cloud registry (note that it is crucial that the PROJECT_ID was set correctly). Then the kubernetes configurations are uploaded and executed. Apply the following scripts exactly in this order from the root folder advse17

```./scripts/uploadContainers.sh```

```./scripts/k8createApp.sh```

After startup our api service will get an EXTERNAL-IP where our UI runs on Port 80. You can find out the IP using:

```./scripts/k8status.sh```


## Load Generation

To put everything under load, we minded about 1 Mio. Tweets about "trump". To have the load generator work properly, please ad the term "trump" using the UI first. After that you can call:

```./scripts/k8replayStart.sh```

Which will start three replay nodes. The nodes will start pushing tweets onto the queue directly.



# Running it locally (Docker compose)

This runs the services using docker compose.

```bash
docker-compose up
```

Or just launch one container

```bash
docker-compose up ase_compute
```
