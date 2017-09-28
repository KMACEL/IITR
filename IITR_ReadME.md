### Report :

adr | alldevicereport : Bu parametre report kısmına geçtikten sonra eğer bütün cihazlardan rapor alınması isteniyorsa kullanılır. Bu konsol üzerinden cihazlarda bakılması istenilen uygulamaların paket ismi verilmesi gereklidir. Bu isimler aralarında "," ile yazılmalıdır.

Örn :
tr.com.innology.taksipager,com.ardic.android.iot.appblocker,com.android.launcher3,com.google.android.apps.maps,com.streamaxtech.mdvr.direct

dr | devicereport : Bu parametre sizden json formatlı bir dosyanın yolunu ister. Bu dosyada bakılması istenilen cihazların numaraları, uygulamaların paket isimleri istenir. Format şu şekildedir :

#### Report Case
> 	{
	"type" : "report",
	"case" :{
			"name" :"detail-report-devices",
			"devices":["867377020746784","867377020747089"],
			"packages":[
				"tr.com.innology.taksipager",
				"com.ardic.android.iot.appblocker",
				"com.android.launcher3","com.google.android.apps.maps",
				"com.streamaxtech.mdvr.direct"]
		}
	}


### Test
#### Test Case

> 	{
	  "type": "test",
	  "code": "new-test",
	  "name": "drom-mode",
	  "devices": [
	    "867377020746784",
	    "867377020747089"
	  ],
	  "case": [
	    {
	      "loop": -1,
	      "steps": [
	        "drom::d&5",
	        "reboot::d&10",
	        "changemode::default-default",
	        "drom::d&?0-5",
	        "getlog::d&?0-5",
	        "screenshot::d&0",
	        "startapp::a&tr.com.innology.taksipager|d&5",
	        "stopapp::a&tr.com.innology.taksipager|d&?0-3",
	        "cleardata::a&tr.com.innology.taksipager|d&5",
	        "wipe::d&15"
	      ]
	    }
	  ]
	}
