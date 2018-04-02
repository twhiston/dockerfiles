# Dockerfiles

Dockerfiles used for apps in atomic.

To generate ci credentials run
```
docker run -it --rm --env-file=credential.env -v "$(pwd):/opt/data/" -v "/var/run/docker.sock:/var/run/docker.sock" --security-opt label=type:container_runtime_t codeship/dockercfg-generator /opt/data/dockercfg
```

then 
```
jet encrypt dockercfg dockercfg.encrypted
```
