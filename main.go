package main

import (
  "fmt"
  "github.com/kklash/gogpg/encryption"
  "github.com/kklash/gogpg/decryption"
  "github.com/kklash/gogpg/signatures"
  "log"
)

const msg string = `-----BEGIN PGP MESSAGE-----

hQEMAz9PGogu1LRZAQf/YccexRGWEW+XV6kKW5Bom+zMzym+aKoZus6rU83RDfB4
I2DpNV4H/Y6DA++mEt0ZlmWHhkNEZ7kYJiuOt256VuYzlIVWMpE9ikhvHIFbHz/5
1FuzP2kDBi4Y4rislg1MIqpH4i9kHQ1UiMZDFUvB2yoCk1tgAyzgLm0VvQ9UdKVL
blKQ4lhxnyoobm/HLo9P2he+la0YDZwmRhx4Fb8wbS8Ohu4mcDBRoUgZGVx8EROn
DB9W7qYsCHSf+jIWM5274iwWWWO3I+UQi4Fm8Fk1Ti6eys9hl3qiAUs5VG0ulLeF
E+WeZffZItVLS5GtCA85FSVRKfNpTXhWGz18v7wDCdKvAQA6Zs36EQLDSvN6SOu6
jIIPbskqb33wfBN/bEvDUFszCnnzvcEwUp5O0+5vN40AyrX7umlKztISfMx2mMsc
p3paLayRmXMPE8Yod9D0aWlgzPruSaHtC7Q5Ze3ouPYNr0KKou4hwnOHXlNIgXug
XD34qMW9R0vxwJOSiyX1UqXgFEBE5yVS7BBrhY4qobQc/FFzjDVIzzD/c2Lq5KUA
5Hdfb7FVaPlGA1mTopbksw==
=7TRt
-----END PGP MESSAGE-----`


const sig string = `-----BEGIN PGP SIGNATURE-----

iQIzBAABCAAdFiEEoSzb4/FqhMWL7jL/0zQZ+ujtCgkFAlrQfVMACgkQ0zQZ+ujt
CgmL9w//egt0jgY9Qlepg4+arZJ10KaQA3CbXAJA7IVVZmRdnd5WFQPPV1sZx/J/
r62L6UvS9tZ6jYnTgSiAqEZ0TSFdwzhGXTnT4QJzgezCOZ8lHWuONp2eDoxN3uxX
NZDnEyVBwf+iipVlDN3IrtbYXwGnR6F1Poi02pWua/G/i6/ZQOmdjosk4VWi2Aeo
6p/16wdfRa4EBkG218//bm9V/UtWIFOGHwdKQB0OrW8H/cks1P6hSaAAB6QqjfY6
Ny520UZkzb7x5rwszk7bX+M/ZIYiKHEFko1Dob2DHuMRYMDdQM+5FS2dsmMdvOCr
8SEQs4OjsJr7dlPW8xQejgUWBUoRUkUmeKHCOaokVcK3eU5EmbdiUy1HK/2j8Y2M
VOoST3Tys3pgtC9/89TspQsBXcCs0qaJANcC0u8schnAG4wSxvcGndZdoBwetC8X
BjjwfOkCguetIn4AkFnWUC4OocJjgju6e5BY/0bu15Wueh20MGowEcdaGl+gtao9
RFYDLMAuaxNrQpOtP51w7ijJcBL3A8tmjfxEyySxE3nLCOxNwcWvLY6GtT6O71Op
RefAM0/4AkeDrBl4eQDPTiAhU4p10HlV7i4Viv7xiSDO9JObNLfTPzI8DNnpJ8jm
f8ULBdiJ2I4sn3fuVc0qhqHADY6Wbw8jG9wJJOsL1T360nzQNzk=
=8NrG
-----END PGP SIGNATURE-----`




func main() {
  var (
    output string
    err error
    result bool
  )
  
  output, err = encryption.Encrypt("hello world", "ariana")
  if err != nil { log.Fatal(err) }
  fmt.Printf("%d ciphertext bytes encrypted from string\n", len(output))
  
  output, err = encryption.EncryptFile("/Users/admin/.bash_profile", "konnor", "ariana")
  if err != nil { log.Fatal(err) }
  fmt.Printf("%d ciphertext bytes encrypted from file\n", len(output))
  
  output, err = decryption.Decrypt(msg)
  if err != nil { log.Fatal(err) }
  fmt.Printf("%d plaintext chars decrypted from string\n", len(output))
  
  output, err = decryption.DecryptFile("/Users/admin/Desktop/cip.txt")
  if err != nil { log.Fatal(err) }
  fmt.Printf("%d plaintext chars decrypted from file\n", len(output))

  output, err = signatures.SignDetached("Hi there, it's konnor!", "ksecc")
  if err != nil { log.Fatal(err) }
  fmt.Printf("Signature of length %d bytes generated\n", len(output))
  
  result, err = signatures.VerifyDetached("Hi there, it's konnor!", sig)
  if err != nil { log.Fatal(err) }
  fmt.Printf("Above signature verifies as: %v\n", result)
  
  output, err = signatures.Sign("Hi there, it's konnor!", "ksecc")
  if err != nil { log.Fatal(err) }
  fmt.Printf("Signed message of length %d made\n", len(output))
  
  result = signatures.Verify(output)
  fmt.Printf("Above signature verifies as: %v\n", result)
  


}