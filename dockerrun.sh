#!/bin/bash

docker run --name groupie_tracker -v $(pwd):/project -dp 8000:8000 groupie_tracker