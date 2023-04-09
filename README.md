# APIfood
Данное API разрабатывалось на основе https://app.swaggerhub.com/apis-docs/serk87/APIfood/1.0.0 для приложения WSR Food, так как изначальное API перестало работать.

Изображения блюд были утеряны, осталось только одно, оно доступно по запросу: \
**GET /up/images/dish.png**


**POST /auth/register** - регистрация\
**Request body**:
```JSON 
{
    "email": string,
    "password": string,
    "login": string
} 
```

**Responses**:\
200 - OK\
400 - Bad Request

**POST /auth/login** - аутендификация в системе\
**Request body**:
```JSON
{
    "email": string,
    "password": string
}
```
**Responses**:\
200 - OK
```JSON
{
    "token": string
}
```
400 - Bad Request

**GET /dishes/version** - получить список версий\
**Responses**:\
200 - OK
```JSON
{
    "version": []string
}
```
400 - Bad Request

**GET /dishes** - получить список блюд\
**Param: ?version=x**\
**Responses**:\
200 - OK
```JSON
[
  {
    "dishId": string,
    "category": string,
    "nameDish": string,
    "price": string,
    "icon": string,
    "version": string
  }
]
```
400 - Bad Request

**POST /order** - отправить заказ\
**Request body**:
```JSON
{
  "address": string,
  "date": string,
  "dishes": [
    {
      "dishId": string,
      "count": int
    }
  ]
}
```
**Responses**:\
200 - OK\
400 - Bad Request

**GET /histories** - получить список версий\
**Responses**:\
200 - OK
```JSON
[
  {
    "address": string,
    "date": string,
    "dishes": [
      {
        "dishId": string,
        "count": int
      }
    ]
  }
]
```
400 - Bad Request
