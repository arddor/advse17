## Project Presentation Stage 1
### Stefanie Ziltener, Marc Heimgartner, Benjamin Bürgisser, Simon Tännler
#### Advanced Software Engineering FS 2017, University of Zürich

---

![inline](img/architecture.png)


^ High-level Overview: Client-Server
^ we will go through each of these boxes

---

# Containerized Microservices deployed through Kubernetes

![left, fit](img/architecture.png)

* Components embedded in (Docker) Containers
* Containers have (ideally) one single responsibility

^ Containers as the instantiation of a Microservice
^ We've seen the benefits of Containers in the lecture

---

# Frontend and API

![left, fit](img/architecture_frontend.png)

* User inputs term
* Click on register
* Send request to API

^ Explain the Architecture by
^ following the Story of a Request through our Architecture

---

# Timeseries: MongoDB (I)

![left, fit](img/architecture_mongodb.png)

* Request Handler stores Term in MongoDB
* Persistency guaranteed by GCE [Persistence Disk](https://cloud.google.com/compute/docs/disks/)

^ TODO: Img of newly created Term in JSON

---

# Twitter Service

![left, fit](img/architecture_twitter.png)

* Gets notified of newly created Terms
* Streaming stops and restarts with the new Term added for [tracking](https://dev.twitter.com/streaming/overview/request-parameters#track)
* Arriving Tweets are immediately stored into the Queue

^ Restarts are due to Streaming API limitations
^ We're trying to keep this a light as possible
^ b/c we have some limitation which we'll talk about later

---

# Worker Queue

![left, fit](img/architecture_queue.png)

* [Redis](https://redis.io): in-memory data structure store
* A FIFO queue of Strings (Tweets)
* Load generator
	* API endpoint to add Strings to Queue directly

---

# Compute Workers

![left, fit](img/architecture_queue.png)

* Running Workers process the Queue:
	* Assign Tweet to Term (filtering)
	* Calculate Sentiment

* Length of the Queue defined the number of Workers
	* Scaled through Kubernetes.

^ Twitter API does not tell you which term the tweet matched on

---

# Timeseries: MongoDB (II)

![left, fit](img/architecture_mongodb.png)

* Workers store the calculated Sentiment into MongoDB.

^ The amount of data is actually very small
^ TODO: Img of newly created Term in JSON

---

# Displaying Results

![left, fit](img/architecture_frontend.png)

* Rest API gets request
* Collect relevant data from MongoDB
* Browser renders data
* Socket gets opened for continues pushes

---

# Architectural Styles

* Client / Server through Rest API
* Event-Driven notification of Term updates
* Blackboard: Redis Queue
	* Factory: Twitter Service
	* Worker: Compute Workers
* Highly decoupled Microservices

^ TODO: this feels like it needs more work

---

# Do you even scale?

---

# Initial Idea

* Digital Ocean
* Docker Swarm

---

![left](img/kubernetes-logo.png)


#❤️


![right](img/docker-logo.png)

---

# What does [Kubernetes](https://kubernetes.io) do for us

* Every component is potentially scalable through Kubernetes
	* Even [MongoDB](https://cloud.google.com/solutions/deploy-mongodb)!

* Fault Tolerance:
	* Container recovery through Kubernetes
	* Decoupled design and Microservice

* Elasticity
	* Container scaling through Kubernetes

---

# Concerns

* High lock-in to Kubernetes
* [Twitter](https://dev.twitter.com/streaming/reference/post/statuses/filter): 400 keyword tracking per stream
	* "default access level allows up to 400 track keywords"
* Matching: may not match to terms perfectly
* Redis: may become a bottleneck
	* but we highly doubt it

^ TODO: this feels like it needs more work

---

# Technology Zoo - Platform

* Cloud Platform: [Google Container Engine](https://cloud.google.com/container-engine/) (GCE)
 	* Easy support of [Kubernetes](https://kubernetes.io)

* Containerization: [Docker](http://docker.com)
	* Popular Container engine

* Container orchestration: [Kubernetes](https://kubernetes.io)
	* Popular Container orchestration

---

# Technology Zoo - Backend

* Programming Language: Google Go
	* New Programming language 🎉
	* Uniquely suited for Web development
	* Have I mentioned it's fast?

* Terms Storage: [MongoDB](https://www.mongodb.com)
	* Easy data schema

* Queue Storage: [Redis](https://redis.io)
	* High-performant in-memory storage ideally suited for our purpose

---

# Technology Zoo - Frontend

* Frontend: [Vuejs](http://vuejs.org)
	* Similar to Angular and React

* Webserver: [nginx](http://nginx.org)
	* Battle-proven Webserver

---

# Dev Environment

Docker containers using [Docker Compose](https://docs.docker.com/compose/)

---

# Demo

---

# Questions?























