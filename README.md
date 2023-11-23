# MSIB5 Hacktiv8 - Final Project 2: MyGram

Berikut ini merupakan pengerjaan final project ke-2 dari Hacktiv8. Aplikasi ini bernama bernama MyGram, yang mana pengguna aplikasi ini dapat mengunggah foto dan membuat komentar untuk foto dari pengguna lain. Aplikasi ini akan dilengkapi dengan proses CRUD.

## Nama
 - Arrayyan Alaya Azzamachzachrie - GLNG-KS07-024

## User Login
```
Email : arrayyan@gmail.com
Password : arrayyan
```

## Endpoint
Di bawah ini merupakan semua endpoint yang dapat diakses di aplikasi ini.

### Users
 
 Berikut ini merupakan endpoint-endpoint yang dapat diakses untuk tabel Users:
 
 | Method | URL |
| ------ | ------ |
| POST | [https://msib5-hacktiv8-mygram.up.railway.app/users/register] |
| POST | [https://msib5-hacktiv8-mygram.up.railway.app/users/login] |
| PUT | [https://msib5-hacktiv8-mygram.up.railway.app/users] |
| DELETE | [https://msib5-hacktiv8-mygram.up.railway.app/users] |

###### Prosedur request users

POST Register User
 ```sh
{
    "age": integer,
    "email": "string",
    "password": "string",
    "username": "string"
}
```
#
POST Login User
 ```sh
{
    "email": "string",
    "password": "string"
}
```
#

PUT User

-Bearer Token <br />
-Param userId
 ```sh
{
    "email": "string",
    "username": "string"
}
```
#
DELETE User

-Authorization: Bearer Token

> Note: Untuk method PUT dan DELETE diperlukan autorisasi, yang mana memerlukan Bearer Token untuk dimasukkan terlebih dahulu. Token didapatkan melalui response pengguna saat melakukan login.
#


### Photos
  Berikut ini merupakan endpoint-endpoint yang dapat diakses untuk tabel Photos:

 | Method | URL |
| ------ | ------ |
| POST | [https://msib5-hacktiv8-mygram.up.railway.app/photos] |
| GET | [https://msib5-hacktiv8-mygram.up.railway.app/photos] |
| PUT | [https://msib5-hacktiv8-mygram.up.railway.app/photos/:id] |
| DELETE | [https://msib5-hacktiv8-mygram.up.railway.app/photos/:id] |

###### Prosedur request photos

POST Photo

-Bearer Token
 ```sh
{
    "title": "string",
    "caption": "string",
    "photo_url": "string"
}
```
#
GET Photo

-Bearer Token

#
PUT Photo

-Bearer Token  <br />
-Param photoId
 ```sh
{
    "title": "string",
    "caption": "string"
    "photo_url": "string"
}
```
#
DELETE 

-Bearer Token  <br />
-Param PhotoId
> Note: Seluruh method diperlukan autorisasi, yang mana perlu memasukan Bearer Token terlebih dahulu. Token didapatkan melalui response pengguna saat melakukan login. Untuk method PUT dan DELETE hanya bisa dilakukan oleh pengguna yang mengunggah foto dan perlu menyertakan parameter Id foto pada URL.
#


### Comments
  Berikut ini merupakan endpoint-endpoint yang dapat diakses untuk tabel Comments:

 | Method | URL |
| ------ | ------ |
| POST | [https://msib5-hacktiv8-mygram.up.railway.app/comments] |
| GET | [https://msib5-hacktiv8-mygram.up.railway.app/comments] |
| PUT | [https://msib5-hacktiv8-mygram.up.railway.app/comments/:id] |
| DELETE | [https://msib5-hacktiv8-mygram.up.railway.app/comments/:id] |

###### Prosedur request comments

POST Comment

-Bearer Token
 ```sh
{
    "message": "string",
    "photo_id": integer
}
```
#
GET Comment

-Bearer Token

#
PUT Comment

-Bearer Token  <br />
-Param commentId
 ```sh
{
    "message": "string"
}
```
#
DELETE 

-Bearer Token  <br />
-Param commentId

> Note: Seluruh method diperlukan autorisasi, yang mana perlu memasukan Bearer Token terlebih dahulu. Token didapatkan melalui response pengguna saat melakukan login. Untuk method PUT dan DELETE hanya bisa dilakukan oleh pengguna yang mengunggah komentar dan perlu menyertakan parameter Id komentar pada URL. 
#


 ### SocialMedias
  Berikut ini merupakan endpoint-endpoint yang dapat diakses untuk tabel SocialMedias:

 | Method | URL |
| ------ | ------ |
| POST | [https://msib5-hacktiv8-mygram.up.railway.app/socialmedias] |
| GET | [https://msib5-hacktiv8-mygram.up.railway.app/socialmedias] |
| PUT | [https://msib5-hacktiv8-mygram.up.railway.app/socialmedias/:id] |
| DELETE | [https://msib5-hacktiv8-mygram.up.railway.app/socialmedias/:id] |

###### Prosedur request socialmedias

POST SocialMedia

-Bearer Token
 ```sh
{
    "name": "string",
    "social_media_url": "string"
}
```
#
GET SocialMedia

-Bearer Token

#
PUT SocialMedia

-Bearer Token  <br />
-Param socialMediaId
 ```sh
{
    "name": "string",
    "social_media_url": "string"
}
```
#
DELETE 

-Bearer Token  <br />
-Param socialMediaId


> Note: Seluruh method diperlukan autorisasi, yang mana perlu memasukan Bearer Token terlebih dahulu. Token didapatkan melalui response pengguna saat melakukan login. Untuk method PUT dan DELETE hanya bisa dilakukan oleh pengguna yang mengunggah sosial media pada URL dan perlu menyertakan parameter Id sosial media tersebut.
