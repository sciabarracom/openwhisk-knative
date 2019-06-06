using JSON

auth = "Authorization: Basic YmQyNWQ0NTQtZmE2My00Mzk4LWE4M2UtNTZlODg4MDEwZGU0OllVVmhWOGtuS0ZMVWJ2anc1VlJGblFZUk9rYWUwQjZnMzhmYmVlQ3FkbmRQc1pTTnREZ0k0N1hwZHBTRW8wMFg="

hosts  = ["https://openwhisk.eu-gb.bluemix.net",
          "http://localhost:8000"]

// `curl -H $auth $host/api/v1/namespaces/_/actions\?limit=30\&skip=0`

actions(host) = JSON.parse(read(`curl -sH $auth $host/api/v1/namespaces/_/actions`, String))

r = actions(hosts[2])

print(join([ i["name"] for i in resp], "\n"))

print(resp[0])

resp[1]