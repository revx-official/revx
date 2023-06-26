# Load Balancing

## Introduction

What is load balancing? In a world of distributed applications, the ability of scaling services horizontally is one of the most important aspects. When you scale a service horizontally, you actually have multiple running instances of exactly the same service. However, how to you determine which service is going to handle an incoming request? This is where a load balancer comes into play.

Instead of the services themselves, the load balancer receives all incoming requests. The load balancer is previously configured to know exactly how many instances (upstreams) of one service are available. Now, if an incoming request needs to be redirected, the load balancer determines (based on a bunch of metrics), which instance is going to handle this request.

## Load Balancing In revx

*revx* implements a simple and straigt forward load balancing algorithm known as round robin load balancing. Basically, *revx* internally increments the upstream index every time a new request is send. For example, if there are 3 upstreams (`up-1`, `up-2` & `up-3`), the first request is redirected to `up-1`, the second to `up-2`, the third to `up-3` and the fourth to `up-1` again.