fbAlbumBackup
======================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/toomore/gogrs/master/LICENSE)

Facebook album tool that backup all photos in albums and supports concurrent download.

Install
--------------

    go get github.com/howeih/daytimeClient 


Examples
---------------
    daytime server list: http://tf.nist.gov/tf-cgi/servers.cgi
    ex:
	daytimer := daytimeClient.NewDayTimer("132.163.4.103")
	fmt.Println(daytimer.GetTime())
