# Product Catalog
REST API for a catalog of products.  

The application contains:
- Product categories;
- Products which belong to some category (one product may belong to one category);  

The following actions are implemented:
- Getting the list of all categories;
- Getting the list of products of the concrete category;
- Create/update/delete of category;
- Create/update/delete of product;

## Project structure  
```bash
. 
│── api
│   ├── controllers
│   │   ├── base.go
│   │   ├── category_controller.go
│   │   ├── product_controller.go
│   │   └── routes.go
│   ├── models
│   │   ├── Category.go
│   │   └── Product.go
│   └── utils
│       ├── errors.go
│       └── response.go
│── tests
│   ├── controllertests
│   │   ├── category_controller_test.go
│   │   ├── product_controller_test.go
│   │   └── controller_test.go
│   └── modeltests
│       ├── category_model_test.go
│       ├── product_model_test.go
│       └── model_test.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Installation
Get project repository:
```bush
$ go get github.com/stdevi/catalog
```

## How to run tests
```bush
cd go/src/github.com/stdevi/catalog
$ go test ./...
```

## How to build and deploy
```bush
cd go/src/github.com/stdevi/catalog
$ go build
$ ./catalog
```
Now you can access API locally at http://localhost:8080.

## Request & Response Examples

### Category API Resources
Resource | Description
------------ | -------------
[```GET /api/categories```](#GET-apicategories) | Getting the list of all categories
[```POST /api/categories```](#POST-apicategories) | Create category
[```PUT /api/categories/[id]```](#PUT-apicategoriesid) | Update category by id
```DELETE /api/categories/[id]``` | Delete category by id
### Product API Resources
Resource | Description
------------ | -------------
[```GET /api/products```](#GET-apiproducts) | Getting the list of all products
[```GET /api/products/category/[id]```](#GET-apiproductscategoryid) | Getting the list of products of the concrete category
[```POST /api/products```](#POST-apiproducts) | Create product
[```PUT /api/products/[id]```](#PUT-apiproductsid) | Update product by id
```DELETE /api/products/[id]``` | Delete product by id

#### GET /api/categories
- Response body:
    ```json
    [
        {
            "id": 1,
            "name": "Laptops"
        },
        {
            "id": 2,
            "name": "Devices"
        }
    ]
    ```

#### POST /api/categories
- Request body:
    ```json
    {
        "name": "Tools"
    }
    ```
- Response body:
    ```json
    {
        "id": 3,
        "name": "Tools"
    }
    ```

#### PUT /api/categories/[id]
- Request body:
    ```json
    {
        "name": "Pro Tools"
    }
    ```
- Response body:
    ```json
    {
        "id": 1,
        "name": "Pro Tools"
    }
    ```

#### GET /api/products
- Response body:
    ```json
    [
        {
            "id": 1,
            "name": "Laptop X42",
            "description": "X42 description",
            "price": 2000,
            "category_id": 1,
            "category": {
                "id": 1,
                "name": "Laptops"
            }
        },
        {
            "id": 2,
            "name": "Dell Monitor",
            "description": "Monitor description",
            "price": 400,
            "category_id": 2,
            "category": {
                "id": 2,
                "name": "Devices"
            }
        },
        {
            "id": 3,
            "name": "Laptop X41",
            "description": "X41 description",
            "price": 42,
            "category_id": 1,
            "category": {
                "id": 1,
                "name": "Laptops"
            }
        }
    ]
    ```

#### GET /api/products/category/[id]
- Response body:
    ```json
    [
        {
            "id": 1,
            "name": "Laptop X42",
            "description": "X42 description",
            "price": 2000,
            "category_id": 1,
            "category": {
                "id": 1,
                "name": "Laptops"
            }
        },
        {
            "id": 3,
            "name": "Laptop X41",
            "description": "X41 description",
            "price": 42,
            "category_id": 1,
            "category": {
                "id": 1,
                "name": "Laptops"
            }
        }
    ]
    ```

#### POST /api/products
- Request body:
    ```json
    {
        "name": "Go Product",
        "description" : "Go description",
        "price": 42,
        "category_id": 2
    }
    ```
- Response body:
    ```json
    {
        "id": 4,
        "name": "Go Product",
        "description": "Go description",
        "price": 42,
        "category_id": 2,
        "category": {
            "id": 2,
            "name": "Devices"
        }
    }
    ```

#### PUT /api/products/[id]
- Request body:
    ```json
    {
        "name": "Updated X42",
        "description": "Updated X42 description",
        "price": 2000,
        "category_id": 1
    }
    ```
- Response body:
    ```json
    {
        "id": 1,
        "name": "Updated X42",
        "description": "Updated X42 description",
        "price": 2000,
        "category_id": 1,
        "category": {
            "id": 1,
            "name": "Laptops"
        }
    }
    ```