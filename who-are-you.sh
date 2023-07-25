#!/bin/bash
curl https://academy.digifemmes.com/assets/superhero/all.json | jq ' .[] | select ( .id == 70 ) | .name'

