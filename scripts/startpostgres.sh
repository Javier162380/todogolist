#!/bin/bash
docker run -it  -p  5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=javier.llorente postgres