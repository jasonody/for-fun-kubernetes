# Jobs

Playing with kubernetes cron jobs

[Source](https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/)

## Hello world cron job

Create cron job: `kubectl create -f hello-world-cron-job.yaml`

Get cron jobs: `kubectl get cronjobs`

Watch jobs: `kubectl get jobs --watch`

Get the logs from one of the pods for this job: `kubectl logs $(kubectl get pods -o wide | grep hello-world-cron-job | head -n 1 | awk '{print $1}')`

Delete cron job: `kubectl delete cronjob hello-world-cron-jobs` or `kubectl delete -f hello-world-cron-job.yaml`

----

Play with kubernetes jobs

[Source](https://medium.com/better-programming/tutorial-how-to-use-kubernetes-job-and-cronjob-1ef4ffbc8e84)

## Enforcing a time limit

Create job: `kubectl create -f enforcing-a-time-limit.yaml`

Check the deadline was enforced: `kubectl get job enforcing-a-time-limit -o yaml` or `kubectl describe job enforcing-a-time-limit` (check for "DeadlineExceeded")

Delete job: `kubectl delete job enforcing-a-time-limit` or `kubectl delete -f enforcing-a-time-limit.yaml`

## Handling failures

Create job: `kubectl apply -f jobs/handling-failures.yaml`

Watch job's pod status and restarts: `kubectl get pod $(kubectl get pods | grep handling-failures | tail -n 1 | awk '{print $1}') --watch`

Check the backoff limit was enforced: `kubectl get job handling-failures -o yaml` or `kubectl describe job handling-failures` (check for "BackoffLimitExceeded")

Delete job: `kubectl delete job handling-failures` or `kubectl delete -f handling-failures.yaml`

## Running in parallel

Create job: `kubectl apply -f jobs/running-in-parallel.yaml`

Watch job's overlapping completions: `kubectl get job running-in-parallel -w`

Delete job: `kubectl delete job running-in-parallel` or `kubectl delete -f running-in-parallel.yaml`