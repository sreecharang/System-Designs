Databases:

* Database will need to hold certain values. 


TRIP 

    ID 
    DRIVER ID 
    RIDER ID 
    PRICE 
    PICKUP LOCATION 
    DESTINATION 
    DATE 
    STATUS 
    RIDE TYPE 

DRIVER 

    ID 
    RIDE TYPE 
    CAR INFO 
    LOCATION 
    PHONE NUMBER 
    EMAIL 
    PASSWORD HASH 
    NAME 


RIDER 

    ID 
    PAYMENT INFO 
    PHONE 
    EMAIL 
    PASSWORD HASH 
    NAME 


Its really important to scale database horizontally. 

* Global Index will get the RiderID and DriverID, Sort the data by status. 



MAP: 

Problems: 

* Serve Map images 
* Convert street address to lat/long
* Get Directions. 


Bunch of services for Map 


Payments: 

Problems: 
* Rider needs to pay API neets to know when payment finishes processing. 


All services need to attach to client side. 


Pricing: 

* Pricing varies by demand batch processing would be too slow. 
* 