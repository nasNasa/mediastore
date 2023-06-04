# drop-box Application

# use-case

## User use-cases
### Register
user can register to application by 1-phone number 2-email 3-google account 4-github account 

### Login
user can log in to the application by 1-phone number or email and password 2-accounts


## storage use-cases
### Write
user can upload files with a specific types(...) 
until it reaches a certain size(?)

### Read
user can download from her/his folders only
user can make the access of her/his folders public or private for others

### Delete


# entity

## User
- ID
- Name
- Phone number
- isVerified
- Avatar
- Max amount of data
- Folders[]
- accountType

## File
- ID
- Type
- size
- UserId
- UserFolder

## Folder
- ID
- UserId
- FilesType(?)


