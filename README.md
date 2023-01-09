# Learning Fluent

fluent-bit GKE multiline 테스트 프로젝트입니다.  

<br/><br/><br/>

## Run  

```bash
export IMAGE_REPO={Image Registry}
```

> $ROOT_ENV 는 프로젝트 Root 경로입니다.  

```bash
docker build --rm --no-cache --tag $IMAGE_REPO/fluentlogger --file ./assets/docker/Dockerfile .
```

run fluentlogger in the kubernetes:  

```
kubectl run fluentlogger --image=$REPO_AR/fluentlogger --env="LOG_OUT=log.out" --env="LOG_LEN=65536" --env="LOG_COUNT=10"
```

* LOG_OUT: Log 출력 파일 이름 (기본 값은 temp.json)
* LOG_LEN: Log Byte 길이 (기본 값은 64 Byte)
* Count: Log 출력 회수 (기본 값은 10 번)

<br/><br/><br/>

## References  
* [Fluent Bit v2.0](https://docs.fluentbit.io/manual/)  
