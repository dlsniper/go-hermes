# go-hermes
Go-hermes is a Golang app that exposes an HTTP API that receives requests with software metrics.

# About
This repository is under heavy development. Doesn't work yet. I appreciate you're taking a look, please read my thoughts below.

# Main components
1) HTTPS endpoint where all requests will go to (we need to spread the servers across the globe to achieve low latency). There are going to be different endpoints for each type of request. For example for mobile app metrics, `mobile.<host>.com`, server metrics: `server.<host>.com` etc.
2) Analyze data from requests. This depends on the type of data we received and each one will require different analysis. For example, if it's an app error analysis would be to extract stack trace.
3) Present data on a UI dashboard. This will allow users to get a visual understanding on what's under-performing and what's performing well.
4) Agents that push data to endpoint. These are going to be executable files installed on remote host (client installs on their machines), and will collect metrics and push them to our endpoint.
5) API Clients in different programming languages. This will allow users to create custom metrics that matters to them, and push them programmatically to our endpoint (for example profiling for a function in their app).                        

There are many interesting issues to tackle on a project like this:
- coming up with a distributed system architecture
- scalability and fault tolerance

The kind of payload we will work with, is going to be JSON as its more lightweight than XML, and all the incoming requests will be compressed, meaning less bandwidth usage (for customers and for us) and faster communication.

Kind of metrics that we can collect (below list is not exhaustive):
- app data (database, external requests, errors/exceptions)
- server monitoring (disks, memory, cpu)
- mobile data (usage data, errors, device information)                        

We also need to think about app events such as: deployments, or software updates and allow customer to compare how that change impacted the servers/application performance. That will allow them to easily decide if a rollback is needed.                   

