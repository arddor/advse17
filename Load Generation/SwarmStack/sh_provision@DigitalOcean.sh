#!/bin/bash


# Reference
# https://docs.docker.com/machine/drivers/digital-ocean/#options

#echo $1

DOTOKEN=

# determine how many nodes we have
i=42


# Create the machine

docker-machine create --driver digitalocean \
--digitalocean-image  ubuntu-16-04-x64 \
--digitalocean-access-token $DOTOKEN node-$i



# Login
docker-machine ssh node-$i

# Join the swarm

# init the swarm (probably never an issue)
#docker swarm init --advertise-addr node_ip_address

# as a manager


# as a worker
docker swarm join \
--token your_swarm_token \
manager_node_ip_address:2377


# GitHub

# pull in our code

git clone


# Simons stuff

Swarm initialized: current node (wf5unet6cgkaf20vszkdsdbc5) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join \
    --token SWMTKN-1-43qeadafih12dclhzhbpatut126hang92e30614fybpya7pl27-bmj0923s1b4ffee2k3ky3k38y \
    192.168.65.2:2377

