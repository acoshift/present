# Note

1. Create k8s cluster

1. Build and push our app to gcr.io

    ```bash
    $ make docker
    $ docker tag postlist gcr.io/acoshift-k8s/postlist:1.0.0
    $ docker push gcr.io/acoshift-k8s/postlist:1.0.0
    ```

1. Run nginx-ingress in k8s

    ```bash
    $ kubectl create -f nginx.yaml
    ```

1. Setup DNS from LoadBalancer IP

    ```bash
    $ kubectl get svc/nginx-ingress
    ```

1. Run CockroachDB

    ```bash
    $ kubectl create -f cockroachdb.yaml
    $ kubectl create -f cockroachdb-init.yaml
    ```

1. Run our App

    ```bash
    $ kubectl create -f app.yaml
    ```

1. Goto http://postlist.acoshift.me/

1. Run kube-lego

    ```bash
    $ kubectl create -f kube-lego.yaml
    ```

1. Setup Cloud Builder

    - Add Container Engine Developer Role to cloudbuild service account
    - Add Trigger

1. Cleanup

    - Delete k8s resources
    ```bash
    $ kubectl delete \
      -f app.yaml \
      -f kube-lego.yaml \
      -f cockroachdb.yaml \
      -f nginx.yaml
    ```

    - Delete k8s cluster
