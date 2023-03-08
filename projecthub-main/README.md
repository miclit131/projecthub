
# ![logo](http://i.epvpimg.com/l009eab.png) ProjectHub

Every semester, our university creates fabulous projects that, in our opinion, get far too little attention.

Our goal is to find a technical solution to make our projects even more easily and sustainably accessible. We don't want to replace the Stage, but rather expand and enrich our university infrastructure in the future.

To be more specific - our aim is it to do the web development for you. In the use case of the Stage - we enriched the stage by a "play now" button so the users dont have to actually download something!

Techstack: Frontend(Node, React & MUI), Backend(Docker, Kubernetes, Go, Go Gin, ...)

## Setup

### Prerequisites:

- Linux or Windows with WSL2
- Docker (https://www.docker.com/get-started/)
- Node (https://nodejs.org/en/)
- Go (https://go.dev/doc/install / WSL: https://medium.com/@benzbraunstein/how-to-install-and-setup-golang-development-under-wsl-2-4b8ca7720374)
- Kubernetes option in Docker settings must be activated

### Permissions:
For the backend to have write permissions, the following commands must be executed:
```
sudo mkdir ~/var/lib/projecthub/
sudo mkdir ~/var/lib/projecthub/k8sObj
sudo chown -R ~/var/lib/projecthub/k8sObj
```

### Install Kompose:
Execute these commands:
```
wget https://github.com/kubernetes/kompose/releases/download/v1.26.1/kompose_1.26.1_amd64.deb # Replace 1.26.1 with latest tag
sudo apt install ./kompose_1.26.1_amd64.deb
```
Releases: https://github.com/kubernetes/kompose/releases

### Registry:
In order for the backend to access the registry, the following commands must be executed:
```
mkdir -p /opt/docker-registry/cert
cd /opt/docker-registry/cert
sudo openssl req -newkey rsa:2048 -nodes -keyout registry_auth.key -x509 -days 365 -out registry_auth.crt
mkdir -p /opt/docker-registry/auth
docker run --entrypoint htpasswd httpd:2 -Bbn projecthub password > htpasswd
```
If you have permission problems: 
```
chown -R username:groupname .
```

Start the registry as Docker container:
```
docker run -d -p 5000:5000 --restart=always --name projecthub-registry registry:2
```

Login to give Docker/Kubernetes access to registry:
```
docker login -u projecthub -p password localhost:5000
```

References:  
https://docs.docker.com/registry/deploying/#run-an-externally-accessible-registry  
https://medium.com/@ManagedKube/docker-registry-2-setup-with-tls-basic-auth-and-persistent-data-8b98a2a73eec

### Ingress
To set up the ingress controller, execute the following commands:
```
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
helm version
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo list # check if repo exists
helm repo update
helm install nginx ingress-nginx/ingress-nginx 
```
References:  
https://helm.sh/docs/intro/install/  
https://artifacthub.io/packages/helm/ingress-nginx/ingress-nginx

Note:
There are still problems with ingress. Only the last uploaded project will be accessible via the URL http://kubernetes.docker.internal/

### Database
To start the database as Docker container, execute:
```
docker run -d -p 27017:27017 --restart=always --name projecthub-db -d -e MONGO_INITDB_ROOT_USERNAME=projecthub -e MONGO_INITDB_ROOT_PASSWORD=password -e MONGO_INITDB_DATABASE=projecthub
```
Note: The database is not implemented yet.


## Start

To start the project, consisting of front- and backend, either open the repo in VSCode (in WSL if you're on Windows) and make use the `launch.json`.

Alternatively `cd` into the front-/backend directory and start it from the command line.

Backend:
```
go run main.go
```

Frontend:
```
npm i
npm start
```


## User Flow

After starting the backend and then the frontend you can reach the frontend at:
 ```
http://localhost:3000/
```
Example Data you can use to enter:
 ```
Name: Example123
Git Link: https://github.com/tomowatt/unity-docker-example/
StageId: 3908
```
Wait some time and click on the button next to the "?" button or visit:
 ```
http://localhost:3000/browse
```

We also provide you our postman export - just in case you are more comfortable with api calls :)
