
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

IMAGE

If you are a new user, you can signup as an admin for your company for the first time. If an admin already exists, only the admin can signup a new user.

IMAGE


### II - Application 🔥

#### A - Dashboard
Upon successful authentication, you will be redirected to the main page. Here you can see user-management and systems-logs as the user is an admin.

IMAGE

The upper part of the page will be in charge to display the projects while the lower part will display the ticket related statistics to get an immediate idea of the state of the organization.


#### B - Project

#### Creation

From the dashboard, an admin can create a new project and assign team to it. Only the users linked to this project will have access to it and its tickets.

IMAGE


#### Update / Delete

From the same page, by hitting the tripple dots in the actions section of the table, you are able to update or delete the projects. Upon using the corresponding button, you will be prompted with a modal asking you to confirm your action. Deleting a project will as well delete every ticket linked to the project.

IMAGE

#### C - Tickets

Upon clicking one of the purple project name, you will be taken to the tickets linked to the projects. Or you can click on Tickets from the sidebar to see all the tickets.

IMAGE


#### Create

Clicking on the blue "New Ticket" button will open a modal allowing you to create the specific fields of your new ticket and assign the developers that will work on it.

IMAGE

#### Update / Delete

Once again, by clicking on the three dots, you will gain access to the update and delete modal.


#### Assigned Tickets

On the left hand side resides the main tabs of the application. If a user click on the ticket tab, it will take him to a page holding all ticket relatid to him/her. The way tickets would be fetch will depend on the role assigned to the user:
    
- <ins>Admin:</ins> Will fetch all the tickets of the system. Has the possibility to filter the result down to the tickets he/she created.
- <ins>Project Manager:</ins> Will fetch all the tickets linked to the projects the PM is assigned to. 
- <ins>Developer</ins>: Will fetch all the tickets they are assigned to.
- <ins>Submitter</ins>: Will fetch all the tickets they created.

IMAGE

#### D - User Management
if the user is logged in as an Admin, he/she will have the option to click on the User management tab.
This page will fetch all the user in the system, display their name, email and roles, and allow the admin to perform diverse operations.

IMAGE

By clicking on the triple dot of each row, the admin will be able to perform three type of action:

#### Update
Clicking on the update button will open a modal allowing you to update the user's information, including the role.

IMAGE

#### Lock / Unlock User
Users can be both locked and unlocked from the system. Still by clicking the triple dots the admin will be able to access the below modal. If the user is locked, the option will be defaulted to unlock and vice-versa.

IMAGE

#### Delete User
Finally, the admin might want to remove permanently somebody from the organization. If he/she wish to do so, the triple dots also give access to a delete user modal. He/She, will be asked for confirmation and will then be able to validate the action.

IMAGE

#### E - System Logs
All the actions of the users will be logged in the system. The admin will be able to see the logs of the system.

IMAGE

and that wraps it up!

## Author

[@ArmaanKatyal](https://github.com/ArmaanKatyal)


## Show your support

Please ⭐️ this repository if you liked the project!


## License
Copyright © 2022 [@ArmaanKatyal](https://github.com/ArmaanKatyal)

This Project is [MIT](https://choosealicense.com/licenses/mit/) Licensed

