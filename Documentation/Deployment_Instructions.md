ADVSE17

The following commands will assume that the application is run on google container engine (GCE), although it is possible to deploy the application on every kubernetes environement.


# Requirements

1) clone the source code from github https://github.com/arddor/advse17
2) go to https://console.cloud.google.com/project/_/kubernetes/list?_ga=1.236200065.1403830672.1494830303 and create or select a project
3) make sure you have the google cloud console on your workstation https://cloud.google.com/sdk/docs/quickstarts
4) ensure that kubectrl is installed or install it via `gcloud components install kubectl`
5) set the default settings for the project. The PROJECT_ID and the compute zone via `gcloud config set project PROJECT_ID
gcloud config set compute/zone us-central1-b`. The PROJECT_ID has to be equal to the project_id, which was created in step 2

In case you have problems with one of the steps please refer to https://cloud.google.com/container-engine/docs/tutorials/hello-node an complete the initialization of the environment (before step 1)


# Deployment

Given that the previous step was executed successfully the deployment scripts in the script folder can be used. Conceptually they first build for each micro service the docker container and upload them to the google cloud registry (note that it is crucial that the PROJECT_ID was set correctly). Then the kubernetes configurations are uploaded and executed. Apply the following scripts exactly in this order from the root folder advse17:
1) ./scripts/uploadContainers.sh
2) ./scripts/k8createAll.sh
