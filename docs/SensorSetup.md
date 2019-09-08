# Thermal sensor
- sensor type: DS18B20
## Setup RPi
### Setup W-1
Hence the sensor communicates with W-1, it has to be setup on the RPi  
1. Open the Terminal
```sudo nano /boot/config.txt```   
paste the end of the file: 
```dtoverlay=w1-gpio```   
This will add the W-1 interface. Default data pin is BCM 4(7).  
1. Add to the modules.
```sudo modprobe w1-gpio```  
```sudo modprobe w1-therm```
1. Check if it is there
```cd /sys/bus/w1/devices/```
```ls```  
it should show an element with a long serial number. This is our connected sensor.
1. Check if it is working
```cd 28-01144609dcaa``` 
```cat w1_slave ```

It shoudl have an output similar to this: 
```
9a 01 4b 46 7f ff 0c 10 9f : crc=9f YES
9a 01 4b 46 7f ff 0c 10 9f t=25625
```
### Multiple 1-worw bus setup
Type this in the /boot/config.txt

```
dtoverlay=w1-gpio,gpiopin=4
dtoverlay=w1-gpio,gpiopin=17
```
