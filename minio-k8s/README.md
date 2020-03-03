# MinIO k8s Deployment

## Quickstart

```bash
kubectl apply -f https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/minio-distributed-headless-service.yaml?raw=true
kubectl apply -f https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/minio-distributed-statefulset.yaml?raw=true
kubectl apply -f https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/minio-distributed-service.yaml?raw=true
```

## Setting up (admin side)

- Install [mc](https://docs.min.io/docs/minio-client-quickstart-guide.html)

- Create a bucket

```bash
mc mb <TARGET>/<BUCKETNAME>
```

- Create new canned policy for a bucket (e.g.: getonly)

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": [
        "arn:aws:s3:::my-bucketname/*"
      ],
      "Sid": ""
    }
  ]
}
```

```bash
mc admin policy add <TARGET> <POLICYNAME> <POLICYFILE>
```

- Create new user

```bash
mc admin user add <TARGET> <ACCESSKEY> <SECRETKEY>
```

- Apply policy

```bash
mc admin policy set <TARGET> <POLICYNAME> user=<ACCESSKEY>
```

## Cleanup

```bash
kubectl delete -f https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/minio-distributed-headless-service.yaml?raw=true
kubectl delete -f https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/minio-distributed-statefulset.yaml?raw=true
kubectl delete -f https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/minio-distributed-service.yaml?raw=true
```

## References

- [Distributed k8s deployment](https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/k8s-yaml.md#minio-distributed-server-deployment)
- [MinIO Quickstart guide](https://docs.min.io/docs/minio-quickstart-guide)
- [MinIO Client Quickstart guide](https://docs.min.io/docs/minio-client-quickstart-guide.html)
- [MinIO Admin Quickstart guide](https://docs.min.io/docs/minio-admin-complete-guide.html)
- [s3/minio object permissions](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-objects)
<!-- - [AWS Policy Generator](https://awspolicygen.s3.amazonaws.com/policygen.html) -->
