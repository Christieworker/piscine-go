#!/bin/bash
curl -s https://academy.digifemmes.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"jcaitano\"}}){id}}"}' | cut -d ":" -f4 | cut -d '}' -f1

