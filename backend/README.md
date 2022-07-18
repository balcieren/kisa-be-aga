## Kisa Be Aga API Reference

#### Redirect to url

```http
  GET /v1/urls/${short_id}
```

| Parameter  | Type     | Description   |
| :--------- | :------- | :------------ |
| `short_id` | `string` | **Required**. |

#### Shorten url

```http
  POST /v1/urls
```

| Body  | Type     | Description                        |
| :---- | :------- | :--------------------------------- |
| `url` | `string` | **Required**. Includes regex match |
