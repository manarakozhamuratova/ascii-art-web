# ASCII-ART-WEB-STYLIZE

##### Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, ascii-art.                 


#### Authors: Manara & Khamzauly


## Usage:
```
$ docker image build -f Dockerfile -t imagename .

$ docker images

$ docker container run -p 8080:8080 --detach --name containername imagename

$ docker ps -a

$ docker exec -it containername ls -l

$ docker inspect imagename | jq -r .[].Config.Labels
```

## Implementation details:
#### First of all, we need a handler. They are responsible for executing the application logic, as well as for the HTTP response header and body;
#### The second component is the HTTP request router (or servemux in Go terminology). It stores the relationship between the URL templates of the application and the corresponding handlers. In an application, there is usually one servemux that contains all the URL routes;
#### The last element for a web application is a web server.