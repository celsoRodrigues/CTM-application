# CTM-application
Objectives from the assignment pdf:
- Produce a simle cluster web service setup in AWS - Done
```
   Here I have an EKS cluster running, deployed with terraform.
   I have an ingress controller and a few helm charts deployed with providers
```
- Cluster hit with get request returns a simple hello world with data and time - Done
```
   App in golang that refreshes the time and a salutation when hit with GET request
```
- infrastructure and pipelines as code - Done
```
   Pipeline created in github actions, with tests written in go.
```
- Tools that are transferable or platform agnostic - done
```
   All resources definied as code. A possible improvement here would be to make even more use of helm.
```
- Products deployed should be highly available, scalable and maintainable and monitored - Done
```
  cluster using an NLB ingress controller, capable of handling a multitude of requests and balance the load through pods/services. 
  Deployment resources highly used, allowing the system to self heal and scale with ease.
  A possible improvement would be to make use of livelyness and rediness in the deployment definition
  
  For monitorization, I have instrumented my golang application with prometheus, installed prometheus and grafana to scrape my app.
```

This is a go server application, instrummented for prometheus.

I also wrote a few example tests, the aim is to use them in the pipeline.

endpoints:
/metrics - prometheus instrumentation output
/        - root responds with hello world and the timestamp
