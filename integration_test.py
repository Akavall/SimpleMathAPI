
import requests

result = requests.get("http://34.208.248.130:8088/is_prime?num=13")

if result.content == "true":
    print "ok"
else:
    print "FAIL"

result = requests.get("http://34.208.248.130:8088/is_prime?num=14")
if result.content == "false":
    print "ok"
else:
    print "FAIL"

