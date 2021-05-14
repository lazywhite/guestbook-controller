## requirement
```
kubebuilder == 3.0.0
cert-manager
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
1. controller-gen生成的crd apiversion 总是apiextensions.k8s.io/v1, 修改PROJECT里面的resources/api/crdVersion没有用 
2. 生成的ValidateWebhookConfiguration, 名字都叫validating-webhook-configuration, 可以通过设置namePrefix来解决
    config/default/kustomization.yaml
        namePrefix: "guestbook-"
3. controller-gen 命令行工具脱离此项目使用
4. config/certmanager/certificate.yam
    spec.secretName  #注意修改
   config/default/manager_webhook_patch.yaml
    volumes:
        secretName # 注意保持一致
    
```

