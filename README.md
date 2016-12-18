# pwcalc

#### Command line utility to calculate strong cryptographic passwords.

Also available on [Android](
https://play.google.com/store/apps/details?id=de.pmorjan.TPSM
)
and [Google Chrome](
https://chrome.google.com/webstore/detail/gfdaokhcoidlgljipinbbcafapmmdanb
).

## Installation
```
go get github.com/pmorjan/pwcalc
```

## Usage
```
pwcalc [alias [pw_length]]
```

## Example
```
$ pwcalc gmail.com
# enter secret:
SWX
XmMvBXYLVWOh2a2e
```

## Description

Passwords are calculated as follows:
````
password = secret word + alias | SHA1 | Base64
````
You chose a single secret and an alias for the service the password is needed for.
E.g:
````
mysecret + gmail.com
````
This is combined to a single string
````
mysecretgmail.com
````
Next the cryptographic hash function SHA1 is applied to this compound word.
It produces a 160-bit hash value typically rendered  as hexadecimal number:
````
5e63 2f05 760b 5563 a1d9 ad9e e555 540d 7c19 9c65
````

These 40 digits are then taken as (binary) input for the
Base64 function which produces the password:

````
XmMv BXYL VWOh 2a2e 5VVU DXwZ nGU=
````
This unique password cannot by any means be traced back to the secret word or alias.

In plain Linux or MacOS the password can be calculated like this:
````
   echo -n "mysecretgmail.com" \
       | shasum \
       | xxd -r -p \
       | base64
    XmMvBXYLVWOh2a2e5VVUDXwZnGU=
````
