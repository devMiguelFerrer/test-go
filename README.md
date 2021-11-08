# TEST GO SERVER

## ROUTES
```
POST http://localhost:8989/auth HTTP/1.1
Content-Type: application/json

{
    "name": "miguel",
    "password": "123456"
}
```
```
GET http://localhost:8989/data HTTP/1.1
Content-Type: application/json
Authorization: Bearer TheTokenIs465047283MoreOrLess
```

## CHALLENGE
- Create login form to sign in
- Get data with token
- Show data in a table
- Add system to hide NIF
- From table you can switch between different tactics to hide NIF (hide full, hide last 3 digits...) 