## requirement
```
kubebuilder == 3.0.0
```

## step
```
# create project
1. kubebuilder3.0.0 init --skip-go-version-check --repo github.com/lazywhite/guestbook-controller --domain example.org

# create api
2. kubebuilder3.0.0 create api --group webapp --version v1 --kind GuestBook

# create webhook
kubebuilder3.0.0 create webhook --group webapp --version v1 --kind GuestBook --webhook-version v1beta1 --defaulting --programmatic-validation

```

## issue
```
controller-gen生成的crd apiversion 总是apiextensions.k8s.io/v1, 修改PROJECT里面的resources/api/crdVersion没有用 
生成的ValidateWebhookConfiguration, 名字都叫validating-webhook-configuration, 可以通过设置namePrefix来解决
    config/default/kustomization.yaml
        namePrefix: "guestbook-"
```

