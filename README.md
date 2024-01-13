# System-Designs
System design for the application. 


Drivers & Riders 

Problems: 
    Handle large number of users 
    Lots of realtime data 

        HTTP polling 
        Web Sockets 


Problems: 

    Latency 
    Pulls API Drivers location. 


    * Websockets will help to maintain persistent connection, this helps to bring driver location on large scale.


By integrating the problems: 

New updated the System design will be: 


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