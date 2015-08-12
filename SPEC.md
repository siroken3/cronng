# Specification for CronNG

## API

* [POST /jobs](#post-jobs)
* [GET /jobs](#get-jobs)
* [GET /jobs/ID](#get-jobs-id)
* [PUT /jobs/ID](#put-jobs-id)
* [DELETE /jobs/ID](#delete-jobs-id)
* [POST /procs](#post-procs)
* [GET /procs/ID](#get-procs-id)
* [DELETE /procs/ID](#delete-procs-id)
* [GET /jobs/ID/procs](#get-jobs-id-procs)

### POST /jobs

Create a new job definition.

The following JSON is request-body example.

```json

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

### GET /jobs/ID

Get a job definition.

### PUT /jobs/ID

Update a job definition.

### DELETE /jobs/JOB_ID

Delete a job definition.

### POST /procs

Start a process specified by JOB_ID.

### GET /procs/JOB_ID

Get a process status specified by JOB_ID.

### DELETE /procs/PROC_ID

Stop a process specified by PROC_ID.

### GET /jobs/JOB_ID/procs

Get process history specified by JOB_ID.
