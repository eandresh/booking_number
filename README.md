# Booking a number

Project name is a `Booking numbers` 

## Table of Contents

- [Prerequisites](#prerequisites)
- [Quickstart](#quickstart)

## Prerequisites

Before you start, make sure you meet the following requirements:
* You have installed the latest version of [Go](https://go.dev/dl/)
* You have installed [Docker](https://docs.docker.com/desktop/)

## Quickstart
To Run, follow these steps:

```bash
 docker-compose up --build
```
Endpoints
## Save a booking number
```
 POST localhost/digital-shift/booking
```
```
curl --location --request POST 'localhost/digital-shift/booking' \
--header 'Content-Type: application/json' \
--data-raw '{
"client_id": "3",
"number" : 1
}'
```
## Get Bookings
```
 GET localhost/digital-shift/booking-list
```
```
curl --location --request GET 'localhost/digital-shift/booking-list' \

```