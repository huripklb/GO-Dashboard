# GO-MoDa
research project by Huripto Sugandi

## What is GO-MoDa?

GO-MoDa is a digital platform created to show necessary informations in a Monitor or any Display.

Informations provided in this platform should give important informations which can be used to perform action(s) that can make operation goes smooth.

MoDa shall be displayed in a big monitor where people can see and easily notice informations provided by the platform  

## Current version

Current version is 0.0.1

## Features

Currently available features are :

* generate aging pending orders
* generate aging proceed orders
* generate aging shipped orders
* generate aging delivered orders
* generate aging draft PO
* generate aging shipped PO
* generate Top 10 Aging pending orders
* generate Top 10 Aging proceed orders
* generate Top 10 Aging shipped orders
* generate Top 10 Aging draft PO
* generate Top 10 Aging shipped PO

## How to run ?

Execute these commands from your main go directory :

* go build src/worker_aging.go
* go build src/worker_topten.go

Then Run these commands :

* ./worker_aging
* ./worker_topten

# ENJOY!