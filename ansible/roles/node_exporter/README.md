# Установка Prometheus Node_Exporter на Linux

Используются,как RPM, так и DEB дистрибутивы.

## Структура

roles/
├── node_exporter/
│   ├── defaults/
│   │   └── main.yml
│   ├── handlers/
│   │   └── main.yml
│   ├── tasks/
│   │   └── main.yml
│   ├── templates/
│   │   └── node_exporter.service.j2
│   └── vars/
│       └── main.yml

