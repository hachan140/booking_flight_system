# GraphQL Sample API

### Note
- All API need authentication need add below section to header

  ```json{ "Authorization": "token_member_here"}```
- API need authentication will have role [MEMBER, ADMIN]                                                      
## Member
### Sign up new member
```graphql
mutation sign_up{
  member{
    sign_up(
    input: {email: "haha6@gmail.com", password: "234346789", phoneNumber: "03377961666", fullName: "Vu Ha", cid: "012347735645568"}
  ) {
    email
    fullName
  }
  }
}
```

### Sign up new admin
```graphql
mutation sign_up_admin {
  member {
    sign_up(
      input: {email: "admin2@gmail.com", password: "123456789", phoneNumber: "01234563467", fullName: "Admin", cid: "0123456677884545", memberType:ADMIN}
    ) {
      email
      fullName
    }
  }
}
```

### Log in: 
Role: ADMIN
```graphql
mutation admin_login{
  member{
    login(email:"admin1@gmail.com", password:"123456789"){
    	token
    	expired_at
  		}
  }
}
```
Role: MEMBER
```graphql
mutation admin_login{
  member{
    login(email:"admin2@gmail.com", password:"123456789"){
      token
      expired_at
    }
  }
}
```

### Call Self
Role: ADMIN, MEMBER
```graphql
mutation call_self{
  member{
    self{
    	email
    	fullName
    	createdAt
  	}
  }
}
```

### Find member by name
Role: ADMIN
```graphql
mutation find_member_by_name{
  member{
    find_member_by_name(name: "vu"){
      id
      createdAt
      updatedAt
      email
      phoneNumber
      fullName
      dob
      cid
    }
  }
}
```

### Change password
Role: ADMIN, MEMBER
```graphql
mutation change_password{
  member{
    change_password(oldPassword: "123456789", newPassword: "12345678")
  }
}
```

### Update member profile
Role: ADMIN, MEMBER
```graphql
mutation {
  member{
  update_member_profile(input:{fullName:"Hi Admin", cid:"32554675678", phoneNumber:"0987654321" dob:"2002-10-14T15:04:05+07:00"}){
    createdAt
    updatedAt
    email
    fullName
    cid
    phoneNumber
    dob
  }
  }
}
```

### Find member by email
Role: ADMIN
```graphql
mutation find_member_by_email{
  member{
    find_member_by_email(email: "havu14102@gmail.com"){
     		createdAt
        updatedAt
        email
        fullName
        cid
        phoneNumber
        dob
    }
  }
}
```

### List members
Role: ADMIN
```graphql
mutation list_members{
  member{
    list_members(first: 5, orderBy:{direction: ASC, field: FULL_NAME}, after: "gqFp0wAAAAAAAAARoXalQWRtaW4"){
     		edges{
          node{
            fullName
            email
            phoneNumber
            dob
            cid
          }
          cursor
        }
      pageInfo{
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      totalCount
    }
  }
}
```

## Customer
### Create customer
```graphql
mutation create_customer{
  customer{
    create_customer(input:{email: "Hachan14@gmail.com", phoneNumber: "01234567894", fullName:"ha ha", cid:"1354657683453"}){
      createdAt
      email
      fullName
      dob
      cid
    }
  }
}
```

### List customers
Role: ADMIN
```graphql
mutation list_customers{
  customer{
    list_customers(first: 5, orderBy:{direction: ASC, field: FULL_NAME}){
     		edges{
          node{
            fullName
            email
            phoneNumber
            dob
            cid
          }
          cursor
        }
      pageInfo{
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      totalCount
    }
  }
}
```

## PLANE
### Create Plane
Role: ADMIN
```graphql
mutation create_plane{
  plane{
    create_plane(input:{name:"VN Airlines", economyClassSlots: 200, businessClassSlots:50 }){
    name
    createdAt
    updatedAt
    economyClassSlots
    businessClassSlots
    status
  }
  }
}
```

### Find Plane by ID
Role: ADMIN
```graphql
mutation find_plane_by_id{
  plane{
    find_plane_by_id(id:1){
    name
  }
  }
}
```
### Update Plane
Role: ADMIN
```graphql
mutation{
  plane{
    update_plane(id: 1, input:{name:"Updated", economyClassSlots:200, businessClassSlots: 40, status:AVAILABLE}){
      id
      updatedAt
      name
      economyClassSlots
      businessClassSlots
      status
    }
  }
}
```
### List Planes
Role: ADMIN
```graphql
mutation list_planes{
  plane{
    list_planes(first: 2, orderBy:{direction: ASC, field: NAME}){
     		edges{
          node{
            id
            name
            economyClassSlots
            businessClassSlots
          }
          cursor
        }
      pageInfo{
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      totalCount
    }
  }
}
```

## Airport

### Create airport
Role: ADMIN
```graphql
mutation create_airport{
  airport{
    create_airport(input:{name:"Da Nang", lat: 15.8, long: 130.7}){
      createdAt
      updatedAt
      name
      lat
      long
    }
  }
}
```

### Update airport
Role: ADMIN
```graphql
mutation update_airport{
  airport{
    update_airport(id: 1, input:{name: "Tan Son Nhat", lat: 124.5, long: 32.5, }){
      id
      createdAt
      updatedAt
      name
      lat
      long
      
    }
  }
}
```

### Find airport by ID
Role: ADMIN
```graphql
mutation{
  airport{
    find_airport_by_id(id: 5){
    name
  }
  }
}
```

### List airports
Role: ADMIN
```graphql
mutation list_airports{
  airport{
    list_airports(first: 2, orderBy:{direction: ASC, field: NAME}){
     		edges{
          node{
            id
            name
            lat
            long
          }
          cursor
        }
      pageInfo{
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      totalCount
    }
  }
}
```

