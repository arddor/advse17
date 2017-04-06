# API Backend

REST API to store/retrieve Terms from the DB. Plus data for frontend


# Streaming Pod

Stream twitter data. Hook from DB to retrieve new Terms from DB.
Put data into Task queue (Redis).


# Job Queue (Redis)

Kubernetes Jobs work on the twitter data queue (parallel).
These compute the sentiment AND figure out to which term the tweet is associated.


# MongoStore

Timeseries storing of sentiments and terms.


# Frontend

Pod serving the client.
