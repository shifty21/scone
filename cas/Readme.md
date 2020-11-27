

load cas related config
required
previous hash
cert
key
session_name
cas address
session file upload and extend
get cas predecessor session hash and request body for posting session
get cas response json

HTTP/1.1 200 OK
{
  "session": "---\nversion: \"0.3\"\nname: vault-init\npredecessor: ~\nimages:\n  - name: \"sconecuratedimages/www2019:vault-1.5.0-alpine\"\n    injection_files:\n      - path: /resources/vault-init/mykey.pem\n        content: \"-----BEGIN RSA PRIVATE KEY-----\\nMIIEpgIBAAKCAQEA1OSWQCTN/hgtYEyqh5VgngUDI+qBtyTbiVTdMM+ebMdA6YiV\\n+7pWuuR20D5xrx7w1c3RQdjwpzfQkt//kqTBJsQn1/rU2aSzm5XPhILiq1jui0rj\\neyGk1IW+u0mjAEa7jA02m2ozp30qxlxADuLWKSIWz2xvl3WZJ4bEGS7q7gcdjv8Z\\nKG4QibivmzPrENQ8njF86jNIvv0iiyUVR81qzy33+nxfWV9mQpIakBeMPzVniOxF\\nyfPrSSh0Io4w9UAa62CCdXei+6P4aMGTOt8I+4DTyoEcRdHdnuQ9FKmrqdi2NqsJ\\nJyr/zBQAlVX0Gzc3Adm49rQgA3vLX7gaXTIW9wIDAQABAoIBAQC53oEXq4p1Z6Jm\\nS0kvasmZ7QJa3yk1Pua1NfSP4xSMIEKaIffgeUWzkjfxhDM5E6hs4m7qMH+bXu2o\\n7gxyeYlxUR0AQiyHgHaXReqR5LwFoXVTA6UsIamJKuPlHFtFqHuhwP+GHOjQOEWa\\nPXxoAr+71dlYa3HaKH/4tH6NBhtyiAcHiaP23Ta44iSiuevyrANS+I3OeaH7DbjS\\nlaotG6XFvs57fAMNUVuH5GnU5RqETFMjLE6+XBYZBVPQGugnytzubiwjPg/1lvUr\\nlm5NBEnLI7UZGAna78xwG8hlJnHsOJZXxznBU1BUoVT1mlJWuwRgCcJoHP1c6GDG\\nUu6qq1ShAoGBAPuy2tiNeVsEVtQPm9sINB2YwJIYM5txm4YNkzVRklu05PdTGIXQ\\nbehmKCfEqIZfB50MUjO2jq+Ib5gJ8xY8V1mHoC7JbdWgH8etPCXyYguMk3OS5iaV\\nIRvyCAV3mR1f+UEIhm0M1iCTqbQ2uJE3kChZBdA9B+4incgts8kNsAGZAoGBANiH\\n9mv3V1RTK3s4/YhXDEdK+CqUsB2+IV9a1KJb1rNb8diCw0DfQ92DEzAFhXxTfbE4\\nGwhtNIRmCMpj3cI0CEbdWlLJ7KuY7IqPtmIXZUYGILX27CJEFA7RgJN9ZKM3KzYB\\nC7xb2ICDNMEBcKeCMPiv1WkPuLc7/0iPFrg+b1cPAoGBAIgcABx65NVDU3D5v96C\\nYSxgHkLis4Wrud6UGLcMlYjiGa9lUC2MuOKj27Mltbx0Rzm2H/23CxIBRdeCCeJM\\nXzAbF5Q1eR+8p3LjS1N572svac1l8u+KVY03JP9P3Yz1CWURpx/xgRm8wFij/ssI\\nsPwgp/QkDNXKAmjtzfs7W6KBAoGBAJqoclAneI7YYORAjoZFdpWtbJgtX6W+2eNb\\n4yicZDvz3kgBDilVzwl2x8uzBecJU2uzYUuhhLNUlc7JieledNL447ziUVM3hSxq\\n/aAOid18Hv0ZgwvuiE0VQrsWAz249/o4wQMmrvsLvDBnMOnUSdo27T1/ZzYpemt/\\nGIE8xxXtAoGBAIB8pFxFa5/fyAiKEaRsmJM1rzb82kyC8LV4CwjKRa4h91VVOGlg\\nRDDFyHwa9GdPHBjqIn/+2nM9epXhYKhrlMf6b5gFy41xmQNEqx5+rkaiDvX5NIFC\\nXXBsjMK1c7q9mag23FqKlUkQUFZIxagFXm23NGhkaQOHTyyYtX/iXJVm\\n-----END RSA PRIVATE KEY-----\\n\"\nservices:\n  - name: dev\n    image_name: \"sconecuratedimages/www2019:vault-1.5.0-alpine\"\n    mrenclaves:\n      - a2fe7d15964b2a6f9fd62b954d8a797ded84c643a0c3838b0a74fbdec56aa853\n    environment:\n      SCONE_MODE: hw\n    command: /root/go/bin/vault-init grpc\n    pwd: /\n    persistency: None\naccess_policy:\n  read:\n    - CREATOR\n  update:\n    - CREATOR\nsecurity:\n  attestation:\n    mode: hardware\n    tolerate:\n      - debug-mode\n      - outdated-tcb\n    ignore_advisories: \"*\"\ncreator: \"-----BEGIN CERTIFICATE-----\\nMIIFejCCA2KgAwIBAgIJANI/KSgTIwOWMA0GCSqGSIb3DQEBCwUAMGwxCzAJBgNV\\nBAYTAlVTMRAwDgYDVQQIDAdEcmVzZGVuMQ8wDQYDVQQHDAZTYXhvbnkxETAPBgNV\\nBAoMCFNjb250YWluMQwwCgYDVQQLDANPcmcxGTAXBgNVBAMMEHd3dy5zY29udGFp\\nbi5jb20wHhcNMjAwODEzMDgxNTI3WhcNMjAwOTEzMDgxNTI3WjBsMQswCQYDVQQG\\nEwJVUzEQMA4GA1UECAwHRHJlc2RlbjEPMA0GA1UEBwwGU2F4b255MREwDwYDVQQK\\nDAhTY29udGFpbjEMMAoGA1UECwwDT3JnMRkwFwYDVQQDDBB3d3cuc2NvbnRhaW4u\\nY29tMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEArE9Lx6d9eAVIl7UN\\nKcgqLzYyrUxyyg2MAkDxHJ66wyUsMCKe6X5h1//y4kiaIn/jejyR3bW3QYlhYfw3\\n7bw3kd0P9gi71t886VxRGRxFEt2544FrUk7V1DAzS90zBWW1i1i6nT9YyuXpyRYh\\nURjTynLxczKpHi0pM89cEyiv5j/vWjPyjeWx6hUTTE+Z/sJO2zLxralJfQ8OOQfH\\njw0zPdLcGQ8EKLfazYvrCP9+5HgUZwYdyUOQYyfuumxRkxq7iTFmPTSw2unbBUNl\\njiXzcE9JpMOg9Dh2wQ8ceb5D4is+WAW95PxsHIWcLfNyiiXmVRtWh4EspyMMVCPT\\nIpCHlz5D63kj3kzNdXU6G1yIha70l1SdSqzF4+1Kt0NnnFTzOm/E8+GFdnBIb6Yn\\nX5aQscoDvLBAMKY8L7LWEpU+fmJl5QMuDWGTSFD3hfUy4Bg4dx7LXFOGpylCUV6w\\nLadz0v8MKmKockZoW7rzxpIXndJ8xcOR8wrYHE7lYMY3OLXjUTjyHrOlQCbvQMxC\\nCKnj1qJYkYLtaPhefHYPn5tgAEqorR54AZPgOHr43DYnQGXQaWbMJirdwCG4lQH6\\nxtMntt54zNislgKjztptlg5cwcHkfdiRflNgSu/Yzto7d2BHevn9lwACgTXwV0Yh\\nVecvvYCI9JE9QrlSW+ch6V4laQsCAwEAAaMfMB0wGwYDVR0RBBQwEoIQd3d3LnNj\\nb250YWluLmNvbTANBgkqhkiG9w0BAQsFAAOCAgEAbGn0FbtCafrwKzzCUWlEWY9z\\nVZqPszrpUPqvE9cTbwqzfs24BnCvhQm1e11uDurKHkdenWW/ykW4i2x/jMrhremb\\naZmoMw1XYaE5+tqCuWzdvn66JDPVqTflMOKe6OYWqibAZ08scSCbxE6LlI1WcDtH\\nyRW9MmqP2/CCi646+ci+ZxONJ4K7pWaBBrfWGIXePdxUXBllvsN935Jweww21fzt\\nDYhQbdAVle4JRolq3PBWO8qA+qGLYNFe1B3IGurFxH2nTeePzyjo2swcfoN7Rrn6\\nR3IhtjfPxR8BSSF0xjhMnGDgUTddj9MiSe3CxHI40V8z/PFKqqVt/YhqKnDphtZ+\\nIxQozzH79hvApwVSFMHTHJUaLex/jpti/KFLaXlazoxToLH6VLYgqRYA0YtBVYVI\\nSawolVrL00mPgeFEVKHXtcmKPtCGRMqHkZbyNY9baotJf2gfWYOUXoEo2YvxtFDm\\nESDSC91e4G0UUmYo6uu9rzrAh5U+LDTGDn9ptOq//W1EUZmqUj8HY5XKcqvt0ew7\\nKDxeRzqsPB8eZQeaVuiPbWBratmUGHmoNQ6xD3i5BVusaKMM5bgt6HWBuuhewWW9\\nkGsyTBFwU2335np5GVc/ApnWRj7AqgckcdPm3B2WPNRHnSVxwVs+WR/Nz0HgwiE2\\nZgb3OsjIa8g3ZdfPn4Y=\\n-----END CERTIFICATE-----\\n\""
}

HTTP/1.1 404 Not Found
content-type: application/json
content-length: 84
date: Tue, 17 Nov 2020 18:33:22 GMT

{
  "msg": [
    "Failed to read session 'vault-ini'",
    "Session not found"
  ]
}

In case wrong certificate is given

HTTP/1.1 403 Forbidden
content-type: application/json
content-length: 125
date: Tue, 17 Nov 2020 18:35:37 GMT

{
  "msg": [
    "Failed to read session 'vault-init'",
    "You have insufficient rights to access session vault-init"
  ]
}