#!/bin/bash
# my-go-packages destroy-pipeline.sh

fly -t ci destroy-pipeline --pipeline my-go-packages
