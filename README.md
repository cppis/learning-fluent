# Learning Fluent

fluent-bit Kubernetes multiline log 를 테스트하기 위한 프로젝트입니다.  

<br/><br/><br/>

## Objective  

* 환경 변수로 파라미터를 전달하는 Go App 작성하기    
* Go언어로 작성한 App 이미지를 쿠버네티스에서 실행하기  
* fluent-bit Helm chart 배포하기  
* fluent-bit 디버깅하기  
  > fluent-bit 의 `latest-debug` Image Tag를 사용하면 Pod 에 쉘로 접근하여 디버깅할 수 있습니다.  
    


<br/><br/><br/>

## Prerequisites  

* Go +1.17 
* Docker  
* Kubernetes  
* Helm  
  * fluent-bit Helm chart 설치하기:  
    ```bash
    helm repo add fluent https://fluent.github.io/helm-charts
    ```
* kubectl run  
* (로컬 환경의 경우) Kind  

<br/><br/><br/>

## Overview  

*fluentlogger* App 은 쿠버네티스 Pod Logging 을 테스트하는 앱입니다.  

환경변수로 전달받는 파라미터는 다음과 같습니다:  

* `LOG_OUT`: Log 출력 파일 이름 (기본 값은 temp.json)
* `LOG_LEN`: 랜덤 Log Byte 길이 (기본 값은 64 Byte)
* `LOG_COUNT`: 랜덤 Log 출력 회수 (기본 값은 10 번)

<br/><br/><br/>

## Run  

fluent-bit 이 쿠버네티스 클러스터에 배포되지 않았다면, fluent-bit Helm chart 설치하기:  

```bash
helm install fluent-bit fluent/fluent-bit
```

환경변수 정의하기:  

```bash
export IMAGE_REPO={Image Registry}
```

이미지 빌드/게시하기:  

```bash
docker build --rm --no-cache --tag $IMAGE_REPO/fluentlogger --file ./assets/docker/Dockerfile .
docker push $IMAGE_REPO/fluentlogger
```

쿠버네티스에서 *fluentlogger* 실행하기:  

```
kubectl run fluentlogger --image=$IMAGE_REPO/fluentlogger \
  --restart=Never \
  --attach=true \
  --rm=true \
  --env="LOG_OUT=log.out" \
  --env="LOG_LEN=65536" \
  --env="LOG_COUNT=10"
```

> `kubectl run` `attach` flag:  
>   true인 경우, Pod 가 실행될 때까지 기다린 다음 `kubectl attach` 처럼 포드에 연결합니다.  
>   기본값은 *false* 지만, 
>   `-i/--stdin` 이 설정되었다면 *true* 입니다.  
>   `--restart=Never` 를 사용하면 컨테이너 프로세스의 종료 코드가 반환됩니다.  

쿠버네티스에서 실행 중인 *fluent-bit* 컨테이너의 쉘에 접근하기:  

```bash
kubectl exec -it {fluent-bit Pod} -- /bin/bash
```

만약 *fluentlogger* Pod 가 남아있다면 삭제하기:  

```bash
kubectl delete pod fluentlogger
```

<br/><br/><br/>

## References  

* [Fluent Bit v2.0](https://docs.fluentbit.io/manual/)  
* [fluent-bit 0.21.7 helm chart](https://artifacthub.io/packages/helm/fluent/fluent-bit)