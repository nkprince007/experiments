# MinIO k8s Deployment

## Why Object Storage?

Follow through [here](https://blog.minio.io/modern-data-lake-with-minio-part-1-716a49499533).

## Quickstart

```bash
kubectl apply -f minio-distributed-headless-service.yaml
kubectl apply -f minio-distributed-statefulset.yaml
kubectl apply -f minio-distributed-service.yaml
```

## Setting up (admin side)

- Install [mc](https://docs.min.io/docs/minio-client-quickstart-guide.html)

- Configure `mc` - add host configuration

```bash
mc config host add compaas <minio-service url> <accesskey> <secretkey>
# default keys need to be reconfigured
```

- Create a bucket

```bash
mc mb compaas/<bucket-name>
```

- Create new canned policy for specific bucket

```bash
mc admin policy add compaas <policy-name> bucketonly.json
```

- Create new user

```bash
mc admin user add compaas <user-accesskey> <user-secretkey>
```

- Apply policy

```bash
mc admin policy set compaas <policy-name> user=<user-accesskey>
```

## Setting up (user side)

- Install [mc](https://docs.min.io/docs/minio-client-quickstart-guide.html)

- Configure `mc` - add host configuration

```bash
mc config host add compaas <minio-service url> <user-accesskey> <user-secretkey>
# default keys need to be reconfigured
```

- Mirror user's local `data` directory to minio

```bash
mc mirror data <user-accesskey>/<bucketname>/data
```

## Cleanup

```bash
kubectl apply -f minio-distributed-headless-service.yaml
kubectl apply -f minio-distributed-statefulset.yaml
kubectl apply -f minio-distributed-service.yaml
```

## References

- [Distributed k8s deployment](https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/k8s-yaml.md#minio-distributed-server-deployment)
- [MinIO Quickstart guide](https://docs.min.io/docs/minio-quickstart-guide)
- [MinIO Client Quickstart guide](https://docs.min.io/docs/minio-client-quickstart-guide.html)
- [MinIO Admin Quickstart guide](https://docs.min.io/docs/minio-admin-complete-guide.html)
- [s3/minio object permissions](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-objects)
- [AWS Policy Generator](https://awspolicygen.s3.amazonaws.com/policygen.html)
