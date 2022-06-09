
# BugTrack

Track, Manage and resolve features and bugs

Universal bug tracker for everyone! BugTrack allows team members to collaborate, discuss and manage projects effectively.
## Features

- Team management
- Statistics
- User assignement
- Ticket labels
- Ticket history
- Role based organisation
- Change tracker


## Coming Soon
- Docker Support
- Mail Service
- Comments and Reactions
- Ticket History


## Tech Stack

**Frontend:** React, TailwindCSS, Chart.js, React-Icons

**Backend:** Go, Gorilla MUX, MongoDB, JWT, Bcrypt


## Application presentation and flow

### I - Authentication
Since the core application is protected through authentication and authorization, the first page you will directed to will be the login page

![signin](https://user-images.githubusercontent.com/66411529/172742755-4f95f041-b88d-4db0-be65-5124d8ed8274.jpg)

If you are a new user, you can signup as an admin for your company for the first time. If an admin already exists, only the admin can signup a new user.

![signup](https://user-images.githubusercontent.com/66411529/172744028-f1ab2e9c-26ca-4d61-afe7-e8cbbb67f734.jpg)

### II - Application 🔥

#### A - Dashboard
Upon successful authentication, you will be redirected to the main page. Here you can see user-management and systems-logs as the user is an admin.

![dashboard](https://user-images.githubusercontent.com/66411529/172742928-750d372e-5a20-42c4-9d92-6ed3ac80fac4.jpg)

The upper part of the page will be in charge to display the projects while the lower part will display the ticket related statistics to get an immediate idea of the state of the organization.


#### B - Project

#### Creation

From the dashboard, an admin
 can create a new project and assign team to it. Only the users linked to this project will have access to it and its tickets.

![newproject](https://user-images.githubusercontent.com/66411529/172743061-87420c57-4658-4a2d-ab61-3a4fafba7638.jpg)

#### Update / Delete

From the same page, by hitting the tripple dots in the actions section of the table, you are able to update or delete the projects. Upon using the corresponding button, you will be prompted with a modal asking you to confirm your action. Deleting a project will as well delete every ticket linked to the project.
![updateproject](https://user-images.githubusercontent.com/66411529/172743216-e98872cd-b4fe-4720-9b9f-0ffa605b6bbb.jpg)
![deleteproject](https://user-images.githubusercontent.com/66411529/172744176-775ed5c7-f150-409a-a6fa-fb4b3ea4f2e3.jpg)

#### C - Tickets

Upon clicking one of the purple project name, you will be taken to the tickets linked to the projects. Or you can click on Tickets from the sidebar to see all the tickets.

![projecttickets](https://user-images.githubusercontent.com/66411529/172743265-9f75eed9-1ff6-45ac-b3bc-dd5aa427806c.jpg)

#### Create

Clicking on the blue "New Ticket" button will open a modal allowing you to create the specific fields of your new ticket and assign the developers that will work on it.

![newticket](https://user-images.githubusercontent.com/66411529/172743329-5b38aa97-3745-45c4-8d17-f099223de495.jpg)

#### Update / Delete

Once again, by clicking on the three dots, you will gain access to the update and delete modal.

![updateticket](https://user-images.githubusercontent.com/66411529/172743474-ed3d25ed-26f8-4a7b-a66a-86ad5bccd85f.jpg)
![deleteticket](https://user-images.githubusercontent.com/66411529/172743484-45185092-f2bd-4a15-98e6-f273facb73ab.jpg)

#### Assigned Tickets

On the left hand side resides the main tabs of the application. If a user click on the ticket tab, it will take him to a page holding all ticket relatid to him/her. The way tickets would be fetch will depend on the role assigned to the user:
    
- <ins>Admin:</ins> Will fetch all the tickets of the system. Has the possibility to filter the result down to the tickets he/she created.
- <ins>Project Manager:</ins> Will fetch all the tickets linked to the projects the PM is assigned to. 
- <ins>Developer</ins>: Will fetch all the tickets they are assigned to.
- <ins>Submitter</ins>: Will fetch all the tickets they created.

![alltickets](https://user-images.githubusercontent.com/66411529/172743287-f8cc0935-097c-4ee1-8b0c-e05e950eeacb.jpg)

#### D - User Management
if the user is logged in as an Admin, he/she will have the option to click on the User management tab.
This page will fetch all the user in the system, display their name, email and roles, and allow the admin to perform diverse operations.

![usermanage](https://user-images.githubusercontent.com/66411529/172743560-3dfebcc6-ca91-410e-b411-d004c745e15a.jpg)

By clicking on the triple dot of each row, the admin will be able to perform three type of action:

#### Create
By clicking on "new user" button you can create a new user and assign them roles

![newuser](https://user-images.githubusercontent.com/66411529/172743752-669cec28-63ff-48cc-9c44-67603e102c74.jpg)

#### Update
Clicking on the update button will open a modal allowing you to update the user's information, including the role.

![updateuser](https://user-images.githubusercontent.com/66411529/172743767-5700ef01-fcfe-4e0a-8fdf-5f2f5910a0f2.jpg)


#### Lock / Unlock User
Users can be both locked and unlocked from the system. Still by clicking the triple dots the admin will be able to access the below modal. If the user is locked, the option will be defaulted to unlock and vice-versa.

![lockuser](https://user-images.githubusercontent.com/66411529/172743784-00a75714-2918-43f5-81fd-a7ce7b90c338.jpg)
![unlock](https://user-images.githubusercontent.com/66411529/172743859-8730d1ab-5255-4e03-b541-aaf444888f5a.jpg)


#### Delete User
Finally, the admin might want to remove permanently somebody from the organization. If he/she wish to do so, the triple dots also give access to a delete user modal. He/She, will be asked for confirmation and will then be able to validate the action.

![deleteuser](https://user-images.githubusercontent.com/66411529/172743930-dced8be4-1550-4981-8509-89e849c69f6e.jpg)

#### E - System Logs
All the actions of the users will be logged in the system. The admin will be able to see the logs of the system.

![logs](https://user-images.githubusercontent.com/66411529/172743944-9eb48225-d639-4cca-9b14-f21089433bd6.jpg)

and that wraps it up!

## Author

[@ArmaanKatyal](https://github.com/ArmaanKatyal)


## Show your support

Please ⭐️ this repository if you liked the project!


## License
Copyright © 2022 [@ArmaanKatyal](https://github.com/ArmaanKatyal)

This Project is [MIT](https://choosealicense.com/licenses/mit/) Licensed

