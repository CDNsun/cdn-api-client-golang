# Client for CDNsun CDN API

SYSTEM REQUIREMENTS

* Go >= 1.6

CDN API DOCUMENTATION

https://cdnsun.com/knowledgebase/api

BUILD

To build the **test.go** go to the directory containing it and run the following command:
```  
go build
```

CLIENT USAGE

* Initialize the client
```
import "cdnsun"

apiClient, err := cdnsun.New('YOUR_API_USERNAME', 'YOUR_API_PASSWORD')

```

* Get CDN service reports (https://cdnsun.com/knowledgebase/api/documentation/res/cdn/act/reports)
```
  result, err := apiClient.Get(&cdnsun.Options{
    Url: "cdns",
  })

```
* Purge CDN service content (https://cdnsun.com/knowledgebase/api/documentation/res/cdn/act/purge)

```
result, err = apiClient.Post(&cdnsun.Options{
    Url: "cdns/" + ID + "/purge",
    Data: map[string]interface{}{
      "purge_paths": []string{
        "/path1.img",
        "/path2.img",
      },
    },
  })

```

NOTES

* The ID stands for a CDN service ID, it is an integer number, eg. 123, to find your CDN service ID please visit the Services/How-To (https://cdnsun.com/cdn/how-to) page in the CDNsun CDN dashboard.

CONTACT

* W: https://cdnsun.com
* E: info@cdnsun.com  