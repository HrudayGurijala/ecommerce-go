# E-commerce Project with Golang

This project is an e-commerce application built with Golang, featuring various API endpoints for user management, product handling, shopping cart functionality, and order processing.

I will update with a postman collection soon.

## Getting Started

To get the project up and running, follow these steps:

```bash

# Run the main Golang application
go run main.go
```

## API Endpoints

Below is a detailed overview of the available API endpoints, including their HTTP methods, request bodies, and example responses.

### User Management

#### **1. User Signup**

Registers a new user in the system.

  - **URL:** `http://localhost:8000/users/signup`

  - **Method:** `POST`

  - **Request Body (JSON):**

    ```json
    {
      "first_name": "Hruday",
      "last_name": "Gurijala",
      "email": "hruday@gmail.com",
      "password": "hrudaygurijala",
      "phone": "+4534545435"
    }
    ```

  - **Response:** `Successfully Signed Up!!`

#### **2. User Login**

Authenticates a user and returns user details along with authentication tokens.

  - **URL:** `http://localhost:8000/users/login`

  - **Method:** `POST`

  - **Request Body (JSON):**

    ```json
    {
      "email": "hruday@gmail.com",
      "password": "hrudaygurijala"
    }
    ```

  - **Response (JSON):**

    ```json
    {
      "_id": "***********************",
      "first_name": "hruday",
      "last_name": "gurijala",
      "password": "$2a$14$UIYjkTfnFnhg4qhIfhtYnuK9qsBQifPKgu/WPZAYBaaN17j0eTQZa",
      "email": "hruday@gmail.com",
      "phone": "+4534545435",
      "token": "eyJc0Bwcm90b25vbWFpbC5jb20iLCJGaXJzdF9OYW1lIjoiam9zZXBoIiwiTGFzdF9OYW1lIjoiaGVybWlzIiwiVWlkIjoiNjE2MTRmNTM5ZjI5YmU5NDJiZDlkZjhlIiwiZXhwIjoxNjMzODUzNjUxfQ.NbcpVtPLJJqRF44OLwoanynoejsjdJb5_v2qB41SmB8",
      "Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnLCJVaWQiOiIiLCJleHAiOjE2MzQzNzIwNTF9.ocpU8-0gCJsejmCeeEiL8DXhFcZsW7Z3OCN34HgIf2c",
      "created_at": "2022-04-09T08:14:11Z",
      "updtaed_at": "2022-04-09T08:14:11Z",
      "user_id": "61614f539f29be942bd9df8e",
      "usercart": [],
      "address": [],
      "orders": []
    }
    ```

### Product Management (Admin)

#### **3. Add Product (Admin Only)**

Allows an administrator to add a new product to the database.

  - **URL:** `http://localhost:8000/admin/addproduct`

  - **Method:** `POST`

  - **Request Body (JSON):**

    ```json
    {
      "product_name": "Sony XM4 headphones",
      "price": 2500,
      "rating": 10,
      "image": "xm4.jpg"
    }
    ```

  - **Response:** `Successfully added our Product Admin!!`

### Product Browse

#### **4. View All Products**

Retrieves a list of all products available in the database.

  - **URL:** `http://localhost:8000/users/productview`

  - **Method:** `GET`

  - **Response (JSON):**

    ```json
    [
      {
        "Product_ID": "6153ff8edef2c3c0a02ae39a",
        "product_name": "sony XM4 headphones",
        "price": 2500,
        "rating": 10,
        "image": "xm4.jpg"
      },
      {
        "Product_ID": "616152679f29be942bd9df8f",
        "product_name": "logitech g34",
        "price": 900,
        "rating": 5,
        "image": "logi.jpg"
      },
      {
        "Product_ID": "616152ee9f29be942bd9df90",
        "product_name": "iphone 13",
        "price": 1700,
        "rating": 4,
        "image": "ipho.jpg"
      },
      {
        "Product_ID": "616153039f29be942bd9df92",
        "product_name": "acer predator",
        "price": 3000,
        "rating": 10,
        "image": "acer.jpg"
      }
    ]
    ```

#### **5. Search Product by Regex**

Searches for products whose names match a given regex pattern.

  - **URL:** `http://localhost:8000/users/search?name=al` (Example: searches for products containing "al")

  - **Method:** `GET`

  - **Response (JSON):**

    ```json
    [
      {
        "Product_ID": "616152fa9f29be942bd9df91",
        "product_name": "Alienware x15",
        "price": 1500,
        "rating": 10,
        "image": "1.jpg"
      },
      {
        "Product_ID": "616153039f29be942bd9df92",
        "product_name": "sony tv",
        "price": 300,
        "rating": 10,
        "image": "1.jpg"
      }
    ]
    ```

### Shopping Cart

#### **6. Add Product to Cart**

Adds a specified product to a user's shopping cart.

  - **URL:** `http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx`
  - **Method:** `GET`
  - **Note:** This action corresponds to a MongoDB query that updates the user's cart.

#### **7. Remove Item From Cart**

Removes a specified product from a user's shopping cart.

  - **URL:** `http://localhost:8000/removeitem?id=xxxxxxx&userID=xxxxxxxxxxxx`
  - **Method:** `GET`

#### **8. List Items in User's Cart and Total Price**

Retrieves all items currently in a user's cart along with the total price.

  - **URL:** `http://localhost:8000/listcart?id=xxxxxxuser_idxxxxxxxxxx`
  - **Method:** `GET`

### Address Management

#### **9. Add Address**

Adds a new address to a user's profile. The user is limited to two addresses (home and work).

  - **URL:** `http://localhost:8000/addadress?id=user_id`

  - **Method:** `POST`

  - **Request Body (JSON):**

    ```json
    {
      "house_name": "white house",
      "street_name": "white street",
      "city_name": "washington",
      "pin_code": "332423432"
    }
    ```

#### **10. Edit Home Address**

Updates a user's home address.

  - **URL:** `http://localhost:8000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx`
  - **Method:** `PUT`
  - **Request Body (JSON):** (Same as Add Address)

#### **11. Edit Work Address**

Updates a user's work address.

  - **URL:** `http://localhost:8000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx`
  - **Method:** `PUT`
  - **Request Body (JSON):** (Same as Add Address)

#### **12. Delete Addresses**

Deletes both home and work addresses for a given user.

  - **URL:** `http://localhost:8000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx`
  - **Method:** `GET`

### Order Processing

#### **13. Cart Checkout and Place Order**

Processes the items in the user's cart as an order. After placing the order, the items are removed from the cart.

  - **URL:** `http://localhost:8000/cartcheckout?id=xxuser_idxxx`
  - **Method:** `GET`

#### **14. Instantly Buy Product**

Allows a user to directly purchase a single product without adding it to the cart first.

  - **URL:** `http://localhost:8000/instantbuy?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx`
  - **Method:** `GET`