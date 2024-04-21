#Web crawler service
This is a simple web crawler service(server and client) implemented in Go. It provides a RESTful API where the server would receive requests from a client to crawl a URL and should send the response i.e sitemap back to
the client.

#Components
  #Server: The server component is responsible for handling incoming requests from clients and crawling the specified URL. It is implemented in Go and exposes an HTTP endpoint for clients to request #crawling.
  Client: The client component sends requests to the server to initiate crawling of a specified URL. It is implemented in Go and interacts with the server via HTTP requests.


#Usage
  #Server:
    Build  and run the server binary: go run .\main.go
    The server will start listening on port 8080 by default.
  #Client:
    Build  and run the client binary: go run .\client\client.go 
    Follow the prompts to enter the URL to crawl.
    The client will display the sitemap returned by the server.

#Dockerization:
  #Build Docker image :
     docker build -t prantikkumarpatra/carwl-service:v1.0.0 .
  #Run Docker container :
    docker run -d -p 8080:8080 --name my-crawler prantikkumarpatra/carwl-service:v<version>
  #Stop Docker container:
    docker stop my-crawler
  #Delete Docker container:
    docker rm my-crawler
#Accessing the Service:
  The service will be available at http://localhost:8080/crawl?url=<url_to_be_crawl>

#Deployment
#Kubernetes Deployment:
  The server can be deployed to a Kubernetes cluster using the provided Kubernetes manifests (deployment.yaml, service.yaml).
  #Apply the manifests to the Kubernetes cluster: 
    cd k8s_deployment  
    kubectl apply -f .

