# oms

### Prerequisite

- go 1.18+
- docker
- docker-compose

### Install Dependencies
```bash
setup-env
```

### Preparation before run
```bash
make sync-dep
```

```bash
make init-config
```

```bash
make migrate
```

### Read CSV
default location file path = files/csv
```bash
make read-csv SOURCE={{filepath}}
```

### Run Rest
```bash
make run-rest
```

### API
API get detail order
```curl
curl --location --request GET 'http://localhost/v1/order/{{orderId}}/detail'
```
