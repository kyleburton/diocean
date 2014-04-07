#!/usr/bin/env bash
set -eux
VERBOSE=""

# ./diocean regions ls
# id      name    slug
# 3       San Francisco 1 sfo1
# 4       New York 2      nyc2
# 5       Amsterdam 2     ams2
# 6       Singapore 1     sgp1

#REGION=nyc2
REGION=sfo1

# create a droplet
./diocean $VERBOSE -w droplets new test1 512mb ubuntu-13-10-x64 $REGION 20848 false false

# show that we have one
./diocean $VERBOSE droplets ls

DROPLET_ID=$(./diocean droplets ls | tail -n 1 | cut -f1)

# power off the droplet
./diocean $VERBOSE -w droplets power-off $DROPLET_ID
# take a snapshot
./diocean $VERBOSE -w droplets snapshot $DROPLET_ID test-snapshot-1

# destroy the droplet
# get the droplet ID from the previous create 
./diocean $VERBOSE -w droplets destroy $DROPLET_ID false

# delete the snapshot (image)
IMAGE_ID=$(./diocean images ls | sort -n | grep false$ | tail -n 1 | cut -f1)
./diocean $VERBOSE -w images destroy $IMAGE_ID


rm new.output
rm destroy.output

./diocean $VERBOSE droplets ls