## Flight
### Create Flight
Role: ADMIN
```graphql
mutation{
  flight{
    create_flight(input:{name:"HCM-HN", departAt:"2022-10-23T10:00:00+07:00", landAt:"2022-10-23T12:00:00+07:00", fromID:1,  toID:3, planeID:2}){
    name
    createdAt
    fromAirport{
      name
    }
    toAirport{
      name
    }
    departAt
    landAt
    hasPlane{
      name
    }
  }
  }
}
```

### Update Flight status
Role: ADMIN
```graphql
mutation update_flight_status{
  flight{
    update_flight_status(id: 16, input:{status: SCHEDULED}){
      id
      name
      status
    }
  }
}
```

### Search Flight
```graphql
mutation{
  flight{
    search_flight(input:{from_airport:"Noi Bai", to_airport:"Tan Son Nhat", depart_at:"2022-10-16T00:00:00+07:00"}){
    name
    id
    createdAt
    fromAirport{
      name
    }
    toAirport{
      name
    }
    departAt
    landAt
    hasPlane{
      name
    }
  }
  }
}
```

### Cancel Flight
Role: ADMIN
```graphql
mutation cancel_flight{
  flight{
    cancel_flight(id: 3)
  }
}
```

### Find Airport by name (using when search flight)
```graphql
mutation find_airport_by_name{
  flight{
    find_airport_by_name(name: "Tan Son Nhat"){
      name
      createdAt
      updatedAt
      lat
      long
    }
  }
}
```

### Find Flight by ID
```graphql
mutation find_flight_by_id{
  flight{
    find_flight_by_id(id:18){
      name
      departAt
      landAt
      availableEcSlot
      availableBcSlot
      hasPlane{
        name
        economyClassSlots
        businessClassSlots
      }
      hasBooking{
        hasCustomer{
          fullName
          phoneNumber
        }
      }
    }
  }
}
```

### List Flight
Role: Admin
```graphql
mutation list_flights{
  flight{
    list_flights(first: 5, orderBy:{direction: ASC, field: NAME}){
      edges{
        node{
          id
          name
          availableEcSlot
          availableBcSlot
          departAt
          landAt
        }
        cursor
      }
      pageInfo{
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      totalCount
    }
  }
}
```

## BOOKING
### Create Customer Booking (one way) (not login)
After finding suitable flight (search flight), customer can book a flight
```graphql
mutation create_customer_booking{
  booking{
    create_customer_booking(input:{seatType:EC, flight_id:2, customer_id: 1}){
    createdAt
    updatedAt
    code
    flightID
    hasFlight{
      name
      availableEcSlot
      availableBcSlot
    }
    hasCustomer{
      id
      fullName
      memberID
    }
  }
  }
}
```
### Create Customer booking round trip
After choosing round trip, customer find flight departure flight and return flight and can reserve 2 slot
```graphql
mutation create_customer_booking_round_trip{
  booking{
    create_customer_booking_round_trip(input:{
    seat_type_arrive:EC,
    flight_id_arrive:2,
    customer_id:1
    seat_type_comeback:BC
    flight_id_comeback:1
  }){
    seatType
    id
    code
    hasCustomer{
      fullName
    }
    hasFlight{
      id
      name
      availableEcSlot
      availableBcSlot
    }
  }
  }
 
}
```
### Create Member Booking (one way)
Role: MEMBER
```graphql
mutation create_member_booking{
  booking{
    create_member_booking(input:{seatType:EC, flight_id:3}){
    createdAt
    updatedAt
    code
    flightID
    hasFlight{
      name
      availableEcSlot
      availableBcSlot
    }
    hasCustomer{
      id
      fullName
      memberID
    }
  }
  }
}
```
### Create Member Booking Round Trip 
Role: MEMBER
```graphql
mutation create_member_booking_round_trip{
  booking{
    create_member_booking_round_trip(input:{
    seat_type_arrive:EC,
    flight_id_arrive:7,
    seat_type_comeback:BC
    flight_id_comeback:2
  }){
    seatType
    id
    code
    hasCustomer{
      fullName
    }
    hasFlight{
      id
      name
      availableEcSlot
      availableBcSlot
    }
  }
  }
  
}
```

### View Booking History
Role: MEMBER
```graphql
mutation{
  booking{
    view_booking_history{
    id
    code
    status
    seatType
    hasFlight{
      fromAirport{
        name
      }
      toAirport{
        name
      }
      departAt
      landAt
    }
    hasCustomer{
      fullName
    }
  }
  }
  
}
```

### Search Booking (for customer does not log in)
```graphql
mutation search_booking{
  booking{
    search_booking(input: {flight_id:2, booking_code:"52KJZYGK", cid:"1354657683453"}){
    code
    seatType
    hasFlight{
      fromAirport{
        name
      }
      toAirport{
        name
      }
      departAt
      landAt
    }
  }
  }
}
```

### Cancel booking 
```graphql
mutation cancel_booking{
  booking{
    cancel_booking(input:{flight_id: 1, booking_code: "11DNKFGT", cid: "12345234754"})
  }
}
```
### Update Booking Status Cancel when flight had been canceled
Return total booking was canceled
```graphql
mutation update_bookings_status_cancel{
  booking{
    update_bookings_status(flight_id: 1)
  }
}
```

### List Bookings
Role: ADMIN
```graphql
mutation list_bookings{
  booking{
      list_bookings(first: 5, orderBy:{direction: ASC, field: CUSTOMER_ID}){
     		edges{
          node{
            id
            createdAt
            updatedAt
            code
            status
            seatType
            isRound
          }
          cursor
        }
      pageInfo{
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      totalCount
    }
    
  }
}
```
























































