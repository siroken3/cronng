# Specification for CronNG

## API

* [POST /jobs](#post-jobs)
* [GET /jobs](#get-jobs)
* [GET /jobs/:id](#get-jobs-id)
* [PUT /jobs/:id](#put-jobs-id)
* [DELETE /jobs/:id](#delete-jobs-id)
* [POST /procs](#post-procs)
* [GET /procs/:id](#get-procs-id)
* [DELETE /procs/:id](#delete-procs-id)
* [GET /jobs/:id/procs](#get-jobs-id-procs)
* [GET /jobs/:id/history](#get-jobs-id-history)

### POST /jobs

Create a new job definition.

The following JSON is request-body example.

```json
{
  "jobs": [
    {
      "name": "Say Hello",
      "definition": "echo \"Hello World.\""
    },
    {
      "name": "Say Good-bye",
      "definition": "echo \"Good-bye.\""
    }
  ]
}
```

The following JSON is response-body example.

```json
```

### GET /jobs

Get job definitions.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json

```

### GET /jobs/:id

Get a job definition.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### PUT /jobs/:id

Update a job definition.

The following JSON is request-body example.

```json

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
    "name": "Say Hello"
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
    "name": "Say Hello"
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
