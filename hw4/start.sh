minikube start --kubernetes-version="v1.19.0"  --vm=true --cni=flannel --cpus=4 && \
istioctl install --set profile=demo && \
kubectl label namespace default istio-injection=enabled