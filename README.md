# odata-batch
OData REST API batch operation demo

Usage

For single operation:
curl \
    --data '{"Value": 100}' \
    -X POST localhost:3000/calculator/Actions/Calculator.Add

For batch operation:
curl \
     -H 'Content-Type: multipart/mixed;boundary=$boundary' \
     --data-binary batch_1111-2222-3333-4444
     -F 'request={"Value": 1};type=application/http'
     -F 'request={"Value": 2}'
     -X POST localhost:3000/calculator/$batch