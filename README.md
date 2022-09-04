# OpenFoodFacts API wrapper

# Run
``` bash
make run
```

or
``` bash
go run main.go
```


## Query stuff


### List categories
```
http://localhost:8081/categories
```
show all categories, response will contain a to list all products under one single category.

### Products in category
```
http://localhost:8081/category/chocolates
```
show products in category "chocolates", result will contain links to single product detail.

### Product detail
```
http://localhost:8081/products/7614500010310
```
