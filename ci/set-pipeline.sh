#!/bin/bash
# my-go-packages set-pipeline.sh

fly -t ci set-pipeline -p my-go-packages -c pipeline.yml --load-vars-from ../../../.credentials.yml
