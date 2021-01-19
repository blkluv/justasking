
#!/bin/bash

(cd /home/sebastian_chande/startup; sudo python sendemail.py "URGENT: VM HAS BEEN RESTARTED." "THE VM HAS BEEN RESTARTED AND ASSIGNED A NEW IP. UPDATE THE DATABASE PROJECT IN GCP TO AUTHORIZE THE NEW IP OR THE SYNCSERVICE WILL NOT CONNECT TO THE DATABASE.")

lastDay=$(date +"%u")
while true ; do

    #once a day, send us an email notifying us that the service is up
    day=$(date +"%u")
    if [ $day -ne $lastDay ]; then
      (cd /home/sebastian_chande/startup; sudo python sendemail.py "service monitor is up" "service monitor is up")
      lastDay=$(date +"%u")
    fi

    realtimePid=$(pgrep -f realtimehub)
    if [[ -z "$realtimePid" ]]; then
        (cd /home/sebastian_chande/startup; sudo python sendemail.py "Service Restart - realtimehub" "Realtime hub was down.")
        (cd /home/sebastian_chande/go/src/justasking/GO/realtimehub; sudo nohup ./realtimehub >/dev/null 2>&1 &)
    fi

    syncPid=$(pgrep -f syncservice)
    if [[ -z "$syncPid" ]]; then
        (cd /home/sebastian_chande/startup; sudo python sendemail.py "Service Restart - syncservice" "syncservice was down.")
        (cd /home/sebastian_chande/go/src/justasking/GO/syncservice; sudo nohup ./syncservice >/dev/null 2>&1 &)
    fi

    sleep 60
done