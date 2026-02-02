# Golang MySQL Visitor Counter Web App
```
How to build docker visit-counter image and push to docker hub registry

docker login -u shehbab

git clone git@github.com:Shehbab-Kakkar/GolangMySQLCounterApp.git

cd GolangMySQLCounterApp/goappbuilder

docker build -t shehbab/visit-counter .

docker push shehbab/visit-counter:latest

-----
How to deploy:

git clone git@github.com:Shehbab-Kakkar/GolangMySQLCounterApp.git

cd GolangMySQLCounterApp

kubectl apply -f mysqlmanifest/

kubectl apply -f goappmanifest/

------
How to delete visitor counter app from K8s :

kubectl delete statefulset.apps/mysql
kubectl delete deployment.apps/visit-counter
kubectl delete svc,cm,secret --all
kubectl delete pvc mysql-data-mysql-0
```
