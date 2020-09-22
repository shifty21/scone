init_status 
curl \
    http://192.168.0.143:8200/v1/sys/init
{"initialized":false}



initialize
curl \
    --request PUT \
    --data @payload.json \
    http://192.168.0.143:8200/v1/sys/init

    {"keys":["ea256d6985dc9c59941e114abebf7dccede43ea00b505db1f8342203baa85d2232","649fde948d9b15ab4c775a5a6e7922e61bc81209245ef2a08fa404ce2e215c539e","aabb2b3cb10adf461eb10951e747f54ff068ffe359f5a5591add7adbbd6509fab1","2d4d0375cc7d2904cd5d70a08f94bea77492e8d0e2360c78c0c18ed010909565ae","30366d8eb21d7eb333a8fb4b1fdc20a7da1957828d5353649e256ebe549a6b3a60"],"keys_base64":["6iVtaYXcnFmUHhFKvr99zO3kPqALUF2x+DQiA7qoXSIy","ZJ/elI2bFatMd1pabnki5hvIEgkkXvKgj6QEzi4hXFOe","qrsrPLEK30YesQlR50f1T/Bo/+NZ9aVZGt16271lCfqx","LU0Ddcx9KQTNXXCgj5S+p3SS6NDiNgx4wMGO0BCQlWWu","MDZtjrIdfrMzqPtLH9wgp9oZV4KNU1NkniVuvlSaazpg"],"root_token":"s.wNbwzne7A4YoeXqPB5SraNkJ"}



unseal 
curl \
    --request PUT \
    --data @key1_payload.json \
    http://192.168.0.143:8200/v1/sys/unseal

    curl \
    --request PUT \
    --data @key1_payload.json \
    http://192.168.0.143:8200/v1/sys/unseal

        curl \
    --request PUT \
    --data @key3_payload.json \
    http://192.168.0.143:8200/v1/sys/unseal