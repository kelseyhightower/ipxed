# Machines API

## Create a machine

```
curl -X POST http://localhost:8080/api/machines --data-binary @"$HOME/core0.json"
```

core0.json:
```
{
  "name": "core0",
  "macaddress": "00:0c:29:5c:b1:00",
  "profile": "production"
}
```

Response

```
HTTP/1.1 201 Created
Location: http://localhost:8080/machines/core0
Date: Sun, 01 Jun 2014 19:22:25 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```

## Get all machines

```
curl -X GET http://localhost:8080/api/machines
```
