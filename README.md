# sj

## Util - 'compare versions'

## Running command

```
$ go run cmd/compare/main.go
INFO[0000] Enter Version 1:
2
INFO[0003] Enter Version 2:
1.1.2
INFO[0007] Version 2 is bigger than Version 1.1.2
```

## Application - 'math api'

## Running web service

```
$ go run cmd/api/main.go
INFO[0000] Starting the Math API on 8080 ...
```

Sample Call to API's

```
$ curl --request GET 'localhost:8080/min' -H "Content-Type: application/json" -d '{"numbers": [ 2, 6, 3, 4, 1], "quantifier": 2}'
{"description":"Min of [2 6 3 4 1] with quantifier 2","results":[1,2]}

$ curl --request GET 'localhost:8080/max' -H "Content-Type: application/json" -d '{"numbers": [ 2, 6, 3, 4, 1], "quantifier": 2}'
{"description":"Max of [2 6 3 4 1] with quantifier 2","results":[4,6]}

$ curl --request GET 'localhost:8080/avg' -H "Content-Type: application/json" -d '{"numbers": [ 2, 6, 3, 4, 1]}'
{"description":"Avg of [2 6 3 4 1]","results":[3.2]}

$ curl --request GET 'localhost:8080/median' -H "Content-Type: application/json" -d '{"numbers": [ 2, 6, 3, 4, 1]}'
{"description":"Median of [2 6 3 4 1]","results":[3]}

$ curl --request GET 'localhost:8080/percentile' -H "Content-Type: application/json" -d '{"numbers": [ 2, 6, 3, 4, 1], "quantifier": 60}'
{"description":"60th percentile of [2 6 3 4 1]","results":[3]}
```
