# GovTech TAP Meteor - Technical Assessment
This project is a REST API for GovTech TAP Meteor technical assessment.
## Features
- Create household
- Add family member to household
- List households
- Show a unique household
- Search for household and recipents of grant disbursement
- Delete household
- Delete family member

## Tech
The REST API uses the following technologies:
- Golang - To host the REST API
- MySQL - Relational database to store household and family members data
 
## Installation
To run the URL Shortener, the following 
- [Golang](https://golang.org/doc/install)
- [MySQL](https://www.mysql.com/downloads/)

### 1. Edit the config file
```sh
cd config
```
Open config.yml in a text editor and change the values of MySQL credentials
### 2. Install the dependencies and start the backend server
```sh
go mod tidy
go run main.go
```

## Usage and endpoints
### Create household
#### Request

`POST /household/create`

    curl --location --request POST "http://localhost:8081/household/create" -H "Content-Type: application/json" -d "{\"households\" : [{\"housing_type\" : \"hello\"}]}"

`Possible "housing_type" fields: "HDB", "LANDED", "CONDOMINIUM"`
### Response

    {"message":"Household created"}

### 2. Add family member
#### Request
`POST /household/add_family_member`

    curl -L -X POST "http://localhost:8081/household/add_family_member" -H "Content-Type: application/json" --data-raw "{\"household_id\" : 4, \"name\" : \"Bobby\", \"gender\" : \"MALE\", \"marital_status\" : \"MARRIED\", \"spouse\" : \"Alice\", \"occupation_type\" : \"STUDENT\", \"annual_income\" : 1000.0, \"dob\" : \"08-08-1997\"}"

`Possible "gender" fields: "MALE", "FEMALE"`
`Possible "marital_status" fields: "HDB", "LANDED", "CONDOMINIUM"`
`Possible "occupation_type" fields: "HDB", "LANDED", "CONDOMINIUM"`

### Response

    {"message":"Added member"}
    

### 3. List households 
`GET /household/list_households`

    curl -i -H 'Accept: application/json' http://localhost:8081/household/list_households

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Mon, 02 Aug 2021 06:36:23 GMT
    Content-Length: 76
    
    {"message":[{"id":1,"housing_type":"LANDED"},{"id":2,"housing_type":"HDB"}]}
### 4. Show a unique household
`GET /household/query_household?id=x`

    curl -i -H 'Accept: application/json' http://localhost:8081/household/query_household?id=1

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Mon, 02 Aug 2021 06:43:53 GMT
    Content-Length: 44
    
    {"message":{"id":1,"housing_type":"LANDED"}}

### 5. Search for household and recipents of grant disbursement

`GET /grants/list_eligible_households?household_size=x&total_income=y`

    curl -i -H 'Accept: application/json' http://localhost:8081/grants/list_eligible_households?household_size=2&total_income=1000

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Mon, 02 Aug 2021 06:52:22 GMT
    Content-Length: 567
    
    {"message":
        {
            "student_encouragement_bonus":[
                {
                    "house_hold":
                        {
                            "id":3,
                            "housing_type":"HDB",
                            "family_members":[{
                                "id":6,
                                "household_id":3,
                                "name":"Bebe",
                                "gender":"MALE",
                                "marital_status":"SINGLE",
                                "spouse":"",
                                "occupation_type":"STUDENT",
                                "annual_income":100000,
                                "dob":"08-08-2007"
                                }]
                        },
                        "eligible_family_members":[{
                            "id":6,
                            "household_id":3,
                            "name":"Bebe",
                            "gender":"MALE",
                            "marital_status":"SINGLE",
                            "spouse":"",
                            "occupation_type":"STUDENT",
                            "annual_income":100000,
                            "dob":"08-08-2007"}]
                }],
            "family_togetherness_scheme":null,
            "elder_bonus":null,
            "baby_sunshine_grant":null,
            "yolo_gst_grant":null
        }
    }
    
### 6. Delete household
### Request
`DELETE /household/delete_household/:household_id`

    curl -L -X DELETE 'http://localhost:8081/household/delete_household/10'

### Response
    {"message":"Household deleted"}

### 7. Delete family member
### Request
`DELETE /household/delete_member/:member_id`

    curl -L -X DELETE "http://localhost:8081/household/delete_member/8"

### Response
    {"message":"Family member deleted"}
