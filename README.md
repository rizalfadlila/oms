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
