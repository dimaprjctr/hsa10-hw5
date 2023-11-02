#!/bin/bash

# Run stress tests
siege -c10 -t15S -f urls.txt
siege -c25 -t15S -f urls.txt
siege -c50 -t15S -f urls.txt
siege -c70 -t15S -f urls.txt
siege -c100 -t15S -f urls.txt