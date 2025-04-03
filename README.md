# puorg-god
## Setup project chạy local sử dụng docker:
Ở đây sử dụng makefile để quản lí các câu lệnh chạy project
1. Clone the project từ github về 

2. Để chạy project local sử dụng docker cần thực hiện 2 thao tác sau:
- Lệnh này sẽ giúp chạy mysql và redis, tạo database, tạo table 
```bash
  make setup
```

- Run the project
```bash
  make run
```

3. Tạo file migration
```bash
  make create-migration
```

4. Run file migration
```bash
  make migrate-up
```

5. Stop file migration
```bash
  make migrate-down
```


6. Stop the project
```bash
  make teardown
```

## Docs API:
### 1. Api login
- Method: POST
- URL: /api/auth/login
- Body:
```json
{
  "email": "test@gmail.com",
  "password": "123456"
}
```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": {
        "access_token": ""
    }
}
```

### 2. Api create user
- Method: POST
- URL: /api/auth/
- Body:
```json
{
    "username": "",
    "role": "user", //default là user
    "password": "",
    "email": ""
}
```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```

### 3. Forgot password
- Method: POST
- URL: /api/auth/forgot-password
- Body:
```json
{
    "email": ""
}
```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": "OTP has been sent to your email"
}
```

### 4. Verify OTP 
- Method: POST
- URL: /api/auth/verify-otp
- Body:
```json
{
    "email": "",
    "otp": ""
} 
```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```

### 5. Reset password
- Method: POST
- URL: /api/auth/reset-password
- Body:
```json
{
    "email": "",
    "password": ""
}
```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```

### 6. Change password
- Method: POST
- URL: /api/auth/change-password
- Body:
```json
{
    "new_password": "",
    "old_password": ""
}
```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```

### 7. Update user
- Method: PUT
- URL: /api/user/{id}
- Body:
```json
{
    "username": ""
}
```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```
### 8. Delete user
- Method: DELETE
- URL: /api/user/{id}
- Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```

### 9. Create recipe
- Method: POST
- URL: /api/recipes/
- Body:
```json
{
  "title": "Bún Đậu",
  "description": "Một món ăn ngon nhanh",
  "image_url": "https://example.com/bun-bo-hue.jpg",
  "cuisine": "Japan",
  "ingredients": [
    {
      "name": "Đậu",
      "quantity": "200g"
    },
    {
      "name": "Bún",
      "quantity": "300g"
    },
    {
      "name": "hành",
      "quantity": "3 cây"
    }
  ],
  "instructions": [
    {
      "step": 1,
      "content": "Rửa sạch bún"
    },
    {
      "step": 2,
      "content": "Nấu nước dùng với sả, gừng, hành tím trong 1 giờ. "
    },
    {
      "step": 3,
      "content": "Rán đậu"
    }
  ]
}

```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```

### 10. Get all recipes
- Method: GET
- URL: /api/recipes/

Params:
- page: number
- limit: number
- title: string
- cuisine: string
- ingredients: string //example: "Đậu, Bún"

Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": {
    "code": 200000,
    "msg": "Success",
    "data": {
        "recipes": [
            {
                "id": 2,
                "title": "Bún Cá",
                "image_url": "https://example.com/bun-bo-hue.jpg",
                "cuisine": "japan",
                "created_at": "2025-04-03T23:08:45+07:00",
                "ingredients": [
                    {
                        "id": 4,
                        "name": "cá"
                    },
                    {
                        "id": 1,
                        "name": "bún"
                    },
                    {
                        "id": 6,
                        "name": "hành"
                    }
                ]
            },
            {
                "id": 3,
                "title": "Bún Đậu",
                "image_url": "https://example.com/bun-bo-hue.jpg",
                "cuisine": "japan",
                "created_at": "2025-04-03T23:09:30+07:00",
                "ingredients": [
                    {
                        "id": 7,
                        "name": "đậu"
                    },
                    {
                        "id": 1,
                        "name": "bún"
                    },
                    {
                        "id": 6,
                        "name": "hành"
                    }
                ]
            },
            {
                "id": 4,
                "title": "Bún Đậu",
                "image_url": "https://example.com/bun-bo-hue.jpg",
                "cuisine": "japan",
                "created_at": "2025-04-03T23:13:35+07:00",
                "ingredients": [
                    {
                        "id": 7,
                        "name": "đậu"
                    },
                    {
                        "id": 1,
                        "name": "bún"
                    },
                    {
                        "id": 6,
                        "name": "hành"
                    }
                ]
            }
        ],
        "total_page": 1,
        "limit": 10,
        "page": 1
    }
}
```

### 11. Get recipe by id
- Method: GET
- URL: /api/recipes/{id}

Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": {
    "code": 200000,
    "msg": "Success",
    "data": {
        "id": 2,
        "title": "Bún Cá",
        "description": "Một món ăn ngon",
        "image_url": "https://example.com/bun-bo-hue.jpg",
        "cuisine": "japan",
        "created_at": "2025-04-03T23:08:45+07:00",
        "updated_at": "2025-04-03T23:08:45+07:00",
        "ingredients": [
            {
                "id": 1,
                "name": "bún",
                "quantity": "300g"
            },
            {
                "id": 4,
                "name": "cá",
                "quantity": "200g"
            },
            {
                "id": 6,
                "name": "hành",
                "quantity": "3 cây"
            }
        ],
        "instructions": [
            {
                "id": 3,
                "step": 1,
                "content": "Rửa sạch thịt bò và cắt lát."
            },
            {
                "id": 4,
                "step": 2,
                "content": "Nấu nước dùng với sả, gừng, hành tím trong 1 giờ. "
            }
        ]
    }  
}
```

### 12. Delete recipe
- Method: DELETE
- URL: /api/recipes/{id}

Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```

### 13. Update recipe
- Method: PUT
- URL: /api/recipes/{id}

Body:
```json
{
    "title": "Bún Bò Huế",
    "description": "Một món ăn truyền thống của miền Trung Việt Nam.",
    "image_url": "https://example.com/bun-bo-hue.jpg",
    "cuisine": "China",
    "ingredients": [
        {
            "name": "Đậu",
            "quantity": "200g"
        },
        {
            "name": "Bún",
            "quantity": "300g"
        },
        {
            "name": "hành",
            "quantity": "3 cây"
        }
    ],
    "instructions": [
        {
            "step": 1,
            "content": "Rửa sạch bún"
        },
        {
            "step": 2,
            "content": "Nấu nước dùng với sả, gừng, hành tím trong 1 giờ. "
        },
        {
            "step": 3,
            "content": "Rán đậu"
        }
    ]
}
```
Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": null
}
```

### 14. Get all ingredients
- Method: GET
- URL: /api/ingredients/

Params:
- page: number
- limit: number
- name: string

Response:
```json
{
    "code": 200000,
    "msg": "Success",
    "data": {
    "code": 200000,
    "msg": "Success",
    "data": {
        "ingredients": [
            {
                "id": 1,
                "name": "bún"
            },
            {
                "id": 2,
                "name": "thịt bò"
            },
            {
                "id": 3,
                "name": "sả"
            },
            {
                "id": 4,
                "name": "cá"
            },
            {
                "id": 6,
                "name": "hành"
            },
            {
                "id": 7,
                "name": "đậu"
            }
        ],
        "total_page": 1,
        "limit": 10,
        "page": 1
    }
}
}
```

Cần xem lại phần search list của recipe