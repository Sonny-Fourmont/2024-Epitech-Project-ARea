# AREA
# ⚠️ Docker and .env required ⚠️
## Installation
1- Clone the repository:
```bash
git clone git@github.com:EpitechPromo2027/B-DEV-500-LYN-5-1-area-sonny.fourmont.git
```
### Fill the .env.example file with your differents tokens (remember to rename this file to .env)
2- To run dev version:
```bash
docker compose -f docker-compose.test.yml up --build
```
3- To run prod version:
```bash
docker compose -f docker-compose.prod.yml up --build
```

## How to use the API:
### All the API routes are based on http://localhost:8080/
### User routes:
-  **HEADER FOR LOGIN ->** `"Content-Type: application/json"`

- POST `/user/register` Register users in database

- GET `/user/{id}` Get a user by his ID

- POST `/user/login` Login user from the database
### Applet routes (Applet = Action -> Reaction):
- POST `/applet` Create an applet

- GET `/applet` Get an applet

- DELETE `/applet` Delete an applet
### Services Routes:
- ### Google routes:
    - GET `/google` (LOGIN REQUIRED) return the google token

    - GET `/google/register` Get a google register link

    - GET `/google/login` Get a google login link
- ### Github routes:
    - GET `/github` (LOGIN REQUIRED) return the github token

    - GET `/github/register` Get a github register link

    - GET `/github/login` Get a github login link
- ### Microsoft routes:
    - GET `/microsoft` (LOGIN REQUIRED) return the microsoft token

    - GET `/microsoft/register` Get a microsoft register link

    - GET `/microsoft/login` Get a microsoft login link
- ### Spotify routes:
    - GET `/spotify` (LOGIN REQUIRED) return the spotify token

    - GET `/spotify/login` Get a spotify login link

## Technologies Used
- ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
- ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
- ![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
## FrameWork and Libraries Used
### Golang:
- Gin
- GoDotenv

## Credits
Developed by:
<table>
    <tbody>
        <tr>
            <td align="center" valign="top" width="14.28%"><a href="https://github.com/Sonny-Fourmont"><img src="https://avatars.githubusercontent.com/u/114910491?v=4" width="100px;" alt="Sonny FOURMONT" style="border-radius: 50%; !important"/><br /><sub><b>Sonny<br>FOURMONT</b></sub></a><br /></td>
            <td align="center" valign="top" width="14.28%"><a href="https://github.com/Matribuk"><img src="https://avatars.githubusercontent.com/u/75017908?v=4" width="100px;" alt="Antonin LEPREST" style="border-radius: 50%; !important"/><br /><sub><b>Antonin<br>LEPREST</b></sub></a><br /></td>
        </tr>
    </tbody>
</table>
