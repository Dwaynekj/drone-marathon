Use the Marathon plugin to deploy your app on Marathon's framework

The following parameters are used to configuration the notification:

* **config_file** - json payloads are sent herefor the post
* **host** - hostname of the Marathon master

The following is a sample Marathon configuration in your .drone.yml file:

```yaml
deploy:
  marathon:
    host: dev-elasticloadbal-marathon.amazonaws.com/service/marathon
    config_file: marathon-config.json
```
