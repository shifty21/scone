path "pki/issue/scone" {
      capabilities = ["create", "update"]
}

path "pki/certs" {
    capabilities = ["list"]
}

path "pki/revoke" {
    capabilities = ["create", "update"]
}

path "pki/tidy" {
    capabilities = ["create", "update"]
}

path "pki/cert/ca" {
    capabilities = ["read"]
}

path "auth/token/renew" {
    capabilities = ["update"]
}

path "auth/token/renew-self" {
    capabilities = ["update"]
}
