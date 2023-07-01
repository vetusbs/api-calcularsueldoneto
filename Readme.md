### api-calcular sueldo neto

This is a simple golang application that exposes a rest api that allows you to calculate the net monthly salary based on your annual gross salary.

### How to run it

```go
go run .
```

### Example of usage

```bash
curl localhost:8080/netSalary -d '{ "gross_salary": 21000 }'
```