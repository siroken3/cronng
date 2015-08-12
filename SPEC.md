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

### POST /jobs

Create a new job definition.

The following JSON is request-body example.

```json
{
   "jobs": [
      {
         "name": "Say Hello",
         "definition": |
           echo "Hello World."
      },
      {
         "name": "Say Good-bye",
         "definition": |
           echo "Good-bye."
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
         "name": "Say Hello"
      },
      {
         "id": "02345678-0234-0234-0234-023456789012",
         "name": "Say Good-bye"
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
```

### GET /jobs/ID

Get a job definition.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### PUT /jobs/ID

Update a job definition.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### DELETE /jobs/JOB_ID

Delete a job definition.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### POST /procs

Start a process specified by JOB_ID.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### GET /procs/JOB_ID

Get a process status specified by JOB_ID.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### DELETE /procs/PROC_ID

Stop a process specified by PROC_ID.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```

### GET /jobs/JOB_ID/procs

Get process history specified by JOB_ID.

The following JSON is request-body example.

```json

```

The following JSON is response-body example.

```json
```
