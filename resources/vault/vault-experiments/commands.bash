vault operator generate-root -init
vault operator generate-root
vault operator generate-root \
  -decode=PEQKFApwIxY5cFQyLUBHB04gNyUVAyZrDng \
  -otp=OjsNO3DATFaYTw1a7UvdftASTM

# rekey 9k4xPwy3HdHkY1OyKnA6iau98sZdAURtjzv9Kzm3G9s=
vault operator rekey -init -key-shares=3 -key-threshold=2
vault operator rekey -nonce=55aa6ad5-a05c-65f7-4e7b-6fc3c1ea5920

# in case of recovery keys add -target=recovery
vault operator rekey -target=recovery -init -key-shares=1 -key-threshold=1 
vault operator rekey -target=recovery -nonce=55aa6ad5-a05c-65f7-4e7b-6fc3c1ea5920

#gpg keys
gpg --gen-key
gpg --import private.key

gpg --output output.gpg --encrypt --recipient yateenderk@gmail.com input
#jack.gpg should be base64 decode output of the key provided by vault, use the base64 response
echo "" | base64 --decode | > output.gpg
gpg --output decrypted_out --decrypt output.gpg 
#export public and private pgp key armored output 
gpg --export-secret-keys -a -o private.key yateenderk@gmail.com
gpg -a --export yateenderk@gmail.com > public.asc
gpg --list-keys shifty21
#export public and private pgp key base64 encoded output 
gpg --export-secret-keys  yateenderk@gmail.com | base64  > private.asc
gpg --export yateenderk@gmail.com | base64 > public.asc
#to import the base64 key first decode it and then import
cat private.asc | base64 --decode | >private.key
gpg --import private.key 