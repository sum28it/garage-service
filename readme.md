# Garage-Service

This project implements a garage service api that exposes endpoints for adding, buying and selling cars. The aim of this project is to build a backend service that is maintainable and extendable over period of time.

## Up & Running

### Kubernetes Cluster

Clone the repo:

``` shell
git clone github.com/sum28it/garage-service
```

Software installation:

- Kind
- Docker
- Go
- Kubectl
- You can also install [telepresence](https://www.telepresence.io/) to be able to run commands from within the cluster.

__NOTE:__ Change the configurations at app/services/main.go and app/scratch/db/main.go to your cluster host and postgres password.

Start Kubernetes Cluster:

- RUN: ```make dev-up``` to start the cluster
- RUN: ```make dev-update-apply``` to build the docker image, load it into the cluster and run the containers.
- All the commands like to see status, logs, etc. are in the makefile.
- Finally, RUN ```make dev-down``` to delete the cluster.

### Locally

To run the project outside a k8s cluster, you need to have PostgreSQL installed .

- Create a postgres role to be used by the application.
- First run the ```make db``` command to migrate and seed the database.
- Run make run to start the app

## Directory Tree

``` shell
├───app
│   ├───scratch
│   │   ├───db
│   │   └───jwt
│   ├───services
│   │   └───sales-api
│   │       └───handlers
│   │           └───v1
│   │               ├───testgrp
│   │               └───usergrp
│   └───tooling
│       └───logfmt
├───business
│   ├───core
│   │   └───user
│   │       └───stores
│   │           └───userdb
│   ├───data
│   │   ├───dbschema
│   │   │   └───sql
│   │   └───dbtest
│   ├───sys
│   │   ├───database
│   │   └───validate
│   └───web
│       ├───auth
│       │   └───rego
│       ├───keystore
│       ├───metrics
│       └───v1
│           ├───debug
│           │   └───checkgrp
│           └───mid
├───foundation
│   ├───docker
│   ├───logger
│   ├───vault
│   └───web
└───zarf
    ├───docker
    ├───k8s
    │   ├───base
    │   │   └───sales
    │   └───dev
    │       ├───database
    │       └───sales
    └───keys

```

## Layers

![Request Lifecycle image](.images/Request_Life_cycle.svg)

- __App:__ This layer contains all the services and tools that are developed as a part of this project. The tooling directory contains the utility CLI tool that is built during the project. Scratch contains scripts for generating some quick objects and may not be needed in production environment.
- __Business:__ The business layer is a set of packages that implement the core business logic of the application. The packages in this layer behave independent of the fact whether they are being used by a web app or a CLI app. The core layer provides APIs that are called directly by app layer to process user requests. The core layer accesses APIs from sys to perform database operations. The web layer provides auth, metrics and other middleware functionalities.
- __Foundation:__ The foundation layer provides a micro web framework that is lightweight, compatible to standard library and extendable. The framework uses minimal dependencies and provides support for middlewares, response and handlers with context. It contains packages that can be used across different projects. It is decoupled from the business logic of the app and packages in this layer cannot access database and loggger.
- __Vendor:__ This directory contains all the third party libraries imported by our application. Vendoring helps in making projects easy to clone and build. Although vendoring sometimes makes the IDE bit slow but the ability to browse the source code in the IDE feels good.
- __Zarf:__ This layer contains the deployment configuration for Kubernetes, dockerfiles and keys for generating JWT.

### Cluster Architecture

![Cluster Architecture](./.images/Cluster_Architecture.svg)

<!-- The application runs in a seperate namespace to avoid naming conflicts. Inside our cluster, we have two deployments for our service pod and database respectively. -->