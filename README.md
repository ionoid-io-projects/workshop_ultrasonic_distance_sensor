# Introduction


Material needed:

1) [Raspberry Pi 3 or Raspberry Pi Zero WiFi](https://www.raspberrypi.org/products/raspberry-pi-zero-w/)


2) [HC-SR04 Ultrasonic distance
sensor](https://www.amazon.fr/gp/product/B06XSJPVW9/ref=as_li_qf_asin_il_tl?ie=UTF8&tag=liens_hc_sr04-21&creative=6746&linkCode=as2&creativeASIN=B06XSJPVW9&linkId=1afc7a8494e21c3b883cd0b08620fb09)

![HC-SR04](img/HC-SR04.png)



3) [1KΩ et 2KΩ
Resistors](https://www.amazon.fr/gp/product/B071LHFQKD/ref=as_li_qf_asin_il_tl?ie=UTF8&tag=liens_hc_sr04-21&creative=6746&linkCode=as2&creativeASIN=B071LHFQKD&linkId=213100c4a53af021f1e4201c62fc61d2)

4)
[Cables](https://www.amazon.fr/gp/product/B01JD5WCG2/ref=as_li_qf_asin_il_tl?ie=UTF8&tag=liens_hc_sr04-21&creative=6746&linkCode=as2&creativeASIN=B01JD5WCG2&linkId=80fd9f40f891574b5404556a03de1b47)

5) [Power
supply](https://www.amazon.com/s?k=Raspberry+Pi+Power+Supply&language=en_US&linkCode=ll2&linkId=7dc5f7ea56c88b3efd0c35316b16e38e&tag=pimylifeup-20&ref=as_li_ss_tl)



The design

Mesure distance 

![wiring](doc/img/wiring-uds.png)

Follow link below for pin connections 
[https://projects.raspberrypi.org/en/projects/physical-computing/14](https://projects.raspberrypi.org/en/projects/physical-computing/14)

![wiring](doc/img/gpio.png)

# Code

## 1- Power your raspberry

You can achive it with connecting it to your pc trought the Micro USB Port of the raspberry pi

![power](doc/img/1-min.jpg)

## 2- Connect to your raspberry pi
Using putty if you're on windows, Ssh if you're on a linux based os
Follow the following instruction if you dont know how to connect to raspberry pi
[Connect to Raspberry Pi using Putty](https://github.com/ionoid-io-projects/workshop/blob/master/doc/od-iot-raspbian-rpi-zero-windows.md#5-first-boot)

## 3- Download the binary file

Assuming you're connected with... copy and past this command
If you're using Raspberry zero
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_ultrasonic_distance_sensor/master/bin/arm6/udistance
```

If you're using Raspberry 3 b
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_ultrasonic_distance_sensor/master/bin/arm7/udistance
```
## make it executable
```
chmod +x udistance
```

## 4- execute binary to make led blink
```
./udistance
```

## How to stop the program
To quit or stop the program click on **Ctrl+C**


Congratulation.

## For explanation
[https://www.modmypi.com/blog/hc-sr04-ultrasonic-range-sensor-on-the-raspberry-pi](https://www.modmypi.com/blog/hc-sr04-ultrasonic-range-sensor-on-the-raspberry-pi)

## Other references

[https://pimylifeup.com/raspberry-pi-distance-sensor](https://pimylifeup.com/raspberry-pi-distance-sensor/)
[http://espace-raspberry-francais.fr/Composants/](http://espace-raspberry-francais.fr/Composants/Mesure-de-distance-avec-HC-SR04-Raspberry-Francais/)
[https://github.com/rhiller/pi-distance](https://github.com/rhiller/pi-distance)
