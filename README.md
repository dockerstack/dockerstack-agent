# Dockerstack Server Agent
This is a agent written for the dockerstack application inorder to install on the server and get the metrics and feed them on to the dockerstack-dashboard.

The core Framework is written in Golang and the architecutre of the application is as follows

![](https://raw.githubusercontent.com/dockerstack/dockerstack-golang-agent/master/dockerstack-agent.png)


### RoadMap

For the first phase of the agent we will be coming up with the server and container level metrics which will be pulled to the centralized api server.

This includes profiling and benchmarking of the code in making much more efficient in taking minimal memory space to fetch the metrics and send to the API server
