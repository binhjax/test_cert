## Test 1
go run generate.go  --host=localhost,127.0.0.1


### Client test
curl --cacert cert/root.pem https://localhost:7252
openssl s_client -showcerts -servername localhost -CAfile cert/root.pem -connect localhost:7252


## Test 2
go run generate_cert.go  --host=localhost,127.0.0.1
# test_cert
