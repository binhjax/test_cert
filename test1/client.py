import requests

r = requests.get("https://localhost:7252", verify="cert/root.pem")
print(r.content)
print(r.status_code)
