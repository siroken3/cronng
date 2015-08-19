# Specification for CronNG

## Overview

### ClientErrors

```
HTTP/1.1 400 Bad Request
Content-Length: 35

{"message":"Problems parsing JSON"}
```

``` 
HTTP/1.1 422 Unprocessable Entity
Content-Length: 149

{
   "message": "Validation Failed",
   "errors": [
     {
       "resource": "job",
       "field": "command",
       "code": "missing_field"
     }
   ]
}
```

|Code          |Description|
|--------------|-----------|
|missing       |This means a resource does not exist.|
|missing_field |This means a required field on a resource has not been set.|
|invalid       |This means the formatting of a field is invalid.|
|already_exists|This means another resource has the same value as this field. This can happen in resources that must have some unique key|


## API

* [POST /jobs](#post-jobs)
* [GET /jobs](#get-jobs)
* [GET /jobs/:id](#get-jobsid)
* [PUT /jobs/:id](#put-jobsid)
* [DELETE /jobs/:id](#deletejobsid)
* [POST /procs](#post-procs)
* [GET /procs/:id](#get-procsid)
* [DELETE /procs/:id](#deleteprocsid)
* [GET /jobs/:id/procs](#get-jobsidprocs)
* [GET /jobs/:id/log](#get-jobsidlog)

### POST /jobs

Create a new job definition.

The following JSON is request-body example.

```json
{
  "jobs": [
    {
      "name": "command1",
      "definition": "/path/to/command1 arg1 arg2"
    },
    {
      "name": "command2",
      "definition": "/path/to/command2 arg1"
    }
  ]
}
```

The following JSON is response-body example.

```json
{
  "jobs": [
    {
      "id": "12345678-1234-1234-1234-123456789012",
      "name": "command1",
      "definition": "/path/to/command1 arg1 arg2",
      "created_at": "2001-01-01T00:00:00Z",
    },
    {
      "id": "02345678-0234-0234-0234-023456789012",
      "name": "command2",
      "definition": "/path/to/command2 arg1",
      "created_at": "2001-01-01T00:00:00Z",
    }
  ]
}
```

### GET /jobs

Get job definitions.

```
GET /jobs
```

The following JSON is response-body example.

```json
{
  "jobs": [
    {
      "id": "12345678-1234-1234-1234-123456789012",
      "name": "command1",
      "definition": "/path/to/command1 arg1 arg2",
      "created_at": "2001-01-01T00:00:00Z"
    },
    {
      "id": "02345678-0234-0234-0234-023456789012",
      "name": "command2",
      "definition": "/path/to/command2 arg1",
      "created_at": "2001-01-01T00:00:00Z"
    }
  ]
}
```

### GET /jobs/:id

Get a job definition.

```
GET /jobs/12345678-1234-1234-1234-123456789012
```

The following JSON is response-body example.

```json
{
  "job": {
      "id": "12345678-1234-1234-1234-123456789012",
      "name": "command1",
      "definition": "/path/to/command1 arg1 arg2",
      "created_at": "2001-01-01T00:00:00Z",
  }
}
```

### PUT /jobs/:id

Update a job definition.

```
PUT /jobs/12345678-1234-1234-1234-123456789012
```

The following JSON is request-body example.

```json
{
  "job": {
      "name": "command1-2",
  }
}
```

The following JSON is response-body example.

```json
{
  "job": {
      "id": "12345678-1234-1234-1234-123456789012",
      "name": "command1-2",
      "definition": "/path/to/command1-2 arg1 arg2",
      "created_at": "2001-01-01T00:00:00Z",
  }
}
```

### DELETE /jobs/:id

Delete a job definition.

```
DELETE /jobs/12345678-1234-1234-1234-123456789012
```

### POST /procs

Start a process specified by id.

The following JSON is request-body example.

```json
{
  "proc": {
      "job_id": "12345678-1234-1234-1234-123456789012",
      "trigger_type": "CRON",
      "crontab": "1 * * * *"
  }
}
```

|name            |type        |description                        |required|note   |
|----------------|------------|-----------------------------------|--------|----------------|
|job_id          |string      |The job id for start process.      |Y       |                |
|trigger_type    |string      |How this process will be started.  |Y       |CRON or ADHOC   |
|crontab         |string      |Crontab-style description          |N       |When trigger_type = "CRON" only.|

The following JSON is response-body example.

```json
{
  "proc": {
      "id": "a2345678-a234-a234-a234-a23456789012",
      "job_id": "12345678-1234-1234-1234-123456789012",
      "created_at": "2001-01-01T00:00:00Z"
  }
}
```

`start_at` is the nearest time which process is started or will be started. 

### GET /procs/:id

Get a process status specified by :id.

```
GET /procs/a2345678-a234-a234-a234-a23456789012
```

The following JSON is response example.

header
```
Status: 200 OK
```
body
```json
{
  "proc": {
      "id": "a2345678-a234-a234-a234-a23456789012",
      "job_id": "12345678-1234-1234-1234-123456789012",
      "trigger_type": "CRON",
      "crontab": "1 * * * *",
      "start_at": "2001-01-01T12:00:00Z"
  }
}
```

### DELETE /procs/:id

Stop a process specified by :id.

```
DELETE /procs/a2345678-a234-a234-a234-a23456789012
```

### GET /jobs/:id/procs

Get current processes specified by :id.

```
GET /jobs/12345678-1234-1234-1234-123456789012/procs
```

The following JSON is response example.

heaer
```
Status: 200 OK
```

body
```json
{
  "job": {
    "id": "12345678-1234-1234-1234-123456789012",
    "name": "command1"
  },
  "procs": [
    {
      "id": "a2345678-a234-a234-a234-a23456789012",
      "trigger_type": "CRON",
      "start_at": "2001-01-01T12:00:00Z"
    },
    {
      "id": "b2345678-b234-b234-b234-b23456789012",
      "trigger_type": "ADHOC",
      "start_at": "2001-01-01T12:00:00Z",
    }
  ]
}
```

### GET /jobs/:id/log

Get current processes history specified by :id.

```
GET /jobs/12345678-1234-1234-1234-123456789012/log
```

The following JSON is response-body example.

```json
{
  "job": {
    "id": "12345678-1234-1234-1234-123456789012",
    "name": "command1"
  },
  "log": [
    {
      "proc_id": "c2345678-c234-c234-c234-c23456789012",
      "result": "SUCCESS",
      "code": 0,
      "start_at": "2001-01-01T12:00:00Z",
      "end_at": "2001-01-01T12:00:01Z"
    },
    {
      "proc_id": "d2345678-d234-d234-d234-d23456789012",
      "result": "FAILED",
      "code": 255,
      "start_at": "2001-01-01T11:00:00Z",
      "end_at": "2001-01-01T11:00:01Z"
    }
  ]
}
```
