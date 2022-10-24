cd deploy-kubernetes || exit

kubectl create -f namespace-cnetdisk.yaml
kubectl create -f pv-mongo.yaml
kubectl create -f pv-cnetdisk.yaml
kubectl create -f pvc-mongo.yaml
kubectl create -f pvc-cnetdisk.yaml
kubectl create -f deployment-mongo.yaml
kubectl create -f service-mongo.yaml
kubectl create -f deployment-cnetdiskserver.yaml
kubectl create -f service-cnetdiskserver.yaml
