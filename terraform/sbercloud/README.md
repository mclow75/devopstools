# Подключение к CloudRU

## Зеркало Terraform Cloud.ru Advanced

Официальное зеркало Terraform от Cloud.ru — `https://terraform.cloud.ru/`
[URL](https://cloud.ru/docs/terraform/ug/topics/guides__mirrors?source-platform=Advanced)
В конфигурационном файле с расширением tf для переменной source присвойте значение "sbercloud-terraform/sbercloud".

Блок с инициализацией провайдера:

```tf
terraform {
  required_providers {
    sbercloud = {
      source  = "sbercloud-terraform/sbercloud" # Initialize Advanced provider
    }
  }
}
```
