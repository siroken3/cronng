# Specification for CronNG

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
* [GET /jobs/:id/history](#get-jobsidhistory)

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
      "definition": "/path/to/command2 arg1"
      "created_at": "2001-01-01T00:00:00Z",
    }
  ]
}
```

### GET /jobs

Get job definitions.

The following JSON is request-body example.

```json

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

The following JSON is request-body example.

```json

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

The following JSON is request-body example.

```json
{
  "job": {
      "name": "command1-2",
      "definition": "/path/to/command1-2 arg1 arg2",
      "created_at": "2001-01-01T00:00:00Z",
  }
}
```

The following JSON is response-body example.

```json
```

### DELETE /jobs/:id

Delete a job definition.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### POST /procs

Start a process specified by id.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### GET /procs/:id

Get a process status specified by :id.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### DELETE /procs/:id

Stop a process specified by :id.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### GET /jobs/:id/procs

Get current processes specified by :id.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
{
  "job": {
    "id": "12345678-1234-1234-1234-123456789012",
    "name": "command1"
  },
  "procs": [
    {
      "status": "RUNNING",
      "start_at": "2001-01-01T12:00:00Z"
    },
    {
      "status": "RUNNING",
      "start_at": "2001-01-01T12:00:00Z",
    }
  ]
}
```

### GET /jobs/:id/history

Get current processes history specified by :id.

The following JSON is request-body example.

```json
```

The following JSON is response-body example.

```json
{
  "job": {
    "id": "12345678-1234-1234-1234-123456789012",
    "name": "command1"
  },
  "history": [
    {
      "status": "SUCCESS",
      "code": 0,
      "start_at": "2001-01-01T12:00:00Z",
      "end_at": "2001-01-01T12:00:01Z"
    },
    {
      "status": "FAILED",
      "code": 255,
      "start_at": "2001-01-01T12:00:00Z",
      "end_at": "2001-01-01T12:00:01Z"
    }
  ]
}
```
