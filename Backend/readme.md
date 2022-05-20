# Backend

this readme gives an overview of the backend of the project. As it is lists all routes present in the backend

## Routes

### Project
+ > [GET] /api/v1/project
+ > [GET] /api/v1/project/:id
+ > [POST] /api/v1/project/create
+ > [PUT] /api/v1/project/update/:id
+ > [DELETE] /api/v1/project/delete/:id

### Ticket
+ > [GET] /api/v1/ticket
+ > [GET] /api/v1/ticket/:id
+ > [POST] /api/v1/ticket/create
+ > [PUT] /api/v1/ticket/update/:id
+ > [DELETE] /api/v1/ticket/delete/:id
+ > [GET] /api/v1/ticket/project/:projectID
+ > [GET] /api/v1/ticket/filter/status?type=
+ > [GET] /api/v1/ticket/filter/priority?type=

### User
+ > [GET] /api/v1/user?role=
+ > [GET] /api/v1/user/:username
+ > [POST] /api/v1/user/create
+ > [PUT] /api/v1/user/update/:username
+ > [DELETE] /api/v1/user/delete/:username
+ > [GET] /api/v1/user/validUsername/:username
+ > [GET] /api/v1/user/profile/:username

### Auth
+ > [POST] /api/v1/auth/login
+ > [POST] /api/v1/auth/logout
+ > [POST] /api/v1/auth/signup
+ > [POST] /api/v1/auth/changePassword
+ > [POST] /api/v1/auth/forgotPassword